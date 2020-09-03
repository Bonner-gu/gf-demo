package validation

import "gf_demo_api/library/ecode"

func CheckFileFormat(key string) error {
	switch key {
	case ".png":
		return nil
	case ".jpg":
		return nil
	case ".jpeg":
		return nil
	case ".pdf":
		return nil
	case ".apk":
		return nil
	case ".ipa":
		return nil
	}
	return ecode.UploadFileErr
}
