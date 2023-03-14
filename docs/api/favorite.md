# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/favorite.proto](#api_v1_favorite-proto)
    - [DeleteFavoriteRequest](#api-v1-DeleteFavoriteRequest)
    - [Favorite](#api-v1-Favorite)
    - [GetFavoriteRequest](#api-v1-GetFavoriteRequest)
    - [ListFavoriteReply](#api-v1-ListFavoriteReply)
    - [ListFavoriteRequest](#api-v1-ListFavoriteRequest)
  
    - [FavoriteAPI](#api-v1-FavoriteAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_favorite-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/favorite.proto



<a name="api-v1-DeleteFavoriteRequest"></a>

### DeleteFavoriteRequest
取消收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Favorite"></a>

### Favorite
文档收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| document_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| ext | [string](#string) |  |  |
| score | [int32](#int32) |  |  |
| size | [int64](#int64) |  |  |
| pages | [int32](#int32) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-GetFavoriteRequest"></a>

### GetFavoriteRequest
根据文章id，查询用户是否有收藏某篇文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  |  |






<a name="api-v1-ListFavoriteReply"></a>

### ListFavoriteReply
查询用户的收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| favorite | [Favorite](#api-v1-Favorite) | repeated |  |






<a name="api-v1-ListFavoriteRequest"></a>

### ListFavoriteRequest
查询用户的收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |





 

 

 


<a name="api-v1-FavoriteAPI"></a>

### FavoriteAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateFavorite | [Favorite](#api-v1-Favorite) | [Favorite](#api-v1-Favorite) | 添加收藏 |
| DeleteFavorite | [DeleteFavoriteRequest](#api-v1-DeleteFavoriteRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 取消收藏 |
| GetFavorite | [GetFavoriteRequest](#api-v1-GetFavoriteRequest) | [Favorite](#api-v1-Favorite) | 根据文章id，查询用户是否有收藏某篇文档 |
| ListFavorite | [ListFavoriteRequest](#api-v1-ListFavoriteRequest) | [ListFavoriteReply](#api-v1-ListFavoriteReply) | 查询用户的收藏 |

 



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

