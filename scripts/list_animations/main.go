package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

func main() {

	var talk = flag.String("talk", "", "Vector's Talk")
	flag.Parse()

	if *talk == "" {
		log.Fatal("please use the -talk argument and set it to your robots talk")
	}

	v, err := vector.New(
		vector.WithTarget(os.Getenv("BOT_TARGET")),
		vector.WithToken(os.Getenv("BOT_TOKEN")),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	start := make(chan bool)
	stop := make(chan bool)

	go func() {
		_ = v.BehaviorControl(ctx, start, stop)
	}()

	for {
		select {
		case <-start:
			animations, err := v.Conn.ListAnimations(ctx, &vectorpb.ListAnimationsRequest{})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(animations)

			stop <- true
			return
		}
	}

}
