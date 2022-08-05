package utils

type MetaStorage struct {
	B *BaseStorage
}

func (m *MetaStorage) SetPicList(paths []string) {
	m.B = NewBaseStorage(paths)
}

func NewMetaStorage() *MetaStorage {
	return &MetaStorage{&BaseStorage{}}
}
