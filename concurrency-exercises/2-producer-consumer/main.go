package main

import (
	"flag"
	"fmt"
	"time"
)

func producer(stream Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return tweets
		}

		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutRhodes() {
			fmt.Println(t.Username, "\ttweets about Rhodes")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about Rhodes")
		}
	}
}

func main() {
	flag.Parse()

	start := time.Now()
	stream := GetMockStream()

	// TODO: Allow producer and consumer to run in parallel.

	// Producer
	tweets := producer(stream)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
