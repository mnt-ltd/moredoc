# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/document.proto](#api_v1_document-proto)
    - [CreateDocumentItem](#api-v1-CreateDocumentItem)
    - [CreateDocumentRequest](#api-v1-CreateDocumentRequest)
    - [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest)
    - [Document](#api-v1-Document)
    - [DocumentScore](#api-v1-DocumentScore)
    - [DownloadDocumentReply](#api-v1-DownloadDocumentReply)
    - [GetDocumentRequest](#api-v1-GetDocumentRequest)
    - [ListDocumentForHomeItem](#api-v1-ListDocumentForHomeItem)
    - [ListDocumentForHomeRequest](#api-v1-ListDocumentForHomeRequest)
    - [ListDocumentForHomeResponse](#api-v1-ListDocumentForHomeResponse)
    - [ListDocumentReply](#api-v1-ListDocumentReply)
    - [ListDocumentRequest](#api-v1-ListDocumentRequest)
    - [RecoverRecycleDocumentRequest](#api-v1-RecoverRecycleDocumentRequest)
    - [SearchDocumentReply](#api-v1-SearchDocumentReply)
    - [SearchDocumentRequest](#api-v1-SearchDocumentRequest)
    - [SetDocumentRecommendRequest](#api-v1-SetDocumentRecommendRequest)
  
    - [DocumentAPI](#api-v1-DocumentAPI)
    - [RecycleAPI](#api-v1-RecycleAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_document-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/document.proto



<a name="api-v1-CreateDocumentItem"></a>

### CreateDocumentItem
创建文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | 文档标题 |
| attachment_id | [int64](#int64) |  | 文档附件ID |
| price | [int32](#int32) |  | 文档价格 |






<a name="api-v1-CreateDocumentRequest"></a>

### CreateDocumentRequest
创建文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| overwrite | [bool](#bool) |  | 是否覆盖。暂时用不到 |
| category_id | [int64](#int64) | repeated | 文档分类ID |
| document | [CreateDocumentItem](#api-v1-CreateDocumentItem) | repeated | 文档列表 |






<a name="api-v1-DeleteDocumentRequest"></a>

### DeleteDocumentRequest
删除文档，放入回收站


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Document"></a>

### Document
文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文档ID |
| title | [string](#string) |  | 文档标题 |
| keywords | [string](#string) |  | 文档关键字 |
| description | [string](#string) |  | 文档描述 |
| user_id | [int64](#int64) |  | 文档作者 |
| cover | [string](#string) |  | 文档封面 |
| width | [int32](#int32) |  | 文档宽度 |
| height | [int32](#int32) |  | 文档高度 |
| preview | [int32](#int32) |  | 文档可预览页数，0表示不限制 |
| pages | [int32](#int32) |  | 文档页数 |
| uuid | [string](#string) |  | 文档UUID |
| download_count | [int32](#int32) |  | 文档下载次数 |
| view_count | [int32](#int32) |  | 文档浏览次数 |
| favorite_count | [int32](#int32) |  | 文档收藏次数 |
| comment_count | [int32](#int32) |  | 文档评论次数 |
| score | [int32](#int32) |  | 文档评分 |
| score_count | [int32](#int32) |  | 文档评分次数 |
| price | [int32](#int32) |  | 文档价格 |
| size | [int64](#int64) |  | 文档大小 |
| status | [int32](#int32) |  | 文档状态，见 web/utils/enum.js |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档更新时间 |
| deleted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档删除时间 |
| recommend_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档推荐时间 |
| deleted_user_id | [int64](#int64) |  | 删除文档的用户 |
| username | [string](#string) |  | 文档作者用户名 |
| category_id | [int64](#int64) | repeated | 文档分类ID |
| deleted_username | [string](#string) |  | 删除文档的用户名 |
| ext | [string](#string) |  | 文档扩展名 |
| attachment | [Attachment](#api-v1-Attachment) |  | 文档附件 |
| user | [User](#api-v1-User) |  | 文档作者 |
| enable_gzip | [bool](#bool) |  | 是否启用gzip压缩 |
| convert_error | [string](#string) |  | 转换错误信息 |






<a name="api-v1-DocumentScore"></a>

### DocumentScore
文档评分


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 评分ID |
| document_id | [int64](#int64) |  | 文档ID |
| user_id | [int64](#int64) |  | 用户ID |
| score | [int32](#int32) |  | 评分，100~500，100为1分，500为5分 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 评分时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DownloadDocumentReply"></a>

### DownloadDocumentReply
文档下载


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="api-v1-GetDocumentRequest"></a>

### GetDocumentRequest
查询文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文档ID |
| with_author | [bool](#bool) |  | 是否查询作者信息 |






<a name="api-v1-ListDocumentForHomeItem"></a>

### ListDocumentForHomeItem
首页文档查询返回项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [int64](#int64) |  | 分类ID |
| category_cover | [string](#string) |  | 分类封面 |
| category_name | [string](#string) |  | 分类名称 |
| document | [Document](#api-v1-Document) | repeated | 文档列表 |






<a name="api-v1-ListDocumentForHomeRequest"></a>

### ListDocumentForHomeRequest
查询文档（针对首页的查询）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |






<a name="api-v1-ListDocumentForHomeResponse"></a>

### ListDocumentForHomeResponse
查询文档（针对首页的查询）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document | [ListDocumentForHomeItem](#api-v1-ListDocumentForHomeItem) | repeated | 文档列表 |






<a name="api-v1-ListDocumentReply"></a>

### ListDocumentReply
文档列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文档总数 |
| document | [Document](#api-v1-Document) | repeated | 文档列表 |






<a name="api-v1-ListDocumentRequest"></a>

### ListDocumentRequest
文档列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段 |
| order | [string](#string) |  | 排序 |
| category_id | [int64](#int64) | repeated | 分类ID |
| user_id | [int64](#int64) | repeated | 用户ID |
| status | [int32](#int32) | repeated | 文档状态 |
| is_recommend | [bool](#bool) | repeated | 是否推荐 |
| limit | [int64](#int64) |  | 查询数量显示。当该值大于0时，page和size无效 |






<a name="api-v1-RecoverRecycleDocumentRequest"></a>

### RecoverRecycleDocumentRequest
恢复文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-SearchDocumentReply"></a>

### SearchDocumentReply
文档搜索响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文档总数 |
| spend | [string](#string) |  | 搜索耗时 |
| document | [Document](#api-v1-Document) | repeated | 文档列表 |






<a name="api-v1-SearchDocumentRequest"></a>

### SearchDocumentRequest
文档搜索


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 页码 |
| size | [int32](#int32) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| category_id | [int64](#int64) | repeated | 分类 |
| sort | [string](#string) |  | 排序 |
| ext | [string](#string) |  | 类型 |






<a name="api-v1-SetDocumentRecommendRequest"></a>

### SetDocumentRecommendRequest
设置文档推荐


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated | 文档ID |
| type | [int32](#int32) |  | 0, 取消推荐，1:推荐 2:重新推荐 |





 

 

 


<a name="api-v1-DocumentAPI"></a>

### DocumentAPI
文档服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListDocumentForHome | [ListDocumentForHomeRequest](#api-v1-ListDocumentForHomeRequest) | [ListDocumentForHomeResponse](#api-v1-ListDocumentForHomeResponse) | 针对首页的文档查询 |
| SetDocumentRecommend | [SetDocumentRecommendRequest](#api-v1-SetDocumentRecommendRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置文档推荐 |
| CreateDocument | [CreateDocumentRequest](#api-v1-CreateDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建文档 |
| UpdateDocument | [Document](#api-v1-Document) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新文档 |
| DeleteDocument | [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除文档 |
| GetDocument | [GetDocumentRequest](#api-v1-GetDocumentRequest) | [Document](#api-v1-Document) | 查询文档 |
| GetRelatedDocuments | [Document](#api-v1-Document) | [ListDocumentReply](#api-v1-ListDocumentReply) | 根据文档ID查询当前文档的相关文档 |
| DownloadDocument | [Document](#api-v1-Document) | [DownloadDocumentReply](#api-v1-DownloadDocumentReply) | 根据文档ID，获取文档下载链接 |
| ListDocument | [ListDocumentRequest](#api-v1-ListDocumentRequest) | [ListDocumentReply](#api-v1-ListDocumentReply) | 文档列表查询 |
| SearchDocument | [SearchDocumentRequest](#api-v1-SearchDocumentRequest) | [SearchDocumentReply](#api-v1-SearchDocumentReply) | 文档搜索 |
| SetDocumentScore | [DocumentScore](#api-v1-DocumentScore) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置文档评分 |
| GetDocumentScore | [DocumentScore](#api-v1-DocumentScore) | [DocumentScore](#api-v1-DocumentScore) | 获取当前登录用户的文档评分 |
| SetDocumentReconvert | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 将文档一键设置为重转 |


<a name="api-v1-RecycleAPI"></a>

### RecycleAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListRecycleDocument | [ListDocumentRequest](#api-v1-ListDocumentRequest) | [ListDocumentReply](#api-v1-ListDocumentReply) | 文档回收站列表 |
| RecoverRecycleDocument | [RecoverRecycleDocumentRequest](#api-v1-RecoverRecycleDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 恢复回收站文档，支持恢复单个文档或者是批量恢复 |
| DeleteRecycleDocument | [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除回收站文档 |
| ClearRecycleDocument | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 清空回收站文档 |

 



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

