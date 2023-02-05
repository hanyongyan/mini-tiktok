package ftp

import (
	"fmt"
	"time"

	ftpPkg "github.com/jlaffaye/ftp"
	"mini_tiktok/pkg/configs/config"
)

var FtpClient *ftpPkg.ServerConn

func Init() {
	// 连接到 FTP 服务器
	cfg := config.GlobalConfigs.FtpConfig
	conn, err := ftpPkg.Dial(
		fmt.Sprintf("%v:%v", cfg.Host, cfg.Port),
		ftpPkg.DialWithTimeout(5*time.Second),
	)
	// ftp.example.com:21是一个模拟的FTP服务器地址，实际使用中需要替换成真正的FTP服务器的地址。
	if err != nil {
		panic(err)
	}

	// 登录到 FTP 服务器
	// 需要提前设置好username and password
	if err = conn.Login(cfg.Username, cfg.Password); err != nil {
		panic(err)
	}

	FtpClient = conn
}
