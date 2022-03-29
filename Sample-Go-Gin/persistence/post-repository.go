package persistence

import (
	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostRepository interface {
	GetPosts() []Post
	AddPost(Post)
}

type postRepository struct {
	database *gorm.DB
}

func Init() PostRepository {
	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
	database, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		QueryFields: true,
	})
	util.IsNilThrow(error != nil, "Connection Failed Database")

	database.AutoMigrate(&Post{})
	return &postRepository{
		database: database,
	}
}

func (this *postRepository) GetPosts() []Post {
	var posts []Post
	error := this.database.Find(&posts).Error
	util.IsNilThrow(error != nil, "GetPosts Failed!")
	return posts
}

func (this *postRepository) AddPost(post Post) {
	error := this.database.Save(&post).Error
	util.IsNilThrow(error != nil, "AddPost Failed!")
}
