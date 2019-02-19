package main

import (
	"regexp"
)

func urlIsValid(url string) (valid bool, err error) {
	if len(url) > 2048 {
		valid = false
		return
	}
	protocol := "^(?:https?|ftp)://"
	auth := "(?:\\S+(?::\\S*)?@)"
	privateIps := "(?:10\\.\\d{1,3}|192\\.168|172\\.(?:1[6-9]|2\\d|3[0-1]))(?:\\.\\d{1,3}){2}"
	localIps := "(?:0|127)(?:\\.\\d{1,3}){3}"
	reg := protocol + auth + "?(?:" + privateIps + "|" + localIps + ")"
	invalid, err := regexp.MatchString(reg, url)
	if invalid {
		valid = false
		return
	}
	ipPattern := "\\d{1,3}(?:\\.\\d{1,3}){3}"
	domainPattern := "(?:(?:[a-z\\x{00a1}-\\x{ffff}0-9]+-?)*[a-z\\x{00a1}-\\x{ffff}0-9]+)(?:\\.(?:[a-z\\x{00a1}-\\x{ffff}0-9]+-?)*[a-z\\x{00a1}-\\x{ffff}0-9]+)*(?:\\.(?:[a-z\\x{00a1}-\\x{ffff}]{2,}))"
	portPattern := "(?::\\d{2,5})"
	uriPattern := "(?:/[^\\s]*)"
	urlPattern := protocol + auth + "?(?:" + ipPattern + "|" + domainPattern + ")" + portPattern + "?" + uriPattern + "?$"
	valid, err = regexp.MatchString(urlPattern, url)
	return
}

func shortUrlIsValid(link string) (valid bool, err error) {
	if len(link) > 5 {
		valid = false
		return
	}
	valid, err = regexp.MatchString("[A-Za-z0-9]", link)
	return
}
