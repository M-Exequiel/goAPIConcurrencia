package main

import (
	"net/http"
	"testing"
)

func BenchmarkGetVariasVeces(b *testing.B) {
	for n:=0;n<b.N ;n++  {
		http.Get("localhost:8080/users/12345")
	}
}

