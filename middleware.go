package fate_go_sdk

import (
	"github.com/zm-dev/go-httputils"
	"net/http"
)

// 这个中间件用来保存 LoginCheckRes
func (f *Fate) CheckResMiddleware() httputils.APPMiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, next httputils.AppHandleFunc) httputils.HTTPError {
		if res, err := f.Check(r); err != nil {
			return httputils.InternalServerError("login check 失败").WithError(err)
		} else {
			return next(w, r.WithContext(newLoginCheckResContext(r.Context(), res)))
		}
	}
}

func (f *Fate) AuthMiddleware() httputils.APPMiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, next httputils.AppHandleFunc) httputils.HTTPError {
		if !GetIsLogin(r.Context()) {
			f.Login(w, r)
			return nil
		}
		return next(w, r)
	}
}
