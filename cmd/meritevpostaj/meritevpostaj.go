package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MedejaGersak/meritevpostaj/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vremenskaPostaja",
		Usage: "Vremenska postaja hrani in izpisuje teperaturo postaj",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Usage: "Najnižja dovoljena temperatura meritve",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Usage: "Najvišja dovoljena temperatura meritve",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev, potrebnih za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			return runPostaja(minTemp, maxTemp, stMeritev)

		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func runPostaja(minTemp float64, maxTemp float64, stMeritev int) error {

	var slovarPostaj map[string]vreme.Postaja

	slovarPostaj = make(map[string]vreme.Postaja)

	slovarPostaj["LJ01"] = vreme.DodajPostajo("Bavarski dvor", "Ljubljana", []float64{35.7, 12.9, 17.5, 27.5})
	slovarPostaj["KP01"] = vreme.DodajPostajo("Koper Luka", "Koper", []float64{38.7, 21.3, 29.5, 31.2})
	slovarPostaj["KR01"] = vreme.DodajPostajo("Kredarica", "Triglav", []float64{9.2, 7.8, 19.4, 22.2})
	slovarPostaj["MB01"] = vreme.DodajPostajo("Maribor Tabor", "Maribor", []float64{24.2, 16.9})

	vreme.DodajMeritev(slovarPostaj, "LJ01", 7.5, maxTemp, minTemp)
	fmt.Print("\n")
	vreme.DodajMeritev(slovarPostaj, "LJ01", -60.6, maxTemp, minTemp)
	fmt.Print("\n")
	vreme.IzpisMeritev(slovarPostaj)
	fmt.Print("\n")
	vreme.IzpisPorocila(slovarPostaj, stMeritev)
	fmt.Print("\n")

	return nil
}
