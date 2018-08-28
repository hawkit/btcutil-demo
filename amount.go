package btcutil

import (
	"math"
	"strconv"
)

// AmountUnit describes a method of covertin an Amount to something
// other than the base unit of a bition. The value of the AmountUnit
// is the exponent  component of the decadic multiple to convert from
// an amount in bitcoin to an amount counted in units.
type AmountUnit int


// Amount represents the base bitcion monetary unit (colloquially referred)
// to as a `Satoshi`). A single Amount is equal to 1e-8 of a bitcoin
type Amount int64

// These constants define various units used when describing a bition
// monetary amount
const (
	AmountMegaBTC AmountUnit = 6
	AmountKiloBTC AmountUnit = 3
	AmountBTC AmountUnit = 0
	AmountMilliBTC AmountUnit = -3
	AmountMicroBTC AmountUnit = -6
	AmountSatoshi AmountUnit = -8
)

func (u AmountUnit) String() string {
	switch u {
	case AmountMegaBTC:
		return "MBTC"
	case AmountKiloBTC:
		return "KBTC"
	case AmountBTC:
		return "BTC"
	case AmountMilliBTC:
		return "mBTC"
	case AmountMicroBTC:
		return "μBTC"
	case AmountSatoshi:
		return "Satoshi"
	default:
		return "1e" + strconv.FormatInt(int64(u), 10) + "BTC"
	}

}
// ToUnit converts a monetary amount counted in bitcoin base units to
// a floating point value representing an amount of bitcoin
func (a Amount) ToUnit(u AmountUnit) float64  {
	return float64(a)/ math.Pow10(int(u+8))
}

// ToBTC is the equivalent of calling ToUnit with AmountBTC.
func (a Amount) ToBTC() float64 {
	return a.ToUnit(AmountBTC)
}

// Format formats a moneytary amount counted in bitcoin base units as
// a string for a given unit. The conversion will succeed for any unit,
// however, known units will be formated with an appended label describing
// the units with SI notation, or "Sataoshi" for the base unit.
func (a Amount) Format(u AmountUnit) string {
	units := " " + u.String()
	return  strconv.FormatFloat(a.ToUnit(u), 'f', -int(u+8), 64) + units
}

// String is the equivalent of calling Format with AmountBTC
func (a Amount) String() string  {
	return a.Format(AmountBTC)
}
