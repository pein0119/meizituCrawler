package meizituCrawler

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
