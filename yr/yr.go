package yr

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"


	"github.com/liahra/funtemps/conv"
)

func CountLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("kunne ikke åpne filen: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("feil under lesing av fil: %v", err)
	}
	return count, nil
}

func CalculateAverageFahr(filename string, scale string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("kunne ikke åpne filen: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	sum := 0.0

	for scanner.Scan() {
		line := scanner.Text()
		dividedString := strings.Split(line, ";")
		if len(dividedString) != 4 {
			return -1.0, errors.New("linje har ikke forventet format")
		}
		if dividedString[3] == "" || dividedString[3] == "Lufttemperatur" {
			continue
		}

		count++

		if scale == "c" {
			celsius, err := strconv.ParseFloat(dividedString[3], 64)
			if err != nil {
				return -1.0, fmt.Errorf("feil ved konvertering av Celsius: %v", err)
			}
			sum += celsius
		} else {
			fahrString, err := CelsiusToFahrenheitString(dividedString[3])
			if err != nil {
				return -1.0, fmt.Errorf("feil ved konvertering av Fahrenheit: %v", err)
			}
			fahrenheit, err := strconv.ParseFloat(fahrString, 64)
			if err != nil {
				return -1.0, fmt.Errorf("feil ved konvertering av Fahrenheit: %v", err)
			}
			sum += fahrenheit
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("feil under lesing av fil: %v", err)
	}

	if count == 0 {
		return -1.0, errors.New("ingen temperaturer funnet i filen")
	}

	average := sum / float64(count)
	return math.Round(average*10) / 10, nil
}

func CalculateAverageCels(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("kunne ikke åpne filen: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	sum := 0.0

	for scanner.Scan() {
		line := scanner.Text()
		dividedString := strings.Split(line, ";")
		if len(dividedString) != 4 {
			return -1.0, errors.New("linje har ikke forventet format")
		}
		if dividedString[3] == "" || dividedString[3] == "Lufttemperatur" {
			continue
		}

		count++

		celsius, err := strconv.ParseFloat(dividedString[3], 64)
		if err != nil {
			return -1.0, fmt.Errorf("feil ved konvertering av Celsius: %v", err)
		}
		sum += celsius
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("feil under lesing av fil: %v", err)
	}

	if count == 0 {
		return -1.0, errors.New("ingen temperaturer funnet i filen")
	}

	average := sum / float64(count)
	return math.Round(average*10) / 10, nil
}

func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

func CelsiusToFahrenheitLine(line string) (string, error) {
	dividedString := strings.Split(line, ";")
	var err error

	if len(dividedString) == 4 {
		if dividedString[3] == "" {
			return "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av C. Nilsen", nil
		} else if dividedString[3] == "Lufttemperatur" {
			return "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur", nil
		} else {
			dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
			if err != nil {
				return "", err
			}
		}

	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil
}
