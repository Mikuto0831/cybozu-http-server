package output

import "github.com/mikuto0831/cybozu-http-server/internal/domain/entity"

type GetObjectByID struct {
	Object *entity.Object `json:"object"`
}
