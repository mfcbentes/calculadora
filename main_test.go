package main

import "testing"

func TestSoma(t *testing.T) {
	// arrange
	teste := soma(3, 2, 1)
	// act
	resultado := 6
	// assert
	if teste != resultado { //assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubtrai(t *testing.T) {
	// arrange
	teste := subtrai(3, 2, 1)
	// act
	resultado := 0
	// assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	// arrange
	teste := multiplica(3, 2, 1)
	// act
	resultado := 6
	// assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivide(t *testing.T) {
	// arrange
	teste := divide(6, 2, 1)
	// act
	resultado := 3
	// assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
