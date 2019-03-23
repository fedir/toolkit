package iptool

import (
	"fmt"
	"github.com/fedir/toolkit/stringtool"
	"strings"

	"github.com/fedir/toolkit/vartool"
)

// ValidateIP

func ValidateIP(IPs interface{}, IP string, logging bool) bool {
	var IPsSliced []string
	var ipValidated bool
	switch vartool.Type(IPs) {
	case "string":
		IPsSliced = strings.Split(IPs.(string), ",")
	case "array":
		IPsSliced = IPs.([]string)
	default:
		fmt.Println("IPs format is not recognized")
		return false
	}
	for _, ip := range IPsSliced {
		if stringtool.TrimSpace(ip) == IP {
			ipValidated = true

		}
	}
	if logging == true {
		switch ipValidated {
		case true:
			fmt.Println("IP Validated : ", IP)
		default:
			fmt.Println("IP NOT Validated : ", IP)
		}
	}
	return ipValidated
}
