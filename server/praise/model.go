package praise

import "github.com/astaxie/beego/orm"

type Praise struct {
	ID         int    `orm:"column(target_id);pk"`
	TargetType int    `orm:"column(target_type);null"`
	LikerID    string `orm:"column(liker_id);size(36)"`
	CreatedAt  int    `orm:"column(created_at)"`
	DeleteAt   int    `orm:"column(delete_at);null"`
}

func (t *Praise) TableName() string {
	return "praise"
}

// AddPraise insert a new Praise into database and returns
// last inserted Id on success.
func AddPraise(m *Praise) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPraiseById retrieves Praise by Id. Returns error if
// Id doesn't exist
func GetPraiseById(id int) (v *Praise, err error) {
	o := orm.NewOrm()
	v = &Praise{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
