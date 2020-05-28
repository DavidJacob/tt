package geography

import "github.com/gookit/color"

type CountryFact struct {
	Name        string
	CapitalCity string
	Population  int
	Area        int
	Flag        Flag
}
type Flag struct {
	Colors []color.RGBColor
	Chars  []string
}

type flagElement struct {
	color color.RGBColor
	char  rune
}

var Countries = []CountryFact{
	{
		Name:        "Austria",
		CapitalCity: "Vienna",
		Population:  8902600,
		Area:        83879,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(238, 36, 54),
				color.RGB(255, 255, 255),
				color.RGB(238, 36, 54),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Bulgaria",
		CapitalCity: "Sofia",
		Population:  6951482,
		Area:        110993,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(255, 255, 255),
				color.RGB(0, 128, 0),
				color.RGB(200, 0, 0),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Estonia",
		CapitalCity: "Talinn",
		Population:  1328976,
		Area:        45227,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(0, 115, 207),
				color.RGB(0, 0, 0),
				color.RGB(255, 255, 255),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Germany",
		CapitalCity: "Berlin",
		Population:  83149300,
		Area:        357022,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(0, 0, 0),
				color.RGB(255, 0, 0),
				color.RGB(255, 204, 0),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Hungary",
		CapitalCity: "Budapest",
		Population:  9772756,
		Area:        93030,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(206, 37, 60),
				color.RGB(255, 255, 255),
				color.RGB(65, 112, 76),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Lithuania",
		CapitalCity: "Vilnius",
		Population:  2794329,
		Area:        65300,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(253, 185, 19),
				color.RGB(0, 106, 68),
				color.RGB(193, 39, 45),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Luxembourg",
		CapitalCity: "Luxembourg City",
		Population:  626108,
		Area:        2586,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(238, 36, 54),
				color.RGB(255, 255, 255),
				color.RGB(0, 162, 223),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Netherlands",
		CapitalCity: "Amsterdam",
		Population:  17424978,
		Area:        41865,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(174, 28, 40),
				color.RGB(255, 255, 255),
				color.RGB(30, 70, 139),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
	{
		Name:        "Russia",
		CapitalCity: "Moscow",
		Population:  146745098,
		Area:        17098246,
		Flag: Flag{
			Colors: []color.RGBColor{
				color.RGB(255, 255, 255),
				color.RGB(0, 46, 157),
				color.RGB(215, 31, 15),
			},
			Chars: []string{
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
				"\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588\u2588",
			},
		},
	},
}
