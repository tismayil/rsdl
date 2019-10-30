package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

func main() {

	hostname := flag.String("hostname", "", "Please input hostname")
	list := flag.String("list", "default.txt", "Subdomain List")
	output := flag.Bool("output", false, "Output Result")

	flag.Parse()

	file, err := os.Open(*list)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count = 1
	var outputText = ""
	scanner := bufio.NewScanner(file)

	fmt.Println(`
	██████╗    ███████╗   ██████╗    ██╗     
	██╔══██╗   ██╔════╝   ██╔══██╗   ██║     
	██████╔╝   ███████╗   ██║  ██║   ██║     
	██╔══██╗   ╚════██║   ██║  ██║   ██║     
	██║  ██║██╗███████║██╗██████╔╝██╗███████╗
	╚═╝  ╚═╝╚═╝╚══════╝╚═╝╚═════╝ ╚═╝╚══════╝
                                         
						
	List : ` + *list + `
	Victim Host : ` + *hostname + `
	`)

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	for scanner.Scan() {
		domain := scanner.Text() + "." + *hostname
		rev := getPing(domain)

		if rev == true {
			s.Restart()
			fmt.Printf("\033[1;36m%s\033[0m", strconv.FormatInt(int64(count), 10)+") "+domain+"\n")
			outputText += domain + "\n"
			count = count + 1
		}
	}

	if *output != false {
		d1 := []byte(outputText)
		err := ioutil.WriteFile(*hostname, d1, 0644)
		check(err)
	}
	s.Restart()
	s.Stop()
}
