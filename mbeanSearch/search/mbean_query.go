package search

import (
	"fmt"
	"github.com/newrelic/nrjmx/gojmx"
)

func Query_mbean(client gojmx.Client, mbeanName string) bool {
	mbeanNames, err := client.QueryMBeanNames(mbeanName)
	if err != nil {
		fmt.Printf("No such mbeanName by:%s", mbeanName)
		return false
	}
	//获取attribute
	for _, mbeanname := range mbeanNames {
		mBeanAttrNames, err := client.GetMBeanAttributeNames(mbeanname)
		if err != nil {
			fmt.Printf("%s has no attributes", mbeanname)
		}
		fmt.Printf("%s has attrbutes:%s", mbeanname, mBeanAttrNames)
	}
	return true
}
