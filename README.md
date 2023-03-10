# Vox 

[![Go Reference](https://pkg.go.dev/badge/github.com/gaurishhs/vox.svg)](https://pkg.go.dev/github.com/gaurishhs/vox)
[![Go Report Card](https://goreportcard.com/badge/github.com/gaurishhs/vox)](https://goreportcard.com/report/github.com/gaurishhs/vox)

Simple pub/sub library for Golang 

## Installation

```bash
go get -u github.com/gaurishhs/vox
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/gaurishhs/vox"
)

func main() {
    // Create a new publisher
    pub := vox.NewPublisher()
    sub := vox.NewSubscriber()

    // Subscribe to a topic
    sub.Subscribe("topic")
   
    // Publish a message to a topic
    pub.Publish("topic", "Hello World")

    sub.Listen(func(message *vox.Message) {
        fmt.Println(topic, message)
    })
}
```

## License

This project is licensed under MIT License - see the 
LICENSE file for details
