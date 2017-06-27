package main_test

import (
	pond "github.com/gjvnq/ponderai-me"
	"testing"
)

func TestGetDisciplinaOnline(t *testing.T) {
	códigos := []string{"FCI0210", "FCI0311", "FCI0127", "FCI0139", "FCI0425"}
	nomes := []string{"Acústica Física", "Arquitetura de Computadores I", "Astronomia e Astrofísica", "Biofísica", "Psicologia da Educação"}
	créditos := []int{4, 4, 4, 3, 6}

	for i := 0; i < len(códigos); i++ {
		discp, err := pond.GetDisciplinaOnline(códigos[i])
		if err != nil {
			t.Errorf("Error returned for %s: %+v", códigos[i], err)
		}
		if discp.Código != códigos[i] {
			t.Errorf("Returned wrong code. Wanted: %s Got: %s", códigos[i], discp.Código)
		}
		if discp.Nome != nomes[i] {
			t.Errorf("Returned wrong name. Wanted: %s Got: %s", nomes[i], discp.Nome)
		}
		if discp.Créditos != créditos[i] {
			t.Errorf("Returned wrong number of credits. Wanted: %d Got: %d", créditos[i], discp.Créditos)
		}
	}
}
