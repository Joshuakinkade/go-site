package photos

type ILibrary interface {
	SavePhoto(name string, data []byte) error
	LoadPhoto(name string) ([]byte, error)
}

type Library struct{}

func NewLibrary() Library {
	return Library{}
}

// SavePhoto saves a photo to the storage.
func (p Library) SavePhoto(name string, data []byte) error {
	return nil
}

// Load a photo from storage
func (p Library) LoadPhoto(name string) ([]byte, error) {
	return nil, nil
}
