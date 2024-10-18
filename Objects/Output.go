package ascii_web

import (
	"strings"
)

type Output struct {
	AsciiOutput string
	LastResult  []string
}

// input is data type dyl zakaria */
func (iout *Output) GenAll(Input Inputs) {
	for i := 0; i < len(Input.DataGen); i++ {
		if Input.DataGen[i] == "\n" {
			iout.GenEachLine()
			iout.AsciiOutput += "\n"
			iout.LastResult = nil
		} else {
			iout.LastResult = append(iout.LastResult, strings.Split(Input.DataGen[i], "\n")...)
		}
	}
}

func (iout *Output) GenEachLine() {
	for i := 0; i < 8; i++ {
		for j := 0; j < len(iout.LastResult); j += 9 {
			iout.AsciiOutput += iout.LastResult[i+j]
		}
		iout.AsciiOutput += "\n"
	}
}

func (iout *Output) ResitOutput() {
	iout.AsciiOutput = ""
	iout.LastResult = nil
}
