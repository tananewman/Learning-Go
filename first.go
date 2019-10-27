package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("hello world!")
	viper.AddConfigPath("hi")
}
