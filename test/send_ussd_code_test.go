package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/request"
	"log"
	"testing"
)

func TestSendUssdCode(t *testing.T) {
	a := &alcatel.Alcatel{}
	//loginResp,err := a.Login("admin")
	//t.Logf( "%#v", loginResp)

	response, err := a.SendUssdCode( "*133#", request.Default )
	if err != nil {
		t.Error("Error sending the code" )
	}
	log.Printf( "%#v", response )
}