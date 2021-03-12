package main

import "log"

func main() {

	img, err := loadImage("test")
	if err != nil {
		log.Panicln(err)
	}
	destWidth := 2000
	s := sketch.NewSketch(img, sketch.UserParams{
		StrokeRatio:              0.75,
		DestWidth:                destWidth,
		DestHeight:               2000,
		InitialAlpha:             0.1,
		StrokeReduction:          0.002,
		AlphaIncrease:            0.06,
		StrokeInversionThreshold: 0.05,
		StrokeJitter:             int(0.1 * float64(destWidth)),
		MinEdgeCount:             3,
		MaxEdgeCount:             4,
	})
}
