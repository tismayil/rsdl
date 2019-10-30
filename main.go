package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sparrc/go-ping"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
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

func main() {

	hostname := flag.String("hostname", "", "Please input hostname")
	list := flag.String("list", "default.txt", "Subdomain List")

	flag.Parse()

	file, err := os.Open(*list)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count = 1
	scanner := bufio.NewScanner(file)

	fmt.Println(`
	██████╗    ███████╗   ██████╗ 
	██╔══██╗   ██╔════╝   ██╔══██╗
	██████╔╝   ███████╗   ██║  ██║
	██╔══██╗   ╚════██║   ██║  ██║
	██║  ██║██╗███████║██╗██████╔╝
	╚═╝  ╚═╝╚═╝╚══════╝╚═╝╚═════╝ 
						
			List : ` + *list + `
			Victim Host : ` + *hostname + `
	`)

	for scanner.Scan() {
		domain := scanner.Text() + "." + *hostname
		rev := getPing(domain)

		if rev == true {
			fmt.Printf(NoticeColor, strconv.FormatInt(int64(count), 10)+") "+domain+"\n")
			count = count + 1
		}

	}

}
