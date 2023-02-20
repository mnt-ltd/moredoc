package gse

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg        gse.Segmenter
	posSeg     pos.Segmenter
	loadedDict = false
)

func init() {
	go loadDict()
}

func loadDict() {
	err := seg.LoadDictEmbed()
	if err != nil {
		fmt.Println("seg.LoadDictEmbed", err)
		return
	}
	err = seg.LoadStopEmbed()
	if err != nil {
		fmt.Println("seg.LoadStopEmbed", err)
		return
	}
	loadedDict = true
}

func SegWords(text string) (words []string) {
	if !loadedDict {
		return
	}
	wds := seg.Cut(text)
	for _, wd := range wds {
		// 跳过单字、空格、标点、数字
		if utf8.RuneCountInString(wd) == 1 || unicode.IsSpace(rune(wd[0])) || unicode.IsPunct(rune(wd[0])) || unicode.IsDigit(rune(wd[0])) {
			continue
		}
		words = append(words, wd)
	}
	return
}
