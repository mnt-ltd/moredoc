package biz

import (
	"context"
	"moredoc/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewAuthService(dbModel *model.DBModel, logger *zap.Logger) (service *UserAPIService) {
	return &UserAPIService{dbModel: dbModel, logger: logger.Named("AuthService")}
}

type ContextKey string

func (ck ContextKey) String() string {
	return string(ck)
}

const (
	CtxKeyUserClaims ContextKey = "user"
)

const (
	messageInvalidToken = "您的登录令牌已过期，请重新登录"
)

type ServiceAuthFuncOverride interface {
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}

func (s *AuthService) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := s.AuthGRPC(ctx, info)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

// AuthGRPC 验证 gRPC 请求
// 1. 从权限表中查询API，如果存在该API，则表示该API需要权限才能访问，如果不存在，则跳过
// 2. 如果用户携带有token，则根据token判断是否有效，如果有效，则获取用户信息放到ctx，否则跳过
func (s *AuthService) AuthGRPC(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		// 可忽略错误
		return ctx, err
	}

	claims, err := s.dbModel.CheckUserJWTToken(token)

	// token存在，但是不正确或者已过期，这时需要返回错误，前端清除存储的错误登录信息
	if err != nil || claims == nil || claims.ExpiresAt < time.Now().Unix() || s.dbModel.IsInvalidToken(claims.UUID) {
		return ctx, status.Error(codes.Unauthenticated, messageInvalidToken)
	}

	newCtx := context.WithValue(ctx, CtxKeyUserClaims, claims)
	return newCtx, nil
}

// AuthGin 验证 HTTP 请求
func (s *AuthService) AuthGin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("authorization")
		bearer := strings.Split(auth, " ")

		if auth == "" || len(bearer) != 2 {
			ctx.Next()
			return
		}

		token := bearer[1]
		claims, err := s.dbModel.CheckUserJWTToken(token)
		if err != nil || claims == nil || claims.ExpiresAt < time.Now().Unix() || s.dbModel.IsInvalidToken(claims.UUID) {
			ctx.JSON(http.StatusUnauthorized, status.Error(codes.Unauthenticated, messageInvalidToken))
			ctx.Abort()
			return
		}

		ctx.Set(CtxKeyUserClaims.String(), claims)
		ctx.Next()
	}
}
