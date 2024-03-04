package helpers

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close() // defer execution of this until surrounding functions returns
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &target)
}
