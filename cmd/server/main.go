package main

import (
	"fmt"
	"net/http"

	"github.com/rykroon/verify/internal/server"
)

func main() {
	jsonRpcServer := server.GetJsonRpcServer()
	http.Handle("/jsonrpc", jsonRpcServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
