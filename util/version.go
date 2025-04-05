package util

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	// Version 版本号
	Version = "0.0.1"
	// Hash 版本号
	Hash = "hash"
	// BuildAt 编译时间
	BuildAt = "2006-01-02 00:00:00"
)

type Release struct {
	TagName     string     `json:"tag_name"`
	Name        string     `json:"name"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Body        string     `json:"body"`
}

// 从 Gitee 获取最新版本
// https://gitee.com/api/v5/repos/mnt-ltd/moredoc/releases/latest
func GetLatestVersionFromGitee() (release *Release, err error) {
	return getLatestVersion("https://gitee.com/api/v5/repos/mnt-ltd/moredoc/releases/latest")
}

// 从Github获取最新版本
// https://api.github.com/repos/mnt-ltd/moredoc/releases/latest
func GetLatestVersionFromGithub() (release *Release, err error) {
	return getLatestVersion("https://gitee.com/api/v5/repos/mnt-ltd/moredoc/releases/latest")
}

func getLatestVersion(url string) (release *Release, err error) {
	resp, err := resty.New().R().SetHeader("Accept", "application/json").Get(url)
	if err != nil {
		return nil, err
	}
	release = &Release{}
	err = json.Unmarshal(resp.Body(), release)
	if err != nil {
		return nil, err
	}
	return release, nil
}
