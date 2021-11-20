package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
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