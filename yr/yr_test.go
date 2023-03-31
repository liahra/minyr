package yr

import (
	"testing"
)

// Test 1
func TestCountLinesInFile(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "../kjevik-temp-fahr-20220318-20230318.csv", want: 16756},
	}

	for _, tc := range tests {
		got, _ := CountLinesInFile(tc.input)
		if !(tc.want == got) {
			t.Errorf("Expected %v, got: %v", tc.want, got)
		}
	}
}

// Test 2
func TestCelsiusToFahrenheitString(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitString(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

// Test 3
func TestCelsiusToFahrenheitLine(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av STUDENTENS_NAVN"},
		//{input: "Kjevik;SN39040;18.03.2022 01:50", want: ""},

	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

// Test 4
func TestAverage(t *testing.T) {
	type test struct {
		input1 string
		input2 string
		want   float64
	}
	tests := []test{
		{input1: "../kjevik-temp-celsius-20220318-20230318.csv", input2: "c", want: 8.56},
	}

	for _, tc := range tests {
		got, _ := CalculateAverage(tc.input1, tc.input2)
		if !(tc.want == got) {
			t.Errorf("expected %v, got: %v", tc.want, got)
		}
	}
}
