package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/raphael/recorder/gen/recorder"
	"github.com/raphael/recorder/poster"
	"github.com/raphael/recorder/sieve"
)

func main() {
	p := poster.New("optima-tve.appspot.com")
	lambda.Start(func(ctx context.Context, n int) (float64, error) {
		_, dur := sieve.Eratosthenes(n)
		err := p.Post(ctx, &recordersvc.Datapoint{
			Service: "lambda",
			Name:    fmt.Sprintf("sieve-%d", n),
			Value:   dur,
		})
		return dur, err
	})
}
