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

	s := `<html lang="en" class="no-ie">
	<!--<![endif]-->
	
	<head>
		<title>Gophercises - Coding exercises for budding gophers</title>
	</head>
	
	<body>
		<section class="header-section">
			<div class="jumbo-content">
				<div class="pull-right login-section">
					Already have an account?
					<a href="#" class="btn btn-login">Login <i class="fa fa-sign-in" aria-hidden="true"></i></a>
				</div>
				<center>
					<img src="https://gophercises.com/img/gophercises_logo.png" style="max-width: 85%; z-index: 3;">
					<h1>coding exercises for budding gophers</h1>
					<br/>
					<form action="/do-stuff" method="post">
						<div class="input-group">
							<input type="email" id="drip-email" name="fields[email]" class="btn-input" placeholder="Email Address" required>
							<button class="btn btn-success btn-lg" type="submit">Sign me up!</button>
							<a href="/lost">Lost? Need help?</a>
						</div>
					</form>
					<p class="disclaimer disclaimer-box">Gophercises is 100% FREE, but is currently in beta. There will be bugs, and things will be changing significantly over the coming weeks.</p>
				</center>
			</div>
		</section>
		<section class="footer-section">
			<div class="row">
				<div class="col-md-6 col-md-offset-1 vcenter">
					<div class="quote">
						"Success is no accident. It is hard work, perseverance, learning, studying, sacrifice and most of all, love of what you are doing or learning to do." - Pele
					</div>
				</div>
				<div class="col-md-4 col-md-offset-0 vcenter">
					<center>
						<img src="https://gophercises.com/img/gophercises_lifting.gif" style="width: 80%">
						<br/>
						<br/>
					</center>
				</div>
			</div>
			<div class="row">
				<div class="col-md-10 col-md-offset-1">
					<center>
						<p class="disclaimer">
							Artwork created by Marcus Olsson (<a href="https://twitter.com/marcusolsson">@marcusolsson</a>), animated by Jon Calhoun (that's me!), and inspired by the original Go Gopher created by Renee French.
						</p>
					</center>
				</div>
			</div>
		</section>
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
