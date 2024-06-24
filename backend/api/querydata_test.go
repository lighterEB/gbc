package api

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	products, err := QueryData()
	if err != nil {
		t.Error(err)
	}
	json, err := json.Marshal(products)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(json))
}
