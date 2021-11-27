package main

import (
	"errors"
	"fmt"
	"github.com/alcatel-link-zone/internal/alcatel"
)

type Application interface {
	Run()
}

type App struct{
	*alcatel.Alcatel
}

var app * App
func GetApp() Application{
	if app == nil{
		app = &App{}
		app.Alcatel = alcatel.NewAlcatel()
	}
	return app
}

func (a * App) Run()  {
	a.initGlobalMenu()
}

func (a * App) initGlobalMenu(){
	fmt.Println("Menu: " )
	fmt.Println("\t0-Exit" )
	fmt.Println("\t1-Send ussd" )

	a.requestInformation("Option")
	option := 0
	fmt.Scanf("%d", &option )

	switch option {
	case 0:
		fmt.Print("Thanks!")
	case 1:
		a.sendUssd()
	default:
		a.printError( errors.New("invalid option enter again") )
		a.initGlobalMenu()
	}
}

func (a *App) requestInformation(message string)  {
	fmt.Printf("\n%s: ", message)
}

func (a *App) printError(err error)  {
	fmt.Println("-------------------------------------------------------")
	fmt.Printf(" ! %s\n", err.Error() )
	fmt.Println("-------------------------------------------------------")
}

func (a *App) printMessage(message string)  {
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%s\n", message )
	fmt.Println("-------------------------------------------------------")
}

func (a *App) sendUssd()  {
	var code string
	a.requestInformation("USSD Code")
	_, err := fmt.Scanf("%s", &code )
	if err != nil {
		a.printError(err)
	}else{
		resp, err := a.Alcatel.SendUssd(code)
		if err != nil {
			a.printError(err)
		}else{
			a.printMessage( resp )
			a.printUssdSubmenu()
		}
	}
}

func (a *App) printUssdSubmenu()  {
	fmt.Println(" 1 -> Replay")
	fmt.Println(" 2 -> Cancel")

	a.requestInformation("Option")
	option := 0
	fmt.Scanf("%d", &option )

	switch option {
	case 1:
		a.sendUssd();
	default:
		_, err := a.CancelUssd()
		if err != nil {
			a.printError(err)
		}
		a.initGlobalMenu()
	}

}