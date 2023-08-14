package serialization

import "encoding/json"

func Serialization(obj interface{}) (string, error) {
	b, e := json.Marshal(obj)
	return string(b), e
}

func Deserialization(str string) (interface{}, error) {
	var obj interface{}
	b := []byte(str)
	e := json.Unmarshal(b, &obj)
	return obj, e
}
