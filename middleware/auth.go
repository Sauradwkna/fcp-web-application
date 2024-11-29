package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// TODO: answer here
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if ctx.GetHeader("Content-Type") == "application/json" {
                ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("error unauthorized"))
            } else {
                ctx.Redirect(http.StatusSeeOther, "/client/login")
            }
			return
		}

		claims := &model.Claims{}
		tokenClaim, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
            if err == jwt.ErrSignatureInvalid {
                ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse(err.Error()))
                return
            }
            ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
            return
        }
		
		if !tokenClaim.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		
		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}
