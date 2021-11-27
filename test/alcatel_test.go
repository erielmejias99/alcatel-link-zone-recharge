package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
	"testing"
)

func TestAlcatelSendUssd(t *testing.T) {
	a := alcatel.NewAlcatel()
	content, err  := a.SendUssd("*222#" )
	if err != nil {
		t.Log( err )
	}else{
		t.Log( content )
	}
}
