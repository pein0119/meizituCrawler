package main

import (
	// "io/ioutil"
	// "github.com/parnurzeal/gorequest"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"github.com/qiniu/iconv"
)

func main() {
	// request := gorequest.New()
	// _, body, errs := request.Get("http://mm.howkuai.com/wp-content/uploads/2015a/11/20/01.jpg").End()

	// if errs == nil {
	// 	image := []byte(body)
	// 	ioutil.WriteFile("1.jpg", image, 0644)
	// }

	doc, err := goquery.NewDocument("http://www.meizitu.com/")
	if err != nil {
		fmt.Println(err.Error())
	}

	doc.Find(".tags a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		
		title := s.Text()
		cd, err := iconv.Open("utf-8", "gb2312")
		if err != nil {
			fmt.Println("iconv.Open failed!")
		}
		defer cd.Close()

		title = cd.ConvString(title)
		fmt.Println(title)
	});
}
