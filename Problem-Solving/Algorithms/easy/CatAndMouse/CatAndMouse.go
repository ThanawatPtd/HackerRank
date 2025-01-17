package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*What am thinking when do this problem
1.find distance from Cat and Mouse
2.check if and return
*/
/*Can Improve
1.use Abs instead of using if to check negative value
2. use return "Cat A" instead fmt.Sprintf("Cat A")
3.dont need to use else just return
*/

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	distanceFromA, distanceFromB := x-z, y-z
	if distanceFromA < 0 {
		distanceFromA = distanceFromA * -1
	}
	if distanceFromB < 0 {
		distanceFromB = distanceFromB * -1
	}
	if distanceFromA < distanceFromB {
		return fmt.Sprintf("Cat A")
	} else if distanceFromB < distanceFromA {
		return fmt.Sprintf("Cat B")
	} else {
		return fmt.Sprintf("Mouse C")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		xyz := strings.Split(readLine(reader), " ")

		xTemp, err := strconv.ParseInt(xyz[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(xyz[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		zTemp, err := strconv.ParseInt(xyz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := catAndMouse(x, y, z)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
