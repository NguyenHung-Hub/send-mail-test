package main

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

type MailReq struct {
	Content string `json:"content"`
	To      string `json:"to"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/sendmail", func(ctx *gin.Context) {

		var req MailReq

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		send(req.Content, req.To)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully sended to " + req.To,
		})
	})
	r.Run(":3002")
}

func send(body string, to string) {
	from := "training.dso.xuanhoa@gmail.com"
	pass := "wqnyrtuibzuxrncs"
	// to := "daminhnguyenhung@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sended to " + to)
}
