package main

import (
	. "./keys"
	. "fmt"
	"net/url"
)

fucn main() {
	api := GetTwitterApi()

	v := url.Values{}
	v.Set("count","10")

	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		panic(err)
	}

	for _, tweet := range tweets {
		Println("tweets: ", tweet.text)
	}
}