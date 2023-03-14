# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/attachment.proto](#api_v1_attachment-proto)
    - [Attachment](#api-v1-Attachment)
    - [DeleteAttachmentRequest](#api-v1-DeleteAttachmentRequest)
    - [GetAttachmentRequest](#api-v1-GetAttachmentRequest)
    - [ListAttachmentReply](#api-v1-ListAttachmentReply)
    - [ListAttachmentRequest](#api-v1-ListAttachmentRequest)
  
    - [AttachmentAPI](#api-v1-AttachmentAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_attachment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/attachment.proto



<a name="api-v1-Attachment"></a>

### Attachment
附件


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 附件ID |
| hash | [string](#string) |  | 附件哈希值，MD5 |
| user_id | [int64](#int64) |  | 上传用户ID |
| type_id | [int64](#int64) |  | 附件类型ID，如果是文档类型，则为文档ID |
| type | [int32](#int32) |  | 附件类型，见 web/utils/enum.js |
| enable | [bool](#bool) |  | 是否启用 |
| path | [string](#string) |  | 附件路径 |
| name | [string](#string) |  | 附件名称 |
| size | [int64](#int64) |  | 附件大小，单位：字节 |
| width | [int64](#int64) |  | 附件宽度，单位：像素。针对图片附件 |
| height | [int64](#int64) |  | 附件高度，单位：像素。针对图片附件 |
| ext | [string](#string) |  | 扩展名，如：.docx |
| ip | [string](#string) |  | 上传IP地址 |
| username | [string](#string) |  | 用户名称 |
| type_name | [string](#string) |  | 附件类型名称 |
| description | [string](#string) |  | 附件描述、备注 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteAttachmentRequest"></a>

### DeleteAttachmentRequest
删除附件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetAttachmentRequest"></a>

### GetAttachmentRequest
获取附件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListAttachmentReply"></a>

### ListAttachmentReply
列出附件响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| attachment | [Attachment](#api-v1-Attachment) | repeated |  |






<a name="api-v1-ListAttachmentRequest"></a>

### ListAttachmentRequest
列出附件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| enable | [bool](#bool) | repeated | 是否启用 |
| user_id | [int64](#int64) | repeated | 用户ID |
| type | [int64](#int64) | repeated | 类型 |
| ext | [string](#string) |  | 扩展名 |





 

 

 


<a name="api-v1-AttachmentAPI"></a>

### AttachmentAPI
附件服务。只有管理员才有权限操作

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdateAttachment | [Attachment](#api-v1-Attachment) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新附件 |
| DeleteAttachment | [DeleteAttachmentRequest](#api-v1-DeleteAttachmentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除附件。这里只是软删除，不会真正删除附件，默认24小时候会真正清除附件 |
| GetAttachment | [GetAttachmentRequest](#api-v1-GetAttachmentRequest) | [Attachment](#api-v1-Attachment) | 查询附件 |
| ListAttachment | [ListAttachmentRequest](#api-v1-ListAttachmentRequest) | [ListAttachmentReply](#api-v1-ListAttachmentReply) | 列出附件 |

 



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

