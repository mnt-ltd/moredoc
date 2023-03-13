# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/comment.proto](#api_v1_comment-proto)
    - [CheckCommentRequest](#api-v1-CheckCommentRequest)
    - [Comment](#api-v1-Comment)
    - [CreateCommentRequest](#api-v1-CreateCommentRequest)
    - [DeleteCommentRequest](#api-v1-DeleteCommentRequest)
    - [GetCommentRequest](#api-v1-GetCommentRequest)
    - [ListCommentReply](#api-v1-ListCommentReply)
    - [ListCommentRequest](#api-v1-ListCommentRequest)
  
    - [CommentAPI](#api-v1-CommentAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_comment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/comment.proto



<a name="api-v1-CheckCommentRequest"></a>

### CheckCommentRequest
审核评论，修改评论状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated | 评论ID |
| status | [int32](#int32) |  | 状态，见 web/utils/enum.js 枚举 |






<a name="api-v1-Comment"></a>

### Comment
评论


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| id | [int64](#int64) |  | 评论ID |
| parent_id | [int64](#int64) |  | 父评论ID |
| content | [string](#string) |  | 评论内容 |
| document_id | [int64](#int64) |  | 文档ID |
| status | [int32](#int32) |  | 状态，见 web/utils/enum.js 枚举 |
| comment_count | [int32](#int32) |  | 回复数量 |
| user_id | [int64](#int64) |  | 用户ID |
| user | [User](#api-v1-User) |  | 用户信息 |
| document_title | [string](#string) |  | 文档标题 |






<a name="api-v1-CreateCommentRequest"></a>

### CreateCommentRequest
创建评论请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  | 文档ID |
| parent_id | [int64](#int64) |  | 父评论ID |
| content | [string](#string) |  | 评论内容 |
| captcha_id | [string](#string) |  | 验证码ID |
| captcha | [string](#string) |  | 验证码 |






<a name="api-v1-DeleteCommentRequest"></a>

### DeleteCommentRequest
删除评论请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCommentRequest"></a>

### GetCommentRequest
获取评论请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCommentReply"></a>

### ListCommentReply
获取评论列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| comment | [Comment](#api-v1-Comment) | repeated | 评论列表 |






<a name="api-v1-ListCommentRequest"></a>

### ListCommentRequest
获取评论列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键词 |
| field | [string](#string) | repeated | 查询的数据字段 |
| order | [string](#string) |  | 排序字段 |
| status | [int32](#int32) | repeated | 状态，见 web/utils/enum.js 枚举 |
| document_id | [int64](#int64) |  | 文档ID |
| user_id | [int64](#int64) |  | 用户ID |
| parent_id | [int64](#int64) | repeated | 父评论ID |
| with_document_title | [bool](#bool) |  | 是否返回文档标题 |





 

 

 


<a name="api-v1-CommentAPI"></a>

### CommentAPI
评论服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateComment | [CreateCommentRequest](#api-v1-CreateCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建评论 |
| UpdateComment | [Comment](#api-v1-Comment) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新评论，仅限管理员操作 |
| DeleteComment | [DeleteCommentRequest](#api-v1-DeleteCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 管理员或用户自己删除自己的评论 |
| GetComment | [GetCommentRequest](#api-v1-GetCommentRequest) | [Comment](#api-v1-Comment) | 获取单个评论 |
| ListComment | [ListCommentRequest](#api-v1-ListCommentRequest) | [ListCommentReply](#api-v1-ListCommentReply) | 获取评论列表 |
| CheckComment | [CheckCommentRequest](#api-v1-CheckCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 审核评论 |

 



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

