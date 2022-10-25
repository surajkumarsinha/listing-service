package main

import (
	"fmt"
	"io"
	"listing-service/configs"
	"listing-service/infrastructures"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	http.HandleFunc("/hello", getHello)
	infrastructures.KernelBuilder().Build(configs.App())
	http.ListenAndServe(getAddress(), nil)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
}
