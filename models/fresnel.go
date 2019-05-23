package models

import (
	"fmt"
	"math"

	"github.com/uadmin/uadmin"
)

// FresnelCalc !
type FresnelCalc struct {
	uadmin.Model
	Distance1                     float64 `uadmin:"required;list_exclude"`
	Distance2                     float64 `uadmin:"required"`
	LowerRadioTotalElevation      float64
	HigherRadioTotalElevation     float64
	SuspectedObstructionElevation float64
	Frequency                     float64
	ThirdFresnelRadius            float64 `uadmin:"read_only"`
	AllowanceDueToEarthsCurve     float64 `uadmin:"read_only"`
	DistanceFrom3rdFresnel        float64 `uadmin:"read_only"`
	RadioTilt                     float64 `uadmin:"read_only"`
}

func (f *FresnelCalc) String() string {
	return fmt.Sprintf("%.2f", f.Distance1)
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
