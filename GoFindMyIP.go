package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ipfiResponseRecord struct {
	IP string `json:ip`
}

func main() {
	ip, err := findMyIPbyParsing("http://myexternalip.com/raw", "", "")
	checkerror(err)
	fmt.Println(" Your IP is: " + ip)
	Useipfi()
	UseipfiJson()

}

func findMyIPbyParsing(url, marker1, marker2 string) (myIP string, err error) {
	var r *http.Response
	myIP = ""
	r, err = http.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()
	bufreader := bufio.NewReader(r.Body)
	for err == nil {
		var line string

		line, err = bufreader.ReadString('\n')
		if err != nil {
			return
		}
		if marker1 == "" || marker2 == "" {
			myIP = line
			break
		}
		i1 := strings.Index(line, marker1)
		if i1 > -1 {
			i2 := strings.Index(line, marker2)
			myIP = line[i1+len(marker1) : i2]
			break
		}
	}
	return
}

func UseipfiJson() {
	resp, err := http.Get("https://api.ipify.org?format=json") //client.Do(req)
	checkerror(err)
	defer resp.Body.Close()

	myrecord := new(ipfiResponseRecord)
	err = json.NewDecoder(resp.Body).Decode(&myrecord)
	checkerror(err)
	fmt.Printf(" My IP Address via useipfi JSON is: %v\n", *myrecord)

}

func Useipfi() {

	resp, err := http.Get("https://api.ipify.org")
	checkerror(err)
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	checkerror(err)

	fmt.Printf(" My IP Address via useipfi is: %s\n", content)

}

func checkerror(err error) {
	if err != nil {
		fmt.Printf("err %s detected \n", err)
		os.Exit(1)
	}
}
