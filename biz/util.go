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
	var ok bool
	userClaims, ok = ctx.Value(auth.CtxKeyUserClaims.String()).(*auth.UserClaims)
	if !ok || dbModel.IsInvalidToken(userClaims.UUID) {
		statusCode = http.StatusUnauthorized
		return nil, statusCode, errors.New(ErrorMessageInvalidToken)
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

func checkGRPCPermission(dbModel *model.DBModel, ctx context.Context) (userClaims *auth.UserClaims, err error) {
	var ok bool
	userClaims, ok = ctx.Value(auth.CtxKeyUserClaims).(*auth.UserClaims)
	if !ok || dbModel.IsInvalidToken(userClaims.UUID) {
		return nil, status.Errorf(codes.Unauthenticated, ErrorMessageInvalidToken)
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
