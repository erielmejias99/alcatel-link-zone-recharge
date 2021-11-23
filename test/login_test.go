package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
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

func TestGetLoginStateLoggedOut(t *testing.T) {
	alcatelClient := &alcatel.Alcatel{}
	loginState,err := alcatelClient.GetLoginState()
	if err != nil {
		t.Errorf( "%#v", loginState )
	}

	if loginState != nil && loginState.Result.State != response.LoggedOut {
		t.Error("User should be logged out.")
	}
}

func TestGetLoginStateLoggedIn(t *testing.T) {
	alcatelClient := &alcatel.Alcatel{}

	login, err := alcatelClient.Login("admin" )
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", login )

	loginState,err := alcatelClient.GetLoginState()
	if err != nil {
		t.Errorf( "%#v", loginState )
	}

	t.Logf("loginState: %#v", loginState )

	if loginState != nil && loginState.Result.State != response.LoggedIn {
		t.Error("User should be logged out.")
	}
}

func TestLogout(t *testing.T) {
	a := alcatel.Alcatel{}

	_, err := a.Login("admin")
	if err != nil {
		t.Error("Error in login.")
	}
	loggedOut, err := a.Logout()
	if err != nil || !loggedOut {
		t.Error("Should be logged out")
	}
}