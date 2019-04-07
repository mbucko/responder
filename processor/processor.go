package processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Export struct {
}

func (c Export) Process(rw http.ResponseWriter, request *http.Request) {
	fmt.Printf("Processing message: %s", request.Body)
	processedMsg := process(request.Body)
	rw.Write([]byte(processedMsg))
}

func renameField(jsonMap *map[string]interface{}) {
	for k, v := range *jsonMap {
		switch vCast := v.(type) {
		case string:
			if k == "toCaps" {
				v2, _ := processString(vCast)
				(*jsonMap)[k] = v2
			} else if k == "crash" {
				fmt.Println("Received crash signal, going to panic.")
				panic("I am crashing! The payload told me to.")
			}
		case map[string]interface{}:
			renameField(&vCast)
		case []interface{}:
			for _, u := range vCast {
				var tmp = u.(map[string]interface{})
				renameField(&tmp)
			}
		default:
		}
	}
}

func process(r io.Reader) string {
	decoder := json.NewDecoder(r)

	var decodedInput map[string]interface{}
	err := decoder.Decode(&decodedInput)
	fmt.Printf("Decoding string: %s", decodedInput)

	if err != nil {
		return "Invlid Input\n"
	}

	renameField(&decodedInput)

	jsonString, err := json.Marshal(decodedInput)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Altered JSON string: %s", jsonString)
	return string(jsonString) + "\n"
}

func processString(input string) (string, error) {
	if input == "" {
		return "Invalid Input\n", errors.New("Input string is empty")
	}
	return strings.ToUpper(input), nil
}
