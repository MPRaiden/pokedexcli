package models

type Pokemon struct {
	Name           string
	BaseExperience float64
	Height         float64
	Weight         float64
	Stats          []Stat
	Types          []Type
}

type Stat struct {
	BaseStat int
	Effort   int
	Name     string
}

type Type struct {
	Slot int
	Name string
}

