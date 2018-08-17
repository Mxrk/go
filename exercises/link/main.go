package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {

	// r, err := os.Open("ex1.html")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	s := `<html>
	<body>
	  <h1>Hello!</h1>
	  <a href="/other-page">A link to another page</a>
	</body>
	</html>`

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if !strings.HasPrefix(a.Val, "/") {
						break
					}
					fmt.Printf("Href: %s", a.Val)
					fmt.Println()
					break
				}

			}
		}
		if n.Type == 1 && len(n.Parent.Attr) > 0 {
			check := strings.TrimSpace(n.Data) == ""
			if !check {
				// fmt.Println(n.Type)
				// fmt.Println(html.TextNode)
				// fmt.Println(html.ElementNode)

				fmt.Printf("Text: %s", n.Data)
				fmt.Println()
			}

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
