package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["A"] = 1
	fmt.Println("x", x)
	fmt.Println("x[\"A\"]", x["A"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println("elements [\"He\"] : ", elements["He"])
	fmt.Println("elements", elements)

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	for key, value := range elements {
		fmt.Println("key : ", key, " value : ", value)
	}
}
