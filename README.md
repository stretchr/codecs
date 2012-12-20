#codecs

Provides interfaces, functions and codecs that can be used to encode/decode data to/from various formats.

A proper README will be written in the future. At the moment, we're a bit too busy to flesh everything out. The package is pretty easy to use, however, so have fun!

## How to use the codecs package

	  // make a codec service
    codecService := new(WebCodecService)

    // get the content type (probably from the request)
		var contentType string = "application/json"

		// get the codec
    codec, err := codecService.GetCodec(contentType)

    if err != nil {
    	// handle errors - specifially ErrorContentTypeNotSupported
    }

    /*
    	[]bytes to object
    */

		// get the raw data
		dataBytes := []byte(`{"somedata": true}`)

    // use the codec to unmarshal the dataBytes
    var obj interface{}
    unmarshalErr := codec.Unmarshal(dataBytes, obj) error

    if unmarshalErr != nil {
    	// handle this error
    }

    // obj will now be an object built from the dataBytes

    /*
    	object to []bytes
    */

    // get the data object
    dataObject := map[string]interface{}{"name": "Mat"}

    bytes, marshalErr := codec.Marshal(dataObject, nil)

    if marshalErr != nil {
    	// handle marshalErr
    }

    // bytes would now be a representation of the data object

## Get started

Get the codecs package