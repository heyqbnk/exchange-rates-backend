// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Currency struct {
	// Аббревиатура валюты.
	ID string `json:"id"`
	// Наименование валюты.
	Title string `json:"title"`
	// Символ валюты.
	Sign string `json:"sign"`
	// Значение, которое используется для конвертации из одной валюты в другую.
	ConvertRate float64 `json:"convertRate"`
}