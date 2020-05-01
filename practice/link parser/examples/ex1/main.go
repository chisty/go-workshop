package main

import (
	"fmt"
	"strings"

	link "github.com/chisty/linkparser"
)

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

var exampleHTML = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/first-page">A link to first page</a>
  <a href="/other-page">
	  A link to another page
	  <span> With a Span </span>
	  Thanks
  </a>
</body>
</html>
`
