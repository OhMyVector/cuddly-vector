package main

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/vectorpb"
)

const ID_TAG = 2000001

func main() {
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

			for i := 0; i < 4; i++ {
				_, _ = v.Conn.DriveStraight(ctx, &vectorpb.DriveStraightRequest{
					DistMm:    200,
					SpeedMmps: 50,
					IdTag:     ID_TAG,
				})

				_, _ = v.Conn.TurnInPlace(ctx, &vectorpb.TurnInPlaceRequest{
					AngleRad: 90 * math.Pi / 180,
					IdTag:    ID_TAG,
				})
			}

			stop <- true
			return
		}
	}

}
