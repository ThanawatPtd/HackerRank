package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */
func checkMin(cal int, minTurns *int) {
	if cal < *minTurns {
		*minTurns = cal
	}
}

func findFromStart(turns int, book [][2]int32, p int32, cal *int) {
	for i := 0; i < turns; i++ {
		if book[i][0] == p || book[i][1] == p {
			break
		}
		*cal++
	}
}

func findFromEnd(turns int, book [][2]int32, p int32, cal *int) {
	for i := turns - 1; i >= 0; i-- {
		if book[i][0] == p || book[i][1] == p {
			break
		}
		*cal++
	}
}

func pageCount(n int32, p int32) int32 {
	turns := int(math.Ceil(float64(n/2.0))) + 1
	book := make([][2]int32, turns)
	book[0][1] = 1
	minTurns := 999999
	cal := 0

	var page int32
	if page = 0; n > 1 {
		page = 2
	}
	for i := 1; i < turns; i++ {
		book[i][0] = int32(page)
		page++
		if page <= n {
			book[i][1] = int32(page)
			page++
		}
	}
	findFromStart(turns, book, p, &cal)
	checkMin(cal, &minTurns)
	cal = 0
	findFromEnd(turns, book, p, &cal)
	checkMin(cal, &minTurns)
	return int32(minTurns)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

	fmt.Fprintf(writer, "%d\n", result)

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
