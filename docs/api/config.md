# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/config.proto](#api_v1_config-proto)
    - [Config](#api-v1-Config)
    - [ConfigCaptcha](#api-v1-ConfigCaptcha)
    - [ConfigFooter](#api-v1-ConfigFooter)
    - [ConfigSecurity](#api-v1-ConfigSecurity)
    - [ConfigSystem](#api-v1-ConfigSystem)
    - [Configs](#api-v1-Configs)
    - [EnvDependent](#api-v1-EnvDependent)
    - [Envs](#api-v1-Envs)
    - [ListConfigRequest](#api-v1-ListConfigRequest)
    - [Settings](#api-v1-Settings)
    - [Stats](#api-v1-Stats)
  
    - [ConfigAPI](#api-v1-ConfigAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/config.proto



<a name="api-v1-Config"></a>

### Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| label | [string](#string) |  |  |
| name | [string](#string) |  |  |
| value | [string](#string) |  |  |
| placeholder | [string](#string) |  |  |
| input_type | [string](#string) |  |  |
| category | [string](#string) |  |  |
| sort | [int32](#int32) |  |  |
| options | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-ConfigCaptcha"></a>

### ConfigCaptcha
验证码配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| length | [int32](#int32) |  |  |
| width | [int32](#int32) |  |  |
| height | [int32](#int32) |  |  |
| type | [string](#string) |  |  |






<a name="api-v1-ConfigFooter"></a>

### ConfigFooter
底链配置项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| about | [string](#string) |  |  |
| contact | [string](#string) |  |  |
| agreement | [string](#string) |  |  |
| copyright | [string](#string) |  |  |
| feedback | [string](#string) |  |  |






<a name="api-v1-ConfigSecurity"></a>

### ConfigSecurity
安全配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_close | [bool](#bool) |  |  |
| close_statement | [string](#string) |  |  |
| enable_register | [bool](#bool) |  |  |
| enable_captcha_login | [bool](#bool) |  |  |
| enable_captcha_register | [bool](#bool) |  |  |
| enable_captcha_comment | [bool](#bool) |  |  |
| enable_captcha_find_password | [bool](#bool) |  |  |
| enable_captcha_upload | [bool](#bool) |  |  |
| max_document_size | [int32](#int32) |  |  |
| document_allowed_ext | [string](#string) | repeated |  |
| login_required | [bool](#bool) |  | 是否登录才能访问 |






<a name="api-v1-ConfigSystem"></a>

### ConfigSystem
系统配置项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  |  |
| title | [string](#string) |  |  |
| keywords | [string](#string) |  |  |
| description | [string](#string) |  |  |
| logo | [string](#string) |  |  |
| favicon | [string](#string) |  |  |
| icp | [string](#string) |  |  |
| analytics | [string](#string) |  |  |
| sitename | [string](#string) |  |  |
| copyright_start_year | [string](#string) |  |  |
| register_background | [string](#string) |  |  |
| login_background | [string](#string) |  |  |
| recommend_words | [string](#string) | repeated |  |
| version | [string](#string) |  | 程序版本号 |






<a name="api-v1-Configs"></a>

### Configs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#api-v1-Config) | repeated |  |






<a name="api-v1-EnvDependent"></a>

### EnvDependent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 依赖名称 |
| description | [string](#string) |  | 依赖描述 |
| is_installed | [bool](#bool) |  | 是否已安装 |
| error | [string](#string) |  | 错误信息 |
| checked_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 检测时间 |
| cmd | [string](#string) |  | 检测命令 |
| is_required | [bool](#bool) |  | 是否必须 |






<a name="api-v1-Envs"></a>

### Envs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| envs | [EnvDependent](#api-v1-EnvDependent) | repeated |  |






<a name="api-v1-ListConfigRequest"></a>

### ListConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) | repeated |  |






<a name="api-v1-Settings"></a>

### Settings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system | [ConfigSystem](#api-v1-ConfigSystem) |  |  |
| footer | [ConfigFooter](#api-v1-ConfigFooter) |  |  |
| security | [ConfigSecurity](#api-v1-ConfigSecurity) |  | ConfigCaptcha captcha = 4; |






<a name="api-v1-Stats"></a>

### Stats



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_count | [int64](#int64) |  |  |
| document_count | [int64](#int64) |  |  |
| category_count | [int64](#int64) |  |  |
| article_count | [int64](#int64) |  |  |
| comment_count | [int64](#int64) |  |  |
| banner_count | [int64](#int64) |  |  |
| friendlink_count | [int64](#int64) |  |  |
| os | [string](#string) |  |  |
| version | [string](#string) |  |  |
| hash | [string](#string) |  |  |
| build_at | [string](#string) |  |  |
| report_count | [int64](#int64) |  |  |





 

 

 


<a name="api-v1-ConfigAPI"></a>

### ConfigAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSettings | [.google.protobuf.Empty](#google-protobuf-Empty) | [Settings](#api-v1-Settings) | 获取系统配置（针对所有用户，只读） |
| UpdateConfig | [Configs](#api-v1-Configs) | [.google.protobuf.Empty](#google-protobuf-Empty) | UpdateConfig 更新配置 |
| ListConfig | [ListConfigRequest](#api-v1-ListConfigRequest) | [Configs](#api-v1-Configs) | ListConfig 查询配置项 |
| GetStats | [.google.protobuf.Empty](#google-protobuf-Empty) | [Stats](#api-v1-Stats) | 获取系统配置 |
| GetEnvs | [.google.protobuf.Empty](#google-protobuf-Empty) | [Envs](#api-v1-Envs) | 获取系统环境依赖检测 |
| UpdateSitemap | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新站点地图 |

 



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

