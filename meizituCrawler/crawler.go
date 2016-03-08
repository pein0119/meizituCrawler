package meizituCrawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/iconv"
	"io/ioutil"
	"path/filepath"
	"sync"
	"fmt"
)

type Crawler struct {
	saveDir    string
	tagQueue   chan *Asset
	albumQueue chan *Asset
	imageQueue chan *Image
	stop       chan int
	wg         *sync.WaitGroup
}

func NewCrawler(saveDir string) *Crawler {
	c := new(Crawler)

	c.saveDir = saveDir
	c.tagQueue = make(chan *Asset)
	c.albumQueue = make(chan *Asset)
	c.imageQueue = make(chan *Image)
	c.stop = make(chan int)

	return c
}

func (c *Crawler) run() error {
	defer func() {
		close(c.tagQueue)
		close(c.albumQueue)
		close(c.imageQueue)
		close(c.stop)
	}()

	startUrl := "http://www.meizitu.com/"
	doc, err := goquery.NewDocument(startUrl)
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

		tag := NewAsset(title, url)
		c.tagQueue <- tag
	})

	for {
		select {
		case tag := <-c.tagQueue:
			go c.downloadTag(tag)
		case album := <-c.albumQueue:
			go c.downloadAlbum(album)
		case image := <-c.imageQueue:
			go c.downloadImage(image)
		case <-c.stop:
			return nil
		}
	}
}

func (c *Crawler) downloadTag(tag *Asset) error {
	
	return nil
}

func (c *Crawler) downloadAlbum(album *Asset) error {
	return nil
}

func (c *Crawler) downloadImage(image *Image) error {
	request := gorequest.New()
	_, body, err := request.Get(image.url).End()

	if err != nil {
		return nil
	}

	fileName := filepath.Join(c.saveDir, image.albumName, image.name)
	
	ioutil.WriteFile(fileName, []byte(body), 0644)

	return nil
}
