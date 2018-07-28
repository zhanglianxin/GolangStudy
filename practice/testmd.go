package main

import (
	"github.com/golang-commonmark/markdown"
	"io/ioutil"
	"log"
	"github.com/go-redis/redis"
	"fmt"
)

func main() {
	redisClient := NewRedisClient()
	redisClient.Del("cgt")

	md := markdown.New()
	src, err := ioutil.ReadFile("/home/zhanglianxin/cgt.md")
	if err != nil {
		log.Fatalln(err)
	}
	tokens := md.Parse(src)
	levelTwoTitle, content := "", ""
	for _, token := range tokens {
		switch token := token.(type) {
		case *markdown.HeadingOpen:
			// log.Println("H" + strconv.Itoa(token.HLevel))
		case *markdown.HeadingClose:
			// log.Println("H" + strconv.Itoa(token.HLevel))
		case *markdown.ParagraphOpen:
			// log.Println("P" + strconv.Itoa(token.Lvl))
		case *markdown.ParagraphClose:
			// log.Println("P" + strconv.Itoa(token.Lvl))
		case *markdown.Inline:
			// log.Println(token.Content)
			if 6 == len(token.Content) {
				levelTwoTitle = token.Content
			} else {
				if levelTwoTitle != "" {
					content = fmt.Sprintf("【%s】%s", levelTwoTitle, token.Content)
					redisClient.RPush("cgt", content)
				}
			}
		}
	}
}

func NewRedisClient() (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return
}
