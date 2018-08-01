package fate_go_sdk

import (
	"github.com/zm-dev/go-httputils"
	"net/http"
)

// 这个中间件用来保存 LoginCheckRes
func (f *Fate) CheckResMiddleware() httputils.APPMiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, next httputils.AppHandleFunc) httputils.HTTPError {
		req, err := f.getLoginCheckResContext(r)
		if err != nil {
			return httputils.InternalServerError("login check 失败").WithError(err)
		}
		return next(w, req)
	}
}

func (f *Fate) AuthMiddleware() httputils.APPMiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request, next httputils.AppHandleFunc) httputils.HTTPError {
		r, err := f.getLoginCheckResContext(r)
		if err != nil {
			return httputils.InternalServerError("login check 失败").WithError(err)
		}
		if !GetIsLogin(r.Context()) {
			if httputils.ExpectsJson(r) {
				return httputils.Unauthorized("Unauthorized")
			} else {
				f.Login(w, r)
			}
			return nil
		}
		return next(w, r)
	}
}
