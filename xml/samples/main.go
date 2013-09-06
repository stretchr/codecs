package main

import (
	"github.com/stretchr/codecs/xml"
	"log"
)

var xmlCodec xml.SimpleXmlCodec

func main() {

	// make a big object
	obj := map[string]interface{}{}
	obj["name"] = "Mat"
	obj["age"] = 30
	obj["address"] = map[string]interface{}{
		"street":  "Pearl Street",
		"city":    "Boulder",
		"state":   "CO",
		"country": "USA",
	}
	obj["animals"] = map[string]interface{}{
		"favourite": []string{"Dog", "Cat"},
	}

	bytes, err := xmlCodec.Marshal(obj, nil)

	if err != nil {
		log.Printf("Failed to marshal simple XML: %s", err)
	} else {
		log.Printf("%v", string(bytes))
	}

}
