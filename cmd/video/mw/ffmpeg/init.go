package ffmpeg

import "mini_tiktok/pkg/consts"

func Init() {
	ffchan = make(chan Ffmsg, consts.MaxMsgCount)
	go dispatcher()
}
