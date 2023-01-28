package main

import "github.com/salticon/sasapay-golang"

var clientId = "8mgx3sf4QhfZpN7aG9DIVdrrMVyTFxU89gz5gaur"
var clientSecret = "EWbIcQEhd3acV8vcAAyuldKpp2EaWNpda4GfQHuANW5biExHDLcGLuxJ6BV1UgHNODfXUUsQqwHBSlc9KINFofXQjQ7DuqI124aICYjsz5MiGn5KajTA8F1YbOQMhHtM"

func main() {
	sasapay := sasapay.NewSasaPay(clientId, clientSecret, int(sasapay.Sandbox))
	sasapay.SetAccessToken()

}
