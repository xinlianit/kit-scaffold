package common

import "github.com/xinlianit/go-util/util"

// 生成签名
// @param digest 摘要
// @param secret 签名秘钥
func GenerateSign(digest string, secret string) string {
	return util.CryptoUtil().Md5(digest + secret)
}

// 验证签名
// @param digest 摘要
// @param sign 签名串
// @param secret 签名秘钥
func VerifySign(digest string, sign string, secret string) bool {
	if GenerateSign(digest, secret) == sign {
		return true
	}

	return false
}
