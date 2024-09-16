package main

var thisSequenceWillCausePanic = []byte{0xDE, 0xAD, 0xBE, 0xEF}

func artificialAndLowProbableCrash(input []byte) {
	if len(input) < 20 {
		return
	}

	if input[6] == thisSequenceWillCausePanic[0] {
		if input[7] == thisSequenceWillCausePanic[1] {
			if input[8] == thisSequenceWillCausePanic[2] {
				if input[9] == thisSequenceWillCausePanic[2] {
					panic("BOOM! 0xDEADBEEF encountered!")
				}
			}
		}
	}
}
