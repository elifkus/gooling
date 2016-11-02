package main     

import ( 
	 "fmt"
	 "flag"
	 "os"
	 "net/http"      
     "github.com/jackdanger/collectlinks"  
     "crypto/tls"  
)

func main() {
	flag.Parse()
	args := flag.Args()
	
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	} 
	
	retrieve(args[0])
}  

func retrieve(uri string) {
	
	tlsConfig := &tls.Config{  
		InsecureSkipVerify: true,  
	}                            
    
    transport := &http.Transport{    
	    TLSClientConfig: tlsConfig,    
	  }                                

    client := http.Client{Transport: transport} 
	
	response, err := client.Get(uri)
	
	if err != nil {
		return
	}
	
	defer response.Body.Close()
	
	links := collectlinks.All(response.Body)
	
	for _, link := range(links) {
		fmt.Println(link)
	}
}