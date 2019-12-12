package day12

type Orbit struct {
	Moons []*Moon
}

type Moon struct {
	Name string
	X    int
	Y    int
	Z    int
	Vx int
	Vy int
	Vz int
}

func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func (moon *Moon) ApplyVelocity() {
	moon.X += moon.Vx
	moon.Y += moon.Vy
	moon.Z += moon.Vz
}

func (moon *Moon) KineticEnergy() int {
	return AbsInt(moon.Vx) + AbsInt(moon.Vy) + AbsInt(moon.Vz)
}

func (moon *Moon) PotentialEnergy() int {
	return AbsInt(moon.X) + AbsInt(moon.Y) + AbsInt(moon.Z)
}

func (moon *Moon) TotalEnergy() int {
	return moon.KineticEnergy() * moon.PotentialEnergy()
}

func (orbit *Orbit) Init()  {
	orbit.Moons = append(orbit.Moons, &Moon{
		Name: "Io",
		X:    1,
		Y:    3,
		Z:    -11,
	})
	orbit.Moons = append(orbit.Moons, &Moon{
		Name: "Europa",
		X:    17,
		Y:    -10,
		Z:    -8,
	})
	orbit.Moons = append(orbit.Moons, &Moon{
		Name: "Ganymede",
		X:    -1,
		Y:    -15,
		Z:    2,
	})
	orbit.Moons = append(orbit.Moons, &Moon{
		Name: "Callisto",
		X:    12,
		Y:    -4,
		Z:    -4,
	})
}

func (orbit *Orbit) UpdateMoonPairVelocity(moon, otherMoon *Moon) {
	if moon.X < otherMoon.X {
		moon.Vx++
		otherMoon.Vx--
	}
	if moon.X > otherMoon.X {
		moon.Vx--
		otherMoon.Vx++
	}
	if moon.Y < otherMoon.Y {
		moon.Vy++
		otherMoon.Vy--
	}
	if moon.Y > otherMoon.Y {
		moon.Vy--
		otherMoon.Vy++
	}
	if moon.Z < otherMoon.Z {
		moon.Vz++
		otherMoon.Vz--
	}
	if moon.Z > otherMoon.Z {
		moon.Vz--
		otherMoon.Vz++
	}
}

func (orbit *Orbit) UpdateMoonsVelocity() {
	done := map[string]bool{}
	for _, m := range orbit.Moons {
		for _, om := range orbit.Moons {
			_, pass := done[m.Name+om.Name]
			if m.Name != om.Name && !pass {
				done[m.Name+om.Name] = true
				done[om.Name+m.Name] = true
				orbit.UpdateMoonPairVelocity(m, om)
			}
		}
	}
}

func (orbit *Orbit) ApplyMoonsVelocity() {
	for _, m := range orbit.Moons {
		m.ApplyVelocity()
	}
}

func (orbit *Orbit) TimeStep() {
	orbit.UpdateMoonsVelocity()
	orbit.ApplyMoonsVelocity()
}