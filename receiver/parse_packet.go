package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type galileoParsePacket struct {
	Client              uint32         `json:"client"`
	PacketID            uint32         `json:"packet_id"`
	NavigationTimestamp int64          `json:"navigation_unix_time"`
	ReceivedTimestamp   int64          `json:"received_unix_time"`
	Latitude            float64        `json:"latitude"`
	Longitude           float64        `json:"longitude"`
	Speed               uint16         `json:"speed"`
	Pdop                uint16         `json:"pdop"`
	Hdop                uint16         `json:"hdop"`
	Vdop                uint16         `json:"vdop"`
	Nsat                uint8          `json:"nsat"`
	Ns                  uint16         `json:"ns"`
	Course              uint8          `json:"course"`
	AnSensors           []anSensor     `json:"an_sensors"`
	LiquidSensors       []liquidSensor `json:"liquid_sensors"`
}

func (g galileoParsePacket) Save() error {
	result, err := json.MarshalIndent(g, " ", " ")
	if err != nil {
		return fmt.Errorf("Ошибка парсинга данных: %v", err)
	}

	fmt.Println(string(result))
	
	responseBody := bytes.NewBuffer(result) 
	resp, err := http.Post("https://tracking.gypsick.com/api/v1/log", "application/json", responseBody)  

	if err != nil {
		return fmt.Errorf("Ошибка парсинга данных: %v", err)
	}

	return err
}

type liquidSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	ValueMm      uint32 `json:"value_mm"`
	ValueL       uint32 `json:"value_l"`
}

type anSensor struct {
	SensorNumber uint8  `json:"sensor_number"`
	Value        uint32 `json:"value"`
}
