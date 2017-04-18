package topic

import (
	"codies-server/skeleton/common"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	TopicType PostType = iota
	ReplyType
	CommentType
)

type PostType int

type Post struct {
	ID         string `orm:"column(id);pk"`
	Type       int    `orm:"column(type);null"`
	Title      string `orm:"column(title);size(200);null"`
	Content    string `orm:"column(content);null"`
	AuthorID   string `orm:"column(author_id);size(36)"`
	ToUsers    string `orm:"column(to_users);size(333);null"`
	CreatedAt  int64  `orm:"column(created_at)"`
	ModifiedAt int64  `orm:"column(modified_at)"`
	DeleteAt   int64  `orm:"column(delete_at);null"`
}

func (t PostType) String() string {
	result := "unkown post type"
	switch t {
	case TopicType:
		result = "topic"
	case ReplyType:
		result = "reply"
	case CommentType:
		result = "comment"
	}
	return result
}

func (p *Post) TableName() string {
	return "post"
}

func NewPost(typ int, title, content, author string) *Post {
	now := time.Now()
	return &Post{
		ID:         newTopicID(),
		Type:       typ,
		Title:      title,
		Content:    content,
		AuthorID:   author,
		CreatedAt:  now.Unix(),
		ModifiedAt: now.Unix(),
	}
}

func (p *Post) isValid() bool {
	return p.ID != "" &&
		p.AuthorID != "" &&
		(p.Title != "" || p.Content != "") &&
		p.CreatedAt > 0 &&
		p.ModifiedAt >= p.CreatedAt
}

// Insert insert a new Post into database
func (p *Post) Insert() error {
	if !p.isValid() {
		return common.InvalidArgumentErr("post %#v is not valid", p)
	}

	o := orm.NewOrm()
	_, err := o.Insert(p)
	return err
}

/*// GetPostByID retrieves Post by ID. Returns error if
// ID doesn't exist
func GetPostByID(id int) (v *Post, err error) {
	o := orm.NewOrm()
	v = &Post{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}*/
