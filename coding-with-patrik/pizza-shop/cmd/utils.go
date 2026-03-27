package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Config struct {
	Port string
	DbPath string
	SessionSecretKey string
}

func getEnv(key string, defaultValue string) string {
	var envValue = os.Getenv(key)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}

func loadConfig() Config {
	return Config {
		Port: getEnv("PORT", "5000"),
		DbPath: getEnv("DATABASE_URL", "./data/orders.db"),
		SessionSecretKey: getEnv("SESSION_SECRET_KEY", "pizza-order-secret-key"),
	}
}

func loadTemplates(router *gin.Engine) error {
	var functions = template.FuncMap {
		"add": func (a int, b int) int {
			return a + b
		},
		"json": func (v any) template.JS {
			var b, _ = json.Marshal(v)
			return template.JS(b)
		},
	}

	var templ, err = template.New("").Funcs(functions).ParseGlob("templates/*.tmpl")
	if err != nil {
		return err
	}

	router.SetHTMLTemplate(templ)

	return nil
}

func setupSessionStore(db *gorm.DB, secretKey []byte) sessions.Store {
	var second = 1
	var fullDay = 24 * 60 * 60 * second
	var onlyWorkForOurSite http.SameSite = 3
	var store = gormsessions.NewStore(db, true, secretKey)
	store.Options(sessions.Options{
		Path: "/",
		MaxAge: fullDay,
		HttpOnly: true,
		Secure: true,
		SameSite: onlyWorkForOurSite,
	})

	return store
}

func SetSessionValue(ctx *gin.Context, key string, value any) error {
	var session = sessions.Default(ctx)
	session.Set(key, value)
	return session.Save()
}

func GetSessionString(ctx *gin.Context, key string) string {
	var session = sessions.Default(ctx)
	var value = session.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func ClearSession(ctx *gin.Context) error {
	var session = sessions.Default(ctx)
	session.Clear()
	return session.Save()
}
