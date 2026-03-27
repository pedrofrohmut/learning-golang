package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Error string
}

type AdminDashboardData struct {
	Username string
}

func (this *Handler) HandleLoginGet(ctx *gin.Context) {
	ctx.HTML(200, "login.tmpl", LoginData{})
}

func (this *Handler) HandleLoginPost(ctx *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"min=3"`
	}

	var err = ctx.ShouldBind(&form)
	if err != nil {
		ctx.HTML(200, "login.tmpl", LoginData { Error: "Invalid input: " + err.Error() })
		return
	}

	user, err := this.users.AuthenticateUser(form.Username, form.Password)
	if err != nil {
		ctx.HTML(200, "login.tmpl", LoginData { Error: "Invalid credentials" })
		return
	}

	fmt.Printf("User successfully authenticated. Logging in...\n")
	SetSessionValue(ctx, "userId", fmt.Sprintf("%d", user.Id))
	SetSessionValue(ctx, "username", user.Username)

	fmt.Printf("Authenticated. Redirecting...\n")
	ctx.Redirect(303, "/admin")
}

func (this *Handler) HandleLogout(ctx *gin.Context) {
	var err = ClearSession(ctx)
	if err != nil {
		ctx.String(500, err.Error())
		return
	}

	ctx.Redirect(303, "/login")
}

func (this *Handler) ServeAdminDashboard(ctx *gin.Context) {
	var username = GetSessionString(ctx, "username")

	fmt.Printf("Serving page for username: %s\n", username)
	ctx.HTML(200, "admin.tmpl", AdminDashboardData { Username: username })
}
