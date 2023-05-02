package ejercicios

import (
	"fmt"
	"testing"
)

func TestFarolas(t *testing.T) {
	entrada := []Farolas{
		{1, 3, false},
		{4, 5, false},
		{7, 7, false},
		{11, 6, false},
		{15, 8, false},
		{26, 9, false},
		{37, 10, false},
	}
	x := []int{1, 21, 13, 4, 25, 16, 37, 45}
	//Debe encender las farolas de las posiciones 3, 15 y 37
	salida, _ := EncenderFarolas(entrada, x)
	if len(salida) != 3 {
		t.Error("Error al encender farolas")
	}
	fmt.Println(salida)
}
