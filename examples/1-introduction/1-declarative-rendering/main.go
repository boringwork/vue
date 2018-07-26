package main

import "github.com/norunners/vue"

type Data struct {
	Message string
}

func main() {
	vue.New(
		vue.El("#app"),
		vue.Template("<div>{{ Message }}</div>"),
		vue.Data(Data{Message: "Hello wasm!"}),
	)

	select {}
}
