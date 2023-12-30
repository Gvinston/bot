package searcher

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SearcherService struct {
}

func (SS *SearcherService) Search(find string) Anime {

	url := "https://shikimori.one/animes?search=" + find

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	htmlAnswer := string(body)
	doc, err := html.Parse(strings.NewReader(htmlAnswer))

	if err != nil {
		log.Fatal(err)
	}
	var anime Anime

	tag := getElementByClass(doc, "title left_aligned")

	if tag != nil {
		anime = getAnimeFromNode(tag)
	}

	if anime.Link == "" {
		tag = getElementByClass(doc, "cover anime-tooltip")

		if tag != nil {
			anime = getAnimeFromNode(tag)
		}
	}

	return anime
}

func getAttribute(n *html.Node, key string) (string, bool) {

	for _, attr := range n.Attr {

		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

func renderNode(n *html.Node) string {

	var buf bytes.Buffer
	w := io.Writer(&buf)

	err := html.Render(w, n)

	if err != nil {
		return ""
	}

	return buf.String()
}

func getAnimeFromNode(n *html.Node) Anime {
	link, _ := getAttribute(n, "href")
	title, _ := getAttribute(n, "title")

	return Anime{
		Title: title,
		Link:  link,
	}
}

func checkClass(n *html.Node, class string) bool {

	if n.Type == html.ElementNode {

		s, ok := getAttribute(n, "class")

		if ok && s == class {
			return true
		}
	}

	return false
}

func traverse(n *html.Node, id string) *html.Node {

	if checkClass(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		res := traverse(c, id)

		if res != nil {
			return res
		}
	}

	return nil
}

func getElementByClass(n *html.Node, class string) *html.Node {

	return traverse(n, class)
}
