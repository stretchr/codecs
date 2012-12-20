package services

import (
	"github.com/stretchrcom/codecs"
)

// CodecService is the interface for a service responsible for providing Codecs.
type CodecService interface {
	// Setup is called when the services are first created in order to
	// perform any initialisation work to prepare the service for use.
	Setup() error

	// TearDown is called when the services have come to the end of their
	// life.  They should use this method to release any resources or close
	// any connections before the service shuts down.
	TearDown()

	// GetCodecForResponding gets the codec to use to respond based on the
	// given accept string, the extension provided and whether it has a callback
	// or not.
	GetCodecForResponding(accept, extension string, hasCallback bool) (codecs.Codec, error)

	// GetCodecForRequest gets the codec to use to interpret the request based on the
	// content type.
	GetCodecForRequest(contentType string) (codecs.Codec, error)

	// MarshalWithCodec marshals the specified object with the specified codec and options.
	// If the object implements the Facade interface, the PublicData object should be
	// marshalled instead.
	MarshalWithCodec(codec codecs.Codec, object interface{}, options map[string]interface{}) ([]byte, error)

	// UnmarshalWithCodec unmarshals the specified data into the object with the specified codec.
	UnmarshalWithCodec(codec codecs.Codec, data []byte, object interface{}) error
}
