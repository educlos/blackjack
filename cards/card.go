package cards

import "fmt"

type Card struct {
	value  string // 2 -> 10, J, Q, K, A
	color  string // S, H, D, C
	shadow bool
}

func (c *Card) SetShadowState(newState bool) {
	c.shadow = newState
}

func (c *Card) IsShadow() bool {
	return c.shadow
}

func (c *Card) Get() string {
	return fmt.Sprintf("%s%s", c.value, c.color)
}

func (c *Card) Value() string {
	return c.value
}

func (c *Card) GetPoints() int {
	switch c.value {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "10", "J", "Q", "K":
		return 10
	case "A":
		return 1
	}

	return 0
}
