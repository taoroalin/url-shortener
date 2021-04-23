package main

import (
	"github.com/valyala/fasthttp"
)

func handleRequest(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello world")
}

func main() {
	println("hello world")
	fasthttp.ListenAndServe("localhost:8091", handleRequest)
}
