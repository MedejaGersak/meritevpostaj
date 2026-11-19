// Paket vreme implemetira strukturo Postaja in se uporablja za ustavrjanje postaj oziroma pridobivanje informacij o postajah
package vreme

import (
	"fmt"
	"math"
)

// podatki o postaji, vsebuje vse meritve o postaji
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// funkcija ki vrne zaokroženo vrednost podane sprmenljivke na željeno decimalko natančno
func roundUp(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Ceil(val*ratio) / ratio
}

// funkcija vrne najnižjo, najvišjo in povprečno meritev postaje podane s oznako ter status ok
// če oznaka v slovarju ne obstaja je vrnjen status false
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {

	p, obstaja := postaje[oznaka]
	if !obstaja {
		ok = false
		return
	}

	if len(p.meritve) < stMeritev {
		ok = false
		return
	}

	najnizja = p.meritve[0]
	najvisja = p.meritve[0]
	povprecna = 0

	for _, v := range p.meritve {

		povprecna += v

		if v > najvisja {
			najvisja = v
		}

		if v < najnizja {
			najnizja = v
		}
	}
	povprecna = roundUp(povprecna/float64(len(p.meritve)), 2)
	ok = true
	return
}

// funkcija vrne stuct Postaja s podanimi podatki
func DodajPostajo(imePostaje string, kraj string, meritve []float64) Postaja {
	return Postaja{imePostaje, kraj, meritve}
}

// funkcija v podani slovar doda novo meritev za podano postajo z ključem oznaka
// če oznaka  ne obstaja ali pa če temperatura ni veljavna meritev ni dodana
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, zgornjaMeja float64, spodnjaMeja float64) {

	if temperatura < spodnjaMeja || temperatura > zgornjaMeja {
		fmt.Println("Temperatura ni veljavna.")
		return
	}

	p, ok := postaje[oznaka]
	if !ok {
		fmt.Println("Ta postaja ne obstaja.")
		return
	}

	p.meritve = append(p.meritve, temperatura)
	postaje[oznaka] = p

	fmt.Println("Meritev dodana.")

}

// funkcija izpiše meritve vseh postaj v slovarju
func IzpisMeritev(postaje map[string]Postaja) {

	for i, v := range postaje {

		fmt.Println(i, "-", v.ime, "("+v.kraj+"):", v.meritve)
	}
}

// funkcija izpiše poročilo o vsaki postaji v slovarju
// funkcija pokliče funkcijo statistka s pomočjo katere oblikuje poročilo
// če j emeritev za podano postajo premalo poročila ne oblikuje
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {

	for i, v := range postaje {
		min, max, povpr, ok := statistika(postaje, i, stMeritev)

		if !ok {
			fmt.Println(v.ime+":", "premalo meritev.")
			continue
		}

		var vreme string
		switch {
		case povpr >= 25.0:
			vreme = "Vroce"
		case povpr >= 15.0 && povpr < 25.0:
			vreme = "Toplo"
		case povpr >= 5.0 && povpr < 15.0:
			vreme = "Hladno"
		default:
			vreme = "Mrzlo"
		}

		fmt.Println(v.ime+":", "min", min, "°C,", "max", max, "°C,", "povprecje", povpr, "°C ->", vreme)
	}
}
