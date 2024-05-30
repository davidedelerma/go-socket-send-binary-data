package main

import (
	"hw-api/listener"
	"hw-api/publisher"
	"sync"
)

func generateData(packet publisher.Packet) publisher.Packet {
	newId := packet.Sensid + 1
	return publisher.Packet{Sensid: newId, Locid: packet.Locid, Tstamp: packet.Tstamp, Temp: packet.Temp}
}

func main() {
	dataIn := publisher.Packet{Sensid: 1, Locid: 1233, Tstamp: 123452123, Temp: 12}
	threads := 2
	var wg sync.WaitGroup
	wg.Add(threads)
	go func() {
		defer wg.Done()
		listener.Listener()
	}()
	go func() {
		defer wg.Done()
		for dataIn.Sensid < 100 {
			publisher.Publisher(dataIn)
			dataIn = generateData(dataIn)
		}

	}()
	wg.Wait()
}
