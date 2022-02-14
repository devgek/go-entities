package gekent

import "github.com/devgek/go-entities"

//ContactAddress ...
type ContactAddress struct {
	entities.Entity `entity:"type:ContactAddress;name:contactaddress"`
	ContactID       uint
	Street          string `gorm:"type:varchar(100);not null"`
	StreetNr        string `gorm:"type:varchar(10);not null"`
	StreetExt       string `gorm:"type:varchar(50)"`
	Zip             string `gorm:"type:varchar(10);not null"`
	City            string `gorm:"type:varchar(100);not null"`
}
