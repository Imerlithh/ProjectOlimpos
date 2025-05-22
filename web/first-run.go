package web

import (
	"html/template"
	"net/http"
	"os"
	"strconv"

	"ProjectOlimpos/config"
	"ProjectOlimpos/db/operations"
	"github.com/gin-gonic/gin"
)

func ShowFirstRunForm(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("web/templates/first_run.html"))
	tmpl.Execute(c.Writer, nil)
}

func HandleFirstRunSubmit(c *gin.Context) {
	port, err := strconv.Atoi(c.PostForm("db_port"))
	if err != nil {
		c.String(http.StatusBadRequest, "Port sayısal olmalı")
		return
	}

	cfg := config.Config{}
	cfg.Database.Type = c.PostForm("db_type")
	cfg.Database.Host = c.PostForm("db_host")
	cfg.Database.Port = port
	cfg.Database.User = c.PostForm("db_user")
	cfg.Database.Password = c.PostForm("db_pass")
	cfg.Database.Name = c.PostForm("db_name")
	cfg.Database.Extra = c.PostForm("db_extra")

	if err := os.MkdirAll("config", 0700); err != nil {
		c.String(http.StatusInternalServerError, "Config klasörü oluşturulamadı: %v", err)
		return
	}
	if err := config.Save(cfg); err != nil {
		c.String(http.StatusInternalServerError, "Config dosyası yazılamadı: %v", err)
		return
	}

	operations.InitDB()
	c.Redirect(http.StatusFound, "/")
}
