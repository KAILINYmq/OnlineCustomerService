package captcha

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func DriverMathFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverMath = base64Captcha.NewDriverMath(38, 110, 0, 1, &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverMath.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

func DriverStringFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverString = base64Captcha.NewDriverString(38, 110, 0, 2, 4, "234567890abcdefghjkmnpqrstuvwxyz", &color.RGBA{240, 240, 246, 246}, []string{"wqy-microhei.ttc"})
	driver := e.DriverString.ConvertFonts()
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}

func DriverDigitFunc() (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(38, 110, 4, 0.7, 80)
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, store)
	return cap.Generate()
}
