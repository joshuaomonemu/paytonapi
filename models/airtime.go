package models

import (
	"fmt"

	reloadly "github.com/reloadly/reloadly-sdk-golang/airtime"
)

func Airtime(number, amount string) (string, error) {
	c, err := reloadly.NewClient("LrSJMJirX0iR3MB6AKj9G2P6Ui3YsXjO", "4PYVcPKzSv-SSI7ZKhHEC4fZ6rcF1X-2VCGByiWTjqT1FylCmq0OcIBMCoJAwGu", true)
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
	fmt.Println(id)

	return "success", nil
}
