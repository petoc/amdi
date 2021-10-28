package amdi

var codes = map[string]string{
	"Watch1,1":   "Watch 38mm case",
	"Watch1,2":   "Watch 42mm case",
	"Watch2,3":   "Watch Series 2 38mm case",
	"Watch2,4":   "Watch Series 2 42mm case",
	"Watch2,6":   "Watch Series 1 38mm case",
	"Watch2,7":   "Watch Series 1 42mm case",
	"Watch3,1":   "Watch Series 3 38mm case (GPS+Cellular)",
	"Watch3,2":   "Watch Series 3 42mm case (GPS+Cellular)",
	"Watch3,3":   "Watch Series 3 38mm case (GPS)",
	"Watch3,4":   "Watch Series 3 42mm case (GPS)",
	"Watch4,1":   "Watch Series 4 40mm case (GPS)",
	"Watch4,2":   "Watch Series 4 44mm case (GPS)",
	"Watch4,3":   "Watch Series 4 40mm case (GPS+Cellular)",
	"Watch4,4":   "Watch Series 4 44mm case (GPS+Cellular)",
	"Watch5,1":   "Watch Series 5 40mm case (GPS)",
	"Watch5,2":   "Watch Series 5 44mm case (GPS)",
	"Watch5,3":   "Watch Series 5 40mm case (GPS+Cellular)",
	"Watch5,4":   "Watch Series 5 44mm case (GPS+Cellular)",
	"Watch5,9":   "Watch SE 40mm case (GPS)",
	"Watch5,10":  "Watch SE 44mm case (GPS)",
	"Watch5,11":  "Watch SE 40mm case (GPS+Cellular)",
	"Watch5,12":  "Watch SE 44mm case (GPS+Cellular)",
	"Watch6,1":   "Watch Series 6 40mm case (GPS)",
	"Watch6,2":   "Watch Series 6 44mm case (GPS)",
	"Watch6,3":   "Watch Series 6 40mm case (GPS+Cellular)",
	"Watch6,4":   "Watch Series 6 44mm case (GPS+Cellular)",
	"Watch6,6":   "Watch Series 7 41mm case (GPS)",
	"Watch6,7":   "Watch Series 7 45mm case (GPS)",
	"Watch6,8":   "Watch Series 7 41mm case (GPS+Cellular)",
	"Watch6,9":   "Watch Series 7 45mm case (GPS+Cellular)",
	"arm64":      "iPhone Simulator",
	"i386":       "iPhone Simulator",
	"iPad1,1":    "iPad",
	"iPad1,2":    "iPad 3G",
	"iPad2,1":    "iPad 2nd Gen",
	"iPad2,2":    "iPad 2nd Gen (GSM)",
	"iPad2,3":    "iPad 2nd Gen (CDMA)",
	"iPad2,4":    "iPad 2nd Gen (New Revision)",
	"iPad2,5":    "iPad mini",
	"iPad2,6":    "iPad mini (GSM+LTE)",
	"iPad2,7":    "iPad mini (CDMA+LTE)",
	"iPad3,1":    "iPad 3rd Gen",
	"iPad3,2":    "iPad 3rd Gen (CDMA)",
	"iPad3,3":    "iPad 3rd Gen (GSM)",
	"iPad3,4":    "iPad 4th Gen",
	"iPad3,5":    "iPad 4th Gen (GSM+LTE)",
	"iPad3,6":    "iPad 4th Gen (CDMA+LTE)",
	"iPad4,1":    "iPad Air (WiFi)",
	"iPad4,2":    "iPad Air (GSM+CDMA)",
	"iPad4,3":    "iPad Air (China)",
	"iPad4,4":    "iPad mini Retina (WiFi)",
	"iPad4,5":    "iPad mini Retina (GSM+CDMA)",
	"iPad4,6":    "iPad mini Retina (China)",
	"iPad4,7":    "iPad mini 3 (WiFi)",
	"iPad4,8":    "iPad mini 3 (GSM+CDMA)",
	"iPad4,9":    "iPad mini 3 (China)",
	"iPad5,1":    "iPad mini 4 (WiFi)",
	"iPad5,2":    "iPad mini 4 (WiFi+Cellular)",
	"iPad5,3":    "iPad Air 2 (WiFi)",
	"iPad5,4":    "iPad Air 2 (Cellular)",
	"iPad6,3":    "iPad Pro (9.7 inch, WiFi)",
	"iPad6,4":    "iPad Pro (9.7 inch, WiFi+LTE)",
	"iPad6,7":    "iPad Pro (12.9 inch, WiFi)",
	"iPad6,8":    "iPad Pro (12.9 inch, WiFi+LTE)",
	"iPad6,11":   "iPad (2017)",
	"iPad6,12":   "iPad (2017)",
	"iPad7,1":    "iPad Pro 2nd Gen (WiFi)",
	"iPad7,11":   "iPad 7th Gen 10.2 inch (WiFi)",
	"iPad7,12":   "iPad 7th Gen 10.2 inch (WiFi+Cellular)",
	"iPad7,2":    "iPad Pro 2nd Gen (WiFi+Cellular)",
	"iPad7,3":    "iPad Pro 10.5 inch 2nd Gen",
	"iPad7,4":    "iPad Pro 10.5 inch 2nd Gen",
	"iPad7,5":    "iPad 6th Gen (WiFi)",
	"iPad7,6":    "iPad 6th Gen (WiFi+Cellular)",
	"iPad8,1":    "iPad Pro 11 inch 3rd Gen (WiFi)",
	"iPad8,2":    "iPad Pro 11 inch 3rd Gen (1TB, WiFi)",
	"iPad8,3":    "iPad Pro 11 inch 3rd Gen (WiFi+Cellular)",
	"iPad8,4":    "iPad Pro 11 inch 3rd Gen (1TB, WiFi+Cellular)",
	"iPad8,5":    "iPad Pro 12.9 inch 3rd Gen (WiFi)",
	"iPad8,6":    "iPad Pro 12.9 inch 3rd Gen (1TB, WiFi)",
	"iPad8,7":    "iPad Pro 12.9 inch 3rd Gen (WiFi+Cellular)",
	"iPad8,8":    "iPad Pro 12.9 inch 3rd Gen (1TB, WiFi+Cellular)",
	"iPad8,9":    "iPad Pro 11 inch 4th Gen (WiFi)",
	"iPad8,10":   "iPad Pro 11 inch 4th Gen (WiFi+Cellular)",
	"iPad8,11":   "iPad Pro 12.9 inch 4th Gen (WiFi)",
	"iPad8,12":   "iPad Pro 12.9 inch 4th Gen (WiFi+Cellular)",
	"iPad11,1":   "iPad mini 5th Gen (WiFi)",
	"iPad11,2":   "iPad mini 5th Gen",
	"iPad11,3":   "iPad Air 3rd Gen (WiFi)",
	"iPad11,4":   "iPad Air 3rd Gen",
	"iPad11,6":   "iPad 8th Gen (WiFi)",
	"iPad11,7":   "iPad 8th Gen (WiFi+Cellular)",
	"iPad12,1":   "iPad 9th Gen (WiFi)",
	"iPad12,2":   "iPad 9th Gen (WiFi+Cellular)",
	"iPad13,1":   "iPad Air 4th Gen (WiFi)",
	"iPad13,2":   "iPad Air 4th Gen (WiFi+Cellular)",
	"iPad13,4":   "iPad Pro 11 inch 5rd Gen",
	"iPad13,5":   "iPad Pro 11 inch 5rd Gen",
	"iPad13,6":   "iPad Pro 11 inch 5rd Gen",
	"iPad13,7":   "iPad Pro 11 inch 5rd Gen",
	"iPad13,8":   "iPad Pro 12.9 inch 5th Gen",
	"iPad13,9":   "iPad Pro 12.9 inch 5th Gen",
	"iPad13,10":  "iPad Pro 12.9 inch 5th Gen",
	"iPad13,11":  "iPad Pro 12.9 inch 5th Gen",
	"iPad14,1":   "iPad mini 6th Gen (WiFi)",
	"iPad14,2":   "iPad mini 6th Gen (WiFi+Cellular)",
	"iPhone1,1":  "iPhone",
	"iPhone1,2":  "iPhone 3G",
	"iPhone2,1":  "iPhone 3GS",
	"iPhone3,1":  "iPhone 4",
	"iPhone3,2":  "iPhone 4 (GSM Rev A)",
	"iPhone3,3":  "iPhone 4 (CDMA)",
	"iPhone4,1":  "iPhone 4S",
	"iPhone5,1":  "iPhone 5 (GSM)",
	"iPhone5,2":  "iPhone 5 (GSM+CDMA)",
	"iPhone5,3":  "iPhone 5C (GSM)",
	"iPhone5,4":  "iPhone 5C (Global)",
	"iPhone6,1":  "iPhone 5S (GSM)",
	"iPhone6,2":  "iPhone 5S (Global)",
	"iPhone7,1":  "iPhone 6 Plus",
	"iPhone7,2":  "iPhone 6",
	"iPhone8,1":  "iPhone 6s",
	"iPhone8,2":  "iPhone 6s Plus",
	"iPhone8,4":  "iPhone SE (GSM)",
	"iPhone9,1":  "iPhone 7",
	"iPhone9,2":  "iPhone 7 Plus",
	"iPhone9,3":  "iPhone 7",
	"iPhone9,4":  "iPhone 7 Plus",
	"iPhone10,1": "iPhone 8",
	"iPhone10,2": "iPhone 8 Plus",
	"iPhone10,3": "iPhone X (Global)",
	"iPhone10,4": "iPhone 8",
	"iPhone10,5": "iPhone 8 Plus",
	"iPhone10,6": "iPhone X (GSM)",
	"iPhone11,2": "iPhone XS",
	"iPhone11,4": "iPhone XS Max",
	"iPhone11,6": "iPhone XS Max (Global)",
	"iPhone11,8": "iPhone XR",
	"iPhone12,1": "iPhone 11",
	"iPhone12,3": "iPhone 11 Pro",
	"iPhone12,5": "iPhone 11 Pro Max",
	"iPhone12,8": "iPhone SE 2nd Gen",
	"iPhone13,1": "iPhone 12 mini",
	"iPhone13,2": "iPhone 12",
	"iPhone13,3": "iPhone 12 Pro",
	"iPhone13,4": "iPhone 12 Pro Max",
	"iPhone14,2": "iPhone 13 Pro",
	"iPhone14,3": "iPhone 13 Pro Max",
	"iPhone14,4": "iPhone 13 mini",
	"iPhone14,5": "iPhone 13",
	"iPod1,1":    "iPod",
	"iPod2,1":    "iPod 2nd Gen",
	"iPod3,1":    "iPod 3rd Gen",
	"iPod4,1":    "iPod 4th Gen",
	"iPod5,1":    "iPod 5th Gen",
	"iPod7,1":    "iPod 6th Gen",
	"iPod9,1":    "iPod 7th Gen",
	"x86_64":     "iPhone Simulator",
}

// Is ...
func Is(code string) bool {
	_, ok := codes[code]
	return ok
}

// Name ...
func Name(code string) string {
	return codes[code]
}
