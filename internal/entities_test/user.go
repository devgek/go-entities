package gekent_test

import (
	"github.com/devgek/go-entities"
)

//User ...
type User struct {
	gekentities.Entity `entity:"type:User;name:user"`
	Name               string   `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Pass            string   `gorm:"type:text;not null" form:"gkvPass"`
	Email           string   `gorm:"type:varchar(100);not null" form:"gkvEmail"`
	Role            RoleType `gorm:"type:integer;not null" form:"gkvRole"`
}

//BuildEntityOption ...
func (u User) BuildEntityOption() gekentities.EntityOption {
	o := gekentities.EntityOption{}
	o.ID = u.Entity.ID
	o.Value = u.Name + ":" + u.Email

	return o
}
