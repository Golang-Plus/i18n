package i18n

// Country Codes ISO 3166-1
var countryCodes = [][]string{
	{"AF", "AFG", "004"},
	{"AX", "ALA", "248"},
	{"AL", "ALB", "008"},
	{"DZ", "DZA", "012"},
	{"AS", "ASM", "016"},
	{"AD", "AND", "020"},
	{"AO", "AGO", "024"},
	{"AI", "AIA", "660"},
	{"AQ", "ATA", "010"},
	{"AG", "ATG", "028"},
	{"AR", "ARG", "032"},
	{"AM", "ARM", "051"},
	{"AW", "ABW", "533"},
	{"AU", "AUS", "036"},
	{"AT", "AUT", "040"},
	{"AZ", "AZE", "031"},
	{"BS", "BHS", "044"},
	{"BH", "BHR", "048"},
	{"BD", "BGD", "050"},
	{"BB", "BRB", "052"},
	{"BY", "BLR", "112"},
	{"BE", "BEL", "056"},
	{"BZ", "BLZ", "084"},
	{"BJ", "BEN", "204"},
	{"BM", "BMU", "060"},
	{"BT", "BTN", "064"},
	{"BO", "BOL", "068"},
	{"BQ", "BES", "535"},
	{"BA", "BIH", "070"},
	{"BW", "BWA", "072"},
	{"BV", "BVT", "074"},
	{"BR", "BRA", "076"},
	{"IO", "IOT", "086"},
	{"BN", "BRN", "096"},
	{"BG", "BGR", "100"},
	{"BF", "BFA", "854"},
	{"BI", "BDI", "108"},
	{"CV", "CPV", "132"},
	{"KH", "KHM", "116"},
	{"CM", "CMR", "120"},
	{"CA", "CAN", "124"},
	{"KY", "CYM", "136"},
	{"CF", "CAF", "140"},
	{"TD", "TCD", "148"},
	{"CL", "CHL", "152"},
	{"CN", "CHN", "156"},
	{"CX", "CXR", "162"},
	{"CC", "CCK", "166"},
	{"CO", "COL", "170"},
	{"KM", "COM", "174"},
	{"CG", "COG", "178"},
	{"CD", "COD", "180"},
	{"CK", "COK", "184"},
	{"CR", "CRI", "188"},
	{"CI", "CIV", "384"},
	{"HR", "HRV", "191"},
	{"CU", "CUB", "192"},
	{"CW", "CUW", "531"},
	{"CY", "CYP", "196"},
	{"CZ", "CZE", "203"},
	{"DK", "DNK", "208"},
	{"DJ", "DJI", "262"},
	{"DM", "DMA", "212"},
	{"DO", "DOM", "214"},
	{"EC", "ECU", "218"},
	{"EG", "EGY", "818"},
	{"SV", "SLV", "222"},
	{"GQ", "GNQ", "226"},
	{"ER", "ERI", "232"},
	{"EE", "EST", "233"},
	{"ET", "ETH", "231"},
	{"FK", "FLK", "238"},
	{"FO", "FRO", "234"},
	{"FJ", "FJI", "242"},
	{"FI", "FIN", "246"},
	{"FR", "FRA", "250"},
	{"GF", "GUF", "254"},
	{"PF", "PYF", "258"},
	{"TF", "ATF", "260"},
	{"GA", "GAB", "266"},
	{"GM", "GMB", "270"},
	{"GE", "GEO", "268"},
	{"DE", "DEU", "276"},
	{"GH", "GHA", "288"},
	{"GI", "GIB", "292"},
	{"GR", "GRC", "300"},
	{"GL", "GRL", "304"},
	{"GD", "GRD", "308"},
	{"GP", "GLP", "312"},
	{"GU", "GUM", "316"},
	{"GT", "GTM", "320"},
	{"GG", "GGY", "831"},
	{"GN", "GIN", "324"},
	{"GW", "GNB", "624"},
	{"GY", "GUY", "328"},
	{"HT", "HTI", "332"},
	{"HM", "HMD", "334"},
	{"VA", "VAT", "336"},
	{"HN", "HND", "340"},
	{"HK", "HKG", "344"},
	{"HU", "HUN", "348"},
	{"IS", "ISL", "352"},
	{"IN", "IND", "356"},
	{"ID", "IDN", "360"},
	{"IR", "IRN", "364"},
	{"IQ", "IRQ", "368"},
	{"IE", "IRL", "372"},
	{"IM", "IMN", "833"},
	{"IL", "ISR", "376"},
	{"IT", "ITA", "380"},
	{"JM", "JAM", "388"},
	{"JP", "JPN", "392"},
	{"JE", "JEY", "832"},
	{"JO", "JOR", "400"},
	{"KZ", "KAZ", "398"},
	{"KE", "KEN", "404"},
	{"KI", "KIR", "296"},
	{"KP", "PRK", "408"},
	{"KR", "KOR", "410"},
	{"KW", "KWT", "414"},
	{"KG", "KGZ", "417"},
	{"LA", "LAO", "418"},
	{"LV", "LVA", "428"},
	{"LB", "LBN", "422"},
	{"LS", "LSO", "426"},
	{"LR", "LBR", "430"},
	{"LY", "LBY", "434"},
	{"LI", "LIE", "438"},
	{"LT", "LTU", "440"},
	{"LU", "LUX", "442"},
	{"MO", "MAC", "446"},
	{"MK", "MKD", "807"},
	{"MG", "MDG", "450"},
	{"MW", "MWI", "454"},
	{"MY", "MYS", "458"},
	{"MV", "MDV", "462"},
	{"ML", "MLI", "466"},
	{"MT", "MLT", "470"},
	{"MH", "MHL", "584"},
	{"MQ", "MTQ", "474"},
	{"MR", "MRT", "478"},
	{"MU", "MUS", "480"},
	{"YT", "MYT", "175"},
	{"MX", "MEX", "484"},
	{"FM", "FSM", "583"},
	{"MD", "MDA", "498"},
	{"MC", "MCO", "492"},
	{"MN", "MNG", "496"},
	{"ME", "MNE", "499"},
	{"MS", "MSR", "500"},
	{"MA", "MAR", "504"},
	{"MZ", "MOZ", "508"},
	{"MM", "MMR", "104"},
	{"NA", "NAM", "516"},
	{"NR", "NRU", "520"},
	{"NP", "NPL", "524"},
	{"NL", "NLD", "528"},
	{"NC", "NCL", "540"},
	{"NZ", "NZL", "554"},
	{"NI", "NIC", "558"},
	{"NE", "NER", "562"},
	{"NG", "NGA", "566"},
	{"NU", "NIU", "570"},
	{"NF", "NFK", "574"},
	{"MP", "MNP", "580"},
	{"NO", "NOR", "578"},
	{"OM", "OMN", "512"},
	{"PK", "PAK", "586"},
	{"PW", "PLW", "585"},
	{"PS", "PSE", "275"},
	{"PA", "PAN", "591"},
	{"PG", "PNG", "598"},
	{"PY", "PRY", "600"},
	{"PE", "PER", "604"},
	{"PH", "PHL", "608"},
	{"PN", "PCN", "612"},
	{"PL", "POL", "616"},
	{"PT", "PRT", "620"},
	{"PR", "PRI", "630"},
	{"QA", "QAT", "634"},
	{"RE", "REU", "638"},
	{"RO", "ROU", "642"},
	{"RU", "RUS", "643"},
	{"RW", "RWA", "646"},
	{"BL", "BLM", "652"},
	{"SH", "SHN", "654"},
	{"KN", "KNA", "659"},
	{"LC", "LCA", "662"},
	{"MF", "MAF", "663"},
	{"PM", "SPM", "666"},
	{"VC", "VCT", "670"},
	{"WS", "WSM", "882"},
	{"SM", "SMR", "674"},
	{"ST", "STP", "678"},
	{"SA", "SAU", "682"},
	{"SN", "SEN", "686"},
	{"RS", "SRB", "688"},
	{"SC", "SYC", "690"},
	{"SL", "SLE", "694"},
	{"SG", "SGP", "702"},
	{"SX", "SXM", "534"},
	{"SK", "SVK", "703"},
	{"SI", "SVN", "705"},
	{"SB", "SLB", "090"},
	{"SO", "SOM", "706"},
	{"ZA", "ZAF", "710"},
	{"GS", "SGS", "239"},
	{"SS", "SSD", "728"},
	{"ES", "ESP", "724"},
	{"LK", "LKA", "144"},
	{"SD", "SDN", "729"},
	{"SR", "SUR", "740"},
	{"SJ", "SJM", "744"},
	{"SZ", "SWZ", "748"},
	{"SE", "SWE", "752"},
	{"CH", "CHE", "756"},
	{"SY", "SYR", "760"},
	{"TW", "TWN", "158"},
	{"TJ", "TJK", "762"},
	{"TZ", "TZA", "834"},
	{"TH", "THA", "764"},
	{"TL", "TLS", "626"},
	{"TG", "TGO", "768"},
	{"TK", "TKL", "772"},
	{"TO", "TON", "776"},
	{"TT", "TTO", "780"},
	{"TN", "TUN", "788"},
	{"TR", "TUR", "792"},
	{"TM", "TKM", "795"},
	{"TC", "TCA", "796"},
	{"TV", "TUV", "798"},
	{"UG", "UGA", "800"},
	{"UA", "UKR", "804"},
	{"AE", "ARE", "784"},
	{"GB", "GBR", "826"},
	{"US", "USA", "840"},
	{"UM", "UMI", "581"},
	{"UY", "URY", "858"},
	{"UZ", "UZB", "860"},
	{"VU", "VUT", "548"},
	{"VE", "VEN", "862"},
	{"VN", "VNM", "704"},
	{"VG", "VGB", "092"},
	{"VI", "VIR", "850"},
	{"WF", "WLF", "876"},
	{"EH", "ESH", "732"},
	{"YE", "YEM", "887"},
	{"ZM", "ZMB", "894"},
	{"ZW", "ZWE", "716"},
}
