package converter

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"moredoc/util"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	soffice      = "soffice"
	ebookConvert = "ebook-convert"
	svgo         = "svgo"
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

func NewConverter(logger *zap.Logger, timeout ...time.Duration) *Converter {
	expire := 1 * time.Hour
	if len(timeout) > 0 {
		expire = timeout[0]
	}
	defaultCachePath := "cache/convert"
	os.MkdirAll(defaultCachePath, os.ModePerm)
	return &Converter{
		cachePath: defaultCachePath,
		timeout:   expire,
		logger:    logger.Named("converter"),
	}
}

func (c *Converter) SetCachePath(cachePath string) {
	os.MkdirAll(cachePath, os.ModePerm)
	c.cachePath = cachePath
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
	case ".pdf":
		return c.PDFToPDF(src)
	// case ".doc", ".docx", ".rtf", ".wps", ".odt",
	// 	".xls", ".xlsx", ".et", ".ods",
	// 	".ppt", ".pptx", ".dps", ".odp", ".pps", ".ppsx", ".pot", ".potx":
	// 	return c.ConvertOfficeToPDF(src)
	default:
		return c.ConvertOfficeToPDF(src)
		// return "", fmt.Errorf("不支持的文件类型：%s", ext)
	}
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
	c.logger.Debug("convert pdf to txt", zap.String("cmd", mutool), zap.Strings("args", args))
	_, err = util.ExecCommand(mutool, args, c.timeout)
	if err != nil {
		c.logger.Error("convert pdf to txt", zap.String("cmd", mutool), zap.Strings("args", args), zap.Error(err))
		return
	}
	return dst, nil
}

// ConvertCHMToPDF 将CHM转为PDF
func (c *Converter) ConvertCHMToPDF(src string) (dst string, err error) {
	return c.convertToPDFByCalibre(src)
}

// ConvertPDFToSVG 将PDF转为SVG
func (c *Converter) ConvertPDFToSVG(src string, fromPage, toPage int, enableSVGO, enableGZIP bool) (pages []Page, err error) {
	pages, err = c.convertPDFToPage(src, fromPage, toPage, ".svg")
	if err != nil {
		return
	}

	if len(pages) == 0 {
		return
	}

	if enableSVGO { // 压缩svg
		c.CompressSVGBySVGO(filepath.Dir(pages[0].PagePath))
	}

	if enableGZIP { // gzip 压缩
		for idx, page := range pages {
			if dst, errCompress := c.CompressSVGByGZIP(page.PagePath); errCompress == nil {
				os.Remove(page.PagePath)
				page.PagePath = dst
				pages[idx] = page
			}
		}
	}
	return
}

// ConvertPDFToPNG 将PDF转为PNG
func (c *Converter) ConvertPDFToPNG(src string, fromPage, toPage int) (pages []Page, err error) {
	return c.convertPDFToPage(src, fromPage, toPage, ".png")
}

func (c *Converter) PDFToPDF(src string) (dst string, err error) {
	dst = strings.ReplaceAll(filepath.Join(c.cachePath, time.Now().Format(dirDteFmt), filepath.Base(src)), "\\", "/")
	err = util.CopyFile(src, dst)
	if err != nil {
		c.logger.Error("copy file error", zap.Error(err))
	}
	return
}

// ext 可选值： .png, .svg
func (c *Converter) convertPDFToPage(src string, fromPage, toPage int, ext string) (pages []Page, err error) {
	pageRange := fmt.Sprintf("%d-%d", fromPage, toPage)
	cacheFile := strings.ReplaceAll(filepath.Join(c.cachePath, time.Now().Format(dirDteFmt), strings.TrimSuffix(filepath.Base(src), filepath.Ext(src))+"/%d"+ext), "\\", "/")
	args := []string{
		"convert",
		"-o",
		cacheFile,
		src,
		pageRange,
	}
	os.MkdirAll(filepath.Dir(cacheFile), os.ModePerm)
	c.logger.Debug("convert pdf to page", zap.String("cmd", mutool), zap.Strings("args", args))
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
			PageNum:  fromPage + i,
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
	c.logger.Debug("convert to pdf by soffice", zap.String("cmd", soffice), zap.Strings("args", args))
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
	c.logger.Debug("convert to pdf by calibre", zap.String("cmd", ebookConvert), zap.Strings("args", args))
	_, err = util.ExecCommand(ebookConvert, args, c.timeout)
	if err != nil {
		c.logger.Error("convert to pdf by calibre", zap.String("cmd", ebookConvert), zap.Strings("args", args), zap.Error(err))
	}
	return
}

func (c *Converter) CountPDFPages(file string) (pages int, err error) {
	args := []string{
		"show",
		file,
		"pages",
	}
	c.logger.Debug("count pdf pages", zap.String("cmd", mutool), zap.Strings("args", args))
	var out string
	out, err = util.ExecCommand(mutool, args, c.timeout)
	if err != nil {
		c.logger.Error("count pdf pages", zap.String("cmd", mutool), zap.Strings("args", args), zap.Error(err))
		return
	}

	lines := strings.Split(out, "\n")
	length := len(lines)
	for i := length - 1; i >= 0; i-- {
		line := strings.TrimSpace(strings.ToLower(lines[i]))
		c.logger.Debug("count pdf pages", zap.String("line", line))
		if strings.HasPrefix(line, "page") {
			pages, _ = strconv.Atoi(strings.TrimSpace(strings.TrimLeft(strings.Split(line, "=")[0], "page")))
			if pages > 0 {
				return
			}
		}
	}
	return
}

func (c *Converter) ExistMupdf() (err error) {
	_, err = exec.LookPath(mutool)
	return
}

func (c *Converter) ExistSoffice() (err error) {
	_, err = exec.LookPath(soffice)
	return
}

func (c *Converter) ExistCalibre() (err error) {
	_, err = exec.LookPath(ebookConvert)
	return
}

func (c *Converter) ExistSVGO() (err error) {
	_, err = exec.LookPath(svgo)
	return
}

func (c *Converter) CompressSVGBySVGO(svgFolder string) (err error) {
	args := []string{
		"-f",
		svgFolder,
	}
	c.logger.Debug("compress svg by svgo", zap.String("cmd", svgo), zap.Strings("args", args))
	var out string
	out, err = util.ExecCommand(svgo, args, c.timeout*10)
	if err != nil {
		c.logger.Error("compress svg by svgo", zap.String("cmd", svgo), zap.Strings("args", args), zap.Error(err))
	}
	c.logger.Debug("compress svg by svgo", zap.String("out", out))
	return
}

// CompressSVGByGZIP 将SVG文件压缩为GZIP格式
func (c Converter) CompressSVGByGZIP(svgFile string) (dst string, err error) {
	var svgBytes []byte
	ext := filepath.Ext(svgFile)
	dst = strings.TrimSuffix(svgFile, ext) + ".gzip.svg"
	svgBytes, err = os.ReadFile(svgFile)
	if err != nil {
		c.logger.Error("read svg file", zap.String("svgFile", svgFile), zap.Error(err))
		return
	}
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	w.Write(svgBytes)
	w.Flush()
	err = os.WriteFile(dst, buf.Bytes(), os.ModePerm)
	if err != nil {
		c.logger.Error("write svgz file", zap.String("svgzFile", dst), zap.Error(err))
	}
	return
}
