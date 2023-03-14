# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/user.proto](#api_v1_user-proto)
    - [DeleteUserRequest](#api-v1-DeleteUserRequest)
    - [Dynamic](#api-v1-Dynamic)
    - [FindPasswordRequest](#api-v1-FindPasswordRequest)
    - [GetUserCaptchaReply](#api-v1-GetUserCaptchaReply)
    - [GetUserCaptchaRequest](#api-v1-GetUserCaptchaRequest)
    - [GetUserPermissionsReply](#api-v1-GetUserPermissionsReply)
    - [GetUserRequest](#api-v1-GetUserRequest)
    - [ListUserDynamicReply](#api-v1-ListUserDynamicReply)
    - [ListUserDynamicRequest](#api-v1-ListUserDynamicRequest)
    - [ListUserReply](#api-v1-ListUserReply)
    - [ListUserRequest](#api-v1-ListUserRequest)
    - [LoginReply](#api-v1-LoginReply)
    - [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest)
    - [SetUserRequest](#api-v1-SetUserRequest)
    - [Sign](#api-v1-Sign)
    - [UpdateUserPasswordRequest](#api-v1-UpdateUserPasswordRequest)
    - [User](#api-v1-User)
  
    - [UserAPI](#api-v1-UserAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/user.proto



<a name="api-v1-DeleteUserRequest"></a>

### DeleteUserRequest
删除用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Dynamic"></a>

### Dynamic
用户动态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 动态ID |
| user_id | [int64](#int64) |  | 用户ID |
| content | [string](#string) |  | 内容 |
| type | [int32](#int32) |  | 类型 |
| username | [string](#string) |  | 用户名 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-FindPasswordRequest"></a>

### FindPasswordRequest
找回密码


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | 邮箱 |
| token | [string](#string) |  | 签名token |
| password | [string](#string) |  | 新密码 |
| captcha | [string](#string) |  | 验证码 |
| captcha_id | [string](#string) |  | 验证码ID |






<a name="api-v1-GetUserCaptchaReply"></a>

### GetUserCaptchaReply
验证码响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enable | [bool](#bool) |  | 是否启用验证码 |
| id | [string](#string) |  | 验证码ID |
| captcha | [string](#string) |  | 验证码 |
| type | [string](#string) |  | 验证码类型 |






<a name="api-v1-GetUserCaptchaRequest"></a>

### GetUserCaptchaRequest
查询验证码请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | 验证码类型：register、login、comment、find_password、upload |






<a name="api-v1-GetUserPermissionsReply"></a>

### GetUserPermissionsReply
用户权限信息查询


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [Permission](#api-v1-Permission) | repeated |  |






<a name="api-v1-GetUserRequest"></a>

### GetUserRequest
获取用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListUserDynamicReply"></a>

### ListUserDynamicReply
用户动态列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| dynamic | [Dynamic](#api-v1-Dynamic) | repeated | 动态列表 |






<a name="api-v1-ListUserDynamicRequest"></a>

### ListUserDynamicRequest
用户动态列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| id | [int64](#int64) |  | 用户ID |






<a name="api-v1-ListUserReply"></a>

### ListUserReply
用户列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| user | [User](#api-v1-User) | repeated | 用户列表 |






<a name="api-v1-ListUserRequest"></a>

### ListUserRequest
用户列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键词 |
| sort | [string](#string) |  | 排序字段 |
| id | [int64](#int64) | repeated | 用户ID |
| group_id | [int64](#int64) | repeated | 用户组ID |
| status | [int32](#int32) | repeated | 用户状态 |
| limit | [int64](#int64) |  | 请求数量限制，大于0时，page和size无效 |






<a name="api-v1-LoginReply"></a>

### LoginReply
用户登录响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| user | [User](#api-v1-User) |  |  |






<a name="api-v1-RegisterAndLoginRequest"></a>

### RegisterAndLoginRequest
用户注册登录请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| captcha | [string](#string) |  | 验证码 |
| captcha_id | [string](#string) |  | 验证码ID |
| email | [string](#string) |  | 邮箱 |






<a name="api-v1-SetUserRequest"></a>

### SetUserRequest
管理后台设置用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户ID |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| group_id | [int64](#int64) | repeated | 用户组ID |
| email | [string](#string) |  | 邮箱 |






<a name="api-v1-Sign"></a>

### Sign
用户签到


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 签到ID |
| user_id | [int64](#int64) |  | 用户ID |
| sign_at | [int32](#int32) |  | 签到日期 |
| ip | [string](#string) |  | 签到IP |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| award | [int32](#int32) |  | 签到积分奖励 |






<a name="api-v1-UpdateUserPasswordRequest"></a>

### UpdateUserPasswordRequest
修改用户密码


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户ID |
| old_password | [string](#string) |  | 旧密码 |
| new_password | [string](#string) |  | 新密码 |






<a name="api-v1-User"></a>

### User
用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 最后登录时间 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 注册时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| id | [int64](#int64) |  | 用户ID |
| username | [string](#string) |  | 用户名 |
| mobile | [string](#string) |  | 手机号 |
| email | [string](#string) |  | 邮箱，唯一 |
| address | [string](#string) |  | 地址 |
| signature | [string](#string) |  | 个性签名 |
| last_login_ip | [string](#string) |  | 最后登录IP |
| register_ip | [string](#string) |  | 注册IP |
| doc_count | [int32](#int32) |  | 文档数量 |
| follow_count | [int32](#int32) |  | 关注数量 |
| fans_count | [int32](#int32) |  | 粉丝数量 |
| favorite_count | [int32](#int32) |  | 收藏数量 |
| comment_count | [int32](#int32) |  | 评论数量 |
| status | [int32](#int32) |  | 用户状态，见 web/utils/enum.js，当前没有使用 |
| avatar | [string](#string) |  | 头像 |
| identity | [string](#string) |  | 身份证 |
| realname | [string](#string) |  | 真实姓名 |
| group_id | [int64](#int64) | repeated | 用户组ID |
| credit_count | [int32](#int32) |  | 积分 |





 

 

 


<a name="api-v1-UserAPI"></a>

### UserAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest) | [LoginReply](#api-v1-LoginReply) | 用户注册 |
| Login | [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest) | [LoginReply](#api-v1-LoginReply) | 用户登录 |
| Logout | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 退出登录 |
| GetUser | [GetUserRequest](#api-v1-GetUserRequest) | [User](#api-v1-User) | 查询用户信息。如果传递了Id参数，则表示查询用户的公开信息，否则查询当前用户的私有信息 |
| UpdateUserPassword | [UpdateUserPasswordRequest](#api-v1-UpdateUserPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户密码。如果不传用户ID，则表示更新当前用户的密码； 如果穿了用户ID，则表示更新指定用户的密码，这时需要验证当前用户的权限 |
| UpdateUserProfile | [User](#api-v1-User) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户密码。如果不传用户ID，则表示更新当前用户的密码； 如果穿了用户ID，则表示更新指定用户的密码，这时需要验证当前用户的权限 |
| DeleteUser | [DeleteUserRequest](#api-v1-DeleteUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除用户。需要验证用户权限 |
| AddUser | [SetUserRequest](#api-v1-SetUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 新增用户 |
| SetUser | [SetUserRequest](#api-v1-SetUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置用户 |
| ListUser | [ListUserRequest](#api-v1-ListUserRequest) | [ListUserReply](#api-v1-ListUserReply) | 查询用户列表。对于非管理员，返回相应用户的公开信息； 对于管理员，返回相应用户的绝大部分信息 |
| GetUserCaptcha | [GetUserCaptchaRequest](#api-v1-GetUserCaptchaRequest) | [GetUserCaptchaReply](#api-v1-GetUserCaptchaReply) | GetUserCaptcha 获取用户验证码 |
| GetUserPermissions | [.google.protobuf.Empty](#google-protobuf-Empty) | [GetUserPermissionsReply](#api-v1-GetUserPermissionsReply) | GetUserCaptcha 获取用户验证码 |
| CanIUploadDocument | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 用户是否可以上传文档 |
| ListUserDynamic | [ListUserDynamicRequest](#api-v1-ListUserDynamicRequest) | [ListUserDynamicReply](#api-v1-ListUserDynamicReply) | 获取用户动态，包括获取关注的用户的动态 |
| SignToday | [.google.protobuf.Empty](#google-protobuf-Empty) | [Sign](#api-v1-Sign) | 每日签到 |
| GetSignedToday | [.google.protobuf.Empty](#google-protobuf-Empty) | [Sign](#api-v1-Sign) | 获取今日已签到记录 |
| FindPasswordStepOne | [FindPasswordRequest](#api-v1-FindPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 找回密码：第一步，发送验证码 |
| FindPasswordStepTwo | [FindPasswordRequest](#api-v1-FindPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 找回密码：第二步，修改密码 |

 



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

