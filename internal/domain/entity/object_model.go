package entity

type Object struct {
	ID   string `json:"id"`		// id: オブジェクトのID(path)
	Data []byte `json:"data"`	// data: オブジェクトのデータ
}