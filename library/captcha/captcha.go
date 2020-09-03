package captcha

import (
	"time"

	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type ConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store base64Captcha.Store

func init() {
	store = base64Captcha.NewMemoryStore(10240, 15*time.Minute)
}

// base64Captcha create http handler
func GenerateCaptchaHandler(config *ConfigJsonBody) (error, map[string]interface{}) {
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch config.CaptchaType {
	case "audio":
		driver = config.DriverAudio
	case "string":
		driver = config.DriverString.ConvertFonts()
	case "math":
		driver = config.DriverMath.ConvertFonts()
	case "chinese":
		driver = config.DriverChinese.ConvertFonts()
	default:
		driver = config.DriverDigit
	}

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()

	body := map[string]interface{}{"b64s": b64s, "captchaId": id}
	if err != nil {
		body = map[string]interface{}{"b64s": "", "captchaId": ""}
	}
	return err, body
}

// 校验验证码
func CaptchaVerifyHandle(id string, value string) bool {
	return store.Verify(id, value, true)
}
