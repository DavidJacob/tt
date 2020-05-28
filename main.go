package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gookit/color"
)

type countryFact struct {
	name        string
	capitalCity string
	population  int
	area        int
	flag        flag
}
type flag struct {
	colors []color.RGBColor
	chars  []string
}

var countries = []countryFact{
	{
		name:        "Austria",
		capitalCity: "Vienna",
		population:  8902600,
		area:        83879,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(238, 36, 54),
				color.RGB(255, 255, 255),
				color.RGB(238, 36, 54),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Bulgaria",
		capitalCity: "Sofia",
		population:  6951482,
		area:        110993,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(255, 255, 255),
				color.RGB(0, 128, 0),
				color.RGB(200, 0, 0),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Estonia",
		capitalCity: "Talinn",
		population:  1328976,
		area:        45227,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(0, 115, 207),
				color.RGB(0, 0, 0),
				color.RGB(255, 255, 255),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Germany",
		capitalCity: "Berlin",
		population:  83149300,
		area:        357022,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(0, 0, 0),
				color.RGB(255, 0, 0),
				color.RGB(255, 204, 0),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Hungary",
		capitalCity: "Budapest",
		population:  9772756,
		area:        93030,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(206, 37, 60),
				color.RGB(255, 255, 255),
				color.RGB(65, 112, 76),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Lithuania",
		capitalCity: "Vilnius",
		population:  2794329,
		area:        65300,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(253, 185, 19),
				color.RGB(0, 106, 68),
				color.RGB(193, 39, 45),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Luxembourg",
		capitalCity: "Luxembourg City",
		population:  626108,
		area:        2586,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(238, 36, 54),
				color.RGB(255, 255, 255),
				color.RGB(0, 162, 223),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Netherlands",
		capitalCity: "Amsterdam",
		population:  17424978,
		area:        41865,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(174, 28, 40),
				color.RGB(255, 255, 255),
				color.RGB(30, 70, 139),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		name:        "Russia",
		capitalCity: "Moscow",
		population:  146745098,
		area:        17098246,
		flag: flag{
			colors: []color.RGBColor{
				color.RGB(255, 255, 255),
				color.RGB(0, 46, 157),
				color.RGB(215, 31, 15),
			},
			chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
}

func printCountryFact(countryFact countryFact) {
	fmt.Println(color.FgLightBlue.Sprint("Name: ") + color.FgWhite.Sprint(countryFact.name))
	fmt.Println(color.FgLightBlue.Sprint("Capital: ") + color.FgWhite.Sprint(countryFact.capitalCity))
	fmt.Println(color.FgLightBlue.Sprint("Population: ") + color.FgWhite.Sprint(countryFact.population))
	fmt.Println(color.FgLightBlue.Sprint("Area: ") + color.FgWhite.Sprint(countryFact.area))
	printFlag(countryFact.flag)
}

func printFlag(flag flag) {
	for i := 0; i < len(flag.colors); i++ {
		s := color.NewRGBStyle(flag.colors[i])
		s.Println(flag.chars[i])
	}
}

func main() {

	if color.IsSupportTrueColor() {
		color.ForceOpenColor()
	}

	rand.Seed(time.Now().UnixNano())
	randomN := rand.Intn(len(countries))
	printCountryFact(countries[randomN])
}
