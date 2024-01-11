package data

type Attribute struct {
	Current int32
	Max     int32
}

type User struct {
	Id       int32
	Email    string
	Password string
}

type UserCharacter struct {
	UserId      int32
	CharacterId int32
}

type CharacterSkill struct {
	SkillId         int32
	CurrentLevel    int32
	TotalExperience int32
}

type ItemStack struct {
	ItemId   int32
	Quantity int32
}

type Character struct {
	Id          int32
	Name        string
	UserId      int32
	Inventory   []ItemStack
	Health      Attribute
	Mana        Attribute
	RunStamina  Attribute
	Woodcutting CharacterSkill
	Mining      CharacterSkill
}

type Monster struct {
	Id             int32
	Name           string
	Health         float64
	AttackInterval float64
}
