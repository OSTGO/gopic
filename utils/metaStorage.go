package utils

type MetaStorage struct {
	B *BaseStorage
}

func (m *MetaStorage) SetPicList(paths []string, nameReserve bool) {
	m.B = NewBaseStorage(paths, nameReserve)
}

func NewMetaStorage() *MetaStorage {
	return &MetaStorage{&BaseStorage{}}
}
