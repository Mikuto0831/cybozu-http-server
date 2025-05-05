package input

type GetObjectByID struct {
	ID string
}

type PutObject struct {
	ID   string
	Data []byte
}
