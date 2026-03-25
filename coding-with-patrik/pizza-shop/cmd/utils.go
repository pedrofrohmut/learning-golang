package main

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
	DbPath string
}

func getEnv(key string, defaultValue string) string {
	var envValue = os.Getenv(key)
	if envValue == "" {
		return defaultValue
	}
	return envValue
}

func loadConfig() Config {
	return Config{
		Port: getEnv("PORT", "5000"),
		DbPath: getEnv("DATABASE_URL", "./data/orders.db"),
	}
}

func loadTemplates(router *gin.Engine) error {
	var functions = template.FuncMap{
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
