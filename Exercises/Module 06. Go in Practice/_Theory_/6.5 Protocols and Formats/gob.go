package main

import (
"bytes"
"encoding/gob"
"fmt"
"log"
)


func GobExample() {
	DecodeGob(EncodeGob())
}

func EncodeGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf) //в отличие от остальных форматов, имплементация gob отличается, мы пишем в буфер, в отличие от возврата строки

	m := make(map[string]string)
	m["foo"] = "bar"

	if err := enc.Encode(m); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func DecodeGob(input []byte) {
	buf := bytes.NewBuffer(input)
	dec := gob.NewDecoder(buf)

	m := make(map[string]string)

	if err := dec.Decode(&m); err != nil { //читаем и декодируем
		log.Fatal(err)
	}

	fmt.Println(m["foo"])
}
