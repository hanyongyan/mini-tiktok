package ffmpeg

import (
	"fmt"
	"io"
	"log"
	"mini_tiktok/cmd/video/mw/ftp"
	"mini_tiktok/pkg/configs/config"
	"os"
	"os/exec"
)

type Ffmsg struct {
	Filename string
}

var Ffchan chan Ffmsg

// 通过增加协程，将获取的信息进行派遣，当信息处理失败之后，还会将处理方式放入通道形成的队列中
func dispatcher() {
	for ffmsg := range Ffchan {
		go func(f Ffmsg) {
			err := Ffmpeg(f.Filename, config.GlobalConfigs.StaticConfig.TmpPath)
			if err != nil {
				Ffchan <- f
				log.Fatal("派遣失败：重新派遣")
			}
			log.Printf("视频%v.mp4截图处理成功", f.Filename)
			// 上传视频缩略图
			var file *os.File
			file, err = os.Open(fmt.Sprintf("%s/img/%s.jpg", config.GlobalConfigs.StaticConfig.TmpPath, f.Filename))
			if err != nil {
				return
			}
			defer file.Close()
			err = ftp.FtpClient.Stor(fmt.Sprintf("img/%s.jpg", f.Filename), file)
			if err != nil {
				log.Println(err)
			}
		}(ffmsg)
	}
}

// Ffmpeg 通过远程调用ffmpeg命令来创建视频截图
func Ffmpeg(filename, filepath string) error {
	cmd := exec.Command("ffmpeg",
		"-ss", "00:00:01",
		"-i", fmt.Sprintf("%s/%s.mp4", filepath, filename),
		"-vframes", "1",
		fmt.Sprintf("%s/img/%s.jpg", filepath, filename))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("命令运行失败", err)
		return err
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		return err
	}
	if outRes, err := io.ReadAll(stdout); err != nil {
		return err
	} else {
		log.Println("get res:", outRes)
	}
	return nil
}
