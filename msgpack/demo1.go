package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type s struct {
	data map[string]interface{}
}

// json序列化
// Go语言中的json包在序列化空接口存放的数字类型（整型、浮点型等）都序列化成float64类型。
func jsonDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1
	ret, err := json.Marshal(s1.data)
	if err != nil {
		fmt.Println("marshal failed ", err)
	}
	fmt.Printf("%#v\n", string(ret))

	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	err = json.Unmarshal(ret, &s2.data)
	if err != nil {
		fmt.Println("marshal failed ", err)
	}
	fmt.Println(s2)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

// gob序列化
func gobDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1

	// encode
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s1.data)
	if err != nil {
		fmt.Println("gob encode failed, err: ", err)
		return
	}
	b := buf.Bytes()
	fmt.Println(b)

	var s2 = s{
		data: make(map[string]interface{}, 8),
	}
	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&s2.data)
	if err != nil {
		fmt.Println("gob encode failed, err: ", err)
		return
	}
	fmt.Println(s2.data)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}

func main() {
	gobDemo()
}
