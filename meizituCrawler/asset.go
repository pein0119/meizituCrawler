package meizituCrawler

// 标签和专辑
type Asset struct {
	name       string // 名称
	url        string // 链接
}

func NewAsset(name string, url string) *Asset {
	asset := new(Asset)
	
	asset.name = name
	asset.url = url
	
	return asset
}

type Image struct {
	Asset
	albumName string			// 图片所属专辑
}

func NewImage(name string, url string, albumName string) *Image {
	image := new(Image)

	image.name = name
	image.url = url
	image.albumName = albumName

	return image
}
