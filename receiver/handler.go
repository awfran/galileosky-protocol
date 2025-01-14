package main

import (
	"encoding/binary"
	"github.com/kuznetsovin/galileosky-protocol/receiver/galileo"
	"io"
	"net"
	"time"
)

const headerLen = 3

func handleRecvPkg(conn net.Conn, ttl time.Duration) {
	var (
		recvPacket []byte
	)
	defer conn.Close()

	//packet.GalileoPaket
	logger.Warnf("Установлено соединение с %s", conn.RemoteAddr())
	outPkg := galileoParsePacket{}
	for {
		connTimer := time.NewTimer(ttl)

		// считываем заголовок пакета
		headerBuf := make([]byte, headerLen)
		_, err := conn.Read(headerBuf)

		switch err {
		case nil:
			connTimer.Reset(ttl)

			// если пакет не егтс формата закрываем соединение
			if headerBuf[0] != 0x01 {
				logger.Warnf("The package does not conform to the Galileo format. Closed connection%s", conn.RemoteAddr())
				return
			}

			// вычисляем длину пакета, 2 байта после тега
			pkgLen := binary.LittleEndian.Uint16(headerBuf[1:])
			pkgLen <<= 1
			pkgLen >>= 1
			pkgLen += 2

			// получаем концовку пакета
			buf := make([]byte, pkgLen)
			if _, err := io.ReadFull(conn, buf); err != nil {
				logger.Errorf("Error while getting the package body: %v", err)
				return
			}

			// формируем порлный пакет
			recvPacket = append(headerBuf, buf...)
		case io.EOF:
			<-connTimer.C
			_ = conn.Close()
			logger.Warnf("Compound %s closed by timeout", conn.RemoteAddr())
			return
		default:
			logger.Errorf("Error while getting: %v", err)
			return
		}

		logger.Debugf("Package accepted: %X", recvPacket)
		pkg := galileo.Packet{}
		if err := pkg.Decode(recvPacket); err != nil {
			logger.Warn("Package decryption error ")
			logger.Error(err)
			return
		}

		receivedTime := time.Now().UTC().Unix()
		crc := make([]byte, 2)
		binary.LittleEndian.PutUint16(crc, pkg.Crc16)
		resp := append([]byte{0x02}, crc...)

		if _, err = conn.Write(resp); err != nil {
			logger.Errorf("Error sending confirmation packet: %v", err)
		}

		if len(pkg.Tags) < 1 {
			continue
		}

		outPkg.ReceivedTimestamp = receivedTime
		prevTag := uint8(0)
		isSave := false
		for _, curTag := range pkg.Tags {
			if prevTag > curTag.Tag && isSave {
				if err := outPkg.Save(); err != nil {
					logger.Error(err)
				}
				client := outPkg.Client
				outPkg = galileoParsePacket{
					Client:            client,
					ReceivedTimestamp: receivedTime,
				}
				isSave = false
			}
			switch curTag.Tag {
			case 0x04:
				val := curTag.Value.(*galileo.UintTag)
				outPkg.Client = uint32(val.Val)
			case 0x10:
				val := curTag.Value.(*galileo.UintTag)
				outPkg.PacketID = uint32(val.Val)
			case 0x20:
				val := curTag.Value.(*galileo.TimeTag)
				outPkg.NavigationTimestamp = val.Val.Unix()
			case 0x30:
				val := curTag.Value.(*galileo.CoordTag)
				outPkg.Nsat = val.Nsat
				outPkg.Latitude = val.Latitude
				outPkg.Longitude = val.Longitude
				isSave = true
			case 0x33:
				val := curTag.Value.(*galileo.SpeedTag)
				outPkg.Course = uint8(val.Course)
				outPkg.Speed = uint16(val.Speed)
			case 0x35:
				val := curTag.Value.(*galileo.UintTag)
				outPkg.Hdop = uint16(val.Val)
			}
			prevTag = curTag.Tag
		}

	}
}
