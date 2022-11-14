package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var hashMap = map[string]func(string) string{
	"md5": hashMd5,
}

func main() {
	args := os.Args[1:]
	argStr := strings.Join(args, " ")
	valueMap := map[string]string{}
	for k, f := range hashMap {
		v := f(argStr)
		valueMap[k] = v
	}
	result := renderResult(valueMap)
	marshal, _ := json.Marshal(result)
	fmt.Print(string(marshal))
}
func renderResult(values map[string]string) map[string]interface{} {
	result := map[string]interface{}{}
	var items []interface{}
	for k, v := range values {
		vm := map[string]interface{}{}
		vm["uid"] = "hash-" + k
		vm["title"] = k
		vm["subtitle"] = v
		vm["arg"] = v
		items = append(items, vm)
	}
	result["items"] = items
	return result
}
func hashMd5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
