package service

import (
	"encoding/json"
	"github.com/nautilus/gateway"
	"net/http"
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
	r.Header.Set("Role", user.UserType)
	r.Header.Set("Email", user.Email)
	r.Header.Set("Create", string(create))
	r.Header.Set("Read", string(read))
	r.Header.Set("Update", string(update))
	r.Header.Set("Delete", string(delete))
	// return the modified request
	return nil
})
