package posts

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type PostStorage struct {
	Rdb *redis.Client
}

func (ps PostStorage) GetPosts() []Post {
	posts := make([]Post, 0)

	ctx := context.TODO()
	iter := ps.Rdb.Scan(ctx, 0, "post*", 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		val, _ := ps.Rdb.HGetAll(ctx, key).Result()

		posts = append(posts, FromMap(val))
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	return posts
}

func (ps PostStorage) SavePost(p Post) {
	ctx := context.TODO()

	ps.Rdb.HSet(ctx, p.ID, p.ToMap())
}
