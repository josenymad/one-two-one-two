package main

import "github.com/josenymad/one-two-one-two/server"

func main() {
	http := server.NewHttpServer("localhost", "8000")
	http.Run()
}
