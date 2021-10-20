package pkg

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/xmlquery"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandleOfficialAccountMsg(context string) error {
	fmt.Println("Begin handle msg")
	fmt.Println(context)
	UnmarshalXml(context)
	return nil
}

func UnmarshalXml(s string) (err error) {
	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		return err
	}

	title := xmlquery.FindOne(doc, "//title").InnerText()
	url := xmlquery.FindOne(doc, "//url").InnerText()
	fmt.Printf("title: %s\nurl: %s\n", title, url)

	context, err := GetWeiYuContect(url)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(context)

	return nil
}

func GetWeiYuContect(url string) (context string, err error) {
	fmt.Println("Begin get weiYun")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("3")
		return "", err
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	client := &http.Client{}
	response,err := client.Do(req)
	if err != nil {
		fmt.Println("1")
		return "", err
	}
	// 解析Response
	//res,err := ParseResponse(response)


	doc, err := goquery.NewDocumentFromReader(response.Body)
	// Find the review items
	doc.Find(".rich_media_content").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

	return "res", nil
}

//func GetWeiYuContect(url string) (context string, err error) {
//
//	doc, err := xmlquery.LoadURL(url)
//	if err != nil {
//		return "", err
//	}
//	s := xmlquery.FindOne(doc, "")
//	context = fmt.Sprintf("msg:\n %s", s.InnerText())
//	return context, nil
//}

func ParseResponse(response *http.Response) (string, error){
	body,err := ioutil.ReadAll(response.Body)

	return string(body),err
}
