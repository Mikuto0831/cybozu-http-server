package repo

import "github.com/mikuto0831/cybozu-http-server/internal/domain/entity"

type IObjectsRepo interface {
	GetObjectByID(id string) (*entity.Object, error)
	PutObject(object entity.Object) error
}