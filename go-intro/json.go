// +build ignore,OMIT
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// START structs OMIT
type ChildData struct {
	Domain string  `json:"domain"`
	Author string  `json:"author"`
	Score  int     `json:"score"`
	URL    url.URL `json:"url"`
}

type Child struct {
	Kind string    `json:"kind"`
	Data ChildData `json:"data"`
}

type PageData struct {
	Children []Child `json:"children"`
}

type RedditPage struct {
	Kind string   `json:"kind"`
	Data PageData `json:"data"`
}
// END structs OMIT

func main() {
    // START OMIT
	var page RedditPage
	resp, err := http.Get("http://www.reddit.com/r/golang.json")
    defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	if err = json.NewDecoder(resp.Body).Decode(&page); err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", page)
    // END OMIT
}

