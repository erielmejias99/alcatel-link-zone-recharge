package main

import "github.com/alcatel-link-zone/internal/alcatel"

func main(){

	alcatelClient := alcatel.NewAlcatel()
	alcatelClient.SendUssd("*222#")

}