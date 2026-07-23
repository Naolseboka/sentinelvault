package geoip

func defaultRanges() []Range {
	return []Range{
		{Start: 1000000000, End: 1001000000, CC: "US"},
		{Start: 1100000000, End: 1100500000, CC: "US"},
		{Start: 2000000000, End: 2000200000, CC: "GB"},
		{Start: 2100000000, End: 2100300000, CC: "DE"},
		{Start: 2200000000, End: 2200100000, CC: "FR"},
		{Start: 2300000000, End: 2300400000, CC: "JP"},
		{Start: 2400000000, End: 2400150000, CC: "AU"},
		{Start: 2500000000, End: 2500250000, CC: "CA"},
		{Start: 2600000000, End: 2600200000, CC: "BR"},
		{Start: 2700000000, End: 2700500000, CC: "IN"},
		{Start: 2800000000, End: 2800100000, CC: "NL"},
		{Start: 2900000000, End: 2900080000, CC: "SE"},
		{Start: 3000000000, End: 3000050000, CC: "NO"},
		{Start: 3100000000, End: 3100060000, CC: "SG"},
		{Start: 3200000000, End: 3200040000, CC: "IE"},
		{Start: 3300000000, End: 3300070000, CC: "CH"},
	}
}
