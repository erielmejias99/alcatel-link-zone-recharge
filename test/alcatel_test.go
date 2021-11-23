package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
	"testing"
)

func TestAlcatelSendUssd(t *testing.T) {
	a := alcatel.NewAlcatel()
	content := a.SendUssd("*222#" )
	t.Log( content )
}
