package twofer

// ShareWith Taking in a name, outputting who to share with. If no name is specified, it's you.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
