package main

import (
        "fmt"

        "github.com/tarm/serial"
)

func main() {
        c := &serial.Config{Name: "/dev/rfcomm0", Baud: 9600}
        s, err := serial.OpenPort(c)
        if err != nil {
                fmt.Println(err)
        }

	for { //i := 0; i<100; i++ {
		buf := make([]byte, 1)
		n, err := s.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%x", buf[:n])
	}
}
