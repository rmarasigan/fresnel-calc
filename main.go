package main

import (
	"github.com/rmarasigan/fresnel-calc/models"
	"github.com/uadmin/uadmin"
)

func main() {
	dbSettings := uadmin.DBSettings{}
	dbSettings.Name = "fresnel.db"
	dbSettings.Type = "sqlite"
	uadmin.Database = &dbSettings

	uadmin.SiteName = "Fresnel Clearance Calculator"
	uadmin.RootURL = "/fresnel-clearance/"

	uadmin.Register(
		models.FresnelCalc{},
	)

	uadmin.Port = 1024
	uadmin.StartServer()
}
