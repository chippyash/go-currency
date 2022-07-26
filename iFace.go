package go_currency

/**
 * Currency Support for Go

 * @author Ashley Kitson
 * @copyright Ashley Kitson, 2022, UK
 * @license BSD 3-Clause See LICENSE
 */

//CrcyIface is an interface for Currency operations
type CrcyIface interface {
	//Symbol returns currency symbol
	Symbol() string
	//Code returns currency code - 3 character
	Code() string
	//Name return name of this currency
	Name() string
	//DisplayFormat returns display format for use in display()
	DisplayFormat() string
	//Locale returns locale used for display purposes
	Locale() string
	//Precision returns number of digits of precision for the currency
	Precision() int
	//Value returns currency amount as int
	Value() int64
	//SetValue sets the currency value
	SetValue(value int64) CrcyIface
	//SetAsFloat sets currency value, upscaling into an integer for internal storage
	SetAsFloat(value float64) CrcyIface
	//AsFloat return currency value as downscaled float value
	AsFloat() float64
	//Display Return currency amount formatted for display according to i18n localisation rules
	Display() string
}
