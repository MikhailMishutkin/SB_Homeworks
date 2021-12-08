package handlers

import "github.com/go-chi/chi"

type Handler interface {
	Register(r chi.Router)
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Register(r chi.Router) {
	r.Post("/create", h.NewUserProfile)     // 201, 4хх, Header Location: url
	r.Get(friendsListURL, h.GetFriendsList) // 200, 404, 500
	r.Post("/make_friends", h.MakeFriends)  // 200
	r.Delete(usersURL, h.DeleteUserProfile) //200/204
	r.Put(userURL, h.UpdateAge)             // 200, 4xx, 500
}
