package cos

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sourcegraph/conc"
	"mini_tiktok/cmd/video/global"
	"mini_tiktok/cmd/video/mw/asynqmw"
	"os"
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
	var wg conc.WaitGroup
	// 上传视频到cos
	wg.Go(func() {
		_, e := global.CosClient.Object.Put(ctx, saveVideoPath, bytes.NewReader(file), nil)
		if e != nil {
			err = fmt.Errorf("上传视频失败：%w", e)
		}
	})
	// 截图上传到cos
	wg.Go(func() {
		tempPhotoPath := fmt.Sprintf("%s/test-%s.jpg", os.TempDir(), strings.Split(videoFileName, ".")[0])
		err = asynqmw.NewTask(savePhotoPath, tmpVideoPath, tempPhotoPath)
		if err != nil {
			err = fmt.Errorf("上传视频缩略图失败：%w", err)
		}
	})
	wg.Wait()
	return
}
