package json

import "encoding/json"

// Encode json序列化
func Encode(v interface{}) (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

// Decode json反序列化，obj为指针
func Decode(buf []byte, obj interface{}) error {
	return json.Unmarshal(buf, obj)
}
