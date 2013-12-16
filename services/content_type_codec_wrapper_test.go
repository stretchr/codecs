package services

import (
	"github.com/stretchr/codecs"
	"github.com/stretchr/codecs/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapCodec(t *testing.T) {
	codec := new(json.JsonCodec)
	testContentType := "application/vnd.stretchr.test+json"
	var target interface{} = wrapCodec(codec, testContentType)

	wrappedCodec, ok := target.(codecs.Codec)
	assert.True(t, ok, "A wrapped codec should still be a Codec")

	if ok {
		assert.Equal(t, testContentType, wrappedCodec.ContentType())
	}
}
