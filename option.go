package fate_go_sdk

import (
	"time"
	"github.com/patrickmn/go-cache"
)

type options struct {
	FateUrl, FateRpcHostname                                                   string
	AppID                                                                      int32
	AppSecret, AccessTokenKey, TicketIdCookieKey, UserIdCookieKey, LoginMethod string
	RpcTimeout                                                                 time.Duration
	Cache                                                                      Cache
}

type Option func(*options)

var defaultOptions = options{
	AccessTokenKey:    "access_token",
	TicketIdCookieKey: "ticket_id",
	UserIdCookieKey:   "user_id",
	LoginMethod:       "redirect",
	RpcTimeout:        3 * time.Second,
	Cache:             cache.New(10*time.Minute, 30*time.Minute),
}

func FateUrlOption(fateUrl string) Option {
	return func(o *options) {
		o.FateUrl = fateUrl
	}
}

func FateRpcHostnameOption(fateRpcHostname string) Option {
	return func(o *options) {
		o.FateRpcHostname = fateRpcHostname
	}
}

func AppIDOption(appID int32) Option {
	return func(o *options) {
		o.AppID = appID
	}
}

func AppSecretOption(appSecret string) Option {
	return func(o *options) {
		o.AppSecret = appSecret
	}
}

func AccessTokenKeyOption(accessTokenKey string) Option {
	return func(o *options) {
		o.AccessTokenKey = accessTokenKey
	}
}

func TicketIdCookieKeyOption(ticketIdCookieKey string) Option {
	return func(o *options) {
		o.TicketIdCookieKey = ticketIdCookieKey
	}
}

func UserIdCookieKeyOption(userIdCookieKey string) Option {
	return func(o *options) {
		o.UserIdCookieKey = userIdCookieKey
	}
}

func LoginMethodOption(loginMethod string) Option {
	return func(o *options) {
		o.LoginMethod = loginMethod
	}
}

func RpcTimeoutOption(rpcTimeout time.Duration) Option {
	return func(o *options) {
		o.RpcTimeout = rpcTimeout
	}
}

func CacheOption(c Cache) Option {
	return func(o *options) {
		o.Cache = c
	}
}
