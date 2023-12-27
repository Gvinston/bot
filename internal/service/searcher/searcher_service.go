package searcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type SearcherService struct {
}

func (SS *SearcherService) Search(find string) Anime {

	url := "https://shikimori.one/animes?search=naruto"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return Anime{}
}
