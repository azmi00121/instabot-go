package instabot

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

const (
	root     = "https://www.instagram.com/"
	login    = "https://www.instagram.com/accounts/login/ajax/"
	logout   = "https://www.instagram.com/accounts/logout/"
	tag      = "https://www.instagram.com/explore/tags/"
	photo    = "https://www.instagram.com/p/"
	like     = "https://www.instagram.com/web/likes/%s/like/"
	unlike   = "https://www.instagram.com/web/likes/%s/unlike/"
	comment  = "https://www.instagram.com/web/comments/%s/add/"
	follow   = "https://www.instagram.com/web/friendships/%s/follow/"
	unfollow = "https://www.instagram.com/web/friendships/%s/unfollow/"
)

// Config struct for reading the config file
type Config struct {
	Username     string   `yaml:"USERNAME"`
	Password     string   `yaml:"PASSWORD"`
	Tags         []string `yaml:"TAGS"`
	TotalLikes   int64    `yaml:"TOTAL_LIKES"`
	LikesPerUser int64    `yaml:"LIKES_PER_USER"`
}

// GetConfig method to read the config file
func GetConfig() (Config, error) {
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal("Couldn't find the config file")
	}
	config := Config{}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config, nil
}

func main() {
	fmt.Println("Instabot v0.0.1")
	fmt.Printf("%s", root)
	_, config := GetConfig()
	fmt.Printf("\n%v\n", config)
}
