package fate_go_sdk

import (
	pb "github.com/zm-dev/fate-go-sdk/pb"
	"time"
	"context"
)

type Cache interface {
	Get(k string) (interface{}, bool)
	Set(k string, x interface{}, d time.Duration)
}

type AccessTokenClient struct {
	appID     int32
	appSecret string
	atsc      pb.AccessTokenServiceClient
	timeout   time.Duration
	cache     Cache
}

const (
	accessTokenCacheKey        = "accessTokenCacheKey"
	defaultAccessTokenCacheTTL = 2 * time.Hour
)

// From cache
func (atc *AccessTokenClient) GetToken() (string, error) {
	var (
		ati interface{}
		ok  bool
	)
	if ati, ok = atc.cache.Get(accessTokenCacheKey); !ok {
		at, err := atc.RequestToken()

		if err != nil {
			return "", err
		}
		d := at.ExpiredAt - time.Now().Unix()
		var ttl time.Duration
		if d <= 0 {
			ttl = defaultAccessTokenCacheTTL
		} else {
			ttl = time.Duration(d) * time.Second
		}
		ati = at.Token
		atc.cache.Set(accessTokenCacheKey, ati, ttl)
	}

	if atStr, ok := ati.(string); ok {
		return atStr, nil
	} else {
		return "", nil
	}
}

func (atc *AccessTokenClient) RequestToken() (*pb.AccessToken, error) {
	ctx, _ := context.WithTimeout(context.Background(), atc.timeout)
	return atc.atsc.Token(ctx, &pb.Credential{AppId: atc.appID, AppSecret: atc.appSecret})
}

func NewAccessTokenClient(appID int32, appSecret string, atsc pb.AccessTokenServiceClient,
	timeout time.Duration, cache Cache) *AccessTokenClient {
	return &AccessTokenClient{appID: appID, appSecret: appSecret, atsc: atsc, timeout: timeout, cache: cache}
}
