package biz

import (
	"context"
	"errors"
	"moredoc/middleware/auth"
	"moredoc/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkGinPermission(dbModel *model.DBModel, ctx *gin.Context) (userClaims *auth.UserClaims, statusCode int, err error) {
	var ok bool
	userClaims, ok = ctx.Value(auth.CtxKeyUserClaims.String()).(*auth.UserClaims)
	if !ok || dbModel.IsInvalidToken(userClaims.UUID) {
		statusCode = http.StatusUnauthorized
		return nil, statusCode, errors.New(ErrorMessageInvalidToken)
	}

	if yes := dbModel.CheckPermissionByUserId(userClaims.UserId, ctx.Request.URL.Path, ctx.Request.Method); !yes {
		statusCode = http.StatusForbidden
		return nil, statusCode, errors.New(ErrorMessagePermissionDenied)
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
	if yes := dbModel.CheckPermissionByUserId(userClaims.UserId, fullMethod); !yes {
		return nil, status.Errorf(codes.PermissionDenied, ErrorMessagePermissionDenied)
	}
	return
}
