package data

import "encoding/json"

func (d *LoginDetails) String() string {
	if j, err := json.Marshal(d); err == nil {
		return string(j)
	}
	return "could not convert to string"
}
