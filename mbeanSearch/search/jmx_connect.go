package search

/**
该函数通过指定的网段，扫描开启jmx的ip并输出到指定文件
*/
import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

type JmxScanner struct {
	network    string
	outputFile string
	numThreads int
}

// Create a new JmxScanner object
func NewJmxScanner(network string, outputFile string, numThreads int) *JmxScanner {
	return &JmxScanner{network: network, outputFile: outputFile, numThreads: numThreads}
}

// Scan for JMX services and save the IP addresses with JMX services enabled to a file
func (j *JmxScanner) Scan(port int) error {
	// Parse the IP network segment
	ip, ipNet, err := net.ParseCIDR(j.network)
	if err != nil {
		return err
	}

	// Create a slice to hold the IP addresses with JMX services enabled
	var jmxIps []string

	// Loop over all IP addresses in the network segment and scan for JMX services
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		_, err := CheckJmx(ip.String(), port, "", "", "jmxrmi")
		if err == nil {
			jmxIps = append(jmxIps, ip.String())
			fmt.Printf("the ip address:%s has connected\n", ip)
		} else {
			fmt.Println(err)
			fmt.Println("ip address:%s has no jvm", ip.String())
		}
	}
	fmt.Printf("JMX Server Scanning over!")

	// Write the IP addresses with JMX services enabled to the output file
	var sb strings.Builder
	for _, ip := range jmxIps {
		sb.WriteString(ip)
		sb.WriteString("\n")
	}
	data := []byte(sb.String())
	err = ioutil.WriteFile(j.outputFile, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

// Helper function to increment an IP address
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
