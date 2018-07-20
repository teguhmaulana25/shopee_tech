package controller

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type res struct {
	Code    int         `json:"code"`    // code error for more explanation ex: 20003
	Message string      `json:"message"` // must as verbose as possible ex: failed Authenticate
	Data    interface{} `json:"data"`
}

type Model struct {
	Name string
}

func renderJSON(w http.ResponseWriter, data res, status int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error("Failed marshal data")
		jsonData = nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
	if data.Code == 200 {
		log.Info(data.Code, " ", data.Message)
	} else if data.Code == 404 {
		log.Warn(data.Code, " ", data.Message)
	} else {
		log.Debug(data.Code, " ", data.Message)
	}
}

func GetIP() string {
	var ipAddress string
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		ipAddress = "failed"
	}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip := networkIp.IP.String()
			ipAddress = ip
		}
	}
	return ipAddress
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
