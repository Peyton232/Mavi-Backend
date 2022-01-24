// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type BillDetails struct {
	Email    *string          `json:"email"`
	FullName *string          `json:"fullName"`
	Phone    *string          `json:"phone"`
	BillAdd  *BillingAdddress `json:"billAdd"`
}

type BillingAdddress struct {
	City         *string `json:"city"`
	CountryCode  *string `json:"countryCode"`
	AddressLine1 *string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`
	Zip          *string `json:"zip"`
	State        *string `json:"state"`
}

type FoodItem struct {
	Name          *string `json:"name"`
	Price         *string `json:"price"`
	Modifications *string `json:"modifications"`
}

type NewUser struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Pin   *string `json:"pin"`
}

type Payment struct {
	PaymentType    *PayType     `json:"paymentType"`
	DigitalWallet  *string      `json:"digitalWallet"`
	Paypal         *string      `json:"paypal"`
	CryptoWallet   *string      `json:"cryptoWallet"`
	BillingDetails *BillDetails `json:"billingDetails"`
}

type Prefence struct {
	RidePrefence *RidePref   `json:"ridePrefence"`
	Alergies     *string     `json:"alergies"`
	Restaurant   *Restaurant `json:"restaurant"`
	FavOrders    []*string   `json:"favOrders"`
	FavCuisines  []*string   `json:"favCuisines"`
}

type Restaurant struct {
	Name          *string     `json:"name"`
	FavoriteOrder []*FoodItem `json:"favoriteOrder"`
	TotalPrice    *string     `json:"totalPrice"`
}

type User struct {
	Name           *string    `json:"name"`
	UserID         string     `json:"userID"`
	Prefence       *Prefence  `json:"prefence"`
	Email          *string    `json:"email"`
	Pin            *string    `json:"pin"`
	PaymentMethods []*Payment `json:"paymentMethods"`
}

type PayType string

const (
	PayTypeDigitalwallet PayType = "DIGITALWALLET"
	PayTypeCreditcard    PayType = "CREDITCARD"
	PayTypeCrypto        PayType = "CRYPTO"
)

var AllPayType = []PayType{
	PayTypeDigitalwallet,
	PayTypeCreditcard,
	PayTypeCrypto,
}

func (e PayType) IsValid() bool {
	switch e {
	case PayTypeDigitalwallet, PayTypeCreditcard, PayTypeCrypto:
		return true
	}
	return false
}

func (e PayType) String() string {
	return string(e)
}

func (e *PayType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PayType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid payType", str)
	}
	return nil
}

func (e PayType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RidePref string

const (
	RidePrefTollsallowed RidePref = "TOLLSALLOWED"
	RidePrefAvoidhwy     RidePref = "AVOIDHWY"
	RidePrefNotoll       RidePref = "NOTOLL"
)

var AllRidePref = []RidePref{
	RidePrefTollsallowed,
	RidePrefAvoidhwy,
	RidePrefNotoll,
}

func (e RidePref) IsValid() bool {
	switch e {
	case RidePrefTollsallowed, RidePrefAvoidhwy, RidePrefNotoll:
		return true
	}
	return false
}

func (e RidePref) String() string {
	return string(e)
}

func (e *RidePref) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RidePref(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ridePref", str)
	}
	return nil
}

func (e RidePref) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
