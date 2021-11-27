package alcatel

import (
	"errors"
	"fmt"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"time"
)

type Alcatel struct {
	ussd * Ussd
	system * System
	connection * Connection
	networkChanged bool
}

func NewAlcatel() *Alcatel {
	return &Alcatel{
		ussd: &Ussd{ ussdType: 1 },
		system: &System{},
		connection: &Connection{},
	}
}

func (a *Alcatel) SendUssd(code string) (resp string, err error) {

	//Check the network status
	connectionStatus, err := a.connection.GetConnectionState()
	if err != nil {
		return "", err
	}

	if connectionStatus.Result.ConnectionStatus != response.Connected {
		return "", errors.New("device must be connected" )
	}

	// Send code to the Network
	_, err = a.ussd.sendUSSD( code )
	if err != nil {
		return "", err
	}

	time.Sleep( 5 * time.Second )
	getUssdResponse, err := a.ussd.getUssdResponse()
	if err != nil {
		return "", err
	}

	switch getUssdResponse.Result.SendState {
	case 1:
		if getUssdResponse.Result.UssdContentLen == 0 {
			return "", fmt.Errorf( "An error ocurred, code send State: %d", getUssdResponse.Result.SendState )
		}
		fallthrough
	case 2:
		return getUssdResponse.Result.UssdContent, nil
	case 3:
		return "", errors.New( "Error, try setting Network Mode to 3G or Automatic" )
	default:
		return "", fmt.Errorf( "An error ocurred, code send State: %d", getUssdResponse.Result.SendState )
	}
}

func (a *Alcatel) CancelUssd() ( cancelled bool, err error ) {

	err = a.ussd.SetUssdEnd()
	if err != nil {
		return false, err
	}

	return true, nil
}