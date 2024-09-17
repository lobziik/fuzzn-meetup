package main

var thisSequenceWillCausePanic = []byte{0xDE, 0xAD, 0xBE, 0xEF}

func artificialAndLowProbableCrash(input string) {
	b := []byte(input)
	if len(b) < 20 {
		return
	}

	if b[6] == thisSequenceWillCausePanic[0] {
		if b[7] == thisSequenceWillCausePanic[1] {
			if b[8] == thisSequenceWillCausePanic[2] {
				if b[9] == thisSequenceWillCausePanic[2] {
					panic("BOOM! 0xDEADBEEF encountered!")
				}
			}
		}
	}
}
