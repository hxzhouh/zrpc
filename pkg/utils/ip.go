package utils

import (
	"errors"
	"fmt"
	"net"
)

func GetClientIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:8")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())

	return "", errors.New("Can not find the client ip address!")

}
