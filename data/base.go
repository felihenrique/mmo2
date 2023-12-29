package data

type Portal struct {
	Id int
}

type Skill struct {
	Id                 int
	Name               string
	Description        string
	ExperiencePerLevel int
	ScaleFactor        float32
	MaxLevel           int
}

type Shard struct {
	Id      int
	Name    string
	Portals []Portal
}

type Item struct {
	Id          int
	Name        string
	Description string
	MaxStack    int
}
