package main

import (
	"fmt"
	"unsafe"

	"github.com/georgebent/go-httpclient/gohttp"
)

func main() {
	client := gohttp.Client{}
	client.Get()

	fmt.Println(unsafe.Sizeof(client))
}
