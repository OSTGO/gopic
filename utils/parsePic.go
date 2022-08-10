package utils

func CheckImageType(b []byte) string {
	if checkType(b, []byte{0xff, 0xd8}) {
		return "jpg"
	}
	if checkType(b, []byte{0x00, 0x00, 0x02, 0x00, 0x00}) {
		return "tga"
	}
	if checkType(b, []byte{0x00, 0x00, 0x10, 0x00, 0x00}) {
		return "tga"
	}
	if checkType(b, []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}) {
		return "png"
	}
	if checkType(b, []byte{'G', 'I', 'F', '8', '9', 'a'}) {
		return "gif"
	}
	if checkType(b, []byte{'G', 'I', 'F', '8', '7', 'a'}) {
		return "gif"
	}
	if checkType(b, []byte{'B', 'M'}) {
		return "bmp"
	}
	if checkType(b, []byte{0x0a}) {
		return "pcx"
	}
	if checkType(b, []byte{0x4D, 0x4D}) {
		return "tiff"
	}
	if checkType(b, []byte{0x49, 0x49}) {
		return "tiff"
	}
	if checkType(b, []byte{0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x20, 0x20}) {
		return "ico"
	}
	if checkType(b, []byte{0x00, 0x00, 0x02, 0x00, 0x01, 0x00, 0x20, 0x20}) {
		return "cur"
	}
	if checkType(b, []byte{'F', 'O', 'R', 'M'}) {
		return "iff"
	}
	if checkType(b, []byte{'R', 'I', 'F', 'F'}) {
		return "ani"
	}
	if checkType(b, []byte("<!DOCTYPE svg")) {
		return "svg"
	}
	return "unknown"
}

func checkType(image, types []byte) bool {
	if len(types) >= len(image) {
		return false
	}
	for i, v := range types {
		if image[i] != v {
			return false
		}
	}
	return true
}
