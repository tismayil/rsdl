# rsdl Subdomain Scanner v.2
Subdomain Scan With Ping Method & Check HTTP/S Statuses.


[![asciicast](https://asciinema.org/a/w3PxQLza7NjAIB2nevGf7887C.svg)](https://asciinema.org/a/w3PxQLza7NjAIB2nevGf7887C)


| Flags        | Value         | Description  |
| ------------ |:-------------:| -----:|
| --hostname   | example.com | Domain for scan. |
| --output     |  | Records the output with the domain name. |
| --list       | /tmp/lists/example.txt | Lister for subdomains.


## New Features
* - Checking HTTP/HTTPS statuses
* - Timeout time's updated for HTTP/HTTPS requests
* - Scanning just got faster
* - Scanning IP's
* - New terminal interface

## Installation

* - go get github.com/tismayil/rsdl
* - clone repo and build ( **go build rsdl.go** )


## Used Repos.

* - GO Spinner : github.com/briandowns/spinner - [ **go get github.com/briandowns/spinner** ]
* - GO Ping    : github.com/sparrc/go-ping     - [ **go get github.com/sparrc/go-ping** ]
