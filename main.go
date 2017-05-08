package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	repo   = "moby/moby"
	tocken = "1660941cfa92bb5e73bbdef4c4ef2b1fd8841be6"
)

//User is
type User struct {
	Login string `json:"login,omitempty"`
}

func main() {
	users := make(chan string, 50000)
	//userinfo := make(chan github.User, 1024)

	ctx := context.Background()
	ctx, c := context.WithCancel(ctx)
	count := 0

	go getUserinfo(ctx, users)
	for i := 1; ; i++ {
		us := new([]User)
		url := fmt.Sprintf("https://api.github.com/repos/%s/stargazers?page=%d&per_page=100&?access_token=%s", repo, i, tocken)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error", err)
		}

		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(us)
		if err != nil {
			fmt.Println("json decode error: ", err)
			time.Sleep(time.Second * 10)
			fmt.Println("limit sleep 10...")
			continue
		}
		if len(*us) == 0 {
			fmt.Println("complete total: ", count)
			//cancel other go, wait for other tasks done
			time.Sleep(time.Hour * 70)
			c()
			break
		}
		for _, u := range *us {
			fmt.Println("got user: ", u.Login)
			count++
			users <- u.Login
		}
	}

}

func tostring(s *string) string {
	if s != nil {
		return *s
	}
	return "null"
}

func getUserinfo(gctx context.Context, users chan string) {
	for {
		select {
		case user := <-users:
			fmt.Println("asdfasdfasdf:", user)
			ctx := context.Background()
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: tocken},
			)
			tc := oauth2.NewClient(ctx, ts)

			client := github.NewClient(tc)

			userinfo, _, err := client.Users.Get(ctx, user)
			if _, ok := err.(*github.RateLimitError); ok {
				fmt.Println("hit rate limit")
				/*
					go func(user) {
						users <- user
					}(user)
				*/
				time.Sleep(time.Second * 60)
				continue
			}

			dump := fmt.Sprintf("%s|%s|%s|%s|%s\n", tostring(userinfo.Login), tostring(userinfo.Name), tostring(userinfo.Company), tostring(userinfo.Location), tostring(userinfo.Email))
			fmt.Println("got info:", dump)
			if tostring(userinfo.Email) != "null" {
				err = writeFile(fmt.Sprintf("%s-gitdata.dump", strings.Replace(repo, "/", "-", -1)), []byte(dump), 0644)
			}

		case <-gctx.Done():
			fmt.Println("cancel get user info")
			//wait for task complete
			return
		}
	}
}

func writeFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
