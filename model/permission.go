package model

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Permission struct {
	Id          int64     `form:"id" json:"id,omitempty" gorm:"primaryKey;autoIncrement;column:id;comment:;"`
	Category    string    `form:"category" json:"category,omitempty" gorm:"column:category;type:varchar(64);size:64;index:category;comment:权限类别组;"`
	Title       string    `form:"title" json:"title,omitempty" gorm:"column:title;type:varchar(255);size:255;comment:权限中文名称;"`
	Description string    `form:"description" json:"description,omitempty" gorm:"column:description;type:varchar(255);size:255;comment:权限描述;"`
	Identifier  string    `form:"identifier" json:"identifier,omitempty" gorm:"column:identifier;type:varchar(64);size:64;index:identifier,unique;comment:权限英文标识，如函数名称等;"`
	CreatedAt   time.Time `form:"created_at" json:"created_at,omitempty" gorm:"column:created_at;type:datetime;comment:创建时间;"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at,omitempty" gorm:"column:updated_at;type:datetime;comment:更新时间;"`
}

func (Permission) TableName() string {
	return tablePrefix + "permission"
}

// 这里是proto文件中的结构体，可以根据需要删除或者调整
//message Permission {
// int64 id = 1;
// string category = 2;
// string title = 3;
// string description = 4;
// string identifier = 5;
//   = 0;
//   = 0;
//}

// PermissionCategoryXXX 基本按照数据表来定义
const (
	PermissionCategoryAttachment      = "attachment"      // 附件管理，包括上传等
	PermissionCategoryBanner          = "banner"          // 管理横幅
	PermissionCategoryCategory        = "category"        // 管理文档分类
	PermissionCategoryConfig          = "config"          // 管理系统配置：开启验证码，是否允许上传等
	PermissionCategoryDocument        = "document"        // 管理文档
	PermissionCategoryFriendlink      = "friendlink"      // 管理友情链接
	PermissionCategoryGroup           = "group"           // 管理用户组：创建、修改、删除、查看等
	PermissionCategoryGroupPermission = "groupPermission" // 权限设置
	PermissionCategoryUser            = "user"            // 管理用户：创建、修改、删除、查看、修改密码、禁用、变更分组等
)

const (
	// PermissionCategoryAttachment = "attachment" // 附件管理，包括上传等
	PermissionIdentifierAttachmentList        = "attachmentList"        // 查看附件列表
	PermissionIdentifierAttachmentDelete      = "attachmentDelete"      // 删除单个附件
	PermissionIdentifierAttachmentBatchDelete = "attachmentBatchDelete" // 批量删除附件
	PermissionIdentifierAttachmentDisable     = "attachmentDisable"     // 禁止附件，禁止后无法访问，用于控制非法附件

	// 	PermissionCategoryBanner     = "banner"     // 管理横幅
	PermissionIdentifierBannerList        = "bannerList"        // 查看横幅列表
	PermissionIdentifierBannerCreate      = "bannerCreate"      // 创建横幅
	PermissionIdentifierBannerUpdate      = "bannerUpdate"      // 更新横幅
	PermissionIdentifierBannerDelete      = "bannerDelete"      // 删除横幅
	PermissionIdentifierBannerBatchDelete = "bannerBatchDelete" // 批量删除横幅

	// 	PermissionCategoryCategory   = "category"   // 管理文档分类
	PermissionIdentifierCategoryList        = "categoryList"        // 查看文档分类列表
	PermissionIdentifierCategoryCreate      = "categoryCreate"      // 创建分类
	PermissionIdentifierCategoryUpdate      = "categoryUpdate"      // 更新分类
	PermissionIdentifierCategoryDelete      = "categoryDelete"      // 删除分类
	PermissionIdentifierCategoryBatchDelete = "categoryBatchDelete" // 批量删除分类

	// 	PermissionCategoryConfig     = "config"     // 管理系统配置：开启验证码，是否允许上传等。不允许删除配置
	PermissionIdentifierConfigList   = "configList"   // 查看系统配置列表
	PermissionIdentifierConfigUpdate = "configUpdate" // 更新系统配置

	// 	PermissionCategoryDocument   = "document"   // 管理文档
	PermissionIdentifierDocumentList        = "documentList"        // 查看文档列表
	PermissionIdentifierDocumentCreate      = "documentCreate"      // 创建文档
	PermissionIdentifierDocumentBatchCreate = "documentBatchCreate" // 批量创建文档
	PermissionIdentifierDocumentUpdate      = "documentUpdate"      // 更新文档
	PermissionIdentifierDocumentDelete      = "documentDelete"      // 删除文档
	PermissionIdentifierDocumentBatchDelete = "documentBatchDelete" // 批量删除文档

	// 	PermissionCategoryFriendlink = "friendlink" // 管理友情链接
	PermissionIdentifierFriendlinkList        = "friendlinkList"        // 查看友情链接列表
	PermissionIdentifierFriendlinkCreate      = "friendlinkCreate"      // 创建友情链接
	PermissionIdentifierFriendlinkUpdate      = "friendlinkUpdate"      // 更新友情链接
	PermissionIdentifierFriendlinkDelete      = "friendlinkDelete"      // 删除友情链接
	PermissionIdentifierFriendlinkBatchDelete = "friendlinkBatchDelete" // 批量删除友情链接

	// 	PermissionCategoryGroup      = "group"      // 管理用户组：创建、修改、删除、查看等
	PermissionIdentifierGroupList        = "groupList"        // 查看用户组列表
	PermissionIdentifierGroupCreate      = "groupCreate"      // 创建用户组
	PermissionIdentifierGroupUpdate      = "groupUpdate"      // 更新用户组
	PermissionIdentifierGroupDelete      = "groupDelete"      // 删除用户组
	PermissionIdentifierGroupBatchDelete = "groupBatchDelete" // 批量删除用户组

	// PermissionCategoryGroupPermission = "groupPermission" // 权限设置
	PermissionIdentifierGroupPermissionList   = "groupPermissionList"   // 查看用户组权限列表
	PermissionIdentifierGroupPermissionUpdate = "groupPermissionUpdate" // 更新用户组权限

	// PermissionCategoryUser            = "user"            // 管理用户：创建、修改、删除、查看、修改密码、禁用、变更分组等
	PermissionIdentifierUserCreate         = "userCreate"         // 创建用户
	PermissionIdentifierUserUpdate         = "userUpdate"         // 修改用户
	PermissionIdentifierUserDelete         = "userDelete"         // 删除用户
	PermissionIdentifierUserBatchDelete    = "userBatchDelete"    // 删除用户
	PermissionIdentifierUserList           = "userList"           // 查看用户
	PermissionIdentifierUserChangePassword = "userChangePassword" // 修改用户密码
	PermissionIdentifierUserDisable        = "userDisable"        // 禁用用户
	PermissionIdentifierUserChangeGroup    = "userChangeGroup"    // 变更用户分组
)

func (m *DBModel) initPermission() (err error) {
	permissions := []Permission{
		{Id: 1, Category: PermissionCategoryAttachment, Identifier: PermissionIdentifierAttachmentList, Title: "附件管理"},
		{Id: 2, Category: PermissionCategoryAttachment, Identifier: PermissionIdentifierAttachmentDelete, Title: "删除附件"},
		{Id: 3, Category: PermissionCategoryAttachment, Identifier: PermissionIdentifierAttachmentBatchDelete, Title: "批量删除附件"},
		{Id: 4, Category: PermissionCategoryAttachment, Identifier: PermissionIdentifierAttachmentDisable, Title: "禁用附件"},
		{Id: 42, Category: PermissionCategoryAttachment, Identifier: PermissionIdentifierAttachmentDisable, Title: "禁用附件"},

		{Id: 5, Category: PermissionCategoryBanner, Identifier: PermissionIdentifierBannerList, Title: "横幅管理"},
		{Id: 6, Category: PermissionCategoryBanner, Identifier: PermissionIdentifierBannerCreate, Title: "创建横幅"},
		{Id: 7, Category: PermissionCategoryBanner, Identifier: PermissionIdentifierBannerUpdate, Title: "更新横幅"},
		{Id: 8, Category: PermissionCategoryBanner, Identifier: PermissionIdentifierBannerDelete, Title: "删除横幅"},
		{Id: 9, Category: PermissionCategoryBanner, Identifier: PermissionIdentifierBannerBatchDelete, Title: "批量删除横幅"},

		{Id: 10, Category: PermissionCategoryCategory, Identifier: PermissionIdentifierCategoryList, Title: "分类管理"},
		{Id: 11, Category: PermissionCategoryCategory, Identifier: PermissionIdentifierCategoryCreate, Title: "创建分类"},
		{Id: 12, Category: PermissionCategoryCategory, Identifier: PermissionIdentifierCategoryUpdate, Title: "更新分类"},
		{Id: 13, Category: PermissionCategoryCategory, Identifier: PermissionIdentifierCategoryDelete, Title: "删除分类"},
		{Id: 14, Category: PermissionCategoryCategory, Identifier: PermissionIdentifierCategoryBatchDelete, Title: "批量删除分类"},

		{Id: 15, Category: PermissionCategoryConfig, Identifier: PermissionIdentifierConfigList, Title: "系统配置管理"},
		{Id: 16, Category: PermissionCategoryConfig, Identifier: PermissionIdentifierConfigUpdate, Title: "更新系统配置"},

		{Id: 17, Category: PermissionCategoryDocument, Identifier: PermissionIdentifierDocumentList, Title: "文档管理"},
		{Id: 18, Category: PermissionCategoryDocument, Identifier: PermissionIdentifierDocumentCreate, Title: "创建文档"},
		{Id: 19, Category: PermissionCategoryDocument, Identifier: PermissionIdentifierDocumentBatchCreate, Title: "批量创建文档"},
		{Id: 20, Category: PermissionCategoryDocument, Identifier: PermissionIdentifierDocumentUpdate, Title: "更新文档"},
		{Id: 21, Category: PermissionCategoryDocument, Identifier: PermissionIdentifierDocumentDelete, Title: "删除文档"},
		{Id: 22, Category: PermissionCategoryDocument, Identifier: PermissionIdentifierDocumentBatchDelete, Title: "批量删除文档"},

		{Id: 23, Category: PermissionCategoryGroup, Identifier: PermissionIdentifierGroupList, Title: "用户组管理"},
		{Id: 24, Category: PermissionCategoryGroup, Identifier: PermissionIdentifierGroupCreate, Title: "创建用户组"},
		{Id: 25, Category: PermissionCategoryGroup, Identifier: PermissionIdentifierGroupUpdate, Title: "更新用户组"},
		{Id: 26, Category: PermissionCategoryGroup, Identifier: PermissionIdentifierGroupDelete, Title: "删除用户组"},
		{Id: 27, Category: PermissionCategoryGroup, Identifier: PermissionIdentifierGroupBatchDelete, Title: "批量删除用户组"},

		{Id: 28, Category: PermissionCategoryGroupPermission, Identifier: PermissionIdentifierGroupPermissionList, Title: "权限管理"},
		{Id: 29, Category: PermissionCategoryGroupPermission, Identifier: PermissionIdentifierGroupPermissionUpdate, Title: "设置权限"},

		{Id: 30, Category: PermissionCategoryFriendlink, Identifier: PermissionIdentifierFriendlinkList, Title: "友链管理"},
		{Id: 31, Category: PermissionCategoryFriendlink, Identifier: PermissionIdentifierFriendlinkCreate, Title: "创建友链"},
		{Id: 32, Category: PermissionCategoryFriendlink, Identifier: PermissionIdentifierFriendlinkUpdate, Title: "更新友链"},
		{Id: 33, Category: PermissionCategoryFriendlink, Identifier: PermissionIdentifierFriendlinkDelete, Title: "删除友链"},
		{Id: 34, Category: PermissionCategoryFriendlink, Identifier: PermissionIdentifierFriendlinkBatchDelete, Title: "批量删除友链"},

		{Id: 35, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserCreate, Title: "创建用户"},
		{Id: 36, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserUpdate, Title: "更新用户"},
		{Id: 37, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserDelete, Title: "删除用户"},
		{Id: 38, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserBatchDelete, Title: "批量删除用户"},
		{Id: 39, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserList, Title: "用户管理"},
		{Id: 40, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserChangePassword, Title: "修改用户密码"},
		{Id: 41, Category: PermissionCategoryUser, Identifier: PermissionIdentifierUserChangeGroup, Title: "变更用户分组"},
	}

	sess := m.db.Begin()
	defer func() {
		if err != nil {
			sess.Rollback()
		} else {
			sess.Commit()
		}
	}()

	err = sess.Where("id > ?", 0).Delete(&Permission{}).Error
	if err != nil {
		m.logger.Error("delete permission error", zap.Error(err))
		return
	}

	err = sess.Create(&permissions).Error
	if err != nil {
		m.logger.Error("create permission error", zap.Error(err))
		return
	}

	return
}

// CreatePermission 创建Permission
// TODO: 创建成功之后，注意相关表统计字段数值的增减
func (m *DBModel) CreatePermission(permission *Permission) (err error) {
	err = m.db.Create(permission).Error
	if err != nil {
		m.logger.Error("CreatePermission", zap.Error(err))
		return
	}
	return
}

// UpdatePermission 更新Permission，如果需要更新指定字段，则请指定updateFields参数
func (m *DBModel) UpdatePermission(permission *Permission, updateFields ...string) (err error) {
	db := m.db.Model(permission)

	updateFields = m.FilterValidFields(Permission{}.TableName(), updateFields...)
	if len(updateFields) > 0 { // 更新指定字段
		db = db.Select(updateFields)
	}

	err = db.Where("id = ?", permission.Id).Updates(permission).Error
	if err != nil {
		m.logger.Error("UpdatePermission", zap.Error(err))
	}
	return
}

// GetPermission 根据id获取Permission
func (m *DBModel) GetPermission(id interface{}, fields ...string) (permission Permission, err error) {
	db := m.db

	fields = m.FilterValidFields(Permission{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	err = db.Where("id = ?", id).First(&permission).Error
	return
}

// GetPermissionByIdentifier(identifier string, fields ...string) 根据唯一索引获取Permission
func (m *DBModel) GetPermissionByIdentifier(identifier string, fields ...string) (permission Permission, err error) {
	db := m.db

	fields = m.FilterValidFields(Permission{}.TableName(), fields...)
	if len(fields) > 0 {
		db = db.Select(fields)
	}

	db = db.Where("identifier = ?", identifier)

	err = db.First(&permission).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetPermissionByIdentifier", zap.Error(err))
		return
	}
	return
}

type OptionGetPermissionList struct {
	Page         int
	Size         int
	WithCount    bool                      // 是否返回总数
	Ids          []interface{}             // id列表
	SelectFields []string                  // 查询字段
	QueryRange   map[string][2]interface{} // map[field][]{min,max}
	QueryIn      map[string][]interface{}  // map[field][]{value1,value2,...}
	QueryLike    map[string][]interface{}  // map[field][]{value1,value2,...}
	Sort         []string
}

// GetPermissionList 获取Permission列表
func (m *DBModel) GetPermissionList(opt OptionGetPermissionList) (permissionList []Permission, total int64, err error) {
	db := m.db.Model(&Permission{})

	for field, rangeValue := range opt.QueryRange {
		fields := m.FilterValidFields(Permission{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		if rangeValue[0] != nil {
			db = db.Where(fmt.Sprintf("%s >= ?", field), rangeValue[0])
		}
		if rangeValue[1] != nil {
			db = db.Where(fmt.Sprintf("%s <= ?", field), rangeValue[1])
		}
	}

	for field, values := range opt.QueryIn {
		fields := m.FilterValidFields(Permission{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(fmt.Sprintf("%s in (?)", field), values)
	}

	for field, values := range opt.QueryLike {
		fields := m.FilterValidFields(Permission{}.TableName(), field)
		if len(fields) == 0 {
			continue
		}
		db = db.Where(strings.TrimSuffix(fmt.Sprintf(strings.Join(make([]string, len(values)+1), "%s like ? or"), field), "or"), values...)
	}

	if len(opt.Ids) > 0 {
		db = db.Where("id in (?)", opt.Ids)
	}

	if opt.WithCount {
		err = db.Count(&total).Error
		if err != nil {
			m.logger.Error("GetPermissionList", zap.Error(err))
			return
		}
	}

	opt.SelectFields = m.FilterValidFields(Permission{}.TableName(), opt.SelectFields...)
	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.Sort) > 0 {
		var sorts []string
		for _, sort := range opt.Sort {
			slice := strings.Split(sort, " ")
			if len(m.FilterValidFields(Permission{}.TableName(), slice[0])) == 0 {
				continue
			}

			if len(slice) == 2 {
				sorts = append(sorts, fmt.Sprintf("%s %s", slice[0], slice[1]))
			} else {
				sorts = append(sorts, fmt.Sprintf("%s desc", slice[0]))
			}
		}
		if len(sorts) > 0 {
			db = db.Order(strings.Join(sorts, ","))
		}
	}

	db = db.Offset((opt.Page - 1) * opt.Size).Limit(opt.Size)

	err = db.Find(&permissionList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		m.logger.Error("GetPermissionList", zap.Error(err))
	}
	return
}

// DeletePermission 删除数据
// TODO: 删除数据之后，存在 permission_id 的关联表，需要删除对应数据，同时相关表的统计数值，也要随着减少
func (m *DBModel) DeletePermission(ids []interface{}) (err error) {
	err = m.db.Where("id in (?)", ids).Delete(&Permission{}).Error
	if err != nil {
		m.logger.Error("DeletePermission", zap.Error(err))
	}
	return
}
