package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sparrc/go-ping"
)

func getPing(host string) bool {

	pinger, err := ping.NewPinger(host)
	if err != nil {
		return false
	}

	pinger.OnRecv = func(pkt *ping.Packet) {

	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func httpCheck(domain string) string {

	client := http.Client{
		Timeout: 4 * time.Second,
	}

	check, err := client.Get("http://" + domain + "/")
	if err != nil {
		return "N/A"
	}
	return strconv.Itoa(check.StatusCode)
}

func httpsCheck(domain string) string {

	client := http.Client{
		Timeout: 4 * time.Second,
	}

	check, err := client.Get("https://" + domain + "/")
	if err != nil {
		return "N/A"
	}

	return strconv.Itoa(check.StatusCode)
}

func main() {

	hostname := flag.String("hostname", "", "Please input hostname")
	list := flag.String("list", "lists/mini.txt", "Subdomain List")
	output := flag.Bool("output", false, "Output Result")

	flag.Parse()

	file, err := os.Open(*list)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var outputText = ""
	scanner := bufio.NewScanner(file)

	fmt.Println(`
	██████╗    ███████╗   ██████╗    ██╗     
	██╔══██╗   ██╔════╝   ██╔══██╗   ██║     
	██████╔╝   ███████╗   ██║  ██║   ██║     
	██╔══██╗   ╚════██║   ██║  ██║   ██║     
	██║  ██║██╗███████║██╗██████╔╝██╗███████╗
	╚═╝  ╚═╝╚═╝╚══════╝╚═╝╚═════╝ ╚═╝╚══════╝
                                         
	R.S.D.L V2				
	List : ` + *list + `
	Victim Host : ` + *hostname + `
	`)

	s := spinner.New(spinner.CharSets[31], 100*time.Millisecond)
	s.Start()
	s.Color("red", "bold")

	for scanner.Scan() {
		domain := scanner.Text() + "." + *hostname

		rev := getPing(domain)
		s.Restart()
		s.Prefix = "Creating Requests " + domain

		if rev == true {
			s.Restart()
			ip, _ := net.LookupIP(domain)
			fmt.Printf("\n\033[1;36m%s\033[0m", "Up Domain "+domain+"\nHTTP Response Code: "+httpCheck(domain)+" \nHTTPS Response Code: "+httpsCheck(domain)+" \n")
			fmt.Printf("\033[1;36m%s\033[0m", "Founded IP Adresses\n")
			fmt.Println(ip)
			fmt.Printf("\033[1;36m%s\033[0m", "==========================================================\n")
			outputText += domain + "\n"

		}
	}

	if *output != false {
		d1 := []byte(outputText)
		err := ioutil.WriteFile(*hostname, d1, 0644)
		check(err)
	}
	s.Restart()
	s.Stop()

	log.Fatal("\n\n\n\nScan Finished\n Happy Hunting....")
}
