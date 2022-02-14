package gekent

import (
	"github.com/devgek/go-entities"
	"github.com/jinzhu/gorm"
)

//Contact ...
type Contact struct {
	entities.Entity  `entity:"type:Contact;name:contact"`
	OrgType          OrgType     `gorm:"type:integer;not null" form:"gkvOrgType"`
	Name             string      `gorm:"type:varchar(100);not null" form:"gkvName"`
	NameExt          string      `gorm:"type:varchar(100)" form:"gkvNameExt"`
	ContactType      ContactType `gorm:"type:integer;not null" form:"gkvContactType"`
	ContactAddresses []ContactAddress
}

//BuildEntityOption ...
func (c Contact) BuildEntityOption() entities.EntityOption {
	o := entities.EntityOption{}
	o.ID = c.Entity.ID
	o.Value = c.Name

	return o
}

//LoadRelated load related entities (implements EntityHolder)
func (c *Contact) LoadRelated(db *gorm.DB) error {
	c.ContactAddresses = []ContactAddress{}
	db.Model(c).Related(&c.ContactAddresses)

	return nil
}
