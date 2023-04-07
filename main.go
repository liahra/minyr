package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/liahra/minyr/yr"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Velg ett av følgende alternativer:")
	fmt.Println("'convert' for å gjennomføre en konvertering")
	fmt.Println("'average' for å finne gjennomsnittstemperaturen")
	fmt.Println("'exit' for å avslutte programmet")
	fmt.Print("> " )

	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			_, err = os.Stat("kjevik-temp-fahr-20220318-20230318.csv")
			if err != nil {
				// Filen finnes ikke, gjennomfører konvertering
				outFile, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer outFile.Close()

				writer := bufio.NewWriter(outFile)
				scanner := bufio.NewScanner(file)

				for scanner.Scan() {
					input := scanner.Text()
					output, err := yr.CelsiusToFahrenheitLine(input)
					if err != nil {
						log.Println(err)
					} else {
						fmt.Fprintln(writer, output)
						writer.Flush()
					}
				}

				if err := scanner.Err(); err != nil {
					log.Println(err)
					return
				}

				fmt.Println("Konvertering utført!")
				fmt.Print("> " )
				continue
			} else {
				fmt.Println("Filen eksisterer allerede. Vil du generere filen på nytt? (j/n)")
				fmt.Print("> " )

				scanner.Scan()
				input := scanner.Text()

				if input == "j" {
					outFile, err := os.OpenFile("kjevik-temp-fahr-20220318-20230318.csv", os.O_RDWR|os.O_CREATE, 0666)
					if err != nil {
						fmt.Println(err)
						return
					}
					defer outFile.Close()
	
					writer := bufio.NewWriter(outFile)
					scanner := bufio.NewScanner(file)
	
					for scanner.Scan() {
						input := scanner.Text()
						output, err := yr.CelsiusToFahrenheitLine(input)
						if err != nil {
							log.Println(err)
						} else {
							fmt.Fprintln(writer, output)
							writer.Flush()
						}
					}
					if err := scanner.Err(); err != nil {
						log.Println(err)
						return
					}

					fmt.Println("Konvertering utført!")
					fmt.Print("> " )
					continue
				} else if input == "n" {
					fmt.Println("Konvertering avbrutt.")
					fmt.Print("> " )
					continue
				} else {
					fmt.Println("Venligst velg (j)a eller (n)ei:")
					fmt.Print("> " )
				}
			}
		} else if input == "average" {
			fmt.Println("Ønsker du gjennomsnittet i grader Celsius eller grader Fahrenheit? (c/f)")
			fmt.Print("> " )

			var averageF string
			
			scanner.Scan()
			input = scanner.Text()

			if input == "c" {
				av, err := yr.CalculateAverageCels("kjevik-temp-celsius-20220318-20230318.csv")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Gjennomsnitt i grader Celsius: %v\n", av)
					fmt.Print("> " )
				}
			} else if input == "f" {
					av, err := yr.CalculateAverageFahr("kjevik-temp-celsius-20220318-20230318.csv", averageF)
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("Gjennomsnitt i grader Fahrenheit: %v\n", av)
						fmt.Print("> " )
					}
				} else {
					fmt.Println("Ugyldig valg")
					fmt.Print("> " )
				}
			} else {
			fmt.Println("Venligst velg convert, average eller exit:")
			fmt.Print("> " )
		}
	}
}
