package models

import "time"
import "github.com/astaxie/beego/orm"
import "strconv"

type Blog struct {
	Id              int64
	Categoryid      int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         string `orm:"index"`
	Updated         string `orm:"index"`
	Views           int64  `orm:"index"`
	Author          string
	ReplyTime       string `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	LikeCount       int64
}

func GetAllBlog() ([]*Blog, error) {
	o := orm.NewOrm()
	blogs := make([]*Blog, 0)

	qeury := o.QueryTable("blog").OrderBy("-Created")
	_, err := qeury.All(&blogs)
	return blogs, err
}

func AddBlog(title, content string) error {
	o := orm.NewOrm()

	blog := &Blog{
		Title:     title,
		Content:   content,
		Created:   time.Now().Format("2006-01-02"),
		Updated:   time.Now().Format("2006-01-02"),
		ReplyTime: time.Now().Format("2006-01-02"),
	}

	_, err := o.Insert(blog)
	return err
}

func GetBlogById(blogId string) (*Blog, error) {
	id, err := strconv.ParseInt(blogId, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	blog := &Blog{Id: id}

	err = o.Read(blog)

	return blog, err
}
