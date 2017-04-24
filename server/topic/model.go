package topic

import (
	"codies-server/skeleton/common"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
)

const (
	TopicType PostType = iota
	ReplyType
	CommentType
)

type PostType int

type Post struct {
	ID         string   `orm:"column(id);pk" json:"id"`
	Type       PostType `orm:"column(type);null" json:"-"`
	Title      string   `orm:"column(title);size(200);null" json:"title,omitempty"`
	Content    string   `orm:"column(content);null" json:"content"`
	AuthorID   string   `orm:"column(author_id);size(36)" json:"author_id"`
	ToUsers    string   `orm:"column(to_users);size(333);null" json:"to_users,omitempty"`
	CreatedAt  int64    `orm:"column(created_at)" json:"created_at"`
	ModifiedAt int64    `orm:"column(modified_at)" json:"modified_at"`
	DeleteAt   int64    `orm:"column(delete_at);null" json:"delete_at,omitempty"`
}

type Topic struct {
	Post
	Replies []Reply `json:"replies"`
}
type Reply struct {
	Post
	Comments []Post `json:"comments"`
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

func NewPost(typ PostType, id, title, content, author string) *Post {
	now := time.Now()
	return &Post{
		ID:         id,
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

// GetPostByID retrieves Post by ID. Returns error if
// ID doesn't exist
func GetPostByID(id string) (*Post, error) {
	o := orm.NewOrm()
	p := &Post{ID: id}
	if err := o.Read(p); err != nil {
		return nil, err
	}
	return p, nil
}

// CountPost count reply or comment by id
func CountPost(id string, typ PostType) (int, error) {
	baseSQL := `select count(id) from post where id like "%s%%" and type=%d;`
	sql := fmt.Sprintf(baseSQL, id, typ)
	glog.V(2).Infof("count topic %s reply, sql: %s", id, sql)
	var count int
	err := orm.NewOrm().Raw(sql).QueryRow(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func addPost(typ PostType, id, title, content, author string) (string, error) {
	post := NewPost(typ, id, title, content, author)
	err := post.Insert()
	if err != nil {
		return "", err
	}
	return post.ID, nil
}

// func fullPost(id string)([]Post, error){

// }
