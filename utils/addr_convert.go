package utils

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
)

func CoventHexIpToAddress(hexAddr string) (string, error) {
	if len(hexAddr) != 8 {
		return "", fmt.Errorf("invalid hex ip address format: invalid lenght")
	}

	ipBytes, err := hex.DecodeString(hexAddr)
	if err != nil {
		return "", fmt.Errorf("invalid hex ip address format: %w", err)
	}

	// Переворачиваем байты IP (little-endian to big-endian)
	for i, j := 0, len(ipBytes)-1; i < j; i, j = i+1, j-1 {
		ipBytes[i], ipBytes[j] = ipBytes[j], ipBytes[i]
	}

	return net.IP(ipBytes).String(), nil
}

func CoventHexPortToAddress(hexPort string) (string, error) {
	if len(hexPort) != 4 {
		return "", fmt.Errorf("invalid hex port format: invalid lenght")
	}

	port, err := strconv.ParseInt(hexPort, 16, 32)
	if err != nil {
		return "", fmt.Errorf("invalid hex port format: %w", err)
	}

	return strconv.Itoa(int(port)), nil
}
