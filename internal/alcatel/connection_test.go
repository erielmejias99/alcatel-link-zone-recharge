package alcatel

import "testing"

func TestConnectionStatus(t *testing.T) {
	c := &Connection{}
	resp, err := c.GetConnectionState()
	if err != nil {
		t.Error( err )
	}else{
		t.Log( resp )
	}
}