# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/report.proto](#api_v1_report-proto)
    - [DeleteReportRequest](#api-v1-DeleteReportRequest)
    - [ListReportReply](#api-v1-ListReportReply)
    - [ListReportRequest](#api-v1-ListReportRequest)
    - [Report](#api-v1-Report)
  
    - [ReportAPI](#api-v1-ReportAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_report-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/report.proto



<a name="api-v1-DeleteReportRequest"></a>

### DeleteReportRequest
删除举报请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-ListReportReply"></a>

### ListReportReply
举报列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| report | [Report](#api-v1-Report) | repeated |  |






<a name="api-v1-ListReportRequest"></a>

### ListReportRequest
举报列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| status | [bool](#bool) | repeated |  |






<a name="api-v1-Report"></a>

### Report
举报


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 举报ID |
| document_id | [int64](#int64) |  | 文档ID |
| user_id | [int64](#int64) |  | 举报人ID |
| reason | [int32](#int32) |  | 举报原因 |
| status | [bool](#bool) |  | 举报处理状态 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 举报时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 处理时间 |
| document_title | [string](#string) |  | 文档标题 |
| remark | [string](#string) |  | 处理备注 |
| username | [string](#string) |  | 举报人 |





 

 

 


<a name="api-v1-ReportAPI"></a>

### ReportAPI
举报服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateReport | [Report](#api-v1-Report) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建举报 |
| UpdateReport | [Report](#api-v1-Report) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新举报，审核举报内容 |
| DeleteReport | [DeleteReportRequest](#api-v1-DeleteReportRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除举报 |
| ListReport | [ListReportRequest](#api-v1-ListReportRequest) | [ListReportReply](#api-v1-ListReportReply) | 获取举报列表 |

 



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

