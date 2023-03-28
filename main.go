package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Usage = func() {
	fmt.Println(`
Usage:
  Example: waybackurls -hostFile <file> -outFile
  Example: waybackurls -host <host> -outFile
  Example: waybackurls -host <host>
  
  Use -hostFile to specify a file with a list of domains to check.
  Use -outFile to save the results to a file by domain name. If not used, the results will be printed to the terminal.
	`)
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	outFile := flag.Bool("outFile", false, "This flag will save the results to a file by domain name. If not used, the results will be printed to the terminal.")
	hostFile := flag.String("hostFile", "", "This flag will specify a file with a list of domains to check.")
	host := flag.String("host", "", "This flag will specify a single domain to check.")
	flag.Parse()

	if *hostFile == "" && *host == "" {
		Usage()
		os.Exit(0)
	}

	if *host != "" {
		url := "https://web.archive.org/cdx/search/cdx?url=" + *host + "/*&output=text&fl=original&collapse=urlkey"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if string(body) != "" {
			if *outFile {
				os.WriteFile(*host+".txt", body, 0644)
				log.Printf("Done: " + *host)
			} else {
				fmt.Println(string(body))
			}
		} else {
			log.Printf("No results for: " + *host)
		}
	} else {
		if _, err := os.Stat(*hostFile); os.IsNotExist(err) {
			log.Fatalln("File does not exist")
		} else {
			readFile, err := os.Open(*hostFile)
			if err != nil {
				log.Fatalln(err)
			}
			defer readFile.Close()

			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)

			for fileScanner.Scan() {
				url := "https://web.archive.org/cdx/search/cdx?url=" + fileScanner.Text() + "/*&output=text&fl=original&collapse=urlkey"
				resp, err := http.Get(url)
				if err != nil {
					log.Fatalln(err)
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				defer resp.Body.Close()

				if string(body) != "" {
					if *outFile {
						os.WriteFile(fileScanner.Text()+".txt", body, 0644)
						log.Printf("Done: " + fileScanner.Text())
					} else {
						fmt.Println(string(body))
					}
				} else {
					log.Printf("No results for: " + fileScanner.Text())
				}
			}
		}
	}
}
