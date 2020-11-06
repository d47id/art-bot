package colors

import (
	"image/color"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Random returns a random CSS3 color name
func Random() string {
	return cols[rand.Intn(len(cols))]
}

var cols = []string{
	"AliceBlue",
	"AntiqueWhite",
	"Aqua",
	"Aquamarine",
	"Azure",
	"Beige",
	"Bisque",
	"Black",
	"BlanchedAlmond",
	"Blue",
	"BlueViolet",
	"Brown",
	"BurlyWood",
	"CadetBlue",
	"Chartreuse",
	"Chocolate",
	"Coral",
	"CornflowerBlue",
	"Cornsilk",
	"Crimson",
	"Cyan",
	"DarkBlue",
	"DarkCyan",
	"DarkGoldenrod",
	"DarkGray",
	"DarkGreen",
	"DarkGrey",
	"DarkKhaki",
	"DarkMagenta",
	"DarkOliveGreen",
	"DarkOrange",
	"DarkOrchid",
	"DarkRed",
	"DarkSalmon",
	"DarkSeaGreen",
	"DarkSlateBlue",
	"DarkSlateGray",
	"DarkSlateGrey",
	"DarkTurquoise",
	"DarkViolet",
	"DeepPink",
	"DeepSkyBlue",
	"DimGray",
	"DodgerBlue",
	"FireBrick",
	"FloralWhite",
	"ForestGreen",
	"Fuchsia",
	"Gainsboro",
	"GhostWhite",
	"Gold",
	"Goldenrod",
	"Gray",
	"Green",
	"GreenYellow",
	"Grey",
	"Honeydew",
	"HotPink",
	"IndianRed",
	"Indigo",
	"Ivory",
	"Khaki",
	"Lavender",
	"LavenderBlush",
	"LawnGreen",
	"LemonChiffon",
	"LightBlue",
	"LightCoral",
	"LightCyan",
	"LightGoldenrodYellow",
	"LightGray",
	"LightGreen",
	"LightGrey",
	"LightPink",
	"LightSalmon",
	"LightSeaGreen",
	"LightSkyBlue",
	"LightSlateGray",
	"LightSlateGrey",
	"LightSteelBlue",
	"LightYellow",
	"Lime",
	"LimeGreen",
	"Linen",
	"Magenta",
	"Maroon",
	"MediumAquamarine",
	"MediumBlue",
	"MediumOrchid",
	"MediumPurple",
	"MediumSeaGreen",
	"MediumSlateBlue",
	"MediumSpringGreen",
	"MediumTurquoise",
	"MediumVioletRed",
	"MidnightBlue",
	"MintCream",
	"MistyRose",
	"Moccasin",
	"NavajoWhite",
	"Navy",
	"OldLace",
	"Olive",
	"OliveDrab",
	"Orange",
	"OrangeRed",
	"Orchid",
	"PaleGoldenrod",
	"PaleGreen",
	"PaleTurquoise",
	"PaleVioletRed",
	"PapayaWhip",
	"PeachPuff",
	"Peru",
	"Pink",
	"Plum",
	"PowderBlue",
	"Purple",
	"Rebeccapurple",
	"Red",
	"RosyBrown",
	"RoyalBlue",
	"SaddleBrown",
	"Salmon",
	"SandyBrown",
	"SeaGreen",
	"Seashell",
	"Sienna",
	"Silver",
	"SkyBlue",
	"SlateBlue",
	"SlateGray",
	"SlateGrey",
	"Snow",
	"SpringGreen",
	"SteelBlue",
	"Tan",
	"Teal",
	"Thistle",
	"Tomato",
	"Turquoise",
	"Violet",
	"Wheat",
	"White",
	"WhiteSmoke",
	"Yellow",
	"YellowGreen",
}

// CSS3 color names
const (
	AliceBlue            = "AliceBlue"
	AntiqueWhite         = "AntiqueWhite"
	Aqua                 = "Aqua"
	Aquamarine           = "Aquamarine"
	Azure                = "Azure"
	Beige                = "Beige"
	Bisque               = "Bisque"
	Black                = "Black"
	BlanchedAlmond       = "BlanchedAlmond"
	Blue                 = "Blue"
	BlueViolet           = "BlueViolet"
	Brown                = "Brown"
	BurlyWood            = "BurlyWood"
	CadetBlue            = "CadetBlue"
	Chartreuse           = "Chartreuse"
	Chocolate            = "Chocolate"
	Coral                = "Coral"
	CornflowerBlue       = "CornflowerBlue"
	Cornsilk             = "Cornsilk"
	Crimson              = "Crimson"
	Cyan                 = "Cyan"
	DarkBlue             = "DarkBlue"
	DarkCyan             = "DarkCyan"
	DarkGoldenrod        = "DarkGoldenrod"
	DarkGray             = "DarkGray"
	DarkGreen            = "DarkGreen"
	DarkGrey             = "DarkGrey"
	DarkKhaki            = "DarkKhaki"
	DarkMagenta          = "DarkMagenta"
	DarkOliveGreen       = "DarkOliveGreen"
	DarkOrange           = "DarkOrange"
	DarkOrchid           = "DarkOrchid"
	DarkRed              = "DarkRed"
	DarkSalmon           = "DarkSalmon"
	DarkSeaGreen         = "DarkSeaGreen"
	DarkSlateBlue        = "DarkSlateBlue"
	DarkSlateGray        = "DarkSlateGray"
	DarkSlateGrey        = "DarkSlateGrey"
	DarkTurquoise        = "DarkTurquoise"
	DarkViolet           = "DarkViolet"
	DeepPink             = "DeepPink"
	DeepSkyBlue          = "DeepSkyBlue"
	DimGray              = "DimGray"
	DodgerBlue           = "DodgerBlue"
	FireBrick            = "FireBrick"
	FloralWhite          = "FloralWhite"
	ForestGreen          = "ForestGreen"
	Fuchsia              = "Fuchsia"
	Gainsboro            = "Gainsboro"
	GhostWhite           = "GhostWhite"
	Gold                 = "Gold"
	Goldenrod            = "Goldenrod"
	Gray                 = "Gray"
	Green                = "Green"
	GreenYellow          = "GreenYellow"
	Grey                 = "Grey"
	Honeydew             = "Honeydew"
	HotPink              = "HotPink"
	IndianRed            = "IndianRed"
	Indigo               = "Indigo"
	Ivory                = "Ivory"
	Khaki                = "Khaki"
	Lavender             = "Lavender"
	LavenderBlush        = "LavenderBlush"
	LawnGreen            = "LawnGreen"
	LemonChiffon         = "LemonChiffon"
	LightBlue            = "LightBlue"
	LightCoral           = "LightCoral"
	LightCyan            = "LightCyan"
	LightGoldenrodYellow = "LightGoldenrodYellow"
	LightGray            = "LightGray"
	LightGreen           = "LightGreen"
	LightGrey            = "LightGrey"
	LightPink            = "LightPink"
	LightSalmon          = "LightSalmon"
	LightSeaGreen        = "LightSeaGreen"
	LightSkyBlue         = "LightSkyBlue"
	LightSlateGray       = "LightSlateGray"
	LightSlateGrey       = "LightSlateGrey"
	LightSteelBlue       = "LightSteelBlue"
	LightYellow          = "LightYellow"
	Lime                 = "Lime"
	LimeGreen            = "LimeGreen"
	Linen                = "Linen"
	Magenta              = "Magenta"
	Maroon               = "Maroon"
	MediumAquamarine     = "MediumAquamarine"
	MediumBlue           = "MediumBlue"
	MediumOrchid         = "MediumOrchid"
	MediumPurple         = "MediumPurple"
	MediumSeaGreen       = "MediumSeaGreen"
	MediumSlateBlue      = "MediumSlateBlue"
	MediumSpringGreen    = "MediumSpringGreen"
	MediumTurquoise      = "MediumTurquoise"
	MediumVioletRed      = "MediumVioletRed"
	MidnightBlue         = "MidnightBlue"
	MintCream            = "MintCream"
	MistyRose            = "MistyRose"
	Moccasin             = "Moccasin"
	NavajoWhite          = "NavajoWhite"
	Navy                 = "Navy"
	OldLace              = "OldLace"
	Olive                = "Olive"
	OliveDrab            = "OliveDrab"
	Orange               = "Orange"
	OrangeRed            = "OrangeRed"
	Orchid               = "Orchid"
	PaleGoldenrod        = "PaleGoldenrod"
	PaleGreen            = "PaleGreen"
	PaleTurquoise        = "PaleTurquoise"
	PaleVioletRed        = "PaleVioletRed"
	PapayaWhip           = "PapayaWhip"
	PeachPuff            = "PeachPuff"
	Peru                 = "Peru"
	Pink                 = "Pink"
	Plum                 = "Plum"
	PowderBlue           = "PowderBlue"
	Purple               = "Purple"
	Rebeccapurple        = "Rebeccapurple"
	Red                  = "Red"
	RosyBrown            = "RosyBrown"
	RoyalBlue            = "RoyalBlue"
	SaddleBrown          = "SaddleBrown"
	Salmon               = "Salmon"
	SandyBrown           = "SandyBrown"
	SeaGreen             = "SeaGreen"
	Seashell             = "Seashell"
	Sienna               = "Sienna"
	Silver               = "Silver"
	SkyBlue              = "SkyBlue"
	SlateBlue            = "SlateBlue"
	SlateGray            = "SlateGray"
	SlateGrey            = "SlateGrey"
	Snow                 = "Snow"
	SpringGreen          = "SpringGreen"
	SteelBlue            = "SteelBlue"
	Tan                  = "Tan"
	Teal                 = "Teal"
	Thistle              = "Thistle"
	Tomato               = "Tomato"
	Turquoise            = "Turquoise"
	Violet               = "Violet"
	Wheat                = "Wheat"
	White                = "White"
	WhiteSmoke           = "WhiteSmoke"
	Yellow               = "Yellow"
	YellowGreen          = "YellowGreen"
)

// Color returns a image/color.Color for the named
// CSS3 color
func Color(name string) color.Color {
	if col, ok := colorMap[strings.ToLower(name)]; ok {
		return col
	}
	return nil
}

var colorMap = map[string]color.Color{
	"aliceblue":            color.NRGBA{240, 248, 255, 255},
	"antiquewhite":         color.NRGBA{250, 235, 215, 255},
	"aqua":                 color.NRGBA{0, 255, 255, 255},
	"aquamarine":           color.NRGBA{127, 255, 212, 255},
	"azure":                color.NRGBA{240, 255, 255, 255},
	"beige":                color.NRGBA{245, 245, 220, 255},
	"bisque":               color.NRGBA{255, 228, 196, 255},
	"black":                color.NRGBA{0, 0, 0, 255},
	"blanchedalmond":       color.NRGBA{255, 235, 205, 255},
	"blue":                 color.NRGBA{0, 0, 255, 255},
	"blueviolet":           color.NRGBA{138, 43, 226, 255},
	"brown":                color.NRGBA{165, 42, 42, 255},
	"burlywood":            color.NRGBA{222, 184, 135, 255},
	"cadetblue":            color.NRGBA{95, 158, 160, 255},
	"chartreuse":           color.NRGBA{127, 255, 0, 255},
	"chocolate":            color.NRGBA{210, 105, 30, 255},
	"coral":                color.NRGBA{255, 127, 80, 255},
	"cornflowerblue":       color.NRGBA{100, 149, 237, 255},
	"cornsilk":             color.NRGBA{255, 248, 220, 255},
	"crimson":              color.NRGBA{220, 20, 60, 255},
	"cyan":                 color.NRGBA{0, 255, 255, 255},
	"darkblue":             color.NRGBA{0, 0, 139, 255},
	"darkcyan":             color.NRGBA{0, 139, 139, 255},
	"darkgoldenrod":        color.NRGBA{184, 134, 11, 255},
	"darkgray":             color.NRGBA{169, 169, 169, 255},
	"darkgreen":            color.NRGBA{0, 100, 0, 255},
	"darkgrey":             color.NRGBA{169, 169, 169, 255},
	"darkkhaki":            color.NRGBA{189, 183, 107, 255},
	"darkmagenta":          color.NRGBA{139, 0, 139, 255},
	"darkolivegreen":       color.NRGBA{85, 107, 47, 255},
	"darkorange":           color.NRGBA{255, 140, 0, 255},
	"darkorchid":           color.NRGBA{153, 50, 204, 255},
	"darkred":              color.NRGBA{139, 0, 0, 255},
	"darksalmon":           color.NRGBA{233, 150, 122, 255},
	"darkseagreen":         color.NRGBA{143, 188, 143, 255},
	"darkslateblue":        color.NRGBA{72, 61, 139, 255},
	"darkslategray":        color.NRGBA{47, 79, 79, 255},
	"darkslategrey":        color.NRGBA{47, 79, 79, 255},
	"darkturquoise":        color.NRGBA{0, 206, 209, 255},
	"darkviolet":           color.NRGBA{148, 0, 211, 255},
	"deeppink":             color.NRGBA{255, 20, 147, 255},
	"deepskyblue":          color.NRGBA{0, 191, 255, 255},
	"dimgray":              color.NRGBA{105, 105, 105, 255},
	"dodgerblue":           color.NRGBA{30, 144, 255, 255},
	"firebrick":            color.NRGBA{178, 34, 34, 255},
	"floralwhite":          color.NRGBA{255, 250, 240, 255},
	"forestgreen":          color.NRGBA{34, 139, 34, 255},
	"fuchsia":              color.NRGBA{255, 0, 255, 255},
	"gainsboro":            color.NRGBA{220, 220, 220, 255},
	"ghostwhite":           color.NRGBA{248, 248, 255, 255},
	"gold":                 color.NRGBA{255, 215, 0, 255},
	"goldenrod":            color.NRGBA{218, 165, 32, 255},
	"gray":                 color.NRGBA{128, 128, 128, 255},
	"green":                color.NRGBA{0, 128, 0, 255},
	"greenyellow":          color.NRGBA{173, 255, 47, 255},
	"grey":                 color.NRGBA{128, 128, 128, 255},
	"honeydew":             color.NRGBA{240, 255, 240, 255},
	"hotpink":              color.NRGBA{255, 105, 180, 255},
	"indianred":            color.NRGBA{205, 92, 92, 255},
	"indigo":               color.NRGBA{75, 0, 130, 255},
	"ivory":                color.NRGBA{255, 255, 240, 255},
	"khaki":                color.NRGBA{240, 230, 140, 255},
	"lavender":             color.NRGBA{230, 230, 250, 255},
	"lavenderblush":        color.NRGBA{255, 240, 245, 255},
	"lawngreen":            color.NRGBA{124, 252, 0, 255},
	"lemonchiffon":         color.NRGBA{255, 250, 205, 255},
	"lightblue":            color.NRGBA{173, 216, 230, 255},
	"lightcoral":           color.NRGBA{240, 128, 128, 255},
	"lightcyan":            color.NRGBA{224, 255, 255, 255},
	"lightgoldenrodyellow": color.NRGBA{250, 250, 210, 255},
	"lightgray":            color.NRGBA{211, 211, 211, 255},
	"lightgreen":           color.NRGBA{144, 238, 144, 255},
	"lightgrey":            color.NRGBA{211, 211, 211, 255},
	"lightpink":            color.NRGBA{255, 182, 193, 255},
	"lightsalmon":          color.NRGBA{255, 160, 122, 255},
	"lightseagreen":        color.NRGBA{32, 178, 170, 255},
	"lightskyblue":         color.NRGBA{135, 206, 250, 255},
	"lightslategray":       color.NRGBA{119, 136, 153, 255},
	"lightslategrey":       color.NRGBA{119, 136, 153, 255},
	"lightsteelblue":       color.NRGBA{176, 196, 222, 255},
	"lightyellow":          color.NRGBA{255, 255, 224, 255},
	"lime":                 color.NRGBA{0, 255, 0, 255},
	"limegreen":            color.NRGBA{50, 205, 50, 255},
	"linen":                color.NRGBA{250, 240, 230, 255},
	"magenta":              color.NRGBA{255, 0, 255, 255},
	"maroon":               color.NRGBA{128, 0, 0, 255},
	"mediumaquamarine":     color.NRGBA{102, 205, 170, 255},
	"mediumblue":           color.NRGBA{0, 0, 205, 255},
	"mediumorchid":         color.NRGBA{186, 85, 211, 255},
	"mediumpurple":         color.NRGBA{147, 112, 219, 255},
	"mediumseagreen":       color.NRGBA{60, 179, 113, 255},
	"mediumslateblue":      color.NRGBA{123, 104, 238, 255},
	"mediumspringgreen":    color.NRGBA{0, 250, 154, 255},
	"mediumturquoise":      color.NRGBA{72, 209, 204, 255},
	"mediumvioletred":      color.NRGBA{199, 21, 133, 255},
	"midnightblue":         color.NRGBA{25, 25, 112, 255},
	"mintcream":            color.NRGBA{245, 255, 250, 255},
	"mistyrose":            color.NRGBA{255, 228, 225, 255},
	"moccasin":             color.NRGBA{255, 228, 181, 255},
	"navajowhite":          color.NRGBA{255, 222, 173, 255},
	"navy":                 color.NRGBA{0, 0, 128, 255},
	"oldlace":              color.NRGBA{253, 245, 230, 255},
	"olive":                color.NRGBA{128, 128, 0, 255},
	"olivedrab":            color.NRGBA{107, 142, 35, 255},
	"orange":               color.NRGBA{255, 165, 0, 255},
	"orangered":            color.NRGBA{255, 69, 0, 255},
	"orchid":               color.NRGBA{218, 112, 214, 255},
	"palegoldenrod":        color.NRGBA{238, 232, 170, 255},
	"palegreen":            color.NRGBA{152, 251, 152, 255},
	"paleturquoise":        color.NRGBA{175, 238, 238, 255},
	"palevioletred":        color.NRGBA{219, 112, 147, 255},
	"papayawhip":           color.NRGBA{255, 239, 213, 255},
	"peachpuff":            color.NRGBA{255, 218, 185, 255},
	"peru":                 color.NRGBA{205, 133, 63, 255},
	"pink":                 color.NRGBA{255, 192, 203, 255},
	"plum":                 color.NRGBA{221, 160, 221, 255},
	"powderblue":           color.NRGBA{176, 224, 230, 255},
	"purple":               color.NRGBA{128, 0, 128, 255},
	"rebeccapurple":        color.NRGBA{102, 51, 153, 255},
	"red":                  color.NRGBA{255, 0, 0, 255},
	"rosybrown":            color.NRGBA{188, 143, 143, 255},
	"royalblue":            color.NRGBA{65, 105, 225, 255},
	"saddlebrown":          color.NRGBA{139, 69, 19, 255},
	"salmon":               color.NRGBA{250, 128, 114, 255},
	"sandybrown":           color.NRGBA{244, 164, 96, 255},
	"seagreen":             color.NRGBA{46, 139, 87, 255},
	"seashell":             color.NRGBA{255, 245, 238, 255},
	"sienna":               color.NRGBA{160, 82, 45, 255},
	"silver":               color.NRGBA{192, 192, 192, 255},
	"skyblue":              color.NRGBA{135, 206, 235, 255},
	"slateblue":            color.NRGBA{106, 90, 205, 255},
	"slategray":            color.NRGBA{112, 128, 144, 255},
	"slategrey":            color.NRGBA{112, 128, 144, 255},
	"snow":                 color.NRGBA{255, 250, 250, 255},
	"springgreen":          color.NRGBA{0, 255, 127, 255},
	"steelblue":            color.NRGBA{70, 130, 180, 255},
	"tan":                  color.NRGBA{210, 180, 140, 255},
	"teal":                 color.NRGBA{0, 128, 128, 255},
	"thistle":              color.NRGBA{216, 191, 216, 255},
	"tomato":               color.NRGBA{255, 99, 71, 255},
	"turquoise":            color.NRGBA{64, 224, 208, 255},
	"violet":               color.NRGBA{238, 130, 238, 255},
	"wheat":                color.NRGBA{245, 222, 179, 255},
	"white":                color.NRGBA{255, 255, 255, 255},
	"whitesmoke":           color.NRGBA{245, 245, 245, 255},
	"yellow":               color.NRGBA{255, 255, 0, 255},
	"yellowgreen":          color.NRGBA{154, 205, 50, 255},
}
