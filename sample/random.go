package sample

import (
	"math/rand"

	"github.com/KhanSufiyanMirza/pcbook/pb"
	"github.com/google/uuid"
)

func randomKeyboardLayout() pb.Keyboard_Layout {

	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ

	}
	return pb.Keyboard_AZERTY
}
func randomBool() bool {
	return rand.Intn(2) == 1
}
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}
func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}
func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E",
			"Core i9",
			"Core i7",
			"Core i5",
			"Core i3",
		)
	}

	return randomStringFromSet(
		"Rayze 7",
		"Rayze 5",
		"Rayze 3",
	)
}
func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"RTX 1600-Ti",
			"RTX 1070",
		)
	}

	return randomStringFromSet(
		"Rx 590",
		"Rx 580",
		"Rx 5700-XT",
	)
}
func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func randomPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}
func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	return &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
}
func randomId() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}
