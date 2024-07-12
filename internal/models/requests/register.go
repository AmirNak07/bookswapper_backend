package requests

import "encoding/json"

type RegisterRequest struct {
	Login    string
	Username string
	Password string
	CityId   json.Number
}
