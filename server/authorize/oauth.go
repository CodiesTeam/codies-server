/*
not used for now
*/
package authorize

type Oauth struct {
	UUID       string `orm:"column(uuid);pk"`
	OauthName  string `orm:"column(oauth_name);size(20)"`
	OauthID    string `orm:"column(oauth_id);size(45)"`
	OauthToken string `orm:"column(oauth_token);size(45)"`
}

func (t *Oauth) TableName() string {
	return "oauth"
}

// func init() {
// 	orm.RegisterModel(new(Oauth))
// }
