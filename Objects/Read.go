package ascii_web

import (
	"fmt"
	"os"
	"strings"
)

type Inputs struct {
	UsreText    string
	UserBanner  string
	DataFile    []string
	DataGen     []string
	rowDataFile []byte
}

func (i *Inputs) ReadFile() {
	Err.ResitError()
	i.rowDataFile, Err.bannerError = os.ReadFile("banners/" + i.UserBanner + ".txt")
	if Err.bannerError != nil {
		if os.IsNotExist(Err.bannerError) {
			Err.bannerError = fmt.Errorf("the file %s doesn't exist", i.UserBanner)
			return
		}
		return
	} else if len(i.rowDataFile) != 6623 && i.UserBanner == "standard" || len(i.rowDataFile) != 7463 && i.UserBanner == "shadow" || len(i.rowDataFile) != 5558 && i.UserBanner == "thinkertoy" {
		Err.bannerError = fmt.Errorf("the file %s is empty or somebody have changed the data on purpose;", i.UserBanner)
		return
	}

	newDataFile := strings.ReplaceAll(string(i.rowDataFile), "\r", "")
	newLine := 0
	lasteIndex := 0
	for index, char := range newDataFile {
		if char == '\n' {
			newLine++
		}
		if newLine == 9 {
			i.DataFile = append(i.DataFile, string(newDataFile[lasteIndex+1:index+1]))
			newLine = 0
			lasteIndex = index + 1
		}
	}
}

func (i *Inputs) GenTheArt() {
	Err.ResitError()
	i.UsreText = strings.ReplaceAll(i.UsreText, "\r", "")
	for _, char := range i.UsreText {
		if char < 32 || char > 126 && char != '\n' {
			Err.inputError = fmt.Errorf("thsi char %c is not suported", char)
			return
		}
	}

	inputSliced := strings.Split(i.UsreText, "\n")

	for _, line := range inputSliced {
		for _, char := range line {
			i.DataGen = append(i.DataGen, i.DataFile[int(char)-32])
		}
		i.DataGen = append(i.DataGen, "\n")
	}
}

func (i *Inputs) ResitInput() {
	i.rowDataFile = nil
	i.DataFile = nil
	i.DataGen = nil
	i.UserBanner = ""
	i.UsreText = ""
}
