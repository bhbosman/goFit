package fitDecl

import strconv "strconv"

type FaveroProduct uint16

const (
	FaveroProduct_AssiomaUno FaveroProduct = 10
	FaveroProduct_AssiomaDuo FaveroProduct = 12
	FaveroProduct_Invalid    FaveroProduct = 65535
)

var FaveroProductNames = map[FaveroProduct]string{
	FaveroProduct_AssiomaUno: "AssiomaUno",
	FaveroProduct_AssiomaDuo: "AssiomaDuo",
}

func (k FaveroProduct) String() string {
	if uint(k) < uint(len(FaveroProductNames)) {
		if v, ok := FaveroProductNames[k]; ok {
			return v
		}
	}
	return "FaveroProduct(" + strconv.Itoa(int(k)) + ")"
}

var FaveroProductValues = map[string]FaveroProduct{
	"AssiomaUno": FaveroProduct_AssiomaUno,
	"AssiomaDuo": FaveroProduct_AssiomaDuo,
}

func FaveroProductFromString(s string) FaveroProduct {
	if v, ok := FaveroProductValues[s]; ok {
		return v
	}
	return FaveroProduct(FaveroProduct_Invalid)
}
