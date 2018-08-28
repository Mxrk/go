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

	s := `<html>
	<head>
	  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
	</head>
	<body>
	  <h1>Social stuffs</h1>
	  <div>
		<a href="https://www.twitter.com/joncalhoun">
		  Check me out on twitter
		  <i class="fa fa-twitter" aria-hidden="true"></i>
		</a>
		<a href="https://github.com/gophercises">
		  Gophercises is on <strong>Github</strong>!
		</a>
	  </div>
	</body>
	</html>`

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			//	fmt.Println(n.Attr)
			for _, a := range n.Attr {
				if a.Key == "href" {
					//fmt.Println(a.Val)
					//	fmt.Printf("Href: %s", a.Val)
					if strings.HasPrefix(a.Val, "/") || strings.HasPrefix(a.Val, "https://") {
						fmt.Printf("Href: %s", a.Val)
						fmt.Println()
					} else {
						break
					}
				}

			}
		}
		t := strings.TrimSpace(n.Data)

		if n.Type == html.TextNode && len(t) > 0 {
			fmt.Println("Text: ", t)
			fmt.Printf("Type: %v, ElementNode: %v,", n.Type, html.ElementNode)
			fmt.Println()
		}
		if n.Type != html.ElementNode {

		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
