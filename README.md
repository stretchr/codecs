#codecs

Provides interfaces, functions and codecs that can be used to encode/decode data to/from various formats.

## Documentation

You can get stuck into the API documentation by checking out these API docs:

  * [stretchrcom/codecs](http://godoc.org/github.com/stretchrcom/codecs)
  * [stretchrcom/codecs/services](http://godoc.org/github.com/stretchrcom/codecs/services)

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




------

Contributing
============

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include steps to reproduce the issue so we can see it on our end also!


Licence
=======
Copyright (c) 2012 Mat Ryer and Tyler Bunnell

Please consider promoting this project if you find it useful.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
