package middleware

import (
	"fg-admin/config"
	"fg-admin/constant"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/context"
)

// New returns a new auth middleware,
func NewAuth() context.Handler {
	return func(ctx context.Context) {
		auth := ctx.GetHeader("token")
		if auth == "" && (ctx.Path() != "/auth/users" || ctx.Path() != "auth/login") {
			resp := map[string]interface{}{"msg": "UNAUTHORIZED"}
			ctx.JSON(resp)
			ctx.StopExecution()
		}

		// Validate token
		token := new(jwt.Token)
		var err error
		token, err = jwt.ParseWithClaims(auth, &config.JwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
			return []byte("secret"), nil
		})


		if (err != nil || !token.Valid || token == nil) && ctx.Path() != "/user/login" {
			resp := map[string]interface{}{"msg": "TOKEN INVALID"}
			ctx.JSON(resp)
			ctx.StopExecution()
		}

		// Store user information from token into context.
		if token != nil {
			myClaims := token.Claims.(*config.JwtClaims)
			ctx.Values().Set(constant.CTX_UID, myClaims.Uid)
			ctx.Values().Set(constant.CTX_USERNAME, myClaims.Username)
			ctx.Values().Set(constant.CTX_ROLEID, myClaims.RoleId)
		}

		ctx.Next()
	}
}
