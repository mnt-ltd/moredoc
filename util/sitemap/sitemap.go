package sitemap

import (
	"fmt"
	"os"
	"strings"
)

//sitemap工具
//<lastmod>是用来指定该链接的最后更新时间，这个很重要。Google的机器人会在索引此链接前先和上次索引记录的最后更新时间进行比较，如果时间一样就会跳过不再索引。所以如果你的链接内容基于上次Google索引时的内容有所改变，应该更新该时间，让Google下次索引时会重新对该链接内容进行分析和提取关键字。这里必须用ISO 8601中指定的时间格式进行描述，格式化的时间格式如下：
//年：YYYY(2005)
//年和月：YYYY-MM(2005-06)
//年月日：YYYY-MM-DD(2005-06-04)
//年月日小时分钟：YYYY-MM-DDThh:mmTZD(2005-06-04T10:37+08:00)
//年月日小时分钟秒：YYYY-MM-DDThh:mmTZD(2005-06-04T10:37:30+08:00)
//这里需注意的是TZD，TZD指定就是本地时间区域标记，像中国就是+08:00了

type ChangeFreq string

// "always", "hourly", "daily", "weekly", "monthly", "yearly"
const (
	ALWAYS  ChangeFreq = "always"
	HOURLY  ChangeFreq = "hourly"
	DAILY   ChangeFreq = "daily"
	WEEKLY  ChangeFreq = "weekly"
	MONTHLY ChangeFreq = "monthly"
	YEARLY  ChangeFreq = "yearly"
)

// sitemap索引
type SitemapIndex struct {
	Loc     string // 必填，sitemap链接地址
	Lastmod string // 索引最后更新时间
}

// sitemap链接结构
type SitemapUrl struct {
	Loc        string     // 链接，必填，长度不能超过256个字符
	Lastmod    string     //链接内容最后更新时间，选填
	ChangeFreq ChangeFreq //链接更新频率，选填
	Priority   float32    //权重，0.0-1.0之间
}

type Sitemap struct {
	Version  string //xml版本，默认1.0
	Encoding string //xml字符编码
}

type SitemapOption struct {
	Version  string
	Encoding string
}

func NewSitemap(opt ...SitemapOption) *Sitemap {
	var version = "1.0"
	var encoding = "utf-8"
	if len(opt) > 0 {
		if len(opt[0].Version) > 0 {
			version = opt[0].Version
		}
		if len(opt[0].Encoding) > 0 {
			encoding = opt[0].Encoding
		}
	}
	return &Sitemap{Version: version, Encoding: encoding}
}

// CreateSitemapIndex 生成站点地图索引
func (m *Sitemap) CreateSitemapIndex(si []SitemapIndex, file string) error {
	var xmls = []string{fmt.Sprintf(`<?xml version="%v" encoding="%v"?><sitemapindex>`, m.Version, m.Encoding)}
	for _, v := range si {
		xml := ""
		if len(v.Lastmod) > 0 {
			xml = fmt.Sprintf(`<sitemap><loc>%v</loc><lastmod>%v</lastmod></sitemap>`, v.Loc, v.Lastmod)
		} else {
			xml = fmt.Sprintf(`<sitemap><loc>%v</loc></sitemap>`, v.Loc)
		}
		xmls = append(xmls, xml)
	}
	xmls = append(xmls, `</sitemapindex>`)
	return os.WriteFile(file, []byte(strings.Join(xmls, "")), os.ModePerm)
}

// 生成站点地图内容
func (m *Sitemap) CreateSitemapContent(su []SitemapUrl, file string) error {
	var xmls = []string{fmt.Sprintf(`<?xml version="%v" encoding="%v"?><urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">`,
		m.Version, m.Encoding,
	)}
	for _, v := range su {
		xmls = append(xmls, "<url>")
		xmls = append(xmls, fmt.Sprintf(`<loc>%v</loc>`, v.Loc))
		if v.Priority > 0 {
			xmls = append(xmls, fmt.Sprintf(`<priority>%v</priority>`, v.Priority))
		}
		if len(v.Lastmod) > 0 {
			xmls = append(xmls, fmt.Sprintf(`<lastmod>%v</lastmod>`, v.Lastmod))
		}
		if len(v.ChangeFreq) > 0 {
			xmls = append(xmls, fmt.Sprintf(`<changefreq>%v</changefreq>`, v.ChangeFreq))
		}
		xmls = append(xmls, "</url>")
	}
	xmls = append(xmls, `</urlset>`)
	return os.WriteFile(file, []byte(strings.Join(xmls, "")), os.ModePerm)
}
