package captcha

import (
	"github.com/mojocn/base64Captcha"
	"testing"
)

func TestGenerateCaptchaHandler(t *testing.T) {
	err, rsp := GenerateCaptchaHandler(&ConfigJsonBody{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			Height:          30,
			Width:           180,
			NoiseCount:      3,
			ShowLineOptions: 3,
		},
	})
	t.Log(err)
	t.Log(rsp)
}
