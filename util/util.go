package util

import jsoniter "github.com/json-iterator/go"

// CopyStruct 拷贝。注意：只能拷贝相同类型的结构体，且结构体中有json标签
func CopyStruct(srcPtr, dstPtr interface{}) {
	bytes, _ := jsoniter.Marshal(srcPtr)
	jsoniter.Unmarshal(bytes, dstPtr)
}
