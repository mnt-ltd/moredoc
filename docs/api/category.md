# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/category.proto](#api_v1_category-proto)
    - [Category](#api-v1-Category)
    - [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest)
    - [GetCategoryRequest](#api-v1-GetCategoryRequest)
    - [ListCategoryReply](#api-v1-ListCategoryReply)
    - [ListCategoryRequest](#api-v1-ListCategoryRequest)
  
    - [CategoryAPI](#api-v1-CategoryAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_category-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/category.proto



<a name="api-v1-Category"></a>

### Category
文档分类


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 分类ID |
| parent_id | [int32](#int32) |  | 父分类ID |
| title | [string](#string) |  | 分类标题 |
| doc_count | [int32](#int32) |  | 文档数量 |
| sort | [int32](#int32) |  | 排序，倒序排序，值越大越靠前 |
| enable | [bool](#bool) |  | 是否启用 |
| cover | [string](#string) |  | 分类封面 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteCategoryRequest"></a>

### DeleteCategoryRequest
删除分类请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCategoryRequest"></a>

### GetCategoryRequest
获取分类请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCategoryReply"></a>

### ListCategoryReply
分类列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| category | [Category](#api-v1-Category) | repeated | 分类列表 |






<a name="api-v1-ListCategoryRequest"></a>

### ListCategoryRequest
分类列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| parent_id | [int64](#int64) | repeated | 父分类ID |
| wd | [string](#string) |  | 搜索关键字 |
| enable | [bool](#bool) | repeated | 是否启用 |
| field | [string](#string) | repeated | 查询字段 |





 

 

 


<a name="api-v1-CategoryAPI"></a>

### CategoryAPI
文档分类API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建分类 |
| UpdateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新分类 |
| DeleteCategory | [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除分类 |
| GetCategory | [GetCategoryRequest](#api-v1-GetCategoryRequest) | [Category](#api-v1-Category) | 获取分类 |
| ListCategory | [ListCategoryRequest](#api-v1-ListCategoryRequest) | [ListCategoryReply](#api-v1-ListCategoryReply) | 分类列表 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

