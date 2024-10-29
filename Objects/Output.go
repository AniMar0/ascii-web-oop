package ascii_web

import (
	"net/http"
	"strconv"
	"strings"
)

// Output struct holds the generated ASCII art and related data
type Output struct {
	AsciiOutput string   // The final generated ASCII output
	LastResult  []string // Temporary storage for the last processed lines
	Old         string   // Holds previous output data
}

// GenAll generates the complete ASCII output based on the provided Inputs
func (iout *Output) GenAll(Input Inputs) {
	// Start with a newline because when clicking on button submit it removes a newline
	iout.AsciiOutput += "\n"
	// Iterate through the generated data from Input
	for i := 0; i < len(Input.DataGen); i++ {
		if Input.DataGen[i] == "\n" {
			// If LastResult has data, generate output for the current line
			if iout.LastResult != nil {
				iout.GenEachLine()
			}
			iout.AsciiOutput += "\n" // Add a newline to the output
			iout.LastResult = nil    // Reset LastResult for the next segment
		} else {
			// Split the current data segment by newline and append to LastResult
			iout.LastResult = append(iout.LastResult, strings.Split(Input.DataGen[i], "\n")...)
		}
	}
}

// GenEachLine generates ASCII output
func (iout *Output) GenEachLine() {
	for i := 0; i < 8; i++ {
		// For each line, append corresponding data from LastResult
		for j := 0; j < len(iout.LastResult); j += 9 {
			iout.AsciiOutput += iout.LastResult[i+j]
		}
		// If not the last line, add a newline
		if i != 7 {
			iout.AsciiOutput += "\n"
		}
	}
}

// ResetOutput resets the Output fields to their initial state
func (iout *Output) ResetOutput() {
	iout.AsciiOutput = "" // Clear the ASCII output
	iout.LastResult = nil // Clear the last result data
	iout.Old = ""
}

func (iout *Output) Download(Writer http.ResponseWriter) {
	Writer.Header().Set("Content-Type", "text/plain")
	Writer.Header().Set("Content-Disposition", "attachment; filename=download.txt")
	Writer.Header().Set("Content-Length", strconv.Itoa(len(iout.AsciiOutput[1:])))

	if _, err := Writer.Write([]byte(iout.AsciiOutput[1:])); err != nil {
		Err.RenderErrorPage(Writer, err, http.StatusInternalServerError)
	}
}

func (iout *Output) Noting(Writer http.ResponseWriter) {
	Writer.Header().Set("Content-Type", "text/plain")
	if _, err := Writer.Write([]byte("Noting to download")); err != nil {
		Err.RenderErrorPage(Writer, err, http.StatusInternalServerError)
	}
}
