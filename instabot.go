package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	yaml "gopkg.in/yaml.v2"
)

const (
	root           = "https://www.instagram.com/"
	login          = "https://www.instagram.com/accounts/login/ajax/"
	logout         = "https://www.instagram.com/accounts/logout/"
	tag            = "https://www.instagram.com/explore/tags/"
	photo          = "https://www.instagram.com/p/"
	like           = "https://www.instagram.com/web/likes/%s/like/"
	unlike         = "https://www.instagram.com/web/likes/%s/unlike/"
	comment        = "https://www.instagram.com/web/comments/%s/add/"
	follow         = "https://www.instagram.com/web/friendships/%s/follow/"
	unfollow       = "https://www.instagram.com/web/friendships/%s/unfollow/"
	acceptLanguage = "en-US,en;q=0.8,pt-br,pt;q=0.6,"
	userAgent      = "('Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 ''(KHTML, like Gecko) Chrome/48.0.2564.109 Safari/537.36')"
)

type Config struct {
	Username     string   `yaml:"USERNAME"`
	Password     string   `yaml:"PASSWORD"`
	Tags         []string `yaml:"TAGS"`
	TotalLikes   int64    `yaml:"TOTAL_LIKES"`
	LikesPerUser int64    `yaml:"LIKES_PER_USER"`
}

func GetConfig() (error, Config) {
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal("Couldn't find the config file")
	}
	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return nil, config
}

func main() {
	fmt.Println("Instabot v0.0.1")
	fmt.Printf("%s", root)
	_, config := GetConfig()
	fmt.Printf("\n%v\n", config)
	jar, err := cookiejar.New(nil)
	client := http.Client{Jar: jar}
	httresp, err := client.Get(root)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf(httresp.Status)
	resp, err := client.PostForm(root, url.Values{
		"password": {config.Password},
		"username": {config.Username},
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	fmt.Println(string(data))
}
