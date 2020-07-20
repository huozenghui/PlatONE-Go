package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// regMatch check if string matches the pattern by regular expression
func regMatch(param, pattern string) bool {
	result, _ := regexp.MatchString(pattern, param)
	return result
}

// IsMatch selects different patterns by the paramName
func IsMatch(param, paramName string) bool {
	var pattern string

	switch paramName {
	case "name":
		pattern = `^[a-zA-Z]\w{2,15}$` //english name: Alice_02
	case "num":
		pattern = `^[+-]{0,1}\d+$` //1823..., +1, -123
	case "email":
		pattern = `^[a-zA-Z\d][\w-.]{2,15}@[\w]+(.[a-zA-Z]{2,6}){1,2}$` //alice@wxblockchain.com
	case "mobile":
		pattern = "^1(3[0-9]|4[57]|[0-35-9]|7[06-8])\\d{8}$" //136xxxxxxxx
	case "version":
		pattern = `^([\d]\.){3}[\d]$` //0.0.0.1
	case "address":
		pattern = `^0[x|X][\da-fA-F]{40}$` //0x00...00
	default:
		pattern = `[\s~!@#\$%^&*\(\)\{\}\[\]\|\,\?]` //special char
	}

	return regMatch(param, pattern)
}

// IsUrl check if the input is an Url, for examplt 127.0.0.1:6791
func IsUrl(url string) bool {
	var port string
	var ip string

	array := strings.Split(url, ":")
	if len(array) == 2 {
		port = array[1]
		ip = array[0]
	} else {
		/// Logger.Printf("verify url error, invalid url syntax <ip>:<port> %s", url)
		return false
	}

	if !IsInRange(port, 65535) {
		return false
	}

	arrayIP := strings.Split(ip, ".")
	if len(arrayIP) == 4 {
		for _, data := range arrayIP {
			if !IsInRange(data, 255) {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

// IsInRange check the value is in the range selected
func IsInRange(value string, num int64) bool {

	if IsMatch(value, "num") {
		intValue, _ := strconv.ParseInt(value, 10, 0)
		if intValue <= num && intValue >= 0 {
			return true
		}
	}

	return false
}

// isValidRoles wraps isRoleMatch, it extracts the roles in the array and validates the roles
func IsValidRoles(roles string) bool {
	if roles == "" {
		return false
	}

	if !strings.HasPrefix(roles, "[") || !strings.HasSuffix(roles, "]") {
		return false
	}

	rolesArray := strings.Split(roles, "\"")
	for i := 1; i < len(rolesArray); i = i + 2 {
		if !IsRoleMatch(rolesArray[i]) {
			return false
		}
	}
	return true
}

// isRoleMatch checks if the input role is valid
func IsRoleMatch(role string) bool {
	var roleList = []string{
		"chainCreator",
		"chainAdmin",
		"nodeAdmin",
		"contractAdmin",
		"contractDeployer",
	}
	role = strings.Trim(role, " ")

	for _, value := range roleList {
		if role == value {
			return true
		}
	}

	return false
}
