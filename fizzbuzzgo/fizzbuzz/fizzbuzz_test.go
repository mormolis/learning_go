package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	result := FizzBuzz(1)
	if result != "1" {
		t.Error("FizzBuzz should return 1 if 1 is given")
	}

	result = FizzBuzz(3)
	if result != "Fizz" {
		t.Error("FizzBuzz should return Fizz if 3 is given")
	}

	result = FizzBuzz(4)
	if result != "4" {
		t.Error("FizzBuzz should return 4 if 4 is given")
	}

	result = FizzBuzz(5)
	if result != "Buzz" {
		t.Error("FizzBuzz should return Buzz if 5 is given")
	}

	result = FizzBuzz(6)
	if result != "Fizz" {
		t.Error("FizzBuzz should return Fizz id 6 is given")
	}

	result = FizzBuzz(15)
	if result != "FizzBuzz" {
		t.Error("FizzBuzz should return FizzBuzz id 15 is given")
	}

	result = FizzBuzz(30)
	if result != "FizzBuzz" {
		t.Error("FizzBuzz should return FizzBuzz id 30 is given")
	}

	result = FizzBuzz(31)
	if result != "31" {
		t.Error("FizzBuzz should return 31 id 31 is given")
	}
}
