package main

func thinkertoy(char byte) []string {
	font := make(map[byte][]string, 100)
	font[32] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	font[33] = []string{
		"  ",
		"o ",
		"| ",
		"o ",
		"  ",
		"O ",
		"  ",
		"  ",
	}
	font[34] = []string{
		"o o ",
		"| | ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	font[35] = []string{
		"      ",
		" | |  ",
		"-O-O- ",
		" | |  ",
		"-O-O- ",
		" | |  ",
		"      ",
		"      ",
	}
	font[36] = []string{
		"  | |   ",
		" -O-O-  ",
		"o | |   ",
		" -O-O-  ",
		"  | | o ",
		" -O-O-  ",
		"  | |   ",
		"        ",
	}
	font[37] = []string{
		"      ",
		"    O ",
		"o  /  ",
		"  /   ",
		" /  o ",
		"O     ",
		"      ",
		"      ",
	}
	font[38] = []string{
		"     ",
		"     ",
		"  o  ",
		" /|  ",
		"o-O- ",
		"  |  ",
		"     ",
		"     ",
	}
	font[39] = []string{
		"o ",
		"| ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	}
	font[40] = []string{
		"   ",
		" / ",
		"o  ",
		"|  ",
		"o  ",
		" \\ ",
		"   ",
		"   ",
	}
	font[41] = []string{
		"   ",
		"\\  ",
		" o ",
		" | ",
		" o ",
		"/  ",
		"   ",
		"   ",
	}
	font[42] = []string{
		"      ",
		"o | o ",
		" \\|/  ",
		"--O-- ",
		" /|\\  ",
		"o | o ",
		"      ",
		"      ",
	}
	font[43] = []string{
		"    ",
		"    ",
		" |  ",
		"-o- ",
		" |  ",
		"    ",
		"    ",
		"    ",
	}
	font[44] = []string{
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"o ",
		"| ",
		"  ",
	}
	font[45] = []string{
		"    ",
		"    ",
		"    ",
		"    ",
		"o-o ",
		"    ",
		"    ",
		"    ",
	}
	font[46] = []string{
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
		"O ",
		"  ",
		"  ",
	}
	font[47] = []string{
		"      ",
		"    o ",
		"   /  ",
		"  o   ",
		" /    ",
		"o     ",
		"      ",
		"      ",
	}
	font[48] = []string{
		"      ",
		" o-o  ",
		"o  /o ",
		"| / | ",
		"o/  o ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[49] = []string{
		"      ",
		"  0   ",
		" /|   ",
		"o |   ",
		"  |   ",
		"o-o-o ",
		"      ",
		"      ",
	}
	font[50] = []string{
		"     ",
		" --  ",
		"o  o ",
		"  /  ",
		" /   ",
		"o--o ",
		"     ",
		"     ",
	}
	font[51] = []string{
		"     ",
		"o-o  ",
		"   | ",
		" oo  ",
		"   | ",
		"o-o  ",
		"     ",
		"     ",
	}
	font[52] = []string{
		"     ",
		"o  o ",
		"|  | ",
		"o--O ",
		"   | ",
		"   o ",
		"     ",
		"     ",
	}
	font[53] = []string{
		"     ",
		"o--o ",
		"|    ",
		"o-o  ",
		"   | ",
		"o-o  ",
		"     ",
		"     ",
	}
	font[54] = []string{
		"      ",
		"  o   ",
		" /    ",
		"O--o  ",
		"o   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[55] = []string{
		"      ",
		"o---o ",
		"   /  ",
		"  o   ",
		"  |   ",
		"  o   ",
		"      ",
		"      ",
	}
	font[56] = []string{
		"      ",
		" o-o  ",
		"|   | ",
		" o-o  ",
		"|   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[57] = []string{
		"      ",
		" o-o  ",
		"|   o ",
		" o--O ",
		"   /  ",
		"  o   ",
		"      ",
		"      ",
	}
	font[58] = []string{
		"  ",
		"  ",
		"O ",
		"  ",
		"O ",
		"  ",
		"  ",
		"  ",
	}
	font[59] = []string{
		"  ",
		"  ",
		"o ",
		"  ",
		"o ",
		"| ",
		"  ",
		"  ",
	}
	font[60] = []string{
		"    ",
		"  o ",
		" /  ",
		"O   ",
		" \\  ",
		"  o ",
		"    ",
		"    ",
	}
	font[61] = []string{
		"     ",
		"     ",
		"     ",
		"o--o ",
		"o--o ",
		"     ",
		"     ",
		"     ",
	}
	font[62] = []string{
		"    ",
		"o   ",
		" \\  ",
		"  O ",
		" /  ",
		"o   ",
		"    ",
		"    ",
	}
	font[63] = []string{
		"      ",
		" o-o  ",
		"o   o ",
		"   /  ",
		"  o   ",
		"      ",
		"  O   ",
		"      ",
	}
	font[64] = []string{
		"      ",
		"  o   ",
		" / \\  ",
		"o O-o ",
		" \\    ",
		"  o-  ",
		"      ",
		"      ",
	}
	font[65] = []string{
		"      ",
		"  O   ",
		" / \\  ",
		"o---o ",
		"|   | ",
		"o   o ",
		"      ",
		"      ",
	}
	font[66] = []string{
		"      ",
		"o--o  ",
		"|   | ",
		"O--o  ",
		"|   | ",
		"o--o  ",
		"      ",
		"      ",
	}
	font[67] = []string{
		"      ",
		"  o-o ",
		" /    ",
		"O     ",
		" \\    ",
		"  o-o ",
		"      ",
		"      ",
	}
	font[68] = []string{
		"      ",
		"o-o   ",
		"|  \\  ",
		"|   O ",
		"|  /  ",
		"o-o   ",
		"      ",
		"      ",
	}
	font[69] = []string{
		"     ",
		"o--o ",
		"|    ",
		"O-o  ",
		"|    ",
		"o--o ",
		"     ",
		"     ",
	}
	font[70] = []string{
		"     ",
		"o--o ",
		"|    ",
		"O-o  ",
		"|    ",
		"o    ",
		"     ",
		"     ",
	}
	font[71] = []string{
		"      ",
		" o-o  ",
		"o     ",
		"|  -o ",
		"o   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[72] = []string{
		"     ",
		"o  o ",
		"|  | ",
		"O--O ",
		"|  | ",
		"o  o ",
		"     ",
		"     ",
	}
	font[73] = []string{
		"      ",
		"o-O-o ",
		"  |   ",
		"  |   ",
		"  |   ",
		"o-O-o ",
		"      ",
		"      ",
	}
	font[74] = []string{
		"      ",
		"    o ",
		"    | ",
		"    | ",
		"\\   o ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[75] = []string{
		"     ",
		"o  o ",
		"| /  ",
		"OO   ",
		"| \\  ",
		"o  o ",
		"     ",
		"     ",
	}
	font[76] = []string{
		"      ",
		"o     ",
		"|     ",
		"|     ",
		"|     ",
		"O---o ",
		"      ",
		"      ",
	}
	font[77] = []string{
		"      ",
		"o   o ",
		"|\\ /| ",
		"| O | ",
		"|   | ",
		"o   o ",
		"      ",
		"      ",
	}
	font[78] = []string{
		"      ",
		"o   o ",
		"|\\  | ",
		"| \\ | ",
		"|  \\| ",
		"o   o ",
		"      ",
		"      ",
	}
	font[79] = []string{
		"      ",
		" o-o  ",
		"o   o ",
		"|   | ",
		"o   o ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[80] = []string{
		"      ",
		"o--o  ",
		"|   | ",
		"O--o  ",
		"|     ",
		"o     ",
		"      ",
		"      ",
	}
	font[81] = []string{
		"      ",
		" o-o  ",
		"o   o ",
		"|   | ",
		"o   O ",
		" o-O\\ ",
		"      ",
		"      ",
	}
	font[82] = []string{
		"      ",
		"o--o  ",
		"|   | ",
		"O-Oo  ",
		"|  \\  ",
		"o   o ",
		"      ",
		"      ",
	}
	font[83] = []string{
		"      ",
		" o-o  ",
		"|     ",
		" o-o  ",
		"    | ",
		"o--o  ",
		"      ",
		"      ",
	}
	font[84] = []string{
		"      ",
		"o-O-o ",
		"  |   ",
		"  |   ",
		"  |   ",
		"  o   ",
		"      ",
		"      ",
	}
	font[85] = []string{
		"      ",
		"o   o ",
		"|   | ",
		"|   | ",
		"|   | ",
		" o-o  ",
		"      ",
		"      ",
	}
	font[86] = []string{
		"      ",
		"o   o ",
		"|   | ",
		"o   o ",
		" \\ /  ",
		"  o   ",
		"      ",
		"      ",
	}
	font[87] = []string{
		"          ",
		"o       o ",
		"|       | ",
		"o   o   o ",
		" \\ / \\ /  ",
		"  o   o   ",
		"          ",
		"          ",
	}
	font[88] = []string{
		"      ",
		"o   o ",
		" \\ /  ",
		"  O   ",
		" / \\  ",
		"o   o ",
		"      ",
		"      ",
	}
	font[89] = []string{
		"      ",
		"o   o ",
		" \\ /  ",
		"  O   ",
		"  |   ",
		"  o   ",
		"      ",
		"      ",
	}
	font[90] = []string{
		"      ",
		"o---o ",
		"   /  ",
		" -O-  ",
		" /    ",
		"o---o ",
		"      ",
		"      ",
	}
	font[91] = []string{
		"    ",
		"O-o ",
		"|   ",
		"|   ",
		"|   ",
		"O-o ",
		"    ",
		"    ",
	}
	font[92] = []string{
		"      ",
		"o     ",
		" \\    ",
		"  o   ",
		"   \\  ",
		"    o ",
		"      ",
		"      ",
	}
	font[93] = []string{
		"    ",
		"o-O ",
		"  | ",
		"  | ",
		"  | ",
		"o-O ",
		"    ",
		"    ",
	}
	font[94] = []string{
		"    ",
		" o  ",
		"/ \\ ",
		"    ",
		"    ",
		"    ",
		"    ",
		"    ",
	}
	font[95] = []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"o---o ",
		"      ",
		"      ",
	}
	font[96] = []string{
		"  ",
		"0 ",
		"| ",
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	}
	font[97] = []string{
		"     ",
		"     ",
		"     ",
		" oo  ",
		"| |  ",
		"o-o- ",
		"     ",
		"     ",
	}
	font[98] = []string{
		"     ",
		"o    ",
		"|    ",
		"O-o  ",
		"|  | ",
		"o-o  ",
		"     ",
		"     ",
	}
	font[99] = []string{
		"     ",
		"     ",
		"     ",
		" o-o ",
		"|    ",
		" o-o ",
		"     ",
		"     ",
	}
	font[100] = []string{
		"     ",
		"   o ",
		"   | ",
		" o-O ",
		"|  | ",
		" o-o ",
		"     ",
		"     ",
	}
	font[101] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		"|-' ",
		"o-o ",
		"    ",
		"    ",
	}
	font[102] = []string{
		"     ",
		" o-o ",
		" |   ",
		"-O-  ",
		" |   ",
		" o   ",
		"     ",
		"     ",
	}
	font[103] = []string{
		"     ",
		"     ",
		"     ",
		"o--o ",
		"|  | ",
		"o--O ",
		"   | ",
		"o--o ",
	}
	font[104] = []string{
		"     ",
		"o    ",
		"|    ",
		"O--o ",
		"|  | ",
		"o  o ",
		"     ",
		"     ",
	}
	font[105] = []string{
		"  ",
		"  ",
		"o ",
		"  ",
		"| ",
		"| ",
		"  ",
		"  ",
	}
	font[106] = []string{
		"      ",
		"      ",
		"    o ",
		"      ",
		"    o ",
		"    | ",
		"o   o ",
		" o-o  ",
	}
	font[107] = []string{
		"     ",
		"o    ",
		"| /  ",
		"OO   ",
		"| \\ ",
		"o  o ",
		"     ",
		"     ",
	}
	font[108] = []string{
		"  ",
		"o ",
		"| ",
		"| ",
		"| ",
		"o ",
		"  ",
		"  ",
	}
	font[109] = []string{
		"      ",
		"      ",
		"      ",
		"o-O-o ",
		"| | | ",
		"o o o ",
		"      ",
		"      ",
	}
	font[110] = []string{
		"     ",
		"     ",
		"     ",
		"o-o  ",
		"|  | ",
		"o  o ",
		"     ",
		"     ",
	}
	font[111] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		"| | ",
		"o-o ",
		"    ",
		"    ",
	}
	font[112] = []string{
		"     ",
		"     ",
		"     ",
		"o-o  ",
		"|  | ",
		"O-o  ",
		"|    ",
		"o    ",
	}
	font[113] = []string{
		"     ",
		"     ",
		"     ",
		" o-o ",
		"|  | ",
		" o-O ",
		"   | ",
		"   o ",
	}
	font[114] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		"|   ",
		"o   ",
		"    ",
		"    ",
	}
	font[115] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		" \\  ",
		"o-o ",
		"    ",
		"    ",
	}
	font[116] = []string{
		"    ",
		" o  ",
		" |  ",
		"-o- ",
		" |  ",
		" o  ",
		"    ",
		"    ",
	}
	font[117] = []string{
		"     ",
		"     ",
		"     ",
		"o  o ",
		"|  | ",
		"o--o ",
		"     ",
		"     ",
	}
	font[118] = []string{
		"      ",
		"      ",
		"      ",
		"o   o ",
		" \\ /  ",
		"  o   ",
		"      ",
		"      ",
	}
	font[119] = []string{
		"          ",
		"          ",
		"          ",
		"o   o   o ",
		" \\ / \\ /  ",
		"  o   o   ",
		"          ",
		"          ",
	}
	font[120] = []string{
		"    ",
		"    ",
		"    ",
		"\\ / ",
		" o  ",
		"/ \\",
		"    ",
		"    ",
	}
	font[121] = []string{
		"     ",
		"     ",
		"     ",
		"o  o ",
		"|  | ",
		"o--O ",
		"   | ",
		"o--o ",
	}
	font[122] = []string{
		"    ",
		"    ",
		"    ",
		"o-o ",
		" /  ",
		"o-o ",
		"    ",
		"    ",
	}
	font[123] = []string{
		"      ",
		"  o-o ",
		"  |   ",
		"o-O   ",
		"  |   ",
		"  o-o ",
		"      ",
		"      ",
	}
	font[124] = []string{
		"  ",
		"o ",
		"| ",
		"o ",
		"| ",
		"o ",
		"  ",
		"  ",
	}
	font[125] = []string{
		"      ",
		"o-o   ",
		"  |   ",
		"  O-o ",
		"  |   ",
		"o-o   ",
		"      ",
		"      ",
	}
	font[126] = []string{
		" o_ / ",
		"/  o  ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	if v, ok := font[char]; ok {
		return v
	}
	return font[32]
}
