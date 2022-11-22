package strand

var dnaToRna = map[rune]rune{
	'A': 'U',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

func ToRNA(dna string) string {
	rna := []rune{}
	for _, r := range dna {
		rna = append(rna, dnaToRna[r])
	}
	return string(rna)
}
