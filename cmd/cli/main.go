package main

import (
	"fmt"

	openapi "github.com/dhpollack/go-oas-faker/openapi/v1"
)

func main() {
	fmt.Println("Hello World")
	openapi.RenderExample()
}
