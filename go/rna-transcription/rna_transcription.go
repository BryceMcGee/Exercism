package strand

func ToRNA(dna string) string {
	complement := ""

	for _, v := range dna {
		switch v {
		case 'G':
			complement += string('C')
		case 'C':
			complement += string('G')
		case 'T':
			complement += string('A')
		case 'A':
			complement += string('U')
		}
	}
	return complement
}
