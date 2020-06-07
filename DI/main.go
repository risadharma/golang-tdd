package main

import (
	"fmt"
	"io"
	"net/http"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// DI located at interface (io.Writer) as an parameter, so we can injecting anything as long the type already implement all interface's method
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(GreetHandler))
}
