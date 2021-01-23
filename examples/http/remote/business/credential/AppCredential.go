package credential

import (
	"context"
)

// 应用凭证
type AppCredential struct {
}

func (c AppCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"X-App-Id":     "kit-scaffold.palm.http.api", // TODO: 应用ID
		"X-App-Secret": "e1ea87euy763lo909721ea",     // TODO: 秘钥
	}, nil
}

func (c AppCredential) RequireTransportSecurity() bool {
	return false
}
