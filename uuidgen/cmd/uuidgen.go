package main

import "github.com/aliaksandrrachko/cli-utils/uuidgen/pkg/uuidgenerator"

func main() {
	var generatedUuidStr string = uuidgenerator.NewUuidGenerator().Generate()

	println(generatedUuidStr)
}
