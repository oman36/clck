package main

import (
	"strings"
)

type ShortUriConverter struct {
	intChar map[int]string
	charInt map[string]int
}

func (c *ShortUriConverter) IntToUri(value uint) (string, error) {
	var result []string
	length := uint(len(c.intChar))
	for value > 0 {
		result = append(result, c.intChar[int(value%length)])
		value /= length
	}
	ReverseSlice(result)
	return strings.Join(result, ""), nil
}

func (c *ShortUriConverter) UriToInt(value string) (uint, error) {
	var result uint
	length := uint(len(c.charInt))
	factor := uint(1)
	for i := 0; i < len(value); i++ {
		result += factor * uint(c.charInt[string(value[i])])
		factor *= length
	}
	return result, nil
}

func NewShortUriConverter() *ShortUriConverter {
	c := new(ShortUriConverter)
	c.intChar = make(map[int]string)
	c.charInt = make(map[string]int)
	for i := 0; i < 10; i++ {
		// 0-9
		c.charInt[string(48+i)] = i
		c.intChar[i] = string(48 + i)
	}
	for i := 10; i < 36; i++ {
		// a-z
		c.charInt[string(55+i)] = i
		c.intChar[i] = string(55 + i)

		// A-Z
		c.charInt[string(55+i+32)] = i + 26
		c.intChar[i+26] = string(55 + i + 32)
	}
	return c
}
