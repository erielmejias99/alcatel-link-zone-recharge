package test

import (
	"github.com/alcatel-link-zone/internal/alcatel"
	"testing"
)

func TestGetSystemStatus(t *testing.T) {
	system := alcatel.System{}
	res, err := system.GetSystemStatus()
	if err != nil {
		t.Error( err )
	}else{
		t.Logf("%#v", res )
	}
}