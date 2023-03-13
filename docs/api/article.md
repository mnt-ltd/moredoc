# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/article.proto](#api_v1_article-proto)
    - [Article](#api-v1-Article)
    - [DeleteArticleRequest](#api-v1-DeleteArticleRequest)
    - [GetArticleRequest](#api-v1-GetArticleRequest)
    - [ListArticleReply](#api-v1-ListArticleReply)
    - [ListArticleRequest](#api-v1-ListArticleRequest)
  
    - [ArticleAPI](#api-v1-ArticleAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_article-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/article.proto



<a name="api-v1-Article"></a>

### Article
文章


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文章ID |
| identifier | [string](#string) |  | 文章唯一标识 |
| author | [string](#string) |  | 文章作者。如果为空，则使用网站名称作为作者 |
| view_count | [int64](#int64) |  | 文章浏览次数 |
| title | [string](#string) |  | 文章标题 |
| keywords | [string](#string) |  | 文章关键字 |
| description | [string](#string) |  | 文章描述 |
| content | [string](#string) |  | 文章内容 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章更新时间 |






<a name="api-v1-DeleteArticleRequest"></a>

### DeleteArticleRequest
删除文章请求，传入单个或者多个文章ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetArticleRequest"></a>

### GetArticleRequest
根据ID或者文章标识获取文章，二选一


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文章ID |
| identifier | [string](#string) |  | 文章唯一标识 |






<a name="api-v1-ListArticleReply"></a>

### ListArticleReply
文章列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文章总数 |
| article | [Article](#api-v1-Article) | repeated | 文章列表 |






<a name="api-v1-ListArticleRequest"></a>

### ListArticleRequest
文章列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段 |
| order | [string](#string) |  | 排序字段，根据指定的字段倒序排序 |





 

 

 


<a name="api-v1-ArticleAPI"></a>

### ArticleAPI
文章API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建文章 |
| UpdateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新文章 |
| DeleteArticle | [DeleteArticleRequest](#api-v1-DeleteArticleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除文章 |
| GetArticle | [GetArticleRequest](#api-v1-GetArticleRequest) | [Article](#api-v1-Article) | 获取文章 |
| ListArticle | [ListArticleRequest](#api-v1-ListArticleRequest) | [ListArticleReply](#api-v1-ListArticleReply) | 文章列表 |

 



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

