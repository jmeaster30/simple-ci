package main

import "fmt"

func main() {
	println("Hello :3")
	config, err := LoadConfig("simple-ci-config.yaml")
	if err != nil {
		println("error :(")
		println(err.Error())
	}
	fmt.Printf("%v\n", config)
}
