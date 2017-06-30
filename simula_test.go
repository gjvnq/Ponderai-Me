package main_test

import (
	pond "github.com/gjvnq/ponderai-me"
	"testing"
)

const FísicaScript = "nota_final = (p1+p2+p3)/3; if (nota_final < 5) { nota_final += rec; nota_final /= 2; }"

func TestSimulaPasso1(t *testing.T) {
	cálculo := pond.NovaDisciplina("", "Cálculo I", 4)
	física := pond.NovaDisciplina("", "Física I", 6)
	cálculo.NotaFinal = pond.NovaNotaRange(5, 7)
	física.JSScript = FísicaScript
	física.Vars = []string{"p1", "p2", "p3", "rec"}
	física.NotasAtéAgora["p1"] = pond.NovaNotaFixa(2)
	física.NotasAtéAgora["p2"] = pond.NovaNotaRange(5, 7)
	física.NotasMáximas["rec"] = 12
	sem := pond.NovoSemestre(2017, 1)
	sem.Disciplinas = []*pond.DisciplinaT{cálculo, física}
	hist := pond.NovoHistórico("Zé", "USP")
	hist.Semestres = []*pond.SemestreT{sem}

	err := pond.SimulaPasso(hist, 30, true)
	n_expected := 8.3
	if hist.Média != n_expected {
		t.Errorf("Got: %4.1f Expected: %4.1f", hist.Média, n_expected)
	}
	if err != nil {
		t.Errorf("Unexpected: %+v", err)
	}
}
