package compiler

//oneBytes is a map that contains all one-byte ti-basic tokens
var oneBytes = map[string]byte{
	// "unused": 0x00,
	"DMS":       0x01,
	"Dec":       0x02,
	"Frac":      0x03,
	"->":        0x04, // store;
	"Boxplot":   0x05,
	"[":         0x06,
	"]":         0x07,
	"{":         0x08,
	"}":         0x09,
	"radian":    0x0a,
	"degree":    0x0b,
	"inverse":   0x0c,
	"^2":        0x0d, // squared
	"transpose": 0x0e,
	"^3":        0x0f,
	"(":         0x10,
	")":         0x11,
	"round(":    0x12,
	"pxl-Test(": 0x13,
	"augment(":  0x14,
	"rowSwap(":  0x15,
	"row+(":     0x16,
	"*row(":     0x17,
	"*row+(":    0x18,
	"max(":      0x19,
	"min(":      0x1a,
	"RPr(":      0x1b,
	"RPtheta(":  0x1c,
	"PRx(":      0x1d,
	"PRy(":      0x1e,
	"median(":   0x1f,
	"randM(":    0x20,
	"mean(":     0x21,
	"solve(":    0x22,
	"seq(":      0x23,
	"fnInt(":    0x24,
	"nDeriv(":   0x25,
	// "unused": 0x26,
	"fMin(":    0x27,
	"fMax(":    0x28,
	" ":        0x29,
	"\"":       0x2a,
	",":        0x2b,
	"i":        0x2c,
	"!":        0x2d,
	"CubicReg": 0x2e,
	"QuartReg": 0x2f,
	"0":        0x30,
	"1":        0x31,
	"2":        0x32,
	"3":        0x33,
	"4":        0x34,
	"5":        0x35,
	"6":        0x36,
	"7":        0x37,
	"8":        0x38,
	"9":        0x39,
	".":        0x3A,

	// this 'E' is the same as the ascii 'E'
	"Etothe10":     0x3b,
	"or":           0x3c,
	"xor":          0x3d,
	":":            0x3e,
	"\n":           0x3f,
	"and":          0x40,
	"A":            0x41,
	"B":            0x42,
	"C":            0x43,
	"D":            0x44,
	"E":            0x45,
	"F":            0x46,
	"G":            0x47,
	"H":            0x48,
	"I":            0x49,
	"J":            0x4a,
	"K":            0x4b,
	"L":            0x4c,
	"M":            0x4d,
	"N":            0x4e,
	"O":            0x4f,
	"P":            0x50,
	"Q":            0x51,
	"R":            0x52,
	"S":            0x53,
	"T":            0x54,
	"U":            0x55,
	"V":            0x56,
	"W":            0x57,
	"X":            0x58,
	"Y":            0x59,
	"Z":            0x5a,
	"theta":        0x5b,
	"prgm":         0x5f,
	"Radian":       0x64,
	"Degree":       0x65,
	"Normal":       0x66,
	"Sci":          0x67,
	"Eng":          0x68,
	"Float":        0x69,
	"=":            0x6a,
	"<":            0x6b,
	">":            0x6c,
	"<=":           0x6d, // lessthan or equal
	">=":           0x6e, // greater than or equal
	"!=":           0x6f, // not equal
	"+":            0x70,
	"-":            0x71,
	"Ans":          0x72,
	"Fix":          0x73,
	"Horiz":        0x74,
	"Full":         0x75,
	"Func":         0x76,
	"Param":        0x77,
	"Polar":        0x78,
	"Seq":          0x79,
	"IndpntAuto":   0x7a,
	"IndpntAsk":    0x7b,
	"DependAuto":   0x7c,
	"DependAsk":    0x7d,
	"weirdmark":    0x7f, // seriously what is this thing
	"plusmark":     0x80,
	"dotmark":      0x81,
	"*":            0x82,
	"/":            0x83,
	"Trace":        0x84,
	"CLrDraw":      0x85,
	"ZStandard":    0x86,
	"ZTrig":        0x87,
	"ZBox":         0x88,
	"ZoomIn":       0x89,
	"ZoomOut":      0x8A,
	"ZSquare":      0x8B,
	"ZInteger":     0x8C,
	"ZPrevious":    0x8D,
	"ZDecimal":     0x8E,
	"ZoomStat":     0x8F,
	"ZoomRcl":      0x90,
	"PrintScreen":  0x91,
	"ZoomSto":      0x92,
	"Text(":        0x93,
	"nPr":          0x94,
	"nCr":          0x95,
	"FnOn":         0x96,
	"FnOff":        0x97,
	"StorePic":     0x98,
	"RecallPic":    0x99,
	"StoreGDB":     0x9A,
	"RecallGDB":    0x9B,
	"Line(":        0x9C,
	"Vertical":     0x9D,
	"Pt-On(":       0x9E,
	"Pt-Off(":      0x9F,
	"Pt-Change(":   0xA0,
	"Pxl-On(":      0xA1,
	"Pxl-Off(":     0xA2,
	"Pxl-Change(":  0xA3,
	"Shade(":       0xA4,
	"Circle(":      0xA5,
	"Horizontal":   0xA6,
	"Tangent(":     0xA7,
	"DrawInv":      0xA8,
	"DrawF":        0xA9,
	"rand":         0xAB,
	"pi":           0xAC,
	"getKey":       0xAD,
	"'":            0xAE,
	"?":            0xAF,
	"neg ":         0xB0, // needs some work
	"int(":         0xB1,
	"abs(":         0xB2,
	"det(":         0xB3,
	"identity(":    0xB4,
	"dim(":         0xB5,
	"sum(":         0xB6,
	"prod(":        0xB7,
	"not(":         0xB8,
	"iPart(":       0xB9,
	"fPart(":       0xBA,
	"sqrt(":        0xBC,
	"crt(":         0xBD, // cube root
	"ln(":          0xBE,
	"e^(":          0xBF, // e to a power
	"log(":         0xC0,
	"10^(":         0xC1, // 10 to a power
	"sin(":         0xC2,
	"arcsin(":      0xC3,
	"cos(":         0xC4,
	"arccos(":      0xC5,
	"tan(":         0xC6,
	"arctan(":      0xC7,
	"sinh(":        0xC8,
	"arcsinh(":     0xC9,
	"cosh(":        0xCA,
	"arccosh(":     0xCB,
	"tanh(":        0xCC,
	"arctanh(":     0xCD,
	"If":           0xCE,
	"Then":         0xCF,
	"Else":         0xD0,
	"While":        0xD1,
	"Repeat":       0xD2,
	"For(":         0xD3,
	"End":          0xD4,
	"Return":       0xD5,
	"Lbl":          0xD6,
	"Goto":         0xD7,
	"Pause":        0xD8,
	"Stop":         0xD9,
	"IS>(":         0xDA,
	"DS<(":         0xDB,
	"Input":        0xDC,
	"Prompt ":      0xDD, // prompt incorporates a space
	"Disp ":        0xDE, // the disp token incorporates a space
	"DispGraph":    0xDF,
	"Output(":      0xE0,
	"ClrHome":      0xE1,
	"Fill(":        0xE2,
	"SortA(":       0xE3,
	"SortD(":       0xE4,
	"DispTable":    0xE5,
	"Menu(":        0xE6,
	"Send(":        0xE7,
	"Get(":         0xE8,
	"PlotsOn":      0xE9,
	"PlotsOff":     0xEA,
	"lists":        0xEB,
	"Plot1(":       0xEC,
	"Plot2(":       0xED,
	"Plot3(":       0xEE,
	"^":            0xF0,
	"xRootOf":      0xF1,
	"1-Var Stats":  0xF2,
	"2-Var Stats":  0xF3,
	"LinReg(a+bx)": 0xF4,
	"ExpReg":       0xF5,
	"LnReg":        0xF6,
	"PwrReg":       0xF7,
	"Med-Med":      0xF8,
	"QuadReg":      0xF9,
	"ClrList":      0xFA,
	"ClrTable":     0xFB,
	"Histogram":    0xFC,
	"xyLine":       0xFD,
	"Scatter":      0xFE,
	"LinReg(ax+b)": 0xFF,
}

var twoBytes = map[string]uint16{
	"L1": 0x5d00,
	"L2": 0x5d01,
	"L3": 0x5d02,
	"L4": 0x5d03,
	"L5": 0x5d04,
	"L6": 0x5d05,

	"Y1": 0x5e10,
	"Y2": 0x5e11,
	"Y3": 0x5e12,
	"Y4": 0x5e13,
	"Y5": 0x5e14,
	"Y6": 0x5e15,
	"Y7": 0x5e16,
	"Y8": 0x5e17,
	"Y9": 0x5e18,
	"Y0": 0x5e19,

	"Sequential": 0x7e00,
	"Simul":      0x7e01,
	"PolarGC":    0x7e02,
	"RectGC":     0x7e03,
	"CoordOn":    0x7e04,
	"CoordOff":   0x7e05,
	"Connected":  0x7e06,
	"Dot":        0x7e07,
	"AxesOn":     0x7e08,
	"AxesOff":    0x7e09,
	"GridOn":     0x7e0a,
	"GridOff":    0x7e0b,
	"LabelOn":    0x7e0c,
	"LabelOff":   0x7e0d,
	"Web":        0x7e0e,
	"Time":       0x7e0f,
	"uvAxes":     0x7e10,
	"vwAxes":     0x7e11,
	"uwAxes":     0x7e12,

	"Str1": 0xaa00,
	"Str2": 0xaa01,
	"Str3": 0xaa02,
	"Str4": 0xaa03,
	"Str5": 0xaa04,
	"Str6": 0xaa05,
	"Str7": 0xaa06,
	"Str8": 0xaa07,
	"Str9": 0xaa08,
	"Str0": 0xaa09,

	"npv(":          0xbb00,
	"irr(":          0xbb01,
	"bal(":          0xbb02,
	"sigmaprn(":     0xbb03,
	"sigmaInt(":     0xbb04,
	"Nom(":          0xbb05,
	"Eff(":          0xbb06,
	"dbd(":          0xbb07,
	"lcm(":          0xbb08,
	"gcd(":          0xbb09,
	"randInt(":      0xbb0a,
	"randBin(":      0xbb0b,
	"sub(":          0xbb0c,
	"stdDev(":       0xbb0d,
	"variance(":     0xbb0e,
	"inString(":     0xbb0f,
	"normalcdf(":    0xbb10,
	"invNorm(":      0xbb11,
	"tcdf(":         0xbb12,
	"chisquarecdf(": 0xbb13,
	"Fcdf(":         0xbb14,
	"binompdf(":     0xbb15,
	"binomcdf(":     0xbb16,
	"poissonpdf(":   0xbb17,
	"poissoncdf(":   0xbb18,
	"geometpdf(":    0xbb19,
	"geometcdf(":    0xbb1a,
	"normalpdf(":    0xbb1b,
	"tpdf(":         0xbb1c,
	"chisquarepdf(": 0xbb1d,
	"Fpdf(":         0xbb1e,
	"randNorm(":     0xbb1f,

	"a": 0xbbb0,
	"b": 0xbbb1,
	"c": 0xbbb2,
	"d": 0xbbb3,
	"e": 0xbbb4,
	"f": 0xbbb5,
	"g": 0xbbb6,
	"h": 0xbbb7,
	"i": 0xbbb8,
	"j": 0xbbb9,
	"k": 0xbbba,
	"l": 0xbbbc,
	"m": 0xbbbd,
	"n": 0xbbbe,
	"o": 0xbbbf,
	"p": 0xbbc0,
	"q": 0xbbc1,
	"r": 0xbbc2,
	"s": 0xbbc3,
	"t": 0xbbc4,
	"u": 0xbbc5,
	"v": 0xbbc6,
	"w": 0xbbc7,
	"x": 0xbbc8,
	"y": 0xbbc9,
	"z": 0xbbca,
}

var revOneBytes = reverseByteMap(oneBytes)
var revTwoBytes = reverseUint16Map(twoBytes)

func reverseByteMap(m map[string]byte) map[byte]string {
	n := make(map[byte]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func reverseUint16Map(m map[string]uint16) map[uint16]string {
	n := make(map[uint16]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
