package alcatel

import (
	"fmt"
	"time"
)

type Alcatel struct {
	ussd * Ussd
	system * System
}

func NewAlcatel() *Alcatel {
	return &Alcatel{
		ussd: &Ussd{ ussdType: 1 },
		system: &System{},
	}
}

func (a Alcatel) SendUssd(code string) (resp string) {

	//Check the network status
	//networkStatus, err := a.system.GetSystemStatus()
	//if err != nil {
	//	return err.Error()
	//}

	// if the mobile is in 4G or 4G+ change the network in order to make possible run USSD codes
	//networkChanged := false
	//if networkStatus.Result.NetworkType == response.Net4G || networkStatus.Result.NetworkType == response.Net3GPlus{
	//	networkChanged = true
	//}

	// Send code to the Network
	_, err := a.ussd.sendUSSD( code )
	if err != nil {
		return err.Error()
	}

	time.Sleep( 5 * time.Second )
	getUssdResponse, err := a.ussd.getUssdResponse()
	if err != nil {
		return err.Error()
	}

	switch getUssdResponse.Result.SendState {
	case 2:
		return getUssdResponse.Result.UssdContent
	case 3:
		return "Error, try setting Network Mode to 3G or Automatic"
	default:
		return fmt.Sprintf( "An error ocurred, code send State: %d", getUssdResponse.Result.SendState )
	}

}