# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/permission.proto](#api_v1_permission-proto)
    - [GetPermissionReply](#api-v1-GetPermissionReply)
    - [GetPermissionRequest](#api-v1-GetPermissionRequest)
    - [ListPermissionReply](#api-v1-ListPermissionReply)
    - [ListPermissionRequest](#api-v1-ListPermissionRequest)
    - [Permission](#api-v1-Permission)
  
    - [PermissionAPI](#api-v1-PermissionAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_permission-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/permission.proto



<a name="api-v1-GetPermissionReply"></a>

### GetPermissionReply
权限响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [Permission](#api-v1-Permission) |  |  |






<a name="api-v1-GetPermissionRequest"></a>

### GetPermissionRequest
权限请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListPermissionReply"></a>

### ListPermissionReply
权限列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| permission | [Permission](#api-v1-Permission) | repeated |  |






<a name="api-v1-ListPermissionRequest"></a>

### ListPermissionRequest
权限列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| method | [string](#string) | repeated |  |
| path | [string](#string) |  |  |






<a name="api-v1-Permission"></a>

### Permission
权限


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 权限ID |
| method | [string](#string) |  | 请求方法 |
| path | [string](#string) |  | 请求路径 |
| title | [string](#string) |  | 权限名称 |
| description | [string](#string) |  | 权限描述 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |





 

 

 


<a name="api-v1-PermissionAPI"></a>

### PermissionAPI
权限API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdatePermission | [Permission](#api-v1-Permission) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新权限信息。这里只能操作title和description |
| GetPermission | [GetPermissionRequest](#api-v1-GetPermissionRequest) | [Permission](#api-v1-Permission) | 查询权限信息 |
| ListPermission | [ListPermissionRequest](#api-v1-ListPermissionRequest) | [ListPermissionReply](#api-v1-ListPermissionReply) | 查询权限列表 |

 



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

