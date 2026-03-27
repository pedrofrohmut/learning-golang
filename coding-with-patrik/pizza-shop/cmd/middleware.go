package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (this *Handler) AuthMiddleware() gin.HandlerFunc {
	return func (ctx *gin.Context) {
		var userId = GetSessionString(ctx, "userId")
		if userId == "" {
			fmt.Printf("AuthMiddleware: User id is blank on session\n")
			ctx.Redirect(303, "/login")
			ctx.Abort()
			return
		}

		var _, err = this.users.GetUserById(userId)
		if err != nil {
			fmt.Printf("AuthMiddleware: Error to get user by id\n")
			ClearSession(ctx)
			ctx.Redirect(303, "/login")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
