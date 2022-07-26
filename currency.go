package go_currency

import "math"

/**
 * Currency Support for Go

 * @author Ashley Kitson
 * @copyright Ashley Kitson, 2022, UK
 * @license BSD 3-Clause See LICENSE
 */

const defaultPrecision = 2

//Currency is a currency object that implements the CrcyIface interface
type Currency struct {
	CrcyIface
	value         int64
	code          string
	symbol        string
	name          string
	displayFormat string
	locale        string
	precision     int
}

func newCrcy() *Currency {
	return &Currency{
		value:         0,
		code:          "GBP",
		symbol:        "£",
		name:          "Stirling",
		displayFormat: "£ %f0.2",
		locale:        "en-gb",
		precision:     defaultPrecision,
	}
}

func (c *Currency) Symbol() string {
	return c.symbol
}

func (c *Currency) Code() string {
	return c.code
}

func (c *Currency) Name() string {
	return c.name
}

func (c *Currency) setName(name string) *Currency {
	c.name = name
	return c
}

func (c *Currency) DisplayFormat() string {
	return c.displayFormat
}

func (c *Currency) Locale() string {
	return c.locale
}

func (c *Currency) Value() int64 {
	return c.value
}

func (c *Currency) SetValue(value int64) CrcyIface {
	c.value = value
	return c
}

func (c *Currency) SetAsFloat(value float64) CrcyIface {
	c.value = int64(value * math.Pow(10, float64(c.precision)))
	return c
}

func (c *Currency) AsFloat() float64 {
	return float64(c.value) / math.Pow(10, float64(c.precision))
}

func (c *Currency) Display() string {
	return ""
}

func (c *Currency) setCode(code string) *Currency {
	c.code = code
	return c
}

func (c *Currency) setSymbol(symbol string) *Currency {
	c.symbol = symbol
	return c
}

func (c *Currency) setDisplayFormat(displayFormat string) *Currency {
	c.displayFormat = displayFormat
	return c
}

func (c *Currency) setLocale(locale string) *Currency {
	c.locale = locale
	return c
}

func (c *Currency) setPrecision(precision int) *Currency {
	c.precision = precision
	return c
}

type crcyBuilder struct {
	crcy *Currency
}

//New returns a currency builder
func New() *crcyBuilder {
	b := new(crcyBuilder)
	b.crcy = newCrcy()
	return b
}

//WithCode sets the currency code
func (b *crcyBuilder) WithCode(v string) *crcyBuilder {
	b.crcy = b.crcy.setCode(v)
	return b
}

//WithSymbol sets the currency symbol
func (b *crcyBuilder) WithSymbol(v string) *crcyBuilder {
	b.crcy = b.crcy.setSymbol(v)
	return b
}

//WithPrecision sets the currency precision
func (b *crcyBuilder) WithPrecision(v int) *crcyBuilder {
	b.crcy = b.crcy.setPrecision(v)
	return b
}

//WithName sets the currency long name
func (b *crcyBuilder) WithName(v string) *crcyBuilder {
	b.crcy = b.crcy.setName(v)
	return b
}

//WithDisplayFormat sets the sprintf display template for use with Display()
func (b *crcyBuilder) WithDisplayFormat(v string) *crcyBuilder {
	b.crcy = b.crcy.setDisplayFormat(v)
	return b
}

//WithLocale sets the locale for Display()
func (b *crcyBuilder) WithLocale(v string) *crcyBuilder {
	b.crcy = b.crcy.setLocale(v)
	return b
}

//WithValue sets the currency value from a float
func (b *crcyBuilder) WithValue(v float64) *crcyBuilder {
	c := b.crcy.SetAsFloat(v).(*Currency)
	b.crcy = c
	return b
}

//Build returns a Currency that implements the CrcyIFace interface
func (b *crcyBuilder) Build() CrcyIface {
	return b.crcy
}
