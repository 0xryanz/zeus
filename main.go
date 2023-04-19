package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"zeus/command"
	"zweb"

	uuid "github.com/satori/go.uuid"
)

func Reply(message string, client *Client) {

	platform, err := strconv.Atoi(string(message))
	if err != nil {
		fmt.Println("unknown command")
	}

	commands := command.PlatformCommands[platform]
	downloadUrl := strings.Join(commands[len(commands)-1].Subs, "")
	commands = commands[:len(commands)-1]

	for _, command := range commands {
		cmdMain := command.Subs[0]
		args := command.Subs[1:]

		cmd := exec.Command(cmdMain, args...)
		cmd.Dir = command.ExecDir
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("cmd.Run() failed with %s\n", err)
		}
		WebsocketManager.Send(client.Id, client.Group, []byte(out))
	}
	WebsocketManager.Send(client.Id, client.Group, []byte(downloadUrl))
	fmt.Println("ok")
}

type File struct {
	Name    string
	Size    int64
	ModTime string
}

func main() {
	go WebsocketManager.Start()
	go WebsocketManager.SendService()
	// go WebsocketManager.SendGroupService()
	// go WebsocketManager.SendAllService()

	r := zweb.Default()
	r.Use(zweb.Logger())

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(ctx *zweb.Context) {
		ctx.String(200, "fuck")
	})

	r.GET("/file", func(ctx *zweb.Context) {
		localFiles, err := ioutil.ReadDir("./static/file/")
		if err != nil {
			log.Fatal(err)
		}

		files := []File{}
		for _, file := range localFiles {
			files = append(files, File{
				Name:    file.Name(),
				Size:    (file.Size() / 1024 / 1024),
				ModTime: file.ModTime().Format("2006-01-02 15:04:05"),
			})
		}
		ctx.HTML(http.StatusOK, "file.tmpl", zweb.H{
			"files": files,
		})
	})

	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/:channel", func(ctx *zweb.Context) {
			ctx.UpgradeWebSocket()
			client := &Client{
				Id:      uuid.NewV4().String(),
				Group:   ctx.Param("channel"),
				Socket:  ctx.WsConn,
				Message: make(chan []byte, 1024),
			}

			WebsocketManager.RegisterClient(client)
			go client.Read()
			go client.Write()
		})
	}

	srv := &http.Server{
		Addr:    ":9090",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}

	// r.Run(":9090")
}
