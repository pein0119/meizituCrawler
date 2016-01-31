package meizituCrawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"os"
)

type Crawler struct {
	saveDir string
	tagQueue   chan Tag
	albumQueue chan Album
	imageQueue chan Image
	stop      chan int
	wg        *sync.WaitGroup
}

func NewCrawler(saveDir string) *Crawler {
	c := new(Crawler)

	c.saveDir = saveDir
	c.tagChan = make(chan Tag)
	c.imageChan = make(chan Image)
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

	startUrl = "http://www.meizitu.com/"
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

		tag := NewTag(title, url)
		c.tagQueue <- tag
	})

	for {
		select {
		case tag := <- tagQueue:
			go c.downloadTag(tag)
		case album := <- albumQueue:
			go c.downloadAlbum(album)
		case image := <- imageQueue:
			go c.downloadImage(image)
		case <-c.stop:
			return
		}
	}
}

func (c *Crawler) downloadTag(Tag) error {
	
}

func (c *Crawler) downloadAlbum(album) error {
	
}


func (c *Crawler) downloadImage(image) error {
	request := gorequest.New()
	_, body, err := request.Get(image.url).End()

	if err != nil {
		return 1
	}

	if FileExist(path) {
		return nil
	}
	fileName := c.saveDir + os.PathSeparator +  image.name
	ioutil.WriteFile(fileName, []byte(body), 0644)
}
