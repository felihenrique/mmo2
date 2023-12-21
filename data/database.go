package data

type Attribute struct {
	Current int
	Max     int
}

type User struct {
	Id       int
	Email    string
	Password string
}

type UserCharacter struct {
	UserId      int
	CharacterId int
}

type CharacterSkill struct {
	SkillId         int
	CurrentLevel    int
	TotalExperience int
}

type ItemStack struct {
	ItemId   int
	Quantity int
}

type Character struct {
	Id          int
	Name        string
	UserId      int
	Inventory   []ItemStack
	Health      Attribute
	Mana        Attribute
	RunStamina  Attribute
	Woodcutting CharacterSkill
	Miner       CharacterSkill
}
