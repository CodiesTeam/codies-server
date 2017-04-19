package topic

import (
	"codies-server/skeleton/common"
	"fmt"
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
	ID         string   `orm:"column(id);pk" json:"id"`
	Type       PostType `orm:"column(type);null" json:"-"`
	Title      string   `orm:"column(title);size(200);null" json:"title"`
	Content    string   `orm:"column(content);null" json:"content"`
	AuthorID   string   `orm:"column(author_id);size(36)" json:"author_id"`
	ToUsers    string   `orm:"column(to_users);size(333);null" json:"to_users,omitempty"`
	CreatedAt  int64    `orm:"column(created_at)" json:"created_at"`
	ModifiedAt int64    `orm:"column(modified_at)" json:"modified_at"`
	DeleteAt   int64    `orm:"column(delete_at);null" json:"delete_at,omitempty"`
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

func CountReply(id string) (int, error) {
	sql := fmt.Sprintf(`select count(id) from %s where id like "%s"`, "post", id)
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
