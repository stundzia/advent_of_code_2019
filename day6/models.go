package day6

type Space struct {
	Objects []*SpaceObject
}

type SpaceObject struct {
	Name string
	Orbit *SpaceObject
	OrbitedBy []*SpaceObject
}

func (space *Space) GetObjectByName(name string) *SpaceObject  {
	for _, so := range space.Objects {
		if so.Name == name {
			return so
		}
	}
	return nil
}

func (space *Space) InitAOrbitsB(a, b string) {
	aObj := space.GetObjectByName(a)
	bObj := space.GetObjectByName(b)
	if bObj == nil {
		bObj = &SpaceObject{
			Name:  b,
			Orbit: nil,
		}
		space.Objects = append(space.Objects, bObj)
	}
	if aObj == nil {
		aObj = &SpaceObject{
			Name:  a,
			Orbit: bObj,
		}
		space.Objects = append(space.Objects, aObj)
	} else {
		aObj.Orbit = bObj
	}
	bObj.OrbitedBy = append(bObj.OrbitedBy, aObj)
}

func (space *Space) CountOrbits() (count int) {
	for _, obj := range space.Objects {
		for ;; {
			if obj.Orbit != nil {
				obj = obj.Orbit
				count++
			} else {
				break
			}
		}
	}
	return count
}

func (space *Space) GetCommonOrbitBody(location, destination *SpaceObject) (transfers int, commonBody *SpaceObject) {
	locOrbits := []*SpaceObject{}
	destOrbits := []*SpaceObject{}
	a := location
	b := destination
	for ;; {
		if a.Orbit != nil {
			locOrbits = append(locOrbits, a.Orbit)
			a = a.Orbit
		} else {
			break
		}
	}
	for ;; {
		if b.Orbit != nil {
			destOrbits = append(destOrbits, b.Orbit)
			b = b.Orbit
		} else {
			break
		}
	}
	for i, l := range locOrbits {
		for k, d := range destOrbits {
			if l == d && l != location {
				return i + k, l
			}
		}
	}
	return 0, nil
}

func (space *Space) GetShortestPath(location, destination *SpaceObject) int {
	transfers := 0
	transfers, location = space.GetCommonOrbitBody(location, destination)
	return transfers
}