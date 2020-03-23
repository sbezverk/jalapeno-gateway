package bgpclient

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

// RDValidator validates Route Distinguisher stored in string
func RDValidator(rd string) error {
	parts := strings.Split(rd, ":")
	if len(parts) != 2 {
		return fmt.Errorf("malformed RD, expected 2 fields separated by ':'")
	}
	part1 := strings.Trim(parts[0], " ")
	part2 := strings.Trim(parts[1], " ")

	if net.ParseIP(part1).To4() != nil {
		// Possible RD in format IP:Value, Value cannot exceed uint16 value.
		n, err := strconv.Atoi(part2)
		if err != nil {
			return fmt.Errorf("malformed RD, failed to parse Value field %s with error: %+v", part2, err)
		}
		if n > math.MaxUint16 {
			return fmt.Errorf("malformed RD, Value field %d exceeds maximum allowable %d", n, math.MaxUint16)
		}
		return nil
	}
	n1, err := strconv.Atoi(part1)
	if err != nil {
		return fmt.Errorf("malformed RD, failed to parse ASN field %s with error: %+v", part1, err)
	}
	n2, err := strconv.Atoi(part2)
	if err != nil {
		return fmt.Errorf("malformed RD, failed to parse Value field %s with error: %+v", part2, err)
	}
	// Check for ASN 4 bytes and Value 2 bytes
	if n1 > math.MaxUint16 && n1 <= math.MaxUint32 {
		if n2 > math.MaxUint16 {
			return fmt.Errorf("malformed RD, Value field %d exceeds maximum allowable %d", n2, math.MaxUint16)
		}
		return nil
	}
	// Check for ASN 2 bytes and Value 4 bytes
	if n1 <= math.MaxUint16 {
		if n2 > math.MaxUint32 {
			return fmt.Errorf("malformed RD, Value field %d exceeds maximum allowable %d", n2, math.MaxUint32)
		}
		return nil
	}
	return fmt.Errorf("malformed RD, ASN field %d exceeds maximum allowable %d", n1, math.MaxUint32)
}

// RTValidator validates Route Target stored in string
func RTValidator(rt string) error {
	parts := strings.Split(rt, ":")
	if len(parts) != 2 {
		return fmt.Errorf("malformed RT, expected 2 fields separated by ':'")
	}
	part1 := strings.Trim(parts[0], " ")
	part2 := strings.Trim(parts[1], " ")

	if net.ParseIP(part1).To4() != nil {
		// Possible RD in format IP:Value, Value cannot exceed uint16 value.
		n, err := strconv.Atoi(part2)
		if err != nil {
			return fmt.Errorf("malformed RT, failed to parse Value field %s with error: %+v", part2, err)
		}
		if n > math.MaxUint16 {
			return fmt.Errorf("malformed RT, Value field %d exceeds maximum allowable %d", n, math.MaxUint16)
		}
		return nil
	}
	n1, err := strconv.Atoi(part1)
	if err != nil {
		return fmt.Errorf("malformed RT, failed to parse ASN field %s with error: %+v", part1, err)
	}
	n2, err := strconv.Atoi(part2)
	if err != nil {
		return fmt.Errorf("malformed RT, failed to parse Value field %s with error: %+v", part2, err)
	}
	// Check for ASN 4 bytes and Value 2 bytes
	if n1 > math.MaxUint16 && n1 <= math.MaxUint32 {
		if n2 > math.MaxUint16 {
			return fmt.Errorf("malformed RT, Value field %d exceeds maximum allowable %d", n2, math.MaxUint16)
		}
		return nil
	}
	// Check for ASN 2 bytes and Value 4 bytes
	if n1 <= math.MaxUint16 {
		if n2 > math.MaxUint32 {
			return fmt.Errorf("malformed RT, Value field %d exceeds maximum allowable %d", n2, math.MaxUint32)
		}
		return nil
	}
	return fmt.Errorf("malformed RT, ASN field %d exceeds maximum allowable %d", n1, math.MaxUint32)
}
