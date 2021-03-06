/*
--- Day 1: Chronal Calibration ---
"We've detected some temporal anomalies," one of Santa's Elves at the Temporal Anomaly Research and Detection Instrument Station tells you. She sounded pretty worried when she called you down here. "At 500-year intervals into the past, someone has been changing Santa's history!"

"The good news is that the changes won't propagate to our time stream for another 25 days, and we have a device" - she attaches something to your wrist - "that will let you fix the changes with no such propagation delay. It's configured to send you 500 years further into the past every few days; that was the best we could do on such short notice."

"The bad news is that we are detecting roughly fifty anomalies throughout time; the device will indicate fixed anomalies with stars. The other bad news is that we only have one device and you're the best person for the job! Good lu--" She taps a button on the device and you suddenly feel like you're falling. To save Christmas, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

After feeling like you've been falling for a few minutes, you look at the device's tiny screen. "Error: Device must be calibrated before first use. Frequency drift detected. Cannot maintain destination lock." Below the message, the device shows a sequence of changes in frequency (your puzzle input). A value like +6 means the current frequency increases by 6; a value like -3 means the current frequency decreases by 3.

For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.
In this example, the resulting frequency is 3.

Here are other example situations:

+1, +1, +1 results in  3
+1, +1, -2 results in  0
-1, -2, -3 results in -6
Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Couldn't open file %q: %v", filename, err)
		os.Exit(1)
	}

	acc := newAccumulator(file)

	if err := acc.accumulate(os.Stdout); err != nil {
		fmt.Printf("Couldn't accumulate result: %v", err)
		os.Exit(1)
	}
}

func newAccumulator(r io.Reader) *accumulator {
	return &accumulator{
		scanner: bufio.NewScanner(r),
		counter: 0,
		lineNum: 0,
	}
}

type accumulator struct {
	scanner *bufio.Scanner
	counter int
	lineNum int
}

func (acc *accumulator) accumulate(out io.Writer) error {
	for acc.scanner.Scan() {
		line := acc.scanner.Text()
		if err := acc.accLine(line); err != nil {
			return fmt.Errorf("Couldn't accumulate line %q: %v", err)
		}
		acc.lineNum++
	}

	_, err := fmt.Fprintf(out, "Accumulated Value: %d\n", acc.counter)
	return err
}

func (acc *accumulator) accLine(line string) error {
	if len(line) <= 0 {
		return fmt.Errorf("Line %d is blank", acc.lineNum)
	}
	switch firstChar := string(line[0]); firstChar {
	case "+":
		return acc.addNum(line[1:])
	case "-":
		return acc.subNum(line[1:])
	default:
		return fmt.Errorf("Unrecognized first character %q on line %d", firstChar, acc.lineNum)
	}
}

func (acc *accumulator) addNum(num string) error {
	toAdd, err := strconv.Atoi(num)
	if err != nil {
		return acc.numConvErr(num)
	}
	acc.counter += toAdd
	return nil
}

func (acc *accumulator) subNum(num string) error {
	toSub, err := strconv.Atoi(num)
	if err != nil {
		return acc.numConvErr(num)
	}
	acc.counter -= toSub
	return nil
}

func (acc *accumulator) numConvErr(num string) error {
	return fmt.Errorf("Couldn't convert %q to number on line %d", num, acc.lineNum)
}
