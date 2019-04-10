/*
CHRIS FELLI, 2019
Validates IPv4 and IPv6 addresses
IPv4: 192.168.255.255
IPv6: 2001:0db8:85a3:0:0:8A2E:0370:7334
*/
// package gocode ...
package gocode

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if validIPv4Address(IP) {
		return "IPv4"
	}
	if validIPv6Address(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv6Address(IP string) bool {
	strArr := strings.Split(IP, ":")
	if len(strArr) != 8 {
		return false
	}
	for _, str := range strArr {
		if len(str) == 0 || len(str) > 4 {
			return false
		}
		for i := 0; i < len(str); i++ {
			if !(str[i] >= '0' && str[i] <= '9') &&
				!(str[i] >= 'a' && str[i] <= 'f') &&
				!(str[i] >= 'A' && str[i] <= 'F') {
				return false
			}
		}
	}
	return true
}

func validIPv4Address(IP string) bool {
	strArr := strings.Split(IP, ".")
	if len(strArr) != 4 {
		return false
	}
	for _, str := range strArr {
		num, err := strconv.Atoi(str)
		if err != nil || num > 255 || num < 0 {
			return false
		}
		newStr := fmt.Sprint(num)
		if str != newStr {
			return false
		}
	}
	return true
}
