package gekent

import (
	"github.com/devgek/go-entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm for sqlite3
)

// ...
var (
	MessiName        = "Lionel"
	MessiPass        = "Secret00"
	MessiEmail       = "lionel.messi@fcb.com"
	MessiEmail2      = "lm@barcelona.es"
	MessiID          = uint(0)
	MustermannName   = "Mustermann GesmbH"
	MustermannStreet = "Short Street"
)

//NewInMemoryDatastore ...
func NewInMemoryDatastore() (*entities.GormEntityDatastoreImpl, error) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	ds := &entities.GormEntityDatastoreImpl{db}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Contact{}, &ContactAddress{})

	//init the db with test data
	messi := &User{Name: MessiName, Pass: MessiPass, Email: MessiEmail}
	err = ds.CreateEntity(messi)
	if err == nil {
		MessiID = messi.Entity.ID
	}

	contactAddress := &ContactAddress{Street: MustermannStreet, StreetNr: "11", Zip: "3100", City: "St. Pauls"}
	contact := &Contact{OrgType: OrgTypeOrg, Name: MustermannName, NameExt: "Max Mustermann", ContactType: ContactTypeK, ContactAddresses: []ContactAddress{*contactAddress}}
	if err := ds.CreateEntity(contact); err != nil {
		return nil, err
	}

	return ds, nil
}
