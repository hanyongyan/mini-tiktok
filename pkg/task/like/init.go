package like

import "github.com/jasonlvhit/gocron"

func Init() {
	s := gocron.NewScheduler()
	s.Every(1).Hours().Do(RedisLikeDateToMysql)
	<-s.Start()
}

// RedisLikeDateToMysql 定时将redis的点赞数据存入mysql
func RedisLikeDateToMysql() {
	// todo ...
}
