package utils

import (
	"github.com/nanakura/go-ramda"
	"mini_tiktok/cmd/api/biz/model/api"
	"mini_tiktok/kitex_gen/videoservice"
)

func CastUserserviceVideoToApiVideo(from []*videoservice.Video) []*api.Video {
	return ramda.Map(func(t *videoservice.Video) *api.Video {
		if t == nil {
			return nil
		}
		return &api.Video{
			ID: t.Id,
			Author: &api.User{
				ID:            t.Author.Id,
				Name:          t.Author.Name,
				FollowCount:   t.Author.FollowCount,
				FollowerCount: t.Author.FollowerCount,
				IsFollow:      t.Author.IsFollow,
			},
			PlayURL:       t.PlayUrl,
			CoverURL:      t.CoverUrl,
			FavoriteCount: t.FavoriteCount,
			CommentCount:  t.CommentCount,
			IsFavorite:    t.IsFavorite,
			Title:         t.Title,
		}
	})(from)
}
