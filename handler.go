package fate_go_sdk

import (
	"net/http"
	"github.com/zm-dev/go-httputils"
	"strings"
	"strconv"
	"time"
)

func (f *Fate) CallbackHandler() httputils.AppHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) httputils.HTTPError {
		ticketID := strings.TrimSpace(r.FormValue("ticket_id"))
		if ticketID == "" {
			return httputils.BadRequest("ticket_id 参数不存在")
		}
		userID := strings.TrimSpace(r.FormValue("user_id"))
		if userID == "" {
			return httputils.BadRequest("user_id 参数不存在")
		}

		expiredAt, err := strconv.ParseInt(r.FormValue("expired_at"), 10, 64)
		if err != nil {
			expiredAt = time.Now().Unix() + 3600*24*7
		}
		callback := strings.TrimSpace(r.FormValue("callback"))
		if callback == "" {
			callback = r.Host
		}

		http.SetCookie(w, &http.Cookie{
			Name:     f.opts.TicketIdCookieKey,
			Value:    ticketID,
			Expires:  time.Unix(expiredAt, 0),
			HttpOnly: true,
			Path:     "/",
		})
		http.SetCookie(w, &http.Cookie{
			Name:     f.opts.UserIdCookieKey,
			Value:    userID,
			Expires:  time.Unix(expiredAt, 0),
			HttpOnly: false,
			Path:     "/",
		})
		http.Redirect(w, r, callback, http.StatusFound)
		return nil
	}
}

func (f *Fate) LogoutURL() string {
	return f.opts.FateUrl + "/logout"
}
