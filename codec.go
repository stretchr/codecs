package codecs

// Codec is the interface to which a codec must conform.
type Codec interface {

	// Marshal converts an object to a []byte representation.
	// You can optionally pass additional arguments to further customize this call.
	Marshal(object interface{}, options map[string]interface{}) ([]byte, error)

	// Unmarshal converts a []byte representation into an object.
	Unmarshal(data []byte, obj interface{}) error

	// ContentType gets the default content type for this codec.
	ContentType() string

	// FileExtension returns the file extension by which the codec is represented.
	FileExtension() string

	// CanMarshalWithCallback gets whether the codec is capable of marshalling a response with
	// a callback parameter.
	CanMarshalWithCallback() bool
}

// ContentTypeMatcherCodec is a Codec that has a method to be used for
// matching content types against any content types that the codec can
// support.
type ContentTypeMatcherCodec interface {
	Codec

	// ContentTypeSupported returns true if the passed in ContentType
	// string is supported by this codec, false otherwise.
	ContentTypeSupported(string) bool
}
