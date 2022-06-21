package main

import "github.com/luccasalves/newsletter-go/server"

func main() {
	s := server.CreateServer()
	s.Init()
}
