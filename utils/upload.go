package utils

var StroageMap = make(map[string]Upload, 0)

var StroageHelp = make(map[string]string, 0)

type Upload interface {
	Upload(image *Image) (string, error)
	SetPicList([]string, bool)
}

func GetStringUploadMapKey(m map[string]Upload) []string {
	l := make([]string, 0, len(m))
	for k := range m {
		l = append(l, k)
	}
	return l
}

func GetStringStringMapKey(m map[string]string) []string {
	l := make([]string, 0, len(m))
	for k := range m {
		l = append(l, k)
	}
	return l
}
