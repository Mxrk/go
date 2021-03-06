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
	<body>
		<a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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
			// fix if only
			//	if strings.HasPrefix(ret.Href, "http") {
			fmt.Printf("Link: %q\n", ret.Href)
			fmt.Printf("Text: %q\n", ret.Text)
			//		}

			// fmt.Println("Test1: ", n.Type)
			// fmt.Println("Test2: ", html.ElementNode)
			// fmt.Println("Test3: ", n.Data)
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
