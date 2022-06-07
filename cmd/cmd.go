package main

import (
	"context"
	"fmt"
	"github.com/Yuukiiii/geekbangWeek04/configs"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 启动服务可以添加参数，例子中是指定 http server 端口
var cmd = &cobra.Command{
	Use:          "user_service",
	Short:        "user_service",
	SilenceUsage: true,
	Example:      `./user_service --port 8090 --user user`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		checkParam()
		g, _ := errgroup.WithContext(context.Background())
		// 参数通过命令行获取
		// 依赖项包括一个 http server 和一个简单的 fmt.Println
		server := initServer(configs.GlobalConfig.Port, configs.GlobalConfig.Message)
		g.Go(func() error {
			return server.ListenAndServe()
		})
		g.Go(func() error {
			sigChan := make(chan os.Signal)
			defer close(sigChan)
			// 这里以 sigterm 为例，实际中可以根据不同的信号做不同的处理
			signal.Notify(sigChan, syscall.SIGTERM)
			_ = <-sigChan
			//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			// 需要保证服务器优雅终止，但也不能超过一定时间
			return server.Shutdown(ctx)
		})
		err := g.Wait()
		fmt.Println(err)
	},
}

func init() {
	pf := cmd.PersistentFlags()
	pf.IntVar(&configs.GlobalConfig.Port, "port", 8080, "server port")
	pf.StringVar(&configs.GlobalConfig.Message, "message", "user", "server user message")
}

func checkParam() {
	if configs.GlobalConfig.Port > 65535 || configs.GlobalConfig.Port < 0 {
		logger.Fatal("invalid port value")
	}
}
