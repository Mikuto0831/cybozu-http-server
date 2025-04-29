package repo

import (
	"fmt"
	"sync"

	"github.com/mikuto0831/cybozu-http-server/internal/domain/entity"
)

type ObjectsRepo struct {
	objects map[string]*entity.Object
	mu      sync.RWMutex
}

func NewObjectsRepo() *ObjectsRepo {
	return &ObjectsRepo{
		objects: make(map[string]*entity.Object),
	}
}

func (r *ObjectsRepo) GetObjectByID(id string) (*entity.Object, error) {
	// 排他制御(Read Lock)
	r.mu.RLock()
	defer r.mu.RUnlock()

	object, ok := r.objects[id]

	if !ok {
		return nil, fmt.Errorf("%s object not found", id)
	}

	return object, nil
}

func (r *ObjectsRepo) PutObject(object entity.Object) error {
	// 排他制御(Write Lock)
	r.mu.Lock()
	defer r.mu.Unlock()

	r.objects[object.ID] = &object

	return nil
}
