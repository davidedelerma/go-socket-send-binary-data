package publisher

import (
	"encoding/binary"
	"fmt"
	"net"
)

/*
0       1       2       3       4       5       6       7
0123456701234567012345670123456701234567012345670123456701234567
+-------+-------+-------+-------+-------+-------+-------+------
|          SensorID             |    LocationID |     Timestamp
+-------+-------+-------+-------+-------+-------+-------+------
                |      Temp     |
+-------+-------+-------+-------+
*/

type Packet struct {
	Sensid uint32
	Locid  uint16
	Tstamp uint32
	Temp   int16
}

func (p *Packet) Encode() []byte {
	buf := make([]byte, 12)
	// encoding the data into buf
	binary.BigEndian.PutUint32(buf[0:], p.Sensid)        // sensorID
	binary.BigEndian.PutUint16(buf[4:], p.Locid)         // locationID
	binary.BigEndian.PutUint32(buf[6:], p.Tstamp)        // timestamp
	binary.BigEndian.PutUint16(buf[10:], uint16(p.Temp)) // temp
	return buf
}

func Decode(buf []byte) Packet {
	sensorID := binary.BigEndian.Uint32(buf[0:])
	locID := binary.BigEndian.Uint16(buf[4:])
	tstamp := binary.BigEndian.Uint32(buf[6:])
	temp := binary.BigEndian.Uint16(buf[10:])
	return Packet{sensorID, locID, tstamp, int16(temp)}
}

func Publisher(data Packet) {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send some data to the server
	_, err = conn.Write(data.Encode())
	fmt.Printf(
		"sent new data: sid: %d, locID %d ts: %d, temp:%d\n",
		data.Sensid, data.Locid, data.Tstamp, data.Temp,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Close the connection
	conn.Close()
}
