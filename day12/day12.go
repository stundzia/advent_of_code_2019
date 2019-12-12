package day12

import "fmt"

func DoSilver() {
	orbit := Orbit{}
	orbit.Init()
	for i := 0 ; i < 1000; i++ {
		orbit.TimeStep()
	}
	kinetic := 0
	potential := 0
	total := 0
	for _, m := range orbit.Moons {
		fmt.Println(m)
		kinetic += m.KineticEnergy()
		potential += m.PotentialEnergy()
		total += m.TotalEnergy()
	}
	fmt.Println("Kinetic: ", kinetic)
	fmt.Println("Potential: ", potential)
	fmt.Println("Total: ", total)
}



func DoGold() {

}
