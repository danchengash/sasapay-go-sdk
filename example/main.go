package main

import (
	"fmt"

	"github.com/salticon/sasapay-golang"
)

var clientId = "8mgx3sf4QhfZpN7aG9DIVdrrMVyTFxU89gz5gaur"
var clientSecret = "EWbIcQEhd3acV8vcAAyuldKpp2EaWNpda4GfQHuANW5biExHDLcGLuxJ6BV1UgHNODfXUUsQqwHBSlc9KINFofXQjQ7DuqI124aICYjsz5MiGn5KajTA8F1YbOQMhHtM"

func main() {
	sasapay := sasapay.NewSasaPay(clientId, clientSecret, int(sasapay.Sandbox))
	_, err := sasapay.RegisterCallBackUrl("600980", "https://posthere.io/67df-4d9c-9386")
	if err != nil {
		fmt.Println(err.Error())
	}
}
