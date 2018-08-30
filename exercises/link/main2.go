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

		var ret Link

		ret.Href = (link(n))

		ret.Text = (text(n))
		if ret.Href != "" && n.Data == "a" {
			fmt.Printf("Link: %q\n", ret.Href)
			fmt.Printf("Text: %q\n", ret.Text)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}

	}
	f(doc)

}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func link(n *html.Node) string {
	var ret string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret = attr.Val
		}
	}
	return ret
}
