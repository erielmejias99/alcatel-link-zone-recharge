package alcatel

import "testing"

func TestSendCode(t *testing.T) {
	a := NewAlcatel()
	resp, err := a.SendUssd("*222#" )
	if err != nil {
		t.Error(err)
	}else{
		_, err := a.CancelUssd()
		if err != nil {
			t.Error("Error cancelling the ussdRequest")
		}
		t.Log(resp)
	}
}

func TestSend2DiffUssdCode(t *testing.T) {
	a := NewAlcatel()
	resp, err := a.SendUssd("*133#" )
	if err != nil {
		t.Error(err)
	}else{
		t.Log(resp)
		_, err := a.CancelUssd()
		if err != nil {
			t.Error("Error cancelling the ussdRequest")
		}

		resp2, err := a.SendUssd( "*222#" )
		if err != nil {
			t.Error(err)
		}
		t.Log(resp2)

		if resp == resp2{
			t.Error("Responses must be different")
		}
	}
}

func TestResponseRequiredUssd(t *testing.T) {
	a := NewAlcatel()
	resp, err := a.SendUssd("*133#" )
	if err != nil {
		t.Error(err)
	}else{
		t.Log(resp)

		resp2, err := a.SendUssd( "2" )
		if err != nil {
			t.Error(err)
		}

		_, err = a.CancelUssd()
		if err != nil {
			t.Error("Error cancelling the ussdRequest")
		}

		t.Log(resp2)

		if resp == resp2{
			t.Error("Responses must be different")
		}
	}
}