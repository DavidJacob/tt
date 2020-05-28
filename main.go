package main

import (
	"math/rand"
	"time"

	"github.com/DavidJacob/tt/geography"
	"github.com/gookit/color"
)

func main() {

	if color.IsSupportTrueColor() {
		color.ForceOpenColor()
	}

	rand.Seed(time.Now().UnixNano())
	randomN := rand.Intn(len(geography.Countries))
	geography.PrintCountryFact(geography.Countries[randomN])
}
