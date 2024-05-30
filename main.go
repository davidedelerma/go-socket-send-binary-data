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
	threads := 11
	var wg sync.WaitGroup
	wg.Add(threads)
	go func() {
		defer wg.Done()
		listener.Listener()
	}()
	go func() {
		for dataIn.Sensid < uint32(threads)-1 {
			defer wg.Done()
			publisher.Publisher(dataIn)
			dataIn = generateData(dataIn)
		}

	}()
	wg.Wait()
}
