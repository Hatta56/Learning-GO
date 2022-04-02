package main

import "fmt"

func main() {
	targetHari := 50
	penduduk := 1
	perkembangan := 3

	hari := 1
	hariKeTiga := 0
	for hari <= targetHari {
		if hari == 1 {
			fmt.Printf("Hari ke-%d: ada %d penduduk\n", hari, penduduk)
		} else if hari != 1 && hari%3 != 0 {
			penduduk *= perkembangan
			fmt.Printf("Hari ke-%d: ada %d penduduk  -> Dr Strange muncul\n", hari, penduduk)
		} else {
			hariKeTiga = 0
			penduduk /= 2
			fmt.Printf("Hari ke-%d: ada %d penduduk  -> Thanos muncul \n", hari, penduduk)
		}
		hari += 1
		hariKeTiga += 1
	}
}
