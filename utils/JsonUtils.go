package utils

import (
	"strings"
	"encoding/json"
	"github.com/jmoiron/jsonq"
)

func GetStringValue(jsonStr string, key string) (string) {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	str, err := jq.String(key)
	if err != nil {
		panic(err)
	}
	return str
}

func GetQuery(jsonStr string) (*jsonq.JsonQuery) {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	return jsonq.NewQuery(data)
}


func GetDataMap(jsonStr string) (map[string]interface{}) {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonStr))
	dec.Decode(&data)
	return data
}