package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/liahra/minyr/yr"
)

// Flag variables
var convert string
var average string

func init() {
	// Flags
	flag.StringVar(&convert, "convert", "j", "Convert to Fahrenheit and create new file")
	flag.StringVar(&average, "average", "c", "Get average temperature in (c)elsius or (f)ahrenheit")

}

func main() {
	// Parse flags
	flag.Parse()

	// File to read from
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Check if file exist
	_, statErr := os.Stat("kjevik-temp-fahr-20220318-20230318.csv")

	outFileExists := !os.IsNotExist(statErr)

	// Convert, not overwrite
	if outFileExists && convert == "n" && isFlagPassed("convert") {
		log.Println("File already exists, no action taken")
		return
	}

	//  Convert and overwrite, or convert if file does not exist
	if (convert == "j" && isFlagPassed("convert")) || (convert == "n" && isFlagPassed("convert") && !outFileExists) {
		// File to read to
		outFile, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outFile.Close()

		// Writer to write to file
		writer := bufio.NewWriter(outFile)

		// Scanner to read file
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			// Reading lines one by one
			input := scanner.Text()

			// Using function from yr
			output, err := yr.CelsiusToFahrenheitLine(input)
			if err != nil {
				log.Println(err)
			} else {
				fmt.Fprintln(writer, output)
				writer.Flush()
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}
	}

	// Avarage
	if isFlagPassed("average") {
		av, err := yr.CalculateAverage("kjevik-temp-celsius-20220318-20230318.csv", average)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(av)
		}
	}
}

// Checking if flagg is specified
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
