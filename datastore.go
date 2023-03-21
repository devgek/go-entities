package gekentities

import (
	"database/sql"
)

//EntityDatastore CRUD operations with abstract entity type
type EntityDatastore interface {
	GetOneEntityBy(entity interface{}, key string, val interface{}) error
	GetEntityByID(entity interface{}, id uint) error
	GetAllEntities(entitySlice interface{}) error
	CreateEntity(entity interface{}) error
	SaveEntity(entity interface{}) error
	DeleteEntityByID(entity interface{}, id uint) error
	GetDB() *sql.DB
}
