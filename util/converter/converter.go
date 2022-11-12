package converter

import (
	"fmt"
	"moredoc/util"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	soffice      = "soffice"
	ebookConvert = "ebook-convert"
	mutool       = "mutool"
	dirDteFmt    = "2006/01/02/15"
)

type ConvertCallback func(page int, pagePath string, err error)

type Converter struct {
	cachePath string
	timeout   time.Duration
	logger    *zap.Logger
}

type Page struct {
	PageNum  int
	PagePath string
}

func NewConverter(logger *zap.Logger, cachePath string, timeout ...time.Duration) *Converter {
	expire := 1 * time.Hour
	if len(timeout) > 0 {
		expire = timeout[0]
	}
	os.MkdirAll(cachePath, os.ModePerm)
	return &Converter{
		cachePath: cachePath,
		timeout:   expire,
		logger:    logger.Named("converter"),
	}
}

// ConvertToPDF 将文件转为PDF。
// 自动根据文件类型调用相应的转换函数。
func (c *Converter) ConvertToPDF(src string) (dst string, err error) {
	ext := strings.ToLower(filepath.Ext(src))
	switch ext {
	case ".epub":
		return c.ConvertEPUBToPDF(src)
	case ".umd":
		return c.ConvertUMDToPDF(src)
	case ".txt":
		return c.ConvertTXTToPDF(src)
	case ".mobi":
		return c.ConvertMOBIToPDF(src)
	case ".chm":
		return c.ConvertCHMToPDF(src)
	case ".doc", ".docx", ".rtf", ".wps", ".odt",
		".xls", ".xlsx", ".et", ".ods",
		".ppt", ".pptx", ".dps", ".odp", ".pps", ".ppsx", ".pot", ".potx":
		return c.ConvertOfficeToPDF(src)
	default:
		return "", fmt.Errorf("不支持的文件类型：%s", ext)
	}
	return
}

// ConvertOfficeToPDF 通过soffice将office文档转换为pdf
func (c *Converter) ConvertOfficeToPDF(src string) (dst string, err error) {
	return c.convertToPDFBySoffice(src)
}

// ConvertEPUBToPDF 将 epub 转为PDF
func (c *Converter) ConvertEPUBToPDF(src string) (dst string, err error) {
	return c.convertToPDFByCalibre(src)
}

// ConvertUMDToPDF 将umd转为PDF
func (c *Converter) ConvertUMDToPDF(src string) (dst string, err error) {
	return c.convertToPDFBySoffice(src)
}

// ConvertTXTToPDF 将 txt 转为PDF
func (c *Converter) ConvertTXTToPDF(src string) (dst string, err error) {
	return c.convertToPDFBySoffice(src)
}

// ConvertMOBIToPDF 将 mobi 转为PDF
func (c *Converter) ConvertMOBIToPDF(src string) (dst string, err error) {
	return c.convertToPDFByCalibre(src)
}

// ConvertPDFToTxt 将PDF转为TXT
func (c *Converter) ConvertPDFToTxt(src string) (dst string, err error) {
	dst = strings.ReplaceAll(filepath.Join(c.cachePath, time.Now().Format(dirDteFmt), filepath.Base(src)+".txt"), "\\", "/")
	args := []string{
		"convert",
		"-o",
		dst,
		src,
	}
	os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	c.logger.Info("convert pdf to txt", zap.String("cmd", mutool), zap.Strings("args", args))
	_, err = util.ExecCommand(mutool, args, c.timeout)
	if err != nil {
		return
	}
	return dst, nil
}

// ConvertCHMToPDF 将CHM转为PDF
func (c *Converter) ConvertCHMToPDF(src string) (dst string, err error) {
	return c.convertToPDFByCalibre(src)
}

// ConvertPDFToSVG 将PDF转为SVG
func (c *Converter) ConvertPDFToSVG(src string, fromPage, toPage int) (pages []Page, err error) {
	return c.convertPDFToPage(src, fromPage, toPage, ".svg")
}

// ConvertPDFToPNG 将PDF转为PNG
func (c *Converter) ConvertPDFToPNG(src string, fromPage, toPage int) (pages []Page, err error) {
	return c.convertPDFToPage(src, fromPage, toPage, ".png")
}

// ext 可选值： .png, .svg
func (c *Converter) convertPDFToPage(src string, fromPage, toPage int, ext string) (pages []Page, err error) {
	pageRange := fmt.Sprintf("%d-%d", fromPage, toPage)
	cacheFile := strings.ReplaceAll(filepath.Join(c.cachePath, time.Now().Format(dirDteFmt), filepath.Base(src)+"/%d"+ext), "\\", "/")
	args := []string{
		"convert",
		"-o",
		cacheFile,
		src,
		pageRange,
	}
	os.MkdirAll(filepath.Dir(cacheFile), os.ModePerm)
	c.logger.Info("convert pdf to page", zap.String("cmd", mutool), zap.Strings("args", args))
	_, err = util.ExecCommand(mutool, args, c.timeout)
	if err != nil {
		return
	}

	for i := 0; i <= toPage-fromPage; i++ {
		pagePath := fmt.Sprintf(cacheFile, i+1)
		if _, errPage := os.Stat(pagePath); errPage != nil {
			break
		}
		pages = append(pages, Page{
			PageNum:  fromPage + i + 1,
			PagePath: pagePath,
		})
	}
	return
}

func (c *Converter) convertToPDFBySoffice(src string) (dst string, err error) {
	basename := filepath.Base(src)
	dst = strings.ReplaceAll(filepath.Join(c.cachePath, time.Now().Format(dirDteFmt), basename, strings.TrimSuffix(basename, filepath.Ext(src))+".pdf"), "\\", "/")
	args := []string{
		"--headless",
		"--convert-to",
		"pdf",
		"--outdir",
		filepath.Dir(dst),
	}
	args = append(args, src)
	os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	c.logger.Info("convert to pdf by soffice", zap.String("cmd", soffice), zap.Strings("args", args))
	_, err = util.ExecCommand(soffice, args, c.timeout)
	if err != nil {
		c.logger.Error("convert to pdf by soffice", zap.String("cmd", soffice), zap.Strings("args", args), zap.Error(err))
	}
	return
}

func (c *Converter) convertToPDFByCalibre(src string) (dst string, err error) {
	basename := filepath.Base(src)
	dst = strings.ReplaceAll(filepath.Join(c.cachePath, time.Now().Format(dirDteFmt), basename, basename+".pdf"), "\\", "/")
	args := []string{
		src,
		dst,
		"--paper-size", "a4",
		// "--pdf-default-font-size", "14",
		"--pdf-page-margin-bottom", "36",
		"--pdf-page-margin-left", "36",
		"--pdf-page-margin-right", "36",
		"--pdf-page-margin-top", "36",
	}
	os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	c.logger.Info("convert to pdf by calibre", zap.String("cmd", ebookConvert), zap.Strings("args", args))
	_, err = util.ExecCommand(ebookConvert, args, c.timeout)
	if err != nil {
		c.logger.Error("convert to pdf by calibre", zap.String("cmd", ebookConvert), zap.Strings("args", args), zap.Error(err))
	}
	return
}
