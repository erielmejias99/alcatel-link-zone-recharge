package response

import "fmt"

type LoginResponse struct {
	Token string `json:"Token"`
}

func (l LoginResponse) String() string {
	return fmt.Sprintf("{ token: %s }", l.Token )
}