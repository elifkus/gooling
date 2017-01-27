package main 

import (
	"github.com/google/go-github/github"
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GistFile struct {
	Content		*string					`json:"content,omitempty"`
	RawURL		*string					`json:"raw_url,omitempty"`
	FileType	*string					`json:"file_type,omitempty"`
	Size		*int					`json:"size,omitempty"`
}

type GistFileCollection struct {
    Collection []GistFile
}


func RetrievePublicGists (client *github.Client) ([]*github.Gist, error) {
	
	options := &github.GistListOptions{
	    ListOptions: github.ListOptions{PerPage: 100},
	}
	
	var allGists []*github.Gist
	var err error
	
	for {
		gists, response, err := client.Gists.ListAll(options)
	
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

func ExtractFileFromGists(gists []*github.Gist) ([]GistFile){
	
	var allFileLinks []GistFile
	
	for _, gist := range gists {
		if gist.Files != nil {
			for _, file := range gist.Files {
				gistFile := GistFile{
					Content: file.Content,
					RawURL: file.RawURL,
					FileType: file.Type,
					Size: file.Size }
				
				allFileLinks = append(allFileLinks, gistFile)
			} 
		}
		
	}
	
	return allFileLinks
}

func WriteFileLinksToFileAsJSON(fileLinks []GistFile, filename string) error {
	fileLinksString, err := json.Marshal(fileLinks)
	
	if err != nil {
		log.Printf("%s", err)
	}
	
	err = ioutil.WriteFile(filename, fileLinksString, 0644)
	
	if err != nil {
		log.Printf("%s", err)
	}

	return err
}

func RetrieveLinksInFileLinkData(filename string) ([]string, error) {
	fileContent, err := ioutil.ReadFile(filename)
	
	if err != nil {
		log.Print(err)
	}
	
	gistFiles := make([]GistFile,0)
    json.Unmarshal(fileContent, &gistFiles)
    
    var allLinks []string
    
    for _ , gistFile := range gistFiles {
	    allLinks = append(allLinks, *gistFile.RawURL)
    }
    
	return allLinks, err
}

func RetrieveAndWriteFileContentToFileAsJSON(fileLinks []string, filename string) error {	
	
	var fileContents []byte 
	
	fileContents = append(fileContents, []byte("[")...)
	for _, link := range fileLinks {
		fileContents = append(fileContents, []byte("\"")...)
		content, _ := RetrieveResponseAsStringFromURL(link)
		fileContents = append(fileContents, content...)
		
		fileContents = append(fileContents, []byte("\",\n")...)
	}
	
	fileContents = append(fileContents, []byte("\"]")...)
	
	err := ioutil.WriteFile(filename, fileContents, 0644)
	
	if err != nil {
		log.Printf("%s", err)
	}

	return err

} 

func RetrieveResponseAsStringFromURL(url string) ([]byte, error) {
	
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
	    log.Print(err)
	}
	defer resp.Body.Close()
	
	var bodyBytes []byte
	
	if resp.StatusCode == 200 { // OK
	    bodyBytes, err = ioutil.ReadAll(resp.Body)
	    
	    if err != nil {
		    log.Print(err)
		}
		    
	}
	
	return bodyBytes, err
}

func RetrieveGistDataFromGithubAndSaveToFile(filename string) {
	client := github.NewClient(nil)

	gists, _ := RetrievePublicGists(client)
	
	fileLinks := ExtractFileFromGists(gists)
	
	WriteFileLinksToFileAsJSON(fileLinks, filename)  
}

func RetreiveGistContentFromGithubAndSaveToFile(gistFilename string, contentFilename string) {
	allLinks, _ := RetrieveLinksInFileLinkData(gistFilename)
	RetrieveAndWriteFileContentToFileAsJSON(allLinks, contentFilename)
}

func main() {
	filename := "filelinkdata.json"
	//RetrieveGistDataFromGithubAndSaveToFile(filename)

	contentFilename := "postdata.json"	
	RetreiveGistContentFromGithubAndSaveToFile(filename, contentFilename)
}




