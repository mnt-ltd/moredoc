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
配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 配置ID |
| label | [string](#string) |  | 配置标签 |
| name | [string](#string) |  | 配置名称 |
| value | [string](#string) |  | 配置值 |
| placeholder | [string](#string) |  | 配置占位符 |
| input_type | [string](#string) |  | 输入类型，如：textarea、number、switch等 |
| category | [string](#string) |  | 配置分类，如：system、footer、security等，见 web/utils/enum.js |
| sort | [int32](#int32) |  | 排序，越小越靠前 |
| options | [string](#string) |  | 配置项枚举，一个一行，如select的option选项，用 key=value 的形式 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-ConfigCaptcha"></a>

### ConfigCaptcha
验证码配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| length | [int32](#int32) |  | 验证码长度 |
| width | [int32](#int32) |  | 验证码宽度 |
| height | [int32](#int32) |  | 验证码高度 |
| type | [string](#string) |  | 验证码类型，见 web/utils/enum.js |






<a name="api-v1-ConfigFooter"></a>

### ConfigFooter
底链配置项，为跳转的链接地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| about | [string](#string) |  | 关于我们 |
| contact | [string](#string) |  | 联系我们 |
| agreement | [string](#string) |  | 用户协议、文库协议 |
| copyright | [string](#string) |  | 版权声明 |
| feedback | [string](#string) |  | 意见和建议反馈 |






<a name="api-v1-ConfigSecurity"></a>

### ConfigSecurity
安全配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_close | [bool](#bool) |  | 是否关闭站点 |
| close_statement | [string](#string) |  | 关闭站点的说明，支持HTML |
| enable_register | [bool](#bool) |  | 是否开放注册 |
| enable_captcha_login | [bool](#bool) |  | 是否开启登录验证码 |
| enable_captcha_register | [bool](#bool) |  | 是否开启注册验证码 |
| enable_captcha_comment | [bool](#bool) |  | 是否开启评论验证码 |
| enable_captcha_find_password | [bool](#bool) |  | 是否开启找回密码验证码 |
| enable_captcha_upload | [bool](#bool) |  | 是否开启上传验证码 |
| max_document_size | [int32](#int32) |  | 文档最大大小 |
| document_allowed_ext | [string](#string) | repeated | 文档允许的扩展名 |
| login_required | [bool](#bool) |  | 是否登录才能访问 |






<a name="api-v1-ConfigSystem"></a>

### ConfigSystem
系统配置项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  | 站点域名，如： https://moredoc.mnt.ltd |
| title | [string](#string) |  | 站点标题，首页显示 |
| keywords | [string](#string) |  | 站点关键词，SEO用 |
| description | [string](#string) |  | 站点描述，SEO用 |
| logo | [string](#string) |  | 站点logo |
| favicon | [string](#string) |  | 站点favicon |
| icp | [string](#string) |  | 站点备案号 |
| analytics | [string](#string) |  | 站点统计代码，目前只支持百度统计 |
| sitename | [string](#string) |  | 站点名称 |
| copyright_start_year | [string](#string) |  | 站点版权起始年份，如：2018，则底部显示 2018 - 2023 |
| register_background | [string](#string) |  | 注册页背景图 |
| login_background | [string](#string) |  | 登录页背景图 |
| recommend_words | [string](#string) | repeated | 推荐搜索词，首页展示 |
| version | [string](#string) |  | 程序版本号 |






<a name="api-v1-Configs"></a>

### Configs
配置列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#api-v1-Config) | repeated |  |






<a name="api-v1-EnvDependent"></a>

### EnvDependent
依赖项


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
依赖项列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| envs | [EnvDependent](#api-v1-EnvDependent) | repeated |  |






<a name="api-v1-ListConfigRequest"></a>

### ListConfigRequest
查询配置项请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) | repeated |  |






<a name="api-v1-Settings"></a>

### Settings
系统配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system | [ConfigSystem](#api-v1-ConfigSystem) |  | 系统配置 |
| footer | [ConfigFooter](#api-v1-ConfigFooter) |  | 底链配置 |
| security | [ConfigSecurity](#api-v1-ConfigSecurity) |  | 安全配置 |






<a name="api-v1-Stats"></a>

### Stats
系统状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_count | [int64](#int64) |  | 用户数量 |
| document_count | [int64](#int64) |  | 文档数量 |
| category_count | [int64](#int64) |  | 分类数量 |
| article_count | [int64](#int64) |  | 文章数量 |
| comment_count | [int64](#int64) |  | 评论数量 |
| banner_count | [int64](#int64) |  | banner数量 |
| friendlink_count | [int64](#int64) |  | 友情链接数量 |
| os | [string](#string) |  | 操作系统 |
| version | [string](#string) |  | 程序版本号 |
| hash | [string](#string) |  | 程序构建时的 git hash |
| build_at | [string](#string) |  | 程序构建时间 |
| report_count | [int64](#int64) |  | 举报数量 |





 

 

 


<a name="api-v1-ConfigAPI"></a>

### ConfigAPI
配置服务

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

