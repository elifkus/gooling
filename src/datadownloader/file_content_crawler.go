package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"sync"
	"encoding/json"
	"runtime"
)

var linkChannel = make(chan string)
var retrievedContentChannel = make(chan []byte)
var printBuffer []byte

func RetrieveAndSaveContentToFile(fileLinks []string, filename string) error {
	runtime.GOMAXPROCS(1)
	log.Printf("Number of fileLinks: %d", len(fileLinks))
	
	printBuffer = make([]byte, 0)
	printBuffer = append(printBuffer, []byte("[")...)

	var waitForRetrieversToFinish sync.WaitGroup
	var waitForSliceToBeFilled sync.WaitGroup
    waitForSliceToBeFilled.Add(1)
	
	log.Print("Started retrievers")
	for i:=0; i<20; i++ {
		waitForRetrieversToFinish.Add(1)
		go RetrieveContent(i, &waitForRetrieversToFinish)
	}
	log.Print("Started response collector")
	go addToSlice(&waitForSliceToBeFilled)
	
	log.Print("Sending links to channel")
	var count int
	for i, link := range fileLinks {
		linkChannel <- link
		count = i+1
	}
	
	close(linkChannel)
	log.Printf("Total number of links %d", count)
	waitForRetrieversToFinish.Wait()
	close(retrievedContentChannel)
	log.Print("Waiting for addToSlice")
	
	waitForSliceToBeFilled.Wait()	
	printBuffer = append(printBuffer, []byte("]")...)

	SaveSliceToFile(filename)
	log.Print("Saved file")

	return nil
}

func RetrieveContent(id int, wg *sync.WaitGroup) {
  for link := range linkChannel {
    response, err := RetrieveResponseFromURL(link)

    if err != nil {
	    log.Print(err)
    } 
    
    retrievedContentChannel <- response
  }
  wg.Done()
}

func addToSlice(wg *sync.WaitGroup) {
	for content := range retrievedContentChannel {
		contentString := string(content[:])

		content, err := json.Marshal(contentString)
		if err != nil {
			log.Print(err)
		}	
		printBuffer = append(printBuffer, content...)
		printBuffer = append(printBuffer, []byte(",\n")...)
	}
	
	log.Printf("All slices added")
	wg.Done()
}

func SaveSliceToFile(filename string){
	err := ioutil.WriteFile(filename, printBuffer, 0644)
	if err != nil {
		log.Print(err)
	}
} 
func SaveToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
	    log.Print(err)
	}
	
	defer file.Close()

	for {	
		content := <- retrievedContentChannel
		if _, err = file.Write([]byte("\"")); err != nil {
		    log.Print(err)
		}
		if _, err = file.Write(content); err != nil {
		    log.Print(err)
		}
		if _, err = file.Write([]byte("\",\n")); err != nil {
		    log.Print(err)
		}
	}
}

func RetrieveResponseFromURL(url string) ([]byte, error) {
	
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
