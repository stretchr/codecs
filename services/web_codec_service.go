package services

import (
	"errors"
	"fmt"
	"github.com/stretchrcom/codecs"
	"github.com/stretchrcom/codecs/bson"
	"github.com/stretchrcom/codecs/constants"
	"github.com/stretchrcom/codecs/json"
	"github.com/stretchrcom/codecs/jsonp"
	"github.com/stretchrcom/codecs/msgpack"
	"strings"
)

// ErrorContentTypeNotSupported is the error for when a content type is requested that is not supported by the system
var ErrorContentTypeNotSupported = errors.New("Content type is not supported.")

// InstalledCodecs is an array of installed codec objects, initialized with the provided default codecs.
var InstalledCodecs []codecs.Codec = []codecs.Codec{new(json.JsonCodec), new(jsonp.JsonPCodec), new(msgpack.MsgpackCodec), new(bson.BsonCodec)}

// WebCodecService represents the default implementation for providing access to the
// currently installed web codecs.
type WebCodecService struct{}

func (w *WebCodecService) Setup() error {
	return nil
}

func (w *WebCodecService) TearDown() {

}

// GetCodecForResponding gets the codec to use to respond based on the
// given accept string, the extension provided and whether it has a callback
// or not.
//
// As of now, if hasCallback is true, the JSONP codec will be returned.
// This may be changed if additional callback capable codecs are added.
func (s *WebCodecService) GetCodecForResponding(accept, extension string, hasCallback bool) (codecs.Codec, error) {

	if hasCallback == true {
		for _, codec := range InstalledCodecs {
			if codec.ContentType() == constants.ContentTypeJSONP {
				return codec, nil
			}
		}
	}

	for _, codec := range InstalledCodecs {
		if strings.Contains(strings.ToLower(accept), strings.ToLower(codec.ContentType())) {
			return codec, nil
		} else if strings.ToLower(codec.FileExtension()) == strings.ToLower(extension) {
			return codec, nil
		} else if hasCallback && codec.CanMarshalWithCallback() {
			return codec, nil
		}
	}

	return InstalledCodecs[0], nil
}

// GetCodec gets the codec to use to interpret the request based on the
// content type.
func (s *WebCodecService) GetCodec(contentType string) (codecs.Codec, error) {

	for _, codec := range InstalledCodecs {

		// default codec
		if len(contentType) == 0 && codec.ContentType() == constants.ContentTypeJSON {
			return codec, nil
		}

		// match the content type
		if strings.Contains(strings.ToLower(contentType), strings.ToLower(codec.ContentType())) {
			return codec, nil
		}

	}

	return nil, errors.New(fmt.Sprintf("Content type \"%s\" is not supported.", contentType))

}

// MarshalWithCodec marshals the specified object with the specified codec and options.
// If the object implements the Facade interface, the PublicData object should be
// marshalled instead.
func (s *WebCodecService) MarshalWithCodec(codec codecs.Codec, object interface{}, options map[string]interface{}) ([]byte, error) {

	// get the public data
	publicData, err := codecs.PublicData(object, options)

	// if there was an error - return it
	if err != nil {
		return nil, err
	}

	// let the codec do its work
	return codec.Marshal(publicData, options)
}

// UnmarshalWithCodec unmarshals the specified data into the object with the specified codec.
func (s *WebCodecService) UnmarshalWithCodec(codec codecs.Codec, data []byte, object interface{}) error {
	return codec.Unmarshal(data, object)
}
