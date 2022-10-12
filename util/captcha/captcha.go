package captcha

import (
	"strings"

	"github.com/mojocn/base64Captcha"
)

var (
	store         = base64Captcha.DefaultMemStore
	width         = 240
	height        = 60
	sourceChinese = strings.Join(strings.Split("欢迎使用由深圳市摩枫网络科技有限公司基于阿帕奇开源协议的魔刀文库系统", ""), ",")
	sourceString  = "1234567890qwertyuioplkjhgfdsazxcvbnm"
)

const (
	CaptchaTypeString  = "string"  // 字符串
	CaptchaTypeDigit   = "digit"   // 数字
	CaptchaTypeMath    = "math"    // 数学公式
	CaptchaTypeChinese = "chinese" // 中文字符
	CaptchaTypeAudio   = "audio"   // 音频
)

// GenerateCaptcha 生成验证码
func GenerateCaptcha(captchaType string) (id, b64s string, err error) {
	var driver base64Captcha.Driver
	switch captchaType {
	case "audio":
		driver = &base64Captcha.DriverAudio{
			Length:   6,
			Language: "zh",
		}
	case "string":
		driver = &base64Captcha.DriverString{
			Height:          height,
			Width:           width,
			Source:          sourceString,
			ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowSineLine,
			Length:          6,
		}
	case "math":
		driver = &base64Captcha.DriverMath{
			Height:     height,
			Width:      width,
			NoiseCount: 0,
		}
	case "chinese":
		driver = &base64Captcha.DriverChinese{
			Height: height,
			Width:  width,
			Source: sourceChinese,
			Length: 4, // 4个字符
			Fonts:  []string{"wqy-microhei.ttc"},
		}
	default:
		driver = &base64Captcha.DriverDigit{
			Height:   height,
			Width:    width,
			DotCount: 80,
			MaxSkew:  1,
			Length:   6,
		}
	}
	return base64Captcha.NewCaptcha(driver, store).Generate()
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(id string, captchaValue string) (ok bool) {
	return store.Verify(id, captchaValue, true)
}
