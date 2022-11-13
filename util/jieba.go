package util

import (
	"unicode"

	"github.com/yanyiwu/gojieba"
)

var jieba *gojieba.Jieba

type Jieba struct {
	jieba *gojieba.Jieba
}

func NewJieba(dictDir ...string) *Jieba {
	defaultDir := "dict"
	if len(dictDir) > 0 {
		defaultDir = dictDir[0]
	}

	dicts := []string{
		defaultDir + "/jieba.dict.utf8",
		defaultDir + "/hmm_model.utf8",
		defaultDir + "/user.dict.utf8",
		defaultDir + "/idf.utf8",
		defaultDir + "/stop_words.utf8",
	}
	if jieba == nil {
		jieba = gojieba.NewJieba(dicts...)
	}
	return &Jieba{
		jieba: jieba,
	}
}

func (j *Jieba) AddWord(words ...string) {
	for _, word := range words {
		j.jieba.AddWord(word)
	}
}

func (j *Jieba) SegWords(text string, length ...int) (words []string) {
	topk := 10
	if len(length) > 0 {
		topk = length[0]
	}
	wds := j.jieba.Extract(text, topk)
	for _, wd := range wds {
		// 不是标点且不是空格也不是数字
		if unicode.IsSpace(rune(wd[0])) || unicode.IsPunct(rune(wd[0])) || unicode.IsDigit(rune(wd[0])) {
			continue
		}
		words = append(words, wd)
	}
	return
}
