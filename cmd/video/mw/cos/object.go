package cos

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sourcegraph/conc"
	"os"
	"os/exec"
	"strings"
)

// SaveUploadedFile 保存视频到cos
func SaveUploadedFile(ctx context.Context, file []byte, videoFileName string) (saveVideoPath, savePhotoPath string, err error) {
	saveVideoPath = fmt.Sprintf("%s%s", "/video/", videoFileName)
	savePhotoPath = fmt.Sprintf("%s%s.jpg", "/photo/", strings.Split(videoFileName, ".")[0])
	// 保存视频到临时文件夹
	tmpVideoPath := fmt.Sprintf("%s/dousheng-%s", os.TempDir(), videoFileName)
	err = os.WriteFile(tmpVideoPath, file, 0666)
	if err != nil {
		err = fmt.Errorf("上传失败：%w", err)
		return
	}
	defer os.Remove(tmpVideoPath)
	// TODO 队列防ffmpeg并发冲突
	// 使用 cmd 命令调用 ffmpeg 生成截图 ，传入的参数一为 视频的真实路径，参数二为生成图片保存的真实路径
	tempPhotoPath := fmt.Sprintf("%s/test-%s.jpg", os.TempDir(), strings.Split(videoFileName, ".")[0])
	cmd := exec.Command("ffmpeg", "-i", tmpVideoPath, tempPhotoPath,
		"-ss", "00:00:00", "-r", "1", "-vframes", "1", "-an", "-vcodec", "mjpeg")
	_ = cmd.Run()
	defer os.Remove(tempPhotoPath)
	var wg conc.WaitGroup
	// 上传视频到cos
	wg.Go(func() {
		_, e := client.Object.Put(ctx, saveVideoPath, bytes.NewReader(file), nil)
		if e != nil {
			err = fmt.Errorf("上传失败：%w", e)
		}
	})
	// 截图上传到cos
	wg.Go(func() {
		_, err = client.Object.PutFromFile(ctx, savePhotoPath, tempPhotoPath, nil)
		if err != nil {
			err = fmt.Errorf("上传失败：%w", err)
		}
	})
	wg.Wait()
	return
}
