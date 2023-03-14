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
    - [Envs](#api-v1-Envs)
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
文章


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文章ID |
| identifier | [string](#string) |  | 文章唯一标识 |
| author | [string](#string) |  | 文章作者。如果为空，则使用网站名称作为作者 |
| view_count | [int64](#int64) |  | 文章浏览次数 |
| title | [string](#string) |  | 文章标题 |
| keywords | [string](#string) |  | 文章关键字 |
| description | [string](#string) |  | 文章描述 |
| content | [string](#string) |  | 文章内容 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章更新时间 |






<a name="api-v1-DeleteArticleRequest"></a>

### DeleteArticleRequest
删除文章请求，传入单个或者多个文章ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetArticleRequest"></a>

### GetArticleRequest
根据ID或者文章标识获取文章，二选一


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文章ID |
| identifier | [string](#string) |  | 文章唯一标识 |






<a name="api-v1-ListArticleReply"></a>

### ListArticleReply
文章列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文章总数 |
| article | [Article](#api-v1-Article) | repeated | 文章列表 |






<a name="api-v1-ListArticleRequest"></a>

### ListArticleRequest
文章列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段 |
| order | [string](#string) |  | 排序字段，根据指定的字段倒序排序 |





 

 

 


<a name="api-v1-ArticleAPI"></a>

### ArticleAPI
文章API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建文章 |
| UpdateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新文章 |
| DeleteArticle | [DeleteArticleRequest](#api-v1-DeleteArticleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除文章 |
| GetArticle | [GetArticleRequest](#api-v1-GetArticleRequest) | [Article](#api-v1-Article) | 获取文章 |
| ListArticle | [ListArticleRequest](#api-v1-ListArticleRequest) | [ListArticleReply](#api-v1-ListArticleReply) | 文章列表 |

 



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

 



<a name="api_v1_banner-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/banner.proto



<a name="api-v1-Banner"></a>

### Banner
banner，轮播图


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 主键 |
| title | [string](#string) |  | 标题，名称 |
| path | [string](#string) |  | 图片地址 |
| sort | [int32](#int32) |  | 排序，值越大越靠前 |
| enable | [bool](#bool) |  | 是否启用 |
| type | [int32](#int32) |  | 类型，如PC横幅、小程序横幅等，见 web/utils/enum.js 中的枚举 |
| url | [string](#string) |  | 跳转地址 |
| description | [string](#string) |  | 描述 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteBannerRequest"></a>

### DeleteBannerRequest
删除横幅


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetBannerRequest"></a>

### GetBannerRequest
获取横幅


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListBannerReply"></a>

### ListBannerReply
横幅列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| banner | [Banner](#api-v1-Banner) | repeated | 横幅数组 |






<a name="api-v1-ListBannerRequest"></a>

### ListBannerRequest
横幅列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| type | [int32](#int32) | repeated | 类型 |
| enable | [bool](#bool) | repeated | 是否启用 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段，不指定，则查询全部 |





 

 

 


<a name="api-v1-BannerAPI"></a>

### BannerAPI
横幅API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBanner | [Banner](#api-v1-Banner) | [Banner](#api-v1-Banner) | 创建横幅 |
| UpdateBanner | [Banner](#api-v1-Banner) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新横幅 |
| DeleteBanner | [DeleteBannerRequest](#api-v1-DeleteBannerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除横幅 |
| GetBanner | [GetBannerRequest](#api-v1-GetBannerRequest) | [Banner](#api-v1-Banner) | 查询横幅 |
| ListBanner | [ListBannerRequest](#api-v1-ListBannerRequest) | [ListBannerReply](#api-v1-ListBannerReply) | 横幅列表 |

 



<a name="api_v1_category-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/category.proto



<a name="api-v1-Category"></a>

### Category
文档分类


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 分类ID |
| parent_id | [int32](#int32) |  | 父分类ID |
| title | [string](#string) |  | 分类标题 |
| doc_count | [int32](#int32) |  | 文档数量 |
| sort | [int32](#int32) |  | 排序，倒序排序，值越大越靠前 |
| enable | [bool](#bool) |  | 是否启用 |
| cover | [string](#string) |  | 分类封面 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteCategoryRequest"></a>

### DeleteCategoryRequest
删除分类请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCategoryRequest"></a>

### GetCategoryRequest
获取分类请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCategoryReply"></a>

### ListCategoryReply
分类列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| category | [Category](#api-v1-Category) | repeated | 分类列表 |






<a name="api-v1-ListCategoryRequest"></a>

### ListCategoryRequest
分类列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| parent_id | [int64](#int64) | repeated | 父分类ID |
| wd | [string](#string) |  | 搜索关键字 |
| enable | [bool](#bool) | repeated | 是否启用 |
| field | [string](#string) | repeated | 查询字段 |





 

 

 


<a name="api-v1-CategoryAPI"></a>

### CategoryAPI
文档分类API服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建分类 |
| UpdateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新分类 |
| DeleteCategory | [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除分类 |
| GetCategory | [GetCategoryRequest](#api-v1-GetCategoryRequest) | [Category](#api-v1-Category) | 获取分类 |
| ListCategory | [ListCategoryRequest](#api-v1-ListCategoryRequest) | [ListCategoryReply](#api-v1-ListCategoryReply) | 分类列表 |

 



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

 



<a name="api_v1_favorite-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/favorite.proto



<a name="api-v1-DeleteFavoriteRequest"></a>

### DeleteFavoriteRequest
取消收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Favorite"></a>

### Favorite
文档收藏


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
根据文章id，查询用户是否有收藏某篇文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  |  |






<a name="api-v1-ListFavoriteReply"></a>

### ListFavoriteReply
查询用户的收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| favorite | [Favorite](#api-v1-Favorite) | repeated |  |






<a name="api-v1-ListFavoriteRequest"></a>

### ListFavoriteRequest
查询用户的收藏


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
删除友情链接


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Friendlink"></a>

### Friendlink
友情链接


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 主键 |
| title | [string](#string) |  | 标题 |
| link | [string](#string) |  | 链接 |
| description | [string](#string) |  | 描述 |
| sort | [int32](#int32) |  | 排序 |
| enable | [bool](#bool) |  | 是否启用 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-GetFriendlinkRequest"></a>

### GetFriendlinkRequest
获取友情链接


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListFriendlinkReply"></a>

### ListFriendlinkReply
友情链接列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| friendlink | [Friendlink](#api-v1-Friendlink) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="api-v1-ListFriendlinkRequest"></a>

### ListFriendlinkRequest
友情链接列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| wd | [string](#string) |  |  |
| enable | [bool](#bool) | repeated |  |
| field | [string](#string) | repeated |  |





 

 

 


<a name="api-v1-FriendlinkAPI"></a>

### FriendlinkAPI
友情链接服务

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateFriendlink | [Friendlink](#api-v1-Friendlink) | [Friendlink](#api-v1-Friendlink) | 创建友情链接 |
| UpdateFriendlink | [Friendlink](#api-v1-Friendlink) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新友情链接 |
| DeleteFriendlink | [DeleteFriendlinkRequest](#api-v1-DeleteFriendlinkRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除友情链接 |
| GetFriendlink | [GetFriendlinkRequest](#api-v1-GetFriendlinkRequest) | [Friendlink](#api-v1-Friendlink) | 获取友情链接 |
| ListFriendlink | [ListFriendlinkRequest](#api-v1-ListFriendlinkRequest) | [ListFriendlinkReply](#api-v1-ListFriendlinkReply) | 获取友情链接 |

 



<a name="api_v1_group-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/group.proto



<a name="api-v1-DeleteGroupRequest"></a>

### DeleteGroupRequest
删除用户组，可以批量删除


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetGroupPermissionRequest"></a>

### GetGroupPermissionRequest
获取用户组权限


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
用户组，角色


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户组ID |
| title | [string](#string) |  | 用户组名称 |
| color | [string](#string) |  | 用户组颜色 |
| is_default | [bool](#bool) |  | 是否是默认用户组 |
| is_display | [bool](#bool) |  | 是否显示 |
| description | [string](#string) |  | 用户组描述 |
| user_count | [int32](#int32) |  | 用户组下的用户数量 |
| sort | [int32](#int32) |  | 排序 |
| enable_upload | [bool](#bool) |  | 是否允许上传文档 |
| enable_comment_approval | [bool](#bool) |  | 是否需要审核评论 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-GroupPermissions"></a>

### GroupPermissions
用户组权限


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission_id | [int64](#int64) | repeated |  |






<a name="api-v1-ListGroupReply"></a>

### ListGroupReply
用户组列表


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
更新用户组权限


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
| GetGroup | [GetGroupRequest](#api-v1-GetGroupRequest) | [Group](#api-v1-Group) | 获取用户组 |
| ListGroup | [ListGroupRequest](#api-v1-ListGroupRequest) | [ListGroupReply](#api-v1-ListGroupReply) | 获取用户组列表 |
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

