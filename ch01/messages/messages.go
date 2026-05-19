package greet

func Morning(name string) string {
	return sayGoodMorning(name)
}

func sayGoodMorning(name string) string {
	return "Good morning, " + name + "!"
}
