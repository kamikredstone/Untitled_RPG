package items

type ItemStats struct {
	Strength     int
	Constitution int
	Dexterity    int
	Intelligence int
	Charisma     int
}

type Item interface {
	GetItemLevel() int
	GetItemName() string
	GetItemType() string
	GetItemBonuses() ItemStats
}
