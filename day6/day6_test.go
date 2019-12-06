package day6

import (
	"testing"
)

func TestStuff(t *testing.T) {
	space := &Space{}
	lines := []string{
		"A)B",
		"B)C",
		"C)D",
		"D)E",
		"D)F",
		"D)G",
		"G)H",
		"H)I",
		"G)J",
		"H)K",
	}
	for _, l := range lines {
		a, b := ParseOrbittingLine(l)
		space.InitAOrbitsB(a, b)
	}
	if space.GetObjectByName("B").Orbit != space.GetObjectByName("A") {
		t.Errorf("B should orbit A")
	}
	if space.GetObjectByName("H").Orbit != space.GetObjectByName("G") {
		t.Errorf("H should orbit G")
	}
	if len(space.GetObjectByName("D").OrbitedBy) != 3 {
		t.Errorf("D should have 3 satellites, got %d ", len(space.GetObjectByName("D").OrbitedBy))
	}
	if space.CountOrbits() != 40 {
		t.Errorf("Expected 40 orbits, got %d", space.CountOrbits())
	}
}

func TestStuffByExample(t *testing.T) {
	space := &Space{}
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	for _, l := range lines {
		a, b := ParseOrbittingLine(l)
		space.InitAOrbitsB(a, b)
	}
	if space.GetObjectByName("C").Orbit != space.GetObjectByName("B") {
		t.Errorf("C should orbit B")
	}
	if space.GetObjectByName("H").Orbit != space.GetObjectByName("G") {
		t.Errorf("H should orbit G")
	}
	if len(space.GetObjectByName("E").OrbitedBy) != 2 {
		t.Errorf("E should have 2 satellites, got %d ", len(space.GetObjectByName("E").OrbitedBy))
	}
	if space.CountOrbits() != 42 {
		t.Errorf("Expected 42 orbits, got %d", space.CountOrbits())
	}
}

func TestStuffByExamplePart2(t *testing.T) {
	space := &Space{}
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	for _, l := range lines {
		a, b := ParseOrbittingLine(l)
		space.InitAOrbitsB(a, b)
	}
	if space.GetObjectByName("C").Orbit != space.GetObjectByName("B") {
		t.Errorf("C should orbit B")
	}
	if space.GetObjectByName("H").Orbit != space.GetObjectByName("G") {
		t.Errorf("H should orbit G")
	}
	if len(space.GetObjectByName("E").OrbitedBy) != 2 {
		t.Errorf("E should have 2 satellites, got %d ", len(space.GetObjectByName("E").OrbitedBy))
	}
	if len(space.GetObjectByName("K").OrbitedBy) != 2 {
		t.Errorf("K should have 2 satellites, got %d ", len(space.GetObjectByName("K").OrbitedBy))
	}
	if space.CountOrbits() != 54 {
		t.Errorf("Expected 54 orbits, got %d", space.CountOrbits())
	}
	distance, commonBody := space.GetCommonOrbitBody(space.GetObjectByName("YOU"), space.GetObjectByName("SAN"))
	if distance != 4 {
		t.Errorf("Distance to common body should be 4, got %d", distance)
	}
	if commonBody != space.GetObjectByName("D") {
		t.Errorf("Common body should be D, got %s", space.GetObjectByName("D").Name)
	}
}