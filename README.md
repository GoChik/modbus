go modbus [![Build Status](https://travis-ci.org/gochik/modbus.svg?branch=master)](https://travis-ci.org/gochik/modbus) [![GoDoc](https://godoc.org/github.com/gochik/modbus?status.svg)](https://godoc.org/github.com/goburrow/modbus)
=========
Fault-tolerant, fail-fast implementation of Modbus protocol in Go.

Supported functions
-------------------
Bit access:
*   Read Discrete Inputs
*   Read Coils
*   Write Single Coil
*   Write Multiple Coils

16-bit access:
*   Read Input Registers
*   Read Holding Registers
*   Write Single Register
*   Write Multiple Registers
*   Read/Write Multiple Registers
*   Mask Write Register
*   Read FIFO Queue

Supported formats
-----------------
*   TCP
*   Serial (RTU, ASCII)

Usage
-----
Basic usage:
```go
// Modbus TCP
client := modbus.TCPClient("localhost:502")
// Read input register 9
results, err := client.ReadInputRegisters(8, 1)
```

```go
// Modbus RTU/ASCII
// use your preferred serial library
import "github.com/tarm/serial"

// open the serial connection
port, _ := serial.OpenPort(&serial.Config{
	Name:   "/dev/ttyS0"
	Baud:   115200,
	Parity: serial.ParityNone,
})

// then create the client
client = modbus.RTUClient(port)
results, err = client.ReadCoils(2, 1)
```

Advanced usage:
```go
// Modbus TCP
handler := modbus.NewTCPClientHandler("localhost:502")
handler.Timeout = 10 * time.Second
handler.SlaveId = 0xFF
handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
// Connect manually so that multiple requests are handled in one connection session
err := handler.Connect()
defer handler.Close()

client := modbus.NewClient(handler)
results, err := client.ReadDiscreteInputs(15, 2)
results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})
```

References
----------
-   [Modbus Specifications and Implementation Guides](http://www.modbus.org/specs.php)
