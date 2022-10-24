package main

import (
	"fmt"
)

func flagDescription(flag int) (description []string) {
	// 1 0x1 template having multiple segments in sequencing
	// 2 0x2 each segment properly aligned according to the aligner
	// 4 0x4 segment unmapped
	// 8 0x8 next segment in the template unmapped
	// 16 0x10 SEQ being reverse complemented
	// 32 0x20 SEQ of the next segment in the template being reverse complemented
	// 64 0x40 the first segment in the template
	// 128 0x80 the last segment in the template
	// 256 0x100 secondary alignment
	// 512 0x200 not passing filters, such as platform/vendor quality controls
	// 1024 0x400 PCR or optical duplicate
	// 2048 0x800 supplementary alignment
	if flag&0x1 > 0 {
		description = append(description, "0x001 | the read is part of a pair")
	}
	if flag&0x2 > 0 {
		description = append(description, "0x002 | the read is part of a pair that aligned in a paired-end fashion")
	}
	if flag&0x4 > 0 {
		description = append(description, "0x004 | segment unmapped")
	}
	if flag&0x8 > 0 {
		description = append(description, "0x008 | the read is part of a pair and the other mate in the pair had at least one valid alignment")
	}
	if flag&0x10 > 0 {
		description = append(description, "0x010 | SEQ being reverse complemented")
	}
	if flag&0x20 > 0 {
		description = append(description, "0x020 | the read is part of a pair and the other mate in the pair aligned to the Crick strand")
	}
	if flag&0x40 > 0 {
		description = append(description, "0x040 | the read is mate 1 in a pair")
	}
	if flag&0x80 > 0 {
		description = append(description, "0x080 | the read is mate 2 in a pair")
	}
	if flag&0x100 > 0 {
		description = append(description, "0x100 | secondary alignment")
	}
	if flag&0x200 > 0 {
		description = append(description, "0x200 | not passing filters, such as platform/vendor quality controls")
	}
	if flag&0x400 > 0 {
		description = append(description, "0x400 | PCR or optical duplicate")
	}
	if flag&0x800 > 0 {
		description = append(description, "0x800 | supplementary alignment")
	}
	return
}
func main() {
	var flag int
	for {
		fmt.Print("\n-------------flag number(-1 kill):")
		fmt.Scanln(&flag)
		if flag == -1 {
			break
		}
		desc := flagDescription(flag)
		for ix, d := range desc {
			fmt.Println(ix, d)
		}
	}
}
