package asynqmw

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hibiken/asynq"
	"mini_tiktok/cmd/video/global"
	"os"
	"os/exec"
	"time"
)

const (
	TypeFfmpegTask = "ffmpeg:thumbnail"
)

type FfmpegTaksPayload struct {
	SavePhotoPath string `json:"save_photo_path"`
	TmpVideoPath  string `json:"video_path"`
	TmpPhotoPath  string `json:"photo_path"`
}

func NewTask(savaPhotoPath, videoPath, photoPath string) error {
	payload, err := json.Marshal(FfmpegTaksPayload{
		SavePhotoPath: savaPhotoPath,
		TmpPhotoPath:  photoPath,
		TmpVideoPath:  videoPath,
	})
	if err != nil {
		return err
	}
	t := asynq.NewTask(TypeFfmpegTask, payload, asynq.MaxRetry(5), asynq.Timeout(20*time.Minute))
	info, err := global.AsynqClient.Enqueue(t, asynq.MaxRetry(10))
	if err != nil {
		return err
	}
	klog.Infof("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	return nil
}

func HandleFfmpegTask(ctx context.Context, t *asynq.Task) error {
	var p FfmpegTaksPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	defer os.Remove(p.TmpVideoPath)
	defer os.Remove(p.TmpPhotoPath)
	cmd := exec.Command("ffmpeg", "-i", p.TmpVideoPath, p.TmpPhotoPath,
		"-ss", "00:00:00", "-r", "1", "-vframes", "1", "-an", "-vcodec", "mjpeg")
	_ = cmd.Run()

	_, err := global.CosClient.Object.PutFromFile(ctx, p.SavePhotoPath, p.TmpPhotoPath, nil)
	if err != nil {
		return err
	}
	klog.Infof("generate thumbnail: video_path=%s, photo_path=%s", p.TmpVideoPath, p.TmpPhotoPath)
	return nil
}
