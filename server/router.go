package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) initRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("mail"))

	r.LoadHTMLGlob("./template/*")
	r.Static("/static", "./static")

	r.Use(sessions.Sessions("session", store))
	// 接口路由
	{
		r.POST("/api/reg", func(c *gin.Context) {
			s.register(c)
		})

		r.POST("/api/login", func(c *gin.Context) {
			s.login(c)
		})

		r.GET("/api/logout", func(c *gin.Context) {
			s.logout(c)
		})

		r.POST("/getVulInfo", func(c *gin.Context) {
			s.getVulInfo(c)
		})

		r.POST("/scannerStart", func(c *gin.Context) {
			s.pushStart(c)
		})

		r.GET("/getLastTime", func(c *gin.Context) {
			s.getLastTime(c)
		})

		r.GET("/api/recentList", func(c *gin.Context) {
			if s.getSession(c) {
				s.getVulList(c)
			} else {
				c.String(403, "403")
			}
		})

		r.GET("/api/all", func(c *gin.Context) {
			if s.getSession(c) {
				s.getAllVul(c)
			} else {
				c.String(403, "403")
			}
		})

	}

	// 模板路由
	{
		r.GET("/", func(c *gin.Context) {
			isLogin := s.getSession(c)
			username := s.getNameByEmail(c)
			c.HTML(http.StatusOK, "index.html", gin.H{"isLogin": isLogin, "username": username})
		})

		r.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})

		r.GET("/all", func(c *gin.Context) {
			isLogin := s.getSession(c)
			if isLogin {
				c.HTML(http.StatusOK, "all.html", gin.H{})
			} else {
				c.String(403, "你还没有登录哦~")
			}
		})

		r.GET("/vul", func(c *gin.Context) {
			isLogin := s.getSession(c)
			if isLogin {
				c.HTML(http.StatusOK, "vul.html", gin.H{})
			} else {
				c.String(403, "你还没有登录哦~")
			}
		})

		r.GET("/reg", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})
	}
	return r
}
