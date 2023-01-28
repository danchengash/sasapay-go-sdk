package main

import "github.com/salticon/sasapay-golang"

func main() {
	sasapay := sasapay.NewSasaPay(" ", " ", int(sasapay.Sandbox))
	sasapay.SetAccessToken()

}
