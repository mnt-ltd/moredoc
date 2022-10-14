package util

import (
	"context"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// CopyStruct 拷贝。注意：只能拷贝相同类型的结构体，且结构体中有json标签
func CopyStruct(srcPtr, dstPtr interface{}) {
	bytes, _ := jsoniter.Marshal(srcPtr)
	jsoniter.Unmarshal(bytes, dstPtr)
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
