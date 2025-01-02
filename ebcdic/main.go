package main

import (
	"fmt"
	"os"
)

func main() {
	//data, err := os.ReadFile("/home/jk/projects/go/src/learning_tools/dd/EBCDIC.txt")
	data, err := os.ReadFile("D:\\DATA\\projects\\go\\learning_tools\\ebcdic\\EBCDIC.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(EncodeToString(data))
	fmt.Println(string(data))

	str := Decode("hello,world!")
	fmt.Println(string(str))
	fmt.Println(EncodeToString(str))
}
