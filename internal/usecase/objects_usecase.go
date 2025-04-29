package usecase

import (
	"github.com/mikuto0831/cybozu-http-server/internal/domain/entity"
	"github.com/mikuto0831/cybozu-http-server/internal/repo"
)

type IObjectsUsecase interface {
	GetObjectByID(id string) (*entity.Object, error)
	PutObject(id string, data []byte) error
}

type ObjectsUsecase struct {
	objectsRepo *repo.ObjectsRepo
}

func NewObjectsUsecase() *ObjectsUsecase {
	return &ObjectsUsecase{
		objectsRepo: repo.NewObjectsRepo(),
	}
}

// ObjectsからIDを使用して1つ取得する
func (u *ObjectsUsecase) GetObjectByID(id string) (*entity.Object, error) {
	return u.objectsRepo.GetObjectByID(id)
}

// ObjectsにIDとデータを使用して1つ追加する
func (u *ObjectsUsecase) PutObject(id string, data []byte) error {

	// IDとデータを使用してオブジェクトを作成
	object := entity.Object{
		ID:   id,
		Data: data,
	}

	return u.objectsRepo.PutObject(object)
}
