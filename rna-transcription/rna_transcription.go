package strand

import (
	"strings"
)

// ToRNA returns the RNA complement (per RNA trasnscription)for a given DNA strand
func ToRNA(dna string) string {
	var rnaComplement []string

	for _, v := range dna {
		switch string(v) {
		case "A":
			rnaComplement = append(rnaComplement, "U")
		case "C":
			rnaComplement = append(rnaComplement, "G")
		case "G":
			rnaComplement = append(rnaComplement, "C")
		case "T":
			rnaComplement = append(rnaComplement, "A")
		}
	}
	return strings.Join(rnaComplement, "")
}
