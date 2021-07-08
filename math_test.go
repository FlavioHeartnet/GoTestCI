package main

import "testing"


func TestSoma(t *testing.T){
	total := Soma(20,20)

	if total != 40 {
		t.Errorf("Invalid Result: %d Esperado: %d", total, 30)
	}
}