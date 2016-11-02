package main     

import ( 
	 "fmt"
	 "flag"
	 "os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	
	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	} 
}  