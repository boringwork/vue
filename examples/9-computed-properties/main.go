package main

import (
	"github.com/norunners/vue"
)

const tmpl = `
<div>
  <p>Original message: "{{ Message }}"</p>
  <p>Computed reversed message: "{{ ReversedMessage }}"</p>
</div>
`

type Data struct {
	Message string
}

func ReversedMessage(context vue.Context) string {
	message := context.Get("Message").(string)
	runes := []rune(message)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	vue.New(
		vue.El("#app"),
		vue.Template(tmpl),
		vue.Data(Data{Message: "Hello WebAssembly!"}),
		vue.Computeds(ReversedMessage),
	)

	select {}
}