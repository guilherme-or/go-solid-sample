package main

import (
	"github.com/guilherme-or/go-solid-sample/interfaces"
	"github.com/guilherme-or/go-solid-sample/interfaces/api"
)

func main() {
	var apiInterface interfaces.UserInterface = api.NewApi()
	apiInterface.Run()
}
