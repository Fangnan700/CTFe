package models

type Response struct {
	Code int         `mapstructure:"code"`
	Body interface{} `mapstructure:"body"`
}
