// range3
// Make me compile!
//
package main_test

import (
	"reflect"
	"testing"
)

func TestFilterEvenNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNumbers := []int{} // use range to filter even numbers

    for _, num := range numbers {
        last := num & 1 // Using a mask get the last BIT 
        if last == 0 {
            evenNumbers = append(evenNumbers, num)
        }
	}

	if !reflect.DeepEqual(evenNumbers, []int{2, 4, 6, 8, 10}) {
		t.Errorf("evenNumbers does not contain all the required even numbers, got %v", evenNumbers)
	}
}
