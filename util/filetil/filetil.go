package filetil

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

var imagesExt = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
	".gif":  {},
	// ".bmp":  {},
	// ".webp": {},
}

var documentExt = map[string]struct{}{
	// word
	".doc": {}, ".docx": {}, ".rtf": {}, ".wps": {}, ".odt": {},
	// PPT
	".ppt": {}, ".pptx": {}, ".pps": {}, ".ppsx": {}, ".dps": {}, ".odp": {}, ".pot": {},
	// XLS
	".xls": {}, ".xlsx": {}, ".et": {}, ".ods": {},
	// 其他
	".epub": {}, ".umd": {}, ".chm": {}, ".mobi": {},
	// TXT
	".txt": {},
	// PDF
	".pdf": {},
}

// IsDocument 是否是文档
func IsDocument(ext string) bool {
	_, ok := documentExt[ext]
	return ok
}

// IsImage 判断文件是否是图片
func IsImage(ext string) bool {
	_, ok := imagesExt[ext]
	return ok
}

// GetFileMD5 获取文件MD5值
func GetFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// CropImage 居中裁剪图片
func CropImage(file string, width, height int) (err error) {
	var img image.Image
	img, err = imaging.Open(file)
	if err != nil {
		return
	}
	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".jpeg", ".jpg", ".png", ".gif":
		img = imaging.Fill(img, width, height, imaging.Center, imaging.CatmullRom)
	default:
		err = errors.New("unsupported image format")
		return
	}
	return imaging.Save(img, file)
}

// GetImageSize 获取图片宽高信息
func GetImageSize(file string) (width, height int, err error) {
	var img image.Image
	img, err = imaging.Open(file)
	if err != nil {
		return
	}
	width = img.Bounds().Max.X
	height = img.Bounds().Max.Y
	return
}
