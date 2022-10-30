package util

import (
	"context"
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// CopyStruct 拷贝。注意：只能拷贝相同类型的结构体，且结构体中有json标签
func CopyStruct(srcPtr, dstPtr interface{}) (err error) {
	bytes, err := jsoniter.Marshal(srcPtr)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(bytes, dstPtr)
}

// GetGRPCRemoteIP 获取用户IP
func GetGRPCRemoteIP(ctx context.Context) (ips []string, err error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		headers := []string{"x-real-ip", "x-forwarded-for", "remote-addr"}
		for _, header := range headers {
			if values := md.Get(header); len(values) > 0 {
				for _, item := range values {
					ips = append(ips, strings.Split(item, ",")...)
				}
			}
		}
	}

	// 如果是 grpc client 直接调用该接口，则用这种方式来获取真实的IP地址
	if p, ok := peer.FromContext(ctx); ok {
		arr := strings.Split(p.Addr.String(), ":")
		if l := len(arr); l > 0 {
			ip := strings.Join(arr[0:l-1], ":")
			ip = strings.NewReplacer("[", "", "]", "").Replace(ip)
			ips = append(ips, ip)
		}
	}
	return
}

// 图片缩放居中裁剪
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

// LimitMin 数字最小值限制
func LimitMin(number int, minValue int) int {
	if number >= minValue {
		return number
	}
	return minValue
}

// LimitMax 数字最大值限制
func LimitMax(number int, maxValue int) int {
	if number >= maxValue {
		return maxValue
	}
	return number
}

// LimitRange 数字范围限制
func LimitRange(number int, min, max int) int {
	if number >= max {
		return max
	}
	if number <= min {
		return min
	}
	return number
}

type Any interface {
	~int | ~int64 | ~int32 | ~bool | ~string | ~float32 | ~float64 | ~uint | ~uint64 | ~uint32
}

// Slice2Interface 切片转interface切片
func Slice2Interface[T Any](slice []T) (values []interface{}) {
	for _, item := range slice {
		values = append(values, item)
	}
	return
}

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	inputFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(dst)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	return nil
}
