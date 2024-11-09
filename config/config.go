// config/config.go
package config

var allowedTapeMovements = map[string]int{
	"right": 1,
	"left":  -1,
}

const (
	emptyCharacter    = " "
	visibleTapeLength = 20
)

func AllowedTapeMovements() []string {
	movements := make([]string, 0, len(allowedTapeMovements))
	for k := range allowedTapeMovements {
		movements = append(movements, k)
	}
	return movements
}

func IsAllowedMovement(direction string) bool {
	_, ok := allowedTapeMovements[direction]
	return ok
}

func TapeMovementFor(direction string) int {
	return allowedTapeMovements[direction]
}

func EmptyCharacter() string {
	return emptyCharacter
}

func VisibleTapeLength() int {
	return visibleTapeLength
}
