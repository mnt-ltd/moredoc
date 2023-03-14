package jieba

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/wangbin/jiebago"
)

var (
	chineseChar = map[string]bool{
		"，": true, "。": true, "？": true, "！": true, "、": true, "：": true, "；": true,
		"（": true, "）": true, "《": true, "》": true, "“": true, "”": true, "‘": true,
		"’": true, "【": true, "】": true, "『": true, "』": true, "〔": true, "〕": true,
		"〈": true, "〉": true, "﹑": true, "●": true, "…": true, "—": true, "～": true,
	}
	seg        jiebago.Segmenter
	dictPath   = "dictionary/dict.txt"
	dictLoaded = false
)

func init() {
	err := loadDictionary()
	if err != nil {
		panic(err)
	}
}

func loadDictionary(dict ...string) error {
	if dictLoaded {
		return nil
	}
	if len(dict) > 0 {
		path := dict[0]
		if path != "" {
			dictPath = path
		}
	}
	err := seg.LoadDictionary(dictPath)
	if err == nil {
		dictLoaded = true
	}
	return err
}

func SegWords(text string) []string {
	words := make([]string, 0)
	exist := make(map[string]bool)
	for word := range seg.CutForSearch(text, true) {
		wd := strings.TrimSpace(word)
		if _, ok := exist[wd]; ok {
			continue
		}
		exist[wd] = true
		l := utf8.RuneCountInString(wd)
		if !(l < 2 || (l == 1 && (isChineseChar(wd) || unicode.IsPunct(rune(wd[0]))))) { // 跳过空格和标点符号
			words = append(words, wd)
		}
	}
	return words
}

func isChineseChar(term string) bool {
	_, ok := chineseChar[term]
	return ok
}
