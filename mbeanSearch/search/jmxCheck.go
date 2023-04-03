package search

/**
检查单个url是否开启jmx
*/
import (
	"context"
	"fmt"
	"github.com/newrelic/nrjmx/gojmx"
)

func CheckJmx(target string, port int, username string, pwd string, urlPath string) (bool, error) {
	// Try to connect to the JMX port (1099) on the target IP address
	config := &gojmx.JMXConfig{
		Hostname:         target,
		Port:             int32(port),
		RequestTimeoutMs: 10000,
		Username:         username,
		Password:         pwd,
		UriPath:          &urlPath,
	}

	//connect to jmx endpoint
	client, err := gojmx.NewClient(context.Background()).Open(config)
	if err != nil {
		return false, err
	}

	//判断是否为tomcat
	tomcat, err := client.QueryMBeanNames("Catalina:type=UserDatabase")
	if err != nil {
		fmt.Printf("Query Error by:%s", tomcat)
	}
	if len(tomcat) == 0 {
		fmt.Print("taget has no tomcat!\n")
	}

	//获取attribute
	for _, mbeanname := range tomcat {
		mBeanAttrNames, err := client.GetMBeanAttributeNames(mbeanname)
		if err != nil {
			fmt.Printf("%s has no attributes\n", mbeanname)
		}
		fmt.Printf("%s has attrbutes:%s\n", mbeanname, mBeanAttrNames)
	}

	//判断是否为springboot
	springboot, err := client.QueryMBeanNames("com.*.*:type=*")
	if err != nil {
		fmt.Printf("Query Error by:%s\n", springboot)
	}
	if len(springboot) == 0 {
		fmt.Printf("taget has no springboot!\n")
	} else {
		for _, mbean := range springboot {
			fmt.Println(mbean)
		}
	}
	// If an error occurred, return false
	return true, err
}
