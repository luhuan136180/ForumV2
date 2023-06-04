package mysql

import (
	"furumvv2/models"
	"github.com/jmoiron/sqlx"
)

func CreatePost(post *models.CtreatePost) (err error) {
	sqlStr := `insert into post(
	title, content, author_address, post_key, topic_id, post_id,status)
	values (?,?,?,?,?,?,?)
	`
	_, err = Db.Exec(sqlStr, post.Title, post.Content, post.AuthorAddress, post.PostKey, post.TopicID, post.PostID, post.Status)
	if err != nil {
		return err
	}
	sqlStr2 := `insert into postpicture(post_id, url) values (?,?)`
	_, err = Db.Exec(sqlStr2, post.PostID, post.PictureURL)
	if err != nil {
		return err
	}
	return
}

func GetPostsList(page, size int64) (posts []*models.GetPost, err error) {
	sqlStr := `select title,content,user_name,post.post_id,postpicture.url from post join user on user.user_address=post.author_address join postpicture on postpicture.post_id = post.post_id where status=1 ORDER BY post.id DESC limit ?,?`
	posts = make([]*models.GetPost, 0, size)
	err = Db.Select(&posts, sqlStr, (page-1)*size, size)

	return posts, err
}

func GetPictureList(postIDlist []int64) (data []*models.PostPicture, err error) {
	data = make([]*models.PostPicture, len(postIDlist))
	sqlStr := `select post_id,url from postpicture where`
	query, args, err := sqlx.In(sqlStr, postIDlist)
	if err != nil {
		return nil, err
	}
	err = Db.Select(data, query, args)
	if err != nil {
		return nil, err
	}
	return
}

func GetPostByContentLIKE(word string) (data []*models.GetPost, err error) {
	keyword := "%" + word + "%"
	sqlStr := `select title,content,user_name,post.post_id,postpicture.url,user.user_address from post 
			join user on user.user_address=post.author_address 
			join postpicture on postpicture.post_id = post.post_id
			where post.content LIKE ? and status=1 ORDER BY post.id DESC;`
	data = make([]*models.GetPost, 0)
	err = Db.Select(&data, sqlStr, keyword)
	if err != nil {
		return nil, err
	}
	return
}

func GetPostByTitleLIKE(word string) (data []*models.GetPost, err error) {
	keyword := "%" + word + "%"
	sqlStr := `select title,content,user_name,post.post_id,postpicture.url from post 
			join user on user.user_address=post.author_address 
			join postpicture on postpicture.post_id = post.post_id
			where post.title LIKE ? and status=1 ORDER BY post.id DESC`
	data = make([]*models.GetPost, 0)
	err = Db.Select(&data, sqlStr, keyword)
	if err != nil {
		return nil, err
	}
	return
}

func GetPostByPostID(postid int64) (data []*models.GetPost, err error) {
	sqlStr := `select title,content,user_name,postpicture.url,user.user_address,post.post_id from post 
			join user on user.user_address=post.author_address
      		join postpicture on postpicture.post_id=post.post_id 
			where post.post_id = ? order by post.create_time`
	data = make([]*models.GetPost, 0)
	err = Db.Select(&data, sqlStr, postid)
	if err != nil {
		return nil, err
	}
	return

}

func CreateResponseByPostID(post *models.CtreatePost) (err error) {
	sqlStr := `insert into post(
	 content, author_address, post_key, topic_id, post_id,status)
	values (?,?,?,?,?,?)
	`
	_, err = Db.Exec(sqlStr, post.Content, post.AuthorAddress, post.PostKey, post.TopicID, post.PostID, post.Status)
	if err != nil {
		return err
	}
	return
}
