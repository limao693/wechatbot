package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	jrsc_rensheng_url = "https://v1.jinrishici.com/rensheng.txt"
)

func JinRiShiCiRenSheng() string {
	rsps, err := http.Get(jrsc_rensheng_url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return " "
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return " "
	}
	fmt.Println(string(body))
	return string(body)
}