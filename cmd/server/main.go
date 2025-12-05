package main

import (
	"fmt"
	"net/http"

	"github.com/rykroon/jsonrpc"
	"github.com/rykroon/verify/internal/server"
)

func main() {
	jsonRpcServer := server.GetJsonRpcServer()
	http.Handle("/jsonrpc", jsonrpc.NewHttpHandler(jsonRpcServer))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
