package handler

import "github.com/rs/xid"

type CookieDTO struct {
	ID       xid.ID  `json:"id"`
	Title    string  `json:"title"`
	Category *string `json:"category:omitempty"`
	Fortune  string  `json:"fortune"`
}

type CreateRequestDTO struct {
	Title    string  `json:"title"`
	Category *string `json:"category"`
}

type CreateResponseDTO struct {
	Cookie CookieDTO `json:"cookie"`
}

type ListRequestDTO struct {
}

type ListResponseDTO struct {
	Cookies []CookieDTO `json:"cookies"`
}

type DeleteRequestDTO struct {
	ID xid.ID `json:"id"`
}

type DeleteResponseDTO struct {
}

type ModifyRequestDTO struct {
	ID       xid.ID  `json:"id"`
	Title    string  `json:"title"`
	Category *string `json:"category"`
}

type ModifyResponseDTO struct {
	Cookie CookieDTO `json:"cookie"`
}
