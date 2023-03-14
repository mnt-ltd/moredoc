# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/banner.proto](#api_v1_banner-proto)
    - [Banner](#api-v1-Banner)
    - [DeleteBannerRequest](#api-v1-DeleteBannerRequest)
    - [GetBannerRequest](#api-v1-GetBannerRequest)
    - [ListBannerReply](#api-v1-ListBannerReply)
    - [ListBannerRequest](#api-v1-ListBannerRequest)
  
    - [BannerAPI](#api-v1-BannerAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_banner-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/banner.proto



<a name="api-v1-Banner"></a>

### Banner
banner，轮播图


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 主键 |
| title | [string](#string) |  | 标题，名称 |
| path | [string](#string) |  | 图片地址 |
| sort | [int32](#int32) |  | 排序，值越大越靠前 |
| enable | [bool](#bool) |  | 是否启用 |
| type | [int32](#int32) |  | 类型，如PC横幅、小程序横幅等，见 web/utils/enum.js 中的枚举 |
| url | [string](#string) |  | 跳转地址 |
| description | [string](#string) |  | 描述 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteBannerRequest"></a>

### DeleteBannerRequest
删除横幅


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetBannerRequest"></a>

### GetBannerRequest
获取横幅


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListBannerReply"></a>

### ListBannerReply
横幅列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| banner | [Banner](#api-v1-Banner) | repeated | 横幅数组 |






<a name="api-v1-ListBannerRequest"></a>

### ListBannerRequest
横幅列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| type | [int32](#int32) | repeated | 类型 |
| enable | [bool](#bool) | repeated | 是否启用 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段，不指定，则查询全部 |





 

 

 


<a name="api-v1-BannerAPI"></a>

### BannerAPI
横幅API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBanner | [Banner](#api-v1-Banner) | [Banner](#api-v1-Banner) | 创建横幅 |
| UpdateBanner | [Banner](#api-v1-Banner) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新横幅 |
| DeleteBanner | [DeleteBannerRequest](#api-v1-DeleteBannerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除横幅 |
| GetBanner | [GetBannerRequest](#api-v1-GetBannerRequest) | [Banner](#api-v1-Banner) | 查询横幅 |
| ListBanner | [ListBannerRequest](#api-v1-ListBannerRequest) | [ListBannerReply](#api-v1-ListBannerReply) | 横幅列表 |

 



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

