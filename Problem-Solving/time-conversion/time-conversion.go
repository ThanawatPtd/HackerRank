package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	var timeInterval string = s[8:] // Extract AM or PM
	var hourString string = s[0:2]  // Extract hour as string
	var output string

	// Manual conversion of hourString to integer
	hour := (int(hourString[0]-'0') * 10) + int(hourString[1]-'0')

	if timeInterval == "AM" {
		// Handle AM case
		hour = hour % 12
		if hour == 0 {
			hourString = "00"
		} else {
			hourString = fmt.Sprintf("%02d", hour) // Ensure two digits
		}
		output = fmt.Sprintf("%s:%s:%s", hourString, s[3:5], s[6:8])
	} else {
		// Handle PM case
		if hour != 12 {
			hour += 12
		}
		output = fmt.Sprintf("%02d:%s:%s", hour, s[3:5], s[6:8])
	}

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
