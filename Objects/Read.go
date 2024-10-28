package ascii_web

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Inputs struct holds data related to user input and file processing
type Inputs struct {
	UsreText    string   // Text input provided by the user
	UserBanner  string   // Name of the banner file to read
	DataFile    []string // Lines read from the banner file
	DataGen     []string // Generated art data
	rowDataFile []byte   // Rows of the data read from the banner file
}

// ReadFile reads the banner file corresponding to the UserBanner and processes its content
func (i *Inputs) ReadFile() {
	// Reset any previous errors
	Err.ResetError()
	// Attempt to read the banner file into rowDataFile
	i.rowDataFile, Err.bannerError = os.ReadFile("banners/" + i.UserBanner + ".txt")
	// If the error indicates the file does not exist, set an error message
	if Err.bannerError != nil {
		if os.IsNotExist(Err.bannerError) {
			Err.bannerError = fmt.Errorf("bad request")
			return
		}
		return
	} else if len(i.rowDataFile) != 6623 && i.UserBanner == "standard" || len(i.rowDataFile) != 7463 && i.UserBanner == "shadow" || len(i.rowDataFile) != 4703 && i.UserBanner == "thinkertoy" {
		fmt.Println(len(i.rowDataFile))
		Err.bannerError = fmt.Errorf("internal server error")
		return
	}

	newDataFile := strings.ReplaceAll(string(i.rowDataFile), "\r", "")
	newLine := 0
	lasteIndex := 0
	// Loop through characters in the processed data to separate lines
	for index, char := range newDataFile {
		if char == '\n' {
			newLine++
		}
		// When 9 lines are reached, append the segment to DataFile
		if newLine == 9 {
			i.DataFile = append(i.DataFile, string(newDataFile[lasteIndex+1:index+1]))
			newLine = 0
			// Update last index
			lasteIndex = index + 1
		}
	}
}

// GenTheArt generates art based on the user input and banner data
func (i *Inputs) GenTheArt() {
	// Reset any previous errors
	Err.ResetError()
	i.UsreText = strings.ReplaceAll(i.UsreText, "\r", "")
	// Check if the user text exceeds the maximum length
	if len([]rune(i.UsreText)) > 200 {
		Err.inputError = errors.New("bad request")
		return
	}
	// Validate characters in the user text
	for _, char := range i.UsreText {
		if char != '\n' && (char < 32 || char > 126) {
			Err.inputError = fmt.Errorf("this char %s is not suported", string(char))
			return
		}
	}

	inputSliced := strings.Split(i.UsreText, "\n")

	// Loop through each line to generate corresponding art
	for _, line := range inputSliced {
		for _, char := range line {
			i.DataGen = append(i.DataGen, i.DataFile[int(char)-32])
		}
		i.DataGen = append(i.DataGen, "\n")
	}
}

// ResetInput resets all fields in the Inputs struct to their initial states
func (i *Inputs) ResetInput() {
	i.rowDataFile = nil // Clear rows of data
	i.DataFile = nil    // Clear processed data
	i.DataGen = nil     // Clear generated art data
	i.UserBanner = ""   // Clear user banner
	i.UsreText = ""     // Clear user text input
}
