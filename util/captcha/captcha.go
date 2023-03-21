package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var (
	store = base64Captcha.DefaultMemStore
	// sourceChinese      = strings.Join(strings.Split("欢迎使用由深圳市摩枫网络科技有限公司基于阿帕奇开源协议的魔豆文库系统", ""), ",")
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
			Height:          height,
			Width:           width,
			Source:          sourceString,
			ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,
			Length:          length,
		}
	case "math":
		driver = &base64Captcha.DriverMath{
			Height:     height,
			Width:      width,
			NoiseCount: 0,
		}
	// case "chinese":
	// 	driver = base64Captcha.NewDriverChinese(
	// 		height,
	// 		width,
	// 		0,
	// 		0,
	// 		4,
	// 		sourceChinese,
	// 		nil,
	// 		nil,
	// 		[]string{"wqy-microhei.ttc"},
	// 	).ConvertFonts()
	default:
		driver = &base64Captcha.DriverDigit{
			Height:   height,
			Width:    width,
			DotCount: 80,
			MaxSkew:  1,
			Length:   length,
		}
	}
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, captchaValue string) (ok bool) {
	return store.Verify(id, captchaValue, true)
}
