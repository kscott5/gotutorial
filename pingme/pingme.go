package main

import (
	"fmt"
	"net"
)

// PingMe allows various net/http operations
type PingMe struct {
	// URLPath the web address with format wwww.somedomain.com
	URLPath string
}

// URLPathIPLookup retrieves the ip address for the PingMe.URLPath
func (pm * PingMe) URLPathIPLookup() {
	ips, err := net.LookupIP(pm.URLPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s ip address list: %s\n", pm.URLPath, ips)

}

func main() {
	var pm PingMe
	pm.URLPath = "www.github.com"
	pm.URLPathIPLookup()
}