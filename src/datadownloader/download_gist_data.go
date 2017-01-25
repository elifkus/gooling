package main 

import (
	"github.com/google/go-github/github"
	"fmt"
	"log"
)

func retrievePublicGists (client *github.Client) ([]*github.Gist, error) {
	
	options := &github.GistListOptions{
	    ListOptions: github.ListOptions{PerPage: 100},
	}
	
	var allGists []*github.Gist
	var err error
	
	for {
		gists, response, err := client.Gists.List("", options)
	
		if (err != nil) {
			if _, ok := err.(*github.RateLimitError); ok {
			    log.Println("hit rate limit")
			}
			if _, ok := err.(*github.AcceptedError); ok {
			    log.Println("scheduled on GitHub side")
			}
			
			log.Print(err)
		} else {
			allGists = append(allGists, gists...)
		}
		
		if response.NextPage == 0 {
	        break
	    }
		
	    options.ListOptions.Page = response.NextPage
	
		fmt.Printf("%s\n", response)

	}
			
	return allGists, err
}

func main() {
	client := github.NewClient(nil)

	gists, err := retrievePublicGists(client)
	
	for _, gist := range gists {
		fmt.Println(gist)
	} 

	if err != nil {
		fmt.Printf("%s", err)
	}
}


