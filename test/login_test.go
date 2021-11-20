package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
	"log"
	"testing"
)

func TestLoginCorrectPassowrd(t *testing.T) {

	alcatelClient := alcatel.Alcatel{}
	loginResponse, err := alcatelClient.Login("admin")
	if( err != nil ){
		t.Errorf( "Error login %s", err.Error() )
	}else{
		log.Printf( "Login response %s", loginResponse )
	}

}

func TestLoginInvalidPassowrd(t *testing.T) {

	alcatelClient := alcatel.Alcatel{}
	loginResponse, err := alcatelClient.Login("asdfasdfcaeggfsdf")
	if _, ok := err.(*request_and_response.Error); ok {
		log.Printf( "Wrong Password OK: %s", err )
	}else if err != nil {
		t.Errorf( "Error login %s", err.Error() )
	}else{
		log.Printf( "Already logged in: %s", loginResponse )
	}
}