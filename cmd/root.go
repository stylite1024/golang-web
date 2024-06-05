package cmd

import (
	"context"
	"iris-web/config"
	"iris-web/internal/logger"
	"iris-web/internal/router"
	"log"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

var (
	port    string
	rootCmd = &cobra.Command{
		Use:          config.AppName,
		Short:        config.Description,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func init() {
	// 关闭官方completion命令
	completion := completionCommand()
	completion.Hidden = true
	rootCmd.AddCommand(completion)
	// 智能提示最小位数
	rootCmd.SuggestionsMinimumDistance = 1

	// 自定义log时间格式
	customTimeFormat := "2006-01-02 15:04:05.000"
	// 关闭默认的时间戳
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime)) // 2024-06-03 06:45:53.758 Server exiting
	// log.SetFlags(log.Llongfile) // 2024-06-03 06:44:22.090 C:/Users/OBY/Desktop/douyin-tool/cmd/root.go:102: Server exiting
	// 设置自定义的日志前缀
	log.SetPrefix(time.Now().Format(customTimeFormat) + " ")

	rootCmd.Flags().StringVarP(&port, "port", "p", "8080", "port")
}

func run() error {
	ac := logger.MakeAccessLog()
	defer ac.Close() // Close the underline file.

	app := iris.New()

	// Register the middleware (UseRouter to catch http errors too).
	app.UseRouter(ac.Handler)

	app.Logger().SetLevel("error")

	router.InitRouter(app)

	// IRIS优雅关闭服务
	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		log.Println("Shutdown...")
		// 创建一个10秒的超时上下文
		timeout := 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
		close(idleConnsClosed)
		log.Println("Server exiting")
	})
	//启动服务器
	app.Listen(":"+port, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
	<-idleConnsClosed
	return nil
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
