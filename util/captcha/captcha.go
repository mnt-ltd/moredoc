package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var (
	store              = base64Captcha.DefaultMemStore
	sourceString       = "1234567890qwertyuioplkjhgfdsazxcvbnm"
	CaptchaTypeOptions = "string:字符串\nmath:算术\ndigit:数字\naudio:语音"
)

const (
	CaptchaTypeString  = "string"  // 字符串
	CaptchaTypeDigit   = "digit"   // 数字
	CaptchaTypeMath    = "math"    // 数学公式
	CaptchaTypeChinese = "chinese" // 中文字符
	CaptchaTypeAudio   = "audio"   // 音频
)

// GenerateCaptcha 生成验证码
func GenerateCaptcha(captchaType string, length, width, height int) (id, b64s string, err error) {
	if width <= 0 {
		width = 240
	}
	if height <= 0 {
		height = 60
	}
	var driver base64Captcha.Driver
	switch captchaType {
	case "audio":
		driver = &base64Captcha.DriverAudio{
			Length:   length,
			Language: "zh",
		}
	case "string":
		driver = &base64Captcha.DriverString{
			Height: height,
			Width:  width,
			Source: sourceString,
			Length: length,
		}
	case "math":
		driver = &base64Captcha.DriverMath{
			Height:     height,
			Width:      width,
			NoiseCount: 0,
		}
	default:
		driver = &base64Captcha.DriverDigit{
			Height:   height,
			Width:    width,
			DotCount: 30,
			MaxSkew:  0,
			Length:   length,
		}
	}
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, captchaValue string, clear bool) (ok bool) {
	return store.Verify(id, captchaValue, clear)
}
