package alcatel

import (
	"errors"
	"fmt"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/request"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"time"
)

type Alcatel struct {
	ussd * Ussd
	system * System
	connection * Connection
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
	networkStatus, err := a.system.GetSystemStatus()
	if err != nil {
		return "", err
	}

	// if the mobile is in 4G or 4G+ change the network in order to make possible run USSD codes
	//networkChanged := false
	if networkStatus.Result.NetworkType == response.Net4G || networkStatus.Result.NetworkType == response.Net4GPlus{

		//if the device is connected it must be disconnected to change the network mode
		connectionStatus, err := a.connection.GetConnectionState()
		if err != nil {
			return "", err
		}

		if connectionStatus.Result.ConnectionStatus == response.Connected {
			changed, err := a.changeNetwork( request.Automatic )
			if err != nil {
				_, err = a.connection.connect()
				return "", err
			}
			if !changed {
				//try to reconnect
				return "", errors.New("error changing the network")
			}
		}
		//networkChanged = true
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

	//defer func() {
	//	a.connection.setNetworkType( request.Net4G )
	//	a.connection.connect()
	//}()

	switch getUssdResponse.Result.SendState {
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

func (a *Alcatel) changeNetwork( networkType request.SetNetworkType ) ( changed bool, err error ){
	disconnected, err := a.connection.disconnect()
	if err != nil {
		return false, err
	}
	if !disconnected{
		return false, errors.New("error trying to disconnect from the network")
	}

	// Set network to Automatic
	err = a.connection.setNetworkType( networkType )
	if err != nil {
		return false, errors.New("unable to change the network mode: " + err.Error() )
	}

	connected, err := a.connection.connect()
	if !connected {
		return true, err
	}
	return
}