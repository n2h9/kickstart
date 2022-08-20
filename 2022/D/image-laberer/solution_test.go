package solution

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const EPSILON = 0.000001

type testData struct {
	pathToInput  string
	pathToOutput string
}

var data []testData = []testData{
	{
		pathToInput:  "testset1_input.txt",
		pathToOutput: "testset1_output.txt",
	},
	{
		pathToInput:  "testset2_input.txt",
		pathToOutput: "testset2_output.txt",
	},
	{
		pathToInput:  "testset3_input.txt",
		pathToOutput: "testset3_output.txt",
	},
}

func TestSol(t *testing.T) {
	testCase := func(inputFilePath, outputFilePath string) {
		input, err := os.Open(inputFilePath)
		assert.NoError(t, err, "must open testset input with no error")
		if err != nil {
			t.FailNow()
		}
		defer input.Close()

		output, err := os.Open(outputFilePath)
		assert.NoError(t, err, "must open testset output with no error")
		if err != nil {
			t.FailNow()
		}
		defer output.Close()

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()
		os.Stdin = input

		oldStdout := os.Stdout
		newStdOutr, newStdoutW, _ := os.Pipe()
		defer newStdoutW.Close()
		defer func() { os.Stdout = oldStdout }()
		os.Stdout = newStdoutW

		solution()
		os.Stdout = oldStdout

		scannerOutput := bufio.NewScanner(output)
		scannerOutput.Split(bufio.ScanLines)

		scannerStdout := bufio.NewScanner(newStdOutr)
		scannerStdout.Split(bufio.ScanLines)

		for scannerOutput.Scan() {
			scannerStdout.Scan()
			var expectedB, b byte
			var expectedF, f float64
			fmt.Sscanf(scannerOutput.Text(), "Case #%d: %f", &expectedB, &expectedF)
			fmt.Sscanf(scannerStdout.Text(), "Case #%d: %f", &b, &f)
			assert.Equal(t, expectedB, b, "test case number must be equal")
			delta := f - expectedF
			assert.Truef(t, delta > -1*EPSILON && delta < EPSILON, "float value must be equal %f %f", expectedF, f)
		}
	}

	for i, d := range data {
		t.Run(fmt.Sprintf("test case #%d", i), func(t *testing.T) {
			testCase(d.pathToInput, d.pathToOutput)
		})
	}
}
