package gekent

import (
	"testing"

	"github.com/devgek/go-entities"
	"github.com/stretchr/testify/assert"
)

func createDatastore() entities.EntityDatastore {
	inMemoryDS, err := NewInMemoryDatastore()
	if err != nil {
		panic(err)
	}

	return inMemoryDS
}

func TestGetOneEntityBy(t *testing.T) {
	inMemoryDS := createDatastore()

	var user = User{}
	err := inMemoryDS.GetOneEntityBy(&user, "name", "Lionel")

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, MessiName, user.Name, "Expected", MessiName)
	assert.Equal(t, MessiEmail, user.Email, "Expected", MessiEmail)

	err = inMemoryDS.GetOneEntityBy(&user, "name", "Lionex")
	assert.NotNil(t, err, "Error expected")
	assert.Equal(t, entities.ErrorEntityNotFountBy, err, "ErrorEntityNotFoundBy expected")
}

func TestLoadRelated(t *testing.T) {
	inMemoryDS := createDatastore()

	var contact = Contact{}
	err := inMemoryDS.GetOneEntityBy(&contact, "name", MustermannName)

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, MustermannName, contact.Name, "Expected", MustermannName)
	assert.NotNil(t, contact.ContactAddresses, "Contact should have ContactAddress")
	assert.Equal(t, MustermannStreet, contact.ContactAddresses[0].Street, "Expected", MustermannStreet)

	var contacts []Contact
	err = inMemoryDS.GetAllEntities(&contacts)

	assert.Nil(t, err, "No error expected")
	assert.NotNil(t, contacts, "There should be at least one contact")
	assert.NotNil(t, contacts[0].ContactAddresses, "Contact should have a ContactAddress")
}
func TestGetAllEntities(t *testing.T) {
	inMemoryDS := createDatastore()

	var users []User
	err := inMemoryDS.GetAllEntities(&users)

	assert.Nil(t, err, "No error expected")
	assert.Equal(t, 1, len(users), "Expected %v, but got %v", 2, len(users))
}

func TestCreateEntity(t *testing.T) {
	inMemoryDS := createDatastore()

	roger := &User{Name: "Roger", Pass: "secret", Email: "roger.federer@atp.com", Role: RoleTypeUser}
	err := inMemoryDS.CreateEntity(roger)

	assert.Nil(t, err, "No error expected")
	expectedID := MessiID + 1
	assert.Equal(t, expectedID, roger.ID, "User id not %v", expectedID)
}

func TestSaveEntity(t *testing.T) {
	inMemoryDS := createDatastore()

	var messi = User{}
	err := inMemoryDS.GetOneEntityBy(&messi, "name", "Lionel")

	assert.Nil(t, err, "No error expected")

	oldMessi := messi

	messi.Email = MessiEmail2
	err = inMemoryDS.SaveEntity(&messi)

	assert.Nil(t, err, "No error expected")
	assert.NotEqual(t, oldMessi.Email, messi.Email, "New Email not saved")
	assert.Equal(t, oldMessi.CreatedAt, messi.CreatedAt, "CreatedAt changed")
	assert.NotEqual(t, oldMessi.UpdatedAt, messi.UpdatedAt, "UpdatedAt not saved")
}

func TestDeleteEntityById(t *testing.T) {
	inMemoryDS := createDatastore()

	roger := &User{Name: "Roger", Pass: "secret", Email: "roger.federer@atp.com", Role: RoleTypeUser}
	err := inMemoryDS.CreateEntity(roger)

	assert.Nil(t, err, "No error expected")
	if err = inMemoryDS.DeleteEntityByID(roger, roger.ID); err != nil {
		t.Errorf("Error while deleting entity: %v", err)
	}

	err = inMemoryDS.GetOneEntityBy(&roger, "name", "Roger")
	assert.NotNil(t, err, "Error expected, cause user should be deleted")

	notExistingUser := &User{}
	err = inMemoryDS.DeleteEntityByID(notExistingUser, 99)
	assert.NotNil(t, err, "Error expected")
	assert.Equal(t, entities.ErrorEntityNotDeleted, err, "Expected dedicated error ErrorEntityNotDeleted")
}
