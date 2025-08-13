package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(2, 3)
	if total != 5 {
		t.Errorf("Resultado da soma %d. Esta incorreto, Esperado: %d", total, 5)
	}
}
