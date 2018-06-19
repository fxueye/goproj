package bubblesort

import (
    "testing"
)

func TestBubbleSort(t *testing.T) {
    values := []int{5,4,3,2,1}
    Bubblesort(values)
    if values[0] != 1|| values[1] != 2|| values[2] != 3|| values[3] !=4||values[4] != 5{
        t.Error("Bubblesort(values):",values,"Expected 12345")
    }
}