package main

import "github.com/devsidnei/go-expert/9-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
