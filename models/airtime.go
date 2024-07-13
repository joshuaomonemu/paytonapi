package models

import (
	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
)

func Airtime(number, amount string) (string, error) {
	c, err := reloadly.NewClient("HWLjSX8IsfLJxtM3riD0WheELc4azOrT", "MSUu3ll37b-zxV5XUMLFeiDxhyvSPd-fSZHgIqFFpTIhVPWcYk2T6UCGVerLRwS", true)
	if err != nil {
		return "", err
	}

	rPhone := reloadly.Phone{
		Number:      number,
		CountryCode: "NG",
	}

	t, err := c.Topup(amount, "341", false, rPhone)
	if err != nil {
		return "", err
	}
	id := t.TransactionDate

	return id, nil
}
