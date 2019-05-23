package models

import (
	"math"

	"github.com/uadmin/uadmin"
)

// FresnelCalc !
type FresnelCalc struct {
	uadmin.Model
	Name                          string  `uadmin:"list_exclude"`
	Description                   string  `uadmin:"html"`
	Distance1                     float64 `uadmin:"required;help:meters"`
	Distance2                     float64 `uadmin:"required;help:meters"`
	LowerRadioTotalElevation      float64 `uadmin:"help:Land height + Bldg. height + Tower height (meters)"`
	HigherRadioTotalElevation     float64 `uadmin:"help:Land height + Bldg. height + Tower height (meters)"`
	SuspectedObstructionElevation float64 `uadmin:"help:Land height + Obstruction height (meters)"`
	Frequency                     float64 `uadmin:"required;help:Mhz"`
	ThirdFresnelRadius            float64 `uadmin:"read_only;help:m"`
	AllowanceDueToEarthsCurve     float64 `uadmin:"read_only;help:m"`
	DistanceFrom3rdFresnel        float64 `uadmin:"read_only;help:+ if good, - if bad"`
	RadioTilt                     float64 `uadmin:"read_only"`
}

func (f *FresnelCalc) String() string {
	return f.Name
}

// Save !
func (f *FresnelCalc) Save() {
	const Pi = 3.14159265358979323846264338327950288419716939937510582097494459

	firstDistance := f.Distance1
	secondDistance := f.Distance2
	lowerRadioElevation := f.LowerRadioTotalElevation
	highRadioElevation := f.HigherRadioTotalElevation
	suspectedObstruction := f.SuspectedObstructionElevation
	frequency := f.Frequency

	thirdFresnelRadius := math.Sqrt((3 * (299792458 / (frequency * 1000000)) * firstDistance * secondDistance) / (firstDistance + secondDistance))
	f.ThirdFresnelRadius = math.Round(thirdFresnelRadius*100) / 100

	thirdFresnelRadius = f.ThirdFresnelRadius

	allowanceEathsCurve := (1000 * math.Pow(((firstDistance+secondDistance)/1000), 2) / (8 * 6371))
	f.AllowanceDueToEarthsCurve = math.Round(allowanceEathsCurve*100) / 100

	allowanceEathsCurve = f.AllowanceDueToEarthsCurve

	distanceFrom3rdFresnel := (firstDistance * ((highRadioElevation - lowerRadioElevation) / (firstDistance + secondDistance))) - allowanceEathsCurve - thirdFresnelRadius - 3 - (suspectedObstruction - lowerRadioElevation)
	f.DistanceFrom3rdFresnel = math.Round(distanceFrom3rdFresnel*100) / 100

	radioTilt := (math.Tan((highRadioElevation - lowerRadioElevation) / (firstDistance + secondDistance))) * (180 / Pi)
	f.RadioTilt = math.Round(radioTilt*100) / 100

	uadmin.Save(f)
}
