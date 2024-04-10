package biz

import (
	"context"
	"errors"
	"fmt"
	"moredoc/middleware/auth"
	"moredoc/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errorMessagePermissionDeniedFormat = "您没有权限访问【%s】"

func checkGinPermission(dbModel *model.DBModel, ctx *gin.Context) (userClaims *auth.UserClaims, statusCode int, err error) {
	userClaims, statusCode, err = checkGinLogin(dbModel, ctx)
	if err != nil {
		return
	}

	if permission, yes := dbModel.CheckPermissionByUserId(userClaims.UserId, ctx.Request.URL.Path, ctx.Request.Method); !yes {
		statusCode = http.StatusForbidden
		item := permission.Title
		if permission.Title == "" {
			item = permission.Path
		}
		return userClaims, statusCode, fmt.Errorf(errorMessagePermissionDeniedFormat, item)
	}
	return
}

func checkGinLogin(dbModel *model.DBModel, ctx *gin.Context) (userClaims *auth.UserClaims, statusCode int, err error) {
	var ok bool
	userClaims, ok = ctx.Value(auth.CtxKeyUserClaims.String()).(*auth.UserClaims)
	if !ok || dbModel.IsInvalidToken(userClaims.UUID) {
		statusCode = http.StatusUnauthorized
		return nil, statusCode, errors.New(ErrorMessageInvalidToken)
	}
	return
}

func checkGRPCPermission(dbModel *model.DBModel, ctx context.Context) (userClaims *auth.UserClaims, err error) {
	// 检查权限。
	// 如果userClaims为空，表示未登录(此时err一定不为nil)，否则表示已登录
	// 如果userClaims不为空(已登录)，err==nil表示是有管理权限，否则表示没有权限
	userClaims, err = checkGRPCLogin(dbModel, ctx)
	if err != nil {
		return
	}

	fullMethod, _ := ctx.Value(auth.CtxKeyFullMethod).(string)
	if permission, yes := dbModel.CheckPermissionByUserId(userClaims.UserId, fullMethod); !yes {
		item := permission.Title
		if item == "" {
			item = permission.Path
		}
		return userClaims, fmt.Errorf(errorMessagePermissionDeniedFormat, item)
	}
	return
}

func checkGRPCLogin(dbModel *model.DBModel, ctx context.Context) (userClaims *auth.UserClaims, err error) {
	var ok bool
	userClaims, ok = ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok || dbModel.IsInvalidToken(userClaims.UUID) {
		return nil, status.Errorf(codes.Unauthenticated, ErrorMessageInvalidToken)
	}
	return
}
