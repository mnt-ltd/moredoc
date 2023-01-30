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
  
- [api/v1/category.proto](#api_v1_category-proto)
    - [Category](#api-v1-Category)
    - [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest)
    - [GetCategoryRequest](#api-v1-GetCategoryRequest)
    - [ListCategoryReply](#api-v1-ListCategoryReply)
    - [ListCategoryRequest](#api-v1-ListCategoryRequest)
  
    - [CategoryAPI](#api-v1-CategoryAPI)
  
- [api/v1/comment.proto](#api_v1_comment-proto)
    - [CheckCommentRequest](#api-v1-CheckCommentRequest)
    - [Comment](#api-v1-Comment)
    - [CreateCommentRequest](#api-v1-CreateCommentRequest)
    - [DeleteCommentRequest](#api-v1-DeleteCommentRequest)
    - [GetCommentRequest](#api-v1-GetCommentRequest)
    - [ListCommentReply](#api-v1-ListCommentReply)
    - [ListCommentRequest](#api-v1-ListCommentRequest)
  
    - [CommentAPI](#api-v1-CommentAPI)
  
- [api/v1/config.proto](#api_v1_config-proto)
    - [Config](#api-v1-Config)
    - [ConfigCaptcha](#api-v1-ConfigCaptcha)
    - [ConfigFooter](#api-v1-ConfigFooter)
    - [ConfigSecurity](#api-v1-ConfigSecurity)
    - [ConfigSystem](#api-v1-ConfigSystem)
    - [Configs](#api-v1-Configs)
    - [EnvDependent](#api-v1-EnvDependent)
    - [ListConfigRequest](#api-v1-ListConfigRequest)
    - [Settings](#api-v1-Settings)
    - [Stats](#api-v1-Stats)
  
    - [ConfigAPI](#api-v1-ConfigAPI)
  
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
  
- [api/v1/favorite.proto](#api_v1_favorite-proto)
    - [DeleteFavoriteRequest](#api-v1-DeleteFavoriteRequest)
    - [Favorite](#api-v1-Favorite)
    - [GetFavoriteRequest](#api-v1-GetFavoriteRequest)
    - [ListFavoriteReply](#api-v1-ListFavoriteReply)
    - [ListFavoriteRequest](#api-v1-ListFavoriteRequest)
  
    - [FavoriteAPI](#api-v1-FavoriteAPI)
  
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
  
- [api/v1/report.proto](#api_v1_report-proto)
    - [DeleteReportRequest](#api-v1-DeleteReportRequest)
    - [ListReportReply](#api-v1-ListReportReply)
    - [ListReportRequest](#api-v1-ListReportRequest)
    - [Report](#api-v1-Report)
  
    - [ReportAPI](#api-v1-ReportAPI)
  
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



<a name="api_v1_article-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/article.proto



<a name="api-v1-Article"></a>

### Article



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| identifier | [string](#string) |  |  |
| author | [string](#string) |  |  |
| view_count | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| keywords | [string](#string) |  |  |
| description | [string](#string) |  |  |
| content | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-DeleteArticleRequest"></a>

### DeleteArticleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetArticleRequest"></a>

### GetArticleRequest
根据ID或者文章标识获取文章


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| identifier | [string](#string) |  |  |






<a name="api-v1-ListArticleReply"></a>

### ListArticleReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| article | [Article](#api-v1-Article) | repeated |  |






<a name="api-v1-ListArticleRequest"></a>

### ListArticleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |





 

 

 


<a name="api-v1-ArticleAPI"></a>

### ArticleAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| UpdateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteArticle | [DeleteArticleRequest](#api-v1-DeleteArticleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetArticle | [GetArticleRequest](#api-v1-GetArticleRequest) | [Article](#api-v1-Article) |  |
| ListArticle | [ListArticleRequest](#api-v1-ListArticleRequest) | [ListArticleReply](#api-v1-ListArticleReply) |  |

 



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
| field | [string](#string) | repeated |  |





 

 

 


<a name="api-v1-BannerAPI"></a>

### BannerAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBanner | [Banner](#api-v1-Banner) | [Banner](#api-v1-Banner) |  |
| UpdateBanner | [Banner](#api-v1-Banner) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteBanner | [DeleteBannerRequest](#api-v1-DeleteBannerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetBanner | [GetBannerRequest](#api-v1-GetBannerRequest) | [Banner](#api-v1-Banner) |  |
| ListBanner | [ListBannerRequest](#api-v1-ListBannerRequest) | [ListBannerReply](#api-v1-ListBannerReply) |  |

 



<a name="api_v1_category-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/category.proto



<a name="api-v1-Category"></a>

### Category



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| parent_id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| doc_count | [int32](#int32) |  |  |
| sort | [int32](#int32) |  |  |
| enable | [bool](#bool) |  |  |
| cover | [string](#string) |  | 分类封面 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-DeleteCategoryRequest"></a>

### DeleteCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCategoryRequest"></a>

### GetCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCategoryReply"></a>

### ListCategoryReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| category | [Category](#api-v1-Category) | repeated |  |






<a name="api-v1-ListCategoryRequest"></a>

### ListCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| parent_id | [int64](#int64) | repeated |  |
| wd | [string](#string) |  |  |
| enable | [bool](#bool) | repeated |  |
| field | [string](#string) | repeated |  |





 

 

 


<a name="api-v1-CategoryAPI"></a>

### CategoryAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| UpdateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteCategory | [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetCategory | [GetCategoryRequest](#api-v1-GetCategoryRequest) | [Category](#api-v1-Category) |  |
| ListCategory | [ListCategoryRequest](#api-v1-ListCategoryRequest) | [ListCategoryReply](#api-v1-ListCategoryReply) |  |

 



<a name="api_v1_comment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/comment.proto



<a name="api-v1-CheckCommentRequest"></a>

### CheckCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |
| status | [int32](#int32) |  |  |






<a name="api-v1-Comment"></a>

### Comment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| id | [int64](#int64) |  |  |
| parent_id | [int64](#int64) |  |  |
| content | [string](#string) |  |  |
| document_id | [int64](#int64) |  |  |
| status | [int32](#int32) |  |  |
| comment_count | [int32](#int32) |  |  |
| user_id | [int64](#int64) |  |  |
| user | [User](#api-v1-User) |  |  |
| document_title | [string](#string) |  |  |






<a name="api-v1-CreateCommentRequest"></a>

### CreateCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  |  |
| parent_id | [int64](#int64) |  |  |
| content | [string](#string) |  |  |
| captcha_id | [string](#string) |  |  |
| captcha | [string](#string) |  |  |






<a name="api-v1-DeleteCommentRequest"></a>

### DeleteCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCommentRequest"></a>

### GetCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCommentReply"></a>

### ListCommentReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| comment | [Comment](#api-v1-Comment) | repeated |  |






<a name="api-v1-ListCommentRequest"></a>

### ListCommentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| status | [int32](#int32) | repeated |  |
| document_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| parent_id | [int64](#int64) | repeated |  |
| with_document_title | [bool](#bool) |  |  |





 

 

 


<a name="api-v1-CommentAPI"></a>

### CommentAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateComment | [CreateCommentRequest](#api-v1-CreateCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| UpdateComment | [Comment](#api-v1-Comment) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新评论，仅限管理员操作 |
| DeleteComment | [DeleteCommentRequest](#api-v1-DeleteCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 管理员或用户自己删除自己的评论 |
| GetComment | [GetCommentRequest](#api-v1-GetCommentRequest) | [Comment](#api-v1-Comment) | 获取单个评论 |
| ListComment | [ListCommentRequest](#api-v1-ListCommentRequest) | [ListCommentReply](#api-v1-ListCommentReply) | 获取评论列表 |
| CheckComment | [CheckCommentRequest](#api-v1-CheckCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 审核评论 |

 



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
| envs | [EnvDependent](#api-v1-EnvDependent) | repeated |  |





 

 

 


<a name="api-v1-ConfigAPI"></a>

### ConfigAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSettings | [.google.protobuf.Empty](#google-protobuf-Empty) | [Settings](#api-v1-Settings) | 获取系统配置（针对所有用户，只读） |
| UpdateConfig | [Configs](#api-v1-Configs) | [.google.protobuf.Empty](#google-protobuf-Empty) | UpdateConfig 更新配置 |
| ListConfig | [ListConfigRequest](#api-v1-ListConfigRequest) | [Configs](#api-v1-Configs) | ListConfig 查询配置项 |
| GetStats | [.google.protobuf.Empty](#google-protobuf-Empty) | [Stats](#api-v1-Stats) | 获取系统配置 |
| UpdateSitemap | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新站点地图 |

 



<a name="api_v1_document-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/document.proto



<a name="api-v1-CreateDocumentItem"></a>

### CreateDocumentItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| attachment_id | [int64](#int64) |  |  |
| price | [int32](#int32) |  |  |






<a name="api-v1-CreateDocumentRequest"></a>

### CreateDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| overwrite | [bool](#bool) |  |  |
| category_id | [int64](#int64) | repeated |  |
| document | [CreateDocumentItem](#api-v1-CreateDocumentItem) | repeated |  |






<a name="api-v1-DeleteDocumentRequest"></a>

### DeleteDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Document"></a>

### Document



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| keywords | [string](#string) |  |  |
| description | [string](#string) |  |  |
| user_id | [int64](#int64) |  |  |
| cover | [string](#string) |  |  |
| width | [int32](#int32) |  |  |
| height | [int32](#int32) |  |  |
| preview | [int32](#int32) |  |  |
| pages | [int32](#int32) |  |  |
| uuid | [string](#string) |  |  |
| download_count | [int32](#int32) |  |  |
| view_count | [int32](#int32) |  |  |
| favorite_count | [int32](#int32) |  |  |
| comment_count | [int32](#int32) |  |  |
| score | [int32](#int32) |  |  |
| score_count | [int32](#int32) |  |  |
| price | [int32](#int32) |  |  |
| size | [int64](#int64) |  |  |
| status | [int32](#int32) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| deleted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| recommend_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| deleted_user_id | [int64](#int64) |  |  |
| username | [string](#string) |  |  |
| category_id | [int64](#int64) | repeated |  |
| deleted_username | [string](#string) |  |  |
| ext | [string](#string) |  |  |
| attachment | [Attachment](#api-v1-Attachment) |  |  |
| user | [User](#api-v1-User) |  |  |
| enable_gzip | [bool](#bool) |  |  |






<a name="api-v1-DocumentScore"></a>

### DocumentScore



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| document_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| score | [int32](#int32) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-DownloadDocumentReply"></a>

### DownloadDocumentReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="api-v1-GetDocumentRequest"></a>

### GetDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| with_author | [bool](#bool) |  |  |






<a name="api-v1-ListDocumentForHomeItem"></a>

### ListDocumentForHomeItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [int64](#int64) |  |  |
| category_cover | [string](#string) |  |  |
| category_name | [string](#string) |  |  |
| document | [Document](#api-v1-Document) | repeated |  |






<a name="api-v1-ListDocumentForHomeRequest"></a>

### ListDocumentForHomeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |






<a name="api-v1-ListDocumentForHomeResponse"></a>

### ListDocumentForHomeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document | [ListDocumentForHomeItem](#api-v1-ListDocumentForHomeItem) | repeated |  |






<a name="api-v1-ListDocumentReply"></a>

### ListDocumentReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| document | [Document](#api-v1-Document) | repeated |  |






<a name="api-v1-ListDocumentRequest"></a>

### ListDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| category_id | [int64](#int64) | repeated |  |
| user_id | [int64](#int64) | repeated |  |
| status | [int32](#int32) | repeated |  |
| is_recommend | [bool](#bool) | repeated |  |
| limit | [int64](#int64) |  |  |






<a name="api-v1-RecoverRecycleDocumentRequest"></a>

### RecoverRecycleDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-SearchDocumentReply"></a>

### SearchDocumentReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| spend | [string](#string) |  | 搜索耗时 |
| document | [Document](#api-v1-Document) | repeated |  |






<a name="api-v1-SearchDocumentRequest"></a>

### SearchDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| wd | [string](#string) |  |  |
| category_id | [int64](#int64) | repeated | 分类 |
| sort | [string](#string) |  | 排序 |
| ext | [string](#string) |  | 类型 |






<a name="api-v1-SetDocumentRecommendRequest"></a>

### SetDocumentRecommendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |
| type | [int32](#int32) |  | 0, 取消推荐，1:推荐 2:重新推荐 |





 

 

 


<a name="api-v1-DocumentAPI"></a>

### DocumentAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListDocumentForHome | [ListDocumentForHomeRequest](#api-v1-ListDocumentForHomeRequest) | [ListDocumentForHomeResponse](#api-v1-ListDocumentForHomeResponse) |  |
| SetDocumentRecommend | [SetDocumentRecommendRequest](#api-v1-SetDocumentRecommendRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| CreateDocument | [CreateDocumentRequest](#api-v1-CreateDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| UpdateDocument | [Document](#api-v1-Document) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteDocument | [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetDocument | [GetDocumentRequest](#api-v1-GetDocumentRequest) | [Document](#api-v1-Document) |  |
| GetRelatedDocuments | [Document](#api-v1-Document) | [ListDocumentReply](#api-v1-ListDocumentReply) |  |
| DownloadDocument | [Document](#api-v1-Document) | [DownloadDocumentReply](#api-v1-DownloadDocumentReply) |  |
| ListDocument | [ListDocumentRequest](#api-v1-ListDocumentRequest) | [ListDocumentReply](#api-v1-ListDocumentReply) |  |
| SearchDocument | [SearchDocumentRequest](#api-v1-SearchDocumentRequest) | [SearchDocumentReply](#api-v1-SearchDocumentReply) |  |
| SetDocumentScore | [DocumentScore](#api-v1-DocumentScore) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置文档评分 |
| GetDocumentScore | [DocumentScore](#api-v1-DocumentScore) | [DocumentScore](#api-v1-DocumentScore) | 获取当前登录用户的文档评分 |


<a name="api-v1-RecycleAPI"></a>

### RecycleAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListRecycleDocument | [ListDocumentRequest](#api-v1-ListDocumentRequest) | [ListDocumentReply](#api-v1-ListDocumentReply) | 文档回收站列表 |
| RecoverRecycleDocument | [RecoverRecycleDocumentRequest](#api-v1-RecoverRecycleDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 恢复回收站文档，支持恢复单个文档或者是批量恢复 |
| DeleteRecycleDocument | [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除回收站文档 |
| ClearRecycleDocument | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 清空回收站文档 |

 



<a name="api_v1_favorite-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/favorite.proto



<a name="api-v1-DeleteFavoriteRequest"></a>

### DeleteFavoriteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Favorite"></a>

### Favorite



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| document_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| ext | [string](#string) |  |  |
| score | [int32](#int32) |  |  |
| size | [int64](#int64) |  |  |
| pages | [int32](#int32) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-GetFavoriteRequest"></a>

### GetFavoriteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  |  |






<a name="api-v1-ListFavoriteReply"></a>

### ListFavoriteReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| favorite | [Favorite](#api-v1-Favorite) | repeated |  |






<a name="api-v1-ListFavoriteRequest"></a>

### ListFavoriteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |





 

 

 


<a name="api-v1-FavoriteAPI"></a>

### FavoriteAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateFavorite | [Favorite](#api-v1-Favorite) | [Favorite](#api-v1-Favorite) | 添加收藏 |
| DeleteFavorite | [DeleteFavoriteRequest](#api-v1-DeleteFavoriteRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 取消收藏 |
| GetFavorite | [GetFavoriteRequest](#api-v1-GetFavoriteRequest) | [Favorite](#api-v1-Favorite) | 根据文章id，查询用户是否有收藏某篇文档 |
| ListFavorite | [ListFavoriteRequest](#api-v1-ListFavoriteRequest) | [ListFavoriteReply](#api-v1-ListFavoriteReply) | 查询用户的收藏 |

 



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
| field | [string](#string) | repeated |  |





 

 

 


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
| enable_upload | [bool](#bool) |  |  |
| enable_comment_approval | [bool](#bool) |  |  |
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

 



<a name="api_v1_report-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/report.proto



<a name="api-v1-DeleteReportRequest"></a>

### DeleteReportRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-ListReportReply"></a>

### ListReportReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| report | [Report](#api-v1-Report) | repeated |  |






<a name="api-v1-ListReportRequest"></a>

### ListReportRequest



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
这里是proto文件中的结构体，可以根据需要删除或者调整


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| document_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| reason | [int32](#int32) |  |  |
| status | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| document_title | [string](#string) |  |  |
| remark | [string](#string) |  |  |
| username | [string](#string) |  |  |





 

 

 


<a name="api-v1-ReportAPI"></a>

### ReportAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateReport | [Report](#api-v1-Report) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| UpdateReport | [Report](#api-v1-Report) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| DeleteReport | [DeleteReportRequest](#api-v1-DeleteReportRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| ListReport | [ListReportRequest](#api-v1-ListReportRequest) | [ListReportReply](#api-v1-ListReportReply) |  |

 



<a name="api_v1_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/user.proto



<a name="api-v1-DeleteUserRequest"></a>

### DeleteUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Dynamic"></a>

### Dynamic



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| content | [string](#string) |  |  |
| type | [int32](#int32) |  | 类型 |
| username | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-FindPasswordRequest"></a>

### FindPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| token | [string](#string) |  |  |
| password | [string](#string) |  |  |
| captcha | [string](#string) |  |  |
| captcha_id | [string](#string) |  |  |






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






<a name="api-v1-ListUserDynamicReply"></a>

### ListUserDynamicReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| dynamic | [Dynamic](#api-v1-Dynamic) | repeated |  |






<a name="api-v1-ListUserDynamicRequest"></a>

### ListUserDynamicRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
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
| limit | [int64](#int64) |  |  |






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
| email | [string](#string) |  |  |






<a name="api-v1-SetUserRequest"></a>

### SetUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| group_id | [int64](#int64) | repeated |  |






<a name="api-v1-Sign"></a>

### Sign



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| sign_at | [int32](#int32) |  |  |
| ip | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






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
| credit_count | [int32](#int32) |  |  |





 

 

 


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
| FindPasswordStepOne | [FindPasswordRequest](#api-v1-FindPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| FindPasswordStepTwo | [FindPasswordRequest](#api-v1-FindPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |

 



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

