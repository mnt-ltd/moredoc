package model

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// CanIUploadDocument 判断用户是否有上传文档的权限
// 1. 用户是否被禁用或被处罚禁止上传文档
// 2. 用户所在的用户组是否允许上传文档
func (m *DBModel) CanIAccessUploadDocument(userId int64) (yes bool) {
	if inPunishing, _ := m.isInPunishing(userId, []int{PunishmentTypeDisabled, PunishmentTypeUploadLimited}); inPunishing {
		return false
	}

	var (
		tableGroup     = Group{}.TableName()
		tableUserGroup = UserGroup{}.TableName()
		group          Group
	)
	err := m.db.Select("g.id").Table(tableGroup+" g").Joins(
		"left join "+tableUserGroup+" ug on g.id=ug.group_id",
	).Where("ug.user_id = ? and g.enable_upload = ?", userId, true).Find(&group).Error
	if err != nil {
		m.logger.Error("CanIUploadDocument", zap.Error(err))
		return
	}
	return group.Id > 0
}

// CanIUploadDocument 判断用户是否有上传文档的权限
// 1. 用户是否被禁用或被处罚禁止发布文章
// 2. 用户所在的用户组是否允许发布文章
func (m *DBModel) CanIAccessPublishArticle(userId int64) (yes bool) {
	if inPunishing, _ := m.isInPunishing(userId, []int{PunishmentTypeDisabled, PunishmentTypePublishArticleLimited}); inPunishing {
		return false
	}

	var (
		tableGroup     = Group{}.TableName()
		tableUserGroup = UserGroup{}.TableName()
		group          Group
	)
	err := m.db.Select("g.id").Table(tableGroup+" g").Joins(
		"left join "+tableUserGroup+" ug on g.id=ug.group_id",
	).Where("ug.user_id = ? and g.enable_article = ?", userId, true).Find(&group).Error
	if err != nil {
		m.logger.Error("CanIAccessPublishArticle", zap.Error(err))
		return
	}
	return group.Id > 0
}

// 用户是否可以下载文档：被禁用的账号或被禁止下载的账户不能下载
func (m *DBModel) CanIAccessDownload(userId int64) (yes bool, err error) {
	yes, err = m.isInPunishing(userId, []int{PunishmentTypeDownloadLimited, PunishmentTypeDisabled})
	yes = !yes
	if err != nil {
		m.logger.Error("CanIAccessDownload", zap.Error(err))
		return
	}
	return
}

// 用户是否可以收藏文档
func (m *DBModel) CanIAccessFavorite(userId int64) (yes bool, err error) {
	yes, err = m.isInPunishing(userId, []int{PunishmentTypeFavoriteLimited, PunishmentTypeDisabled})
	yes = !yes
	if err != nil {
		m.logger.Error("CanIAccessFavorite", zap.Error(err))
		return
	}
	return
}

// 用户是否可以评论
func (m *DBModel) CanIAccessComment(userId int64) (yes bool, err error) {
	yes, err = m.isInPunishing(userId, []int{PunishmentTypeCommentLimited, PunishmentTypeDisabled})
	yes = !yes
	if err != nil {
		m.logger.Error("CanIAccessComment", zap.Error(err))
		return
	}

	// 用户没有被禁止评论
	if userId <= 0 {
		err = fmt.Errorf("请先登录")
		return
	}

	var (
		group     Group
		userGroup UserGroup
		comment   Comment
	)

	// 查询用户组。只要用户组中有一个允许评论，就允许评论，以及用户组，有一个评论不需要审核，就不需要审核
	m.db.Select("g.id", "max(g.enable_comment) as enable_comment", "min(g.enable_comment_approval) as enable_comment_approval").Table(Group{}.TableName()+" g").Joins(
		"left join "+userGroup.TableName()+" ug on g.id=ug.group_id",
	).Where("ug.user_id = ?", userId).Find(&group)

	m.logger.Debug("CanIPublishComment", zap.Any("group", group))

	if !group.EnableComment {
		err = fmt.Errorf("您所在用户组不允许评论")
		return
	}

	// 评论时间间隔<=0，表示不限发布频率
	commentInterval := m.GetConfigOfSecurity(ConfigSecurityCommentInterval).CommentInterval
	if commentInterval <= 0 {
		return
	}

	// 获取用户最新的一条评论信息
	m.db.Select("id", "created_at").Where("user_id = ?", userId).Order("id desc").Find(&comment)
	if comment.Id <= 0 { // 用户没有评论过
		return
	}

	seconds := int32(time.Since(*comment.CreatedAt).Seconds())
	left := commentInterval - seconds
	if left > 0 {
		err = fmt.Errorf("您的评论太快了，请等待 %d 秒后再试", left)
	}
	return
}
