package auth

import (
	"context"
	"moredoc/conf"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

type Auth struct {
	jwt *conf.JWT
}

type UserClaims struct {
	UserId     int64
	UUID       string
	HaveAccess bool
	jwt.StandardClaims
}

func NewAuth(jwt *conf.JWT) *Auth {
	return &Auth{
		jwt: jwt,
	}
}

type ContextKey string

func (ck ContextKey) String() string {
	return string(ck)
}

const (
	CtxKeyUserClaims ContextKey = "user"
	CtxKeyFullMethod ContextKey = "fullMethod"
)

func (p *Auth) AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := p.AuthGRPC(ctx, info)
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

func (p *Auth) AuthGRPC(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {
	ctx = context.WithValue(ctx, CtxKeyFullMethod, info.FullMethod)
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return ctx, nil
	}

	claims, err := p.CheckJWTToken(token)
	if err != nil || claims == nil || claims.ExpiresAt < time.Now().Unix() {
		return ctx, nil
	}

	newCtx := context.WithValue(ctx, CtxKeyUserClaims, claims)
	return newCtx, nil
}

func (p *Auth) AuthGin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("authorization")
		bearer := strings.Split(auth, " ")

		if auth == "" || len(bearer) != 2 {
			ctx.Next()
			return
		}

		token := bearer[1]
		claims, err := p.CheckJWTToken(token)
		if err != nil || claims == nil || claims.ExpiresAt < time.Now().Unix() {
			ctx.Next()
			return
		}

		ctx.Set(CtxKeyUserClaims.String(), claims)
		ctx.Next()
	}
}

// CreateUserJWTToken 生成用户JWT Token
func (p *Auth) CreateJWTToken(userId int64) (string, error) {
	expireTime := time.Now().Add(time.Duration(p.jwt.ExpireDays) * 24 * time.Hour).Unix()
	claims := UserClaims{
		UserId: userId,
		UUID:   uuid.Must(uuid.NewV1()).String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "moredoc",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(p.jwt.Secret))
	return token, err
}

// CheckUserJWTToken 验证用户JWT token
func (p *Auth) CheckJWTToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(p.jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*UserClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
