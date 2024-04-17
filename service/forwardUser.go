package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nautilus/gateway"
)

var forwardUser = gateway.RequestMiddleware(func(r *http.Request) error {
	// the initial context of the request is set as the same context
	// provided by net/http

	// we are safe to extract the value we saved in context and set it as the outbound header
	user, _ := r.Context().Value("user").(UserToken)
	create, _ := json.Marshal(user.Create)
	read, _ := json.Marshal(user.Read)
	update, _ := json.Marshal(user.Update)
	delete, _ := json.Marshal(user.Delete)
	r.Header.Set("Accountid", user.AccountID)
	r.Header.Set("Id", user.ID)
	r.Header.Set("Email", user.Email)
	r.Header.Set("Cobrowsing", strconv.FormatBool(user.CoBrowsing))
	r.Header.Set("Screenshare", strconv.FormatBool(user.ScreenShare))
	r.Header.Set("Audiodownload", strconv.FormatBool(user.AudioDownload))
	r.Header.Set("Videodownload", strconv.FormatBool(user.VideoDownload))
	r.Header.Set("Username", user.Username)
	r.Header.Set("Create", string(create))
	r.Header.Set("Read", string(read))
	r.Header.Set("Update", string(update))
	r.Header.Set("Delete", string(delete))
	// return the modified request
	return nil
})
