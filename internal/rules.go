package internal

var rand *RandSource

func init() {
	rand = NewRandSource()
}

func GenerateBool(rule string) bool {
	if rule == "" {
		return rand.Bool()
	}

	return false
}
