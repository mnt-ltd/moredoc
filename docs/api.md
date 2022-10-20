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
  
- [api/v1/banner.proto](#api_v1_banner-proto)
    - [Banner](#api-v1-Banner)
    - [DeleteBannerRequest](#api-v1-DeleteBannerRequest)
    - [GetBannerRequest](#api-v1-GetBannerRequest)
    - [ListBannerReply](#api-v1-ListBannerReply)
    - [ListBannerRequest](#api-v1-ListBannerRequest)
  
    - [BannerAPI](#api-v1-BannerAPI)
  
- [api/v1/config.proto](#api_v1_config-proto)
    - [Config](#api-v1-Config)
    - [Configs](#api-v1-Configs)
    - [ListConfigRequest](#api-v1-ListConfigRequest)
  
    - [ConfigAPI](#api-v1-ConfigAPI)
  
- [api/v1/friendlink.proto](#api_v1_friendlink-proto)
    - [DeleteFriendlinkRequest](#api-v1-DeleteFriendlinkRequest)
    - [Friendlink](#api-v1-Friendlink)
    - [GetFriendlinkRequest](#api-v1-GetFriendlinkRequest)
    - [ListFriendlinkReply](#api-v1-ListFriendlinkReply)
    - [ListFriendlinkRequest](#api-v1-ListFriendlinkRequest)
  
    - [FriendlinkAPI](#api-v1-FriendlinkAPI)
  
- [api/v1/group.proto](#api_v1_group-proto)
    - [DeleteGroupRequest](#api-v1-DeleteGroupRequest)
    - [GetGroupPermissionRequest](#api-v1-GetGroupPermissionRequest)
    - [GetGroupRequest](#api-v1-GetGroupRequest)
    - [Group](#api-v1-Group)
    - [GroupPermissions](#api-v1-GroupPermissions)
    - [ListGroupReply](#api-v1-ListGroupReply)
    - [ListGroupRequest](#api-v1-ListGroupRequest)
    - [UpdateGroupPermissionRequest](#api-v1-UpdateGroupPermissionRequest)
  
    - [GroupAPI](#api-v1-GroupAPI)
  
- [api/v1/health.proto](#api_v1_health-proto)
    - [PingRequest](#-PingRequest)
    - [PongReply](#-PongReply)
  
    - [HealthAPI](#-HealthAPI)
  
- [api/v1/permission.proto](#api_v1_permission-proto)
    - [DeletePermissionRequest](#api-v1-DeletePermissionRequest)
    - [GetPermissionReply](#api-v1-GetPermissionReply)
    - [GetPermissionRequest](#api-v1-GetPermissionRequest)
    - [ListPermissionReply](#api-v1-ListPermissionReply)
    - [ListPermissionRequest](#api-v1-ListPermissionRequest)
    - [Permission](#api-v1-Permission)
  
    - [PermissionAPI](#api-v1-PermissionAPI)
  
- [api/v1/user.proto](#api_v1_user-proto)
    - [DeleteUserRequest](#api-v1-DeleteUserRequest)
    - [GetUserCaptchaReply](#api-v1-GetUserCaptchaReply)
    - [GetUserCaptchaRequest](#api-v1-GetUserCaptchaRequest)
    - [GetUserPermissionsReply](#api-v1-GetUserPermissionsReply)
    - [GetUserRequest](#api-v1-GetUserRequest)
    - [ListUserReply](#api-v1-ListUserReply)
    - [ListUserRequest](#api-v1-ListUserRequest)
    - [LoginReply](#api-v1-LoginReply)
    - [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest)
    - [UpdateUserPasswordRequest](#api-v1-UpdateUserPasswordRequest)
    - [User](#api-v1-User)
  
    - [UserAPI](#api-v1-UserAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_attachment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/attachment.proto



<a name="api-v1-Attachment"></a>

### Attachment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| hash | [string](#string) |  |  |
| user_id | [int64](#int64) |  |  |
| type_id | [int64](#int64) |  |  |
| type | [int32](#int32) |  |  |
| enable | [bool](#bool) |  |  |
| path | [string](#string) |  |  |
| name | [string](#string) |  |  |
| size | [int64](#int64) |  |  |
| width | [int64](#int64) |  |  |
| height | [int64](#int64) |  |  |
| ext | [string](#string) |  |  |
| ip | [string](#string) |  |  |
| username | [string](#string) |  | 用户名称 |
| type_name | [string](#string) |  | 附件类型名称 |
| description | [string](#string) |  | 附件描述、备注 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-DeleteAttachmentRequest"></a>

### DeleteAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetAttachmentRequest"></a>

### GetAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListAttachmentReply"></a>

### ListAttachmentReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| attachment | [Attachment](#api-v1-Attachment) | repeated |  |






<a name="api-v1-ListAttachmentRequest"></a>

### ListAttachmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  | 搜索关键字 |
| enable | [bool](#bool) | repeated |  |
| user_id | [int64](#int64) | repeated | 用户ID |
| type | [int64](#int64) | repeated | 类型 |
| ext | [string](#string) |  | 扩展名 |





 

 

 


<a name="api-v1-AttachmentAPI"></a>

### AttachmentAPI
附件服务。只有管理员才有权限操作

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdateAttachment | [Attachment](#api-v1-Attachment) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteAttachment | [DeleteAttachmentRequest](#api-v1-DeleteAttachmentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetAttachment | [GetAttachmentRequest](#api-v1-GetAttachmentRequest) | [Attachment](#api-v1-Attachment) |  |
| ListAttachment | [ListAttachmentRequest](#api-v1-ListAttachmentRequest) | [ListAttachmentReply](#api-v1-ListAttachmentReply) |  |

 



<a name="api_v1_banner-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/banner.proto



<a name="api-v1-Banner"></a>

### Banner



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| path | [string](#string) |  |  |
| sort | [int32](#int32) |  |  |
| enable | [bool](#bool) |  |  |
| type | [int32](#int32) |  |  |
| url | [string](#string) |  |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-DeleteBannerRequest"></a>

### DeleteBannerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetBannerRequest"></a>

### GetBannerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListBannerReply"></a>

### ListBannerReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| banner | [Banner](#api-v1-Banner) | repeated |  |






<a name="api-v1-ListBannerRequest"></a>

### ListBannerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| type | [int32](#int32) | repeated |  |
| enable | [bool](#bool) | repeated |  |
| wd | [string](#string) |  |  |





 

 

 


<a name="api-v1-BannerAPI"></a>

### BannerAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBanner | [Banner](#api-v1-Banner) | [Banner](#api-v1-Banner) |  |
| UpdateBanner | [Banner](#api-v1-Banner) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteBanner | [DeleteBannerRequest](#api-v1-DeleteBannerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetBanner | [GetBannerRequest](#api-v1-GetBannerRequest) | [Banner](#api-v1-Banner) |  |
| ListBanner | [ListBannerRequest](#api-v1-ListBannerRequest) | [ListBannerReply](#api-v1-ListBannerReply) |  |

 



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






<a name="api-v1-Configs"></a>

### Configs



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#api-v1-Config) | repeated |  |






<a name="api-v1-ListConfigRequest"></a>

### ListConfigRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) | repeated |  |





 

 

 


<a name="api-v1-ConfigAPI"></a>

### ConfigAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdateConfig | [Configs](#api-v1-Configs) | [.google.protobuf.Empty](#google-protobuf-Empty) | UpdateConfig 更新配置 |
| ListConfig | [ListConfigRequest](#api-v1-ListConfigRequest) | [Configs](#api-v1-Configs) | ListConfig 查询配置项 |

 



<a name="api_v1_friendlink-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/friendlink.proto



<a name="api-v1-DeleteFriendlinkRequest"></a>

### DeleteFriendlinkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Friendlink"></a>

### Friendlink



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| link | [string](#string) |  |  |
| description | [string](#string) |  |  |
| sort | [int32](#int32) |  |  |
| enable | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-GetFriendlinkRequest"></a>

### GetFriendlinkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListFriendlinkReply"></a>

### ListFriendlinkReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| friendlink | [Friendlink](#api-v1-Friendlink) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="api-v1-ListFriendlinkRequest"></a>

### ListFriendlinkRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| wd | [string](#string) |  |  |
| enable | [bool](#bool) | repeated |  |





 

 

 


<a name="api-v1-FriendlinkAPI"></a>

### FriendlinkAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateFriendlink | [Friendlink](#api-v1-Friendlink) | [Friendlink](#api-v1-Friendlink) |  |
| UpdateFriendlink | [Friendlink](#api-v1-Friendlink) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteFriendlink | [DeleteFriendlinkRequest](#api-v1-DeleteFriendlinkRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetFriendlink | [GetFriendlinkRequest](#api-v1-GetFriendlinkRequest) | [Friendlink](#api-v1-Friendlink) |  |
| ListFriendlink | [ListFriendlinkRequest](#api-v1-ListFriendlinkRequest) | [ListFriendlinkReply](#api-v1-ListFriendlinkReply) |  |

 



<a name="api_v1_group-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/group.proto



<a name="api-v1-DeleteGroupRequest"></a>

### DeleteGroupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetGroupPermissionRequest"></a>

### GetGroupPermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-GetGroupRequest"></a>

### GetGroupRequest
根据组名或者ID获取用户组


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |






<a name="api-v1-Group"></a>

### Group



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| color | [string](#string) |  |  |
| is_default | [bool](#bool) |  |  |
| is_display | [bool](#bool) |  |  |
| description | [string](#string) |  |  |
| user_count | [int32](#int32) |  |  |
| sort | [int32](#int32) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-GroupPermissions"></a>

### GroupPermissions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission_id | [int64](#int64) | repeated |  |






<a name="api-v1-ListGroupReply"></a>

### ListGroupReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [Group](#api-v1-Group) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="api-v1-ListGroupRequest"></a>

### ListGroupRequest
查询用户组列表。不需要分页，直接返回全部用户组，只是可以指定查询哪些字段


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wd | [string](#string) |  |  |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| sort | [string](#string) |  |  |
| field | [string](#string) | repeated |  |






<a name="api-v1-UpdateGroupPermissionRequest"></a>

### UpdateGroupPermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int64](#int64) |  |  |
| permission_id | [int64](#int64) | repeated |  |





 

 

 


<a name="api-v1-GroupAPI"></a>

### GroupAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateGroup | [Group](#api-v1-Group) | [Group](#api-v1-Group) | 创建用户组 |
| UpdateGroup | [Group](#api-v1-Group) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户组 |
| DeleteGroup | [DeleteGroupRequest](#api-v1-DeleteGroupRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除用户组 |
| GetGroup | [GetGroupRequest](#api-v1-GetGroupRequest) | [Group](#api-v1-Group) | 获取用户组列表 |
| ListGroup | [ListGroupRequest](#api-v1-ListGroupRequest) | [ListGroupReply](#api-v1-ListGroupReply) |  |
| GetGroupPermission | [GetGroupPermissionRequest](#api-v1-GetGroupPermissionRequest) | [GroupPermissions](#api-v1-GroupPermissions) | 获取用户组权限列表 |
| UpdateGroupPermission | [UpdateGroupPermissionRequest](#api-v1-UpdateGroupPermissionRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户组权限，给用户组设置权限 |

 



<a name="api_v1_health-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/health.proto



<a name="-PingRequest"></a>

### PingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="-PongReply"></a>

### PongReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 


<a name="-HealthAPI"></a>

### HealthAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Health | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| Ping | [.PingRequest](#PingRequest) | [.PongReply](#PongReply) |  |

 



<a name="api_v1_permission-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/permission.proto



<a name="api-v1-DeletePermissionRequest"></a>

### DeletePermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetPermissionReply"></a>

### GetPermissionReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [Permission](#api-v1-Permission) |  |  |






<a name="api-v1-GetPermissionRequest"></a>

### GetPermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListPermissionReply"></a>

### ListPermissionReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| permission | [Permission](#api-v1-Permission) | repeated |  |






<a name="api-v1-ListPermissionRequest"></a>

### ListPermissionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| method | [string](#string) | repeated |  |
| path | [string](#string) |  |  |






<a name="api-v1-Permission"></a>

### Permission



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| method | [string](#string) |  |  |
| path | [string](#string) |  |  |
| title | [string](#string) |  |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 


<a name="api-v1-PermissionAPI"></a>

### PermissionAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdatePermission | [Permission](#api-v1-Permission) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetPermission | [GetPermissionRequest](#api-v1-GetPermissionRequest) | [Permission](#api-v1-Permission) |  |
| ListPermission | [ListPermissionRequest](#api-v1-ListPermissionRequest) | [ListPermissionReply](#api-v1-ListPermissionReply) |  |

 



<a name="api_v1_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/user.proto



<a name="api-v1-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetUserCaptchaReply"></a>

### GetUserCaptchaReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enable | [bool](#bool) |  |  |
| id | [string](#string) |  |  |
| captcha | [string](#string) |  |  |
| type | [string](#string) |  |  |






<a name="api-v1-GetUserCaptchaRequest"></a>

### GetUserCaptchaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | 验证码类型：register、login、comment、find_password、upload |






<a name="api-v1-GetUserPermissionsReply"></a>

### GetUserPermissionsReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [Permission](#api-v1-Permission) | repeated |  |






<a name="api-v1-GetUserRequest"></a>

### GetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListUserReply"></a>

### ListUserReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| user | [User](#api-v1-User) | repeated |  |






<a name="api-v1-ListUserRequest"></a>

### ListUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| sort | [string](#string) |  |  |
| id | [int64](#int64) | repeated |  |
| group_id | [int64](#int64) | repeated |  |
| status | [int32](#int32) | repeated |  |






<a name="api-v1-LoginReply"></a>

### LoginReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| user | [User](#api-v1-User) |  |  |






<a name="api-v1-RegisterAndLoginRequest"></a>

### RegisterAndLoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| captcha | [string](#string) |  |  |
| captcha_id | [string](#string) |  |  |






<a name="api-v1-UpdateUserPasswordRequest"></a>

### UpdateUserPasswordRequest
修改用户密码


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| old_password | [string](#string) |  |  |
| new_password | [string](#string) |  |  |






<a name="api-v1-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| id | [int64](#int64) |  |  |
| username | [string](#string) |  |  |
| mobile | [string](#string) |  |  |
| email | [string](#string) |  |  |
| address | [string](#string) |  |  |
| signature | [string](#string) |  |  |
| last_login_ip | [string](#string) |  |  |
| register_ip | [string](#string) |  |  |
| doc_count | [int32](#int32) |  |  |
| follow_count | [int32](#int32) |  |  |
| fans_count | [int32](#int32) |  |  |
| favorite_count | [int32](#int32) |  |  |
| comment_count | [int32](#int32) |  |  |
| status | [int32](#int32) |  |  |
| avatar | [string](#string) |  |  |
| identity | [string](#string) |  |  |
| realname | [string](#string) |  |  |
| group_id | [int64](#int64) | repeated |  |





 

 

 


<a name="api-v1-UserAPI"></a>

### UserAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 用户注册 |
| Login | [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest) | [LoginReply](#api-v1-LoginReply) | 用户登录 |
| Logout | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 退出登录 |
| GetUser | [GetUserRequest](#api-v1-GetUserRequest) | [User](#api-v1-User) | 查询用户信息。如果传递了Id参数，则表示查询用户的公开信息，否则查询当前用户的私有信息 |
| UpdateUserPassword | [UpdateUserPasswordRequest](#api-v1-UpdateUserPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户密码。如果不传用户ID，则表示更新当前用户的密码； 如果穿了用户ID，则表示更新指定用户的密码，这时需要验证当前用户的权限 |
| UpdateUser | [User](#api-v1-User) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户密码。如果不传用户ID，则表示更新当前用户的密码； 如果穿了用户ID，则表示更新指定用户的密码，这时需要验证当前用户的权限 |
| DeleteUser | [DeleteUserRequest](#api-v1-DeleteUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除用户。需要验证用户权限 |
| ListUser | [ListUserRequest](#api-v1-ListUserRequest) | [ListUserReply](#api-v1-ListUserReply) | 查询用户列表。对于非管理员，返回相应用户的公开信息； 对于管理员，返回相应用户的绝大部分信息 |
| GetUserCaptcha | [GetUserCaptchaRequest](#api-v1-GetUserCaptchaRequest) | [GetUserCaptchaReply](#api-v1-GetUserCaptchaReply) | GetUserCaptcha 获取用户验证码 |
| GetUserPermissions | [.google.protobuf.Empty](#google-protobuf-Empty) | [GetUserPermissionsReply](#api-v1-GetUserPermissionsReply) | GetUserCaptcha 获取用户验证码 |

 



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

