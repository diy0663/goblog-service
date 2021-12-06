package test

import (
	"encoding/base64"
	"log"
	"testing"
)

// 测试反向解析token
func TestTokenParse(t *testing.T) {
	payload, _ := base64.StdEncoding.DecodeString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiMjc1NjY4YmE2NTUwNDljZDczOWQxZDllNmIzMWNjZjEiLCJhcHBfc2VjcmV0IjoiN2M5NzI2NjMxNzBkNmJjMTg0ODRkMDViYzk4NzIyZjQiLCJleHAiOjE2Mzg4MDg2OTksImlzcyI6ImJsb2ctc2VydmljZSJ9.tHkJK1PTJEswWIkHA-oLQe811UfEZ13PJYzOUbn7huQ")
	log.Println(string(payload))
	// 输出: {"alg":"HS256","typ":"JWT"}
}

// func TestParseToken(t *testing.T) {
// 	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiMjc1NjY4YmE2NTUwNDljZDczOWQxZDllNmIzMWNjZjEiLCJhcHBfc2VjcmV0IjoiN2M5NzI2NjMxNzBkNmJjMTg0ODRkMDViYzk4NzIyZjQiLCJleHAiOjE2Mzg4MDg2OTksImlzcyI6ImJsb2ctc2VydmljZSJ9.tHkJK1PTJEswWIkHA-oLQe811UfEZ13PJYzOUbn7huQ"
// 	_, _ = app.ParseToken(token)

// 	//fmt.Println(*data)

// }
