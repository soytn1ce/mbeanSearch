package main

import (
	"flag"
	"fmt"
	"mbeanSearch/search"
	"os"
)

func main() {
	var network string
	var outputFile string
	var target string
	var port int
	var file string
	var threadNum int
	var usernamewordlist string
	var passwordwordlist string
	var urlpath string

	//参数指定
	flag.StringVar(&network, "i", "", "For Input a file record ip address")
	flag.StringVar(&file, "f", "", "For Input a file record target list")
	flag.StringVar(&target, "u", "", "For Input target, use ',' to split")
	flag.StringVar(&outputFile, "o", "ip_record.txt", "For Input output result")
	flag.StringVar(&usernamewordlist, "wu", "", "For Input burte username wordlist")
	flag.StringVar(&passwordwordlist, "wp", "", "For Input burte password wordlist")
	flag.StringVar(&urlpath, "up", "jmxrmi", "For Input jmx url path")
	flag.IntVar(&threadNum, "t", 1, "For input a number to make sure scan threads number")
	flag.IntVar(&port, "p", 1099, "For intput jmx port")
	flag.Parse()

	if target != "" {
		hasJmx, err := search.CheckJmx(target, port, usernamewordlist, passwordwordlist, urlpath)
		if err != nil {
			fmt.Printf("error:%s", err)
			return
		}
		if hasJmx {
			fmt.Printf("JMX service is enabled on the target IP address:%s\n", target)
		} else {
			fmt.Printf("no!")
		}
	}
	if network != "" {
		// 扫描网段
		scanner := search.NewJmxScanner(network, outputFile, threadNum)
		// Scan for JMX services and write the results to the output file
		err := scanner.Scan(port)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	}

}
