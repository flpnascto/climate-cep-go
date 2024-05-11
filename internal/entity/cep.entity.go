package entity

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Cep struct {
	Value string
}

func NewCep(value string) (*Cep, error) {
	cep := &Cep{Value: value}
	err := cep.isValid()
	if err != nil {
		return nil, err
	}
	return cep, nil
}

func (c *Cep) isValid() error {
	c.clean()
	if len(c.Value) != 8 {
		return errors.New("invalid cep value")
	}

	if number, _ := strconv.Atoi(c.Value); number < 1001000 {
		return errors.New("invalid cep value")
	}
	return nil
}

func (c *Cep) clean() {
	re := regexp.MustCompile("[^0-9]+")
	c.Value = re.ReplaceAllString(c.Value, "")
}

func (c *Cep) GetCep() string {
	return c.Value
}

func (c *Cep) GetCepFormatted() string {
	return fmt.Sprintf("%s.%s-%s", c.Value[0:2], c.Value[2:5], c.Value[5:])
}
