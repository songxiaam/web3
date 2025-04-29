package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	testRLP()
}

func testRLPSource() {
}

func testRLP() {
	type Person struct {
		Name string
		Age  uint
	}

	p := Person{"Alice", 30}

	var buf bytes.Buffer
	err := rlp.Encode(&buf, p)
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.Bytes())

	//var buf1 bytes.Buffer
	encoded1, err := rlp.EncodeToBytes("Alice")
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded1)

	//var buf2 bytes.Buffer
	encoded2, err := rlp.EncodeToBytes(uint(30))
	if err != nil {
		panic(err)
	}
	fmt.Println(encoded2)

	//data := append(buf1.Bytes(), buf2.Bytes()...)
	//fmt.Println(data)

	//var bufData bytes.Buffer
	//err = rlp.Encode(&bufData, data)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(bufData.Bytes())

	var decoded Person
	err = rlp.Decode(bytes.NewReader(buf.Bytes()), &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Println(decoded)
}
