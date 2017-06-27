package main_test

import (
	pond "github.com/gjvnq/ponderai-me"
	"testing"
)

const EmptyScript = ""
const AverageScript = "nota_final = (p1+p2)/2; if (nota_final < 5) {nota_final += 0.1+rec/2}"
const MaxScript = "nota_final = max(p1, p2, p3, t1+t2)"

func TestRunScriptEmpty(t *testing.T) {
}

func TestRunScriptAverage1(t *testing.T) {
	discp := pond.DisciplinaT{}
	discp.Vars = []string{"p1", "p2"}
	discp.JSScript = AverageScript
	err := discp.RunScript(5.0, false)
	if err != nil {
		t.Errorf("Error returned: %+v", err)
	}
	n_expected := 5.0
	if discp.SugestãoFinal != n_expected {
		t.Errorf("Got: %4.1f Expected: %4.1f", discp.SugestãoFinal, n_expected)
	}
}

func TestRunScriptAverage2(t *testing.T) {
	discp := pond.DisciplinaT{}
	discp.Init()
	discp.Vars = []string{"p1", "p2"}
	discp.NotasAtéAgora["p1"] = pond.NewNotaFixed(6)
	discp.JSScript = AverageScript
	err := discp.RunScript(4.0, false)
	if err != nil {
		t.Errorf("Error returned: %+v", err)
	}
	n_expected := 5.0
	if discp.SugestãoFinal != n_expected {
		t.Errorf("Got: %4.1f Expected: %4.1f", discp.SugestãoFinal, n_expected)
	}
}

func TestRunScriptAverage3(t *testing.T) {
	discp := pond.DisciplinaT{}
	discp.Init()
	discp.Vars = []string{"p1", "p2", "rec"}
	discp.NotasAtéAgora["p1"] = pond.NewNotaFixed(3)
	discp.NotasAtéAgora["p2"] = pond.NewNotaFixed(4)
	discp.JSScript = AverageScript
	err := discp.RunScript(8.0, false)
	if err != nil {
		t.Errorf("Error returned: %+v", err)
	}
	n_expected := 7.6
	if discp.SugestãoFinal != n_expected {
		t.Errorf("Got: %4.1f Expected: %4.1f", discp.SugestãoFinal, n_expected)
	}
}

func TestRunScriptMax(t *testing.T) {
	discp := pond.DisciplinaT{}
	discp.Init()
	discp.Vars = []string{"p1", "p2", "p3", "t1", "t2"}
	discp.NotasAtéAgora["p1"] = pond.NewNotaFixed(2)
	discp.NotasAtéAgora["p2"] = pond.NewNotaFixed(3)
	discp.NotasAtéAgora["p3"] = pond.NewNotaFixed(1)
	discp.NotasAtéAgora["t1"] = pond.NewNotaFixed(4.5)
	discp.JSScript = MaxScript
	err := discp.RunScript(4.0, false)
	if err != nil {
		t.Errorf("Error returned: %+v", err)
	}
	n_expected := 8.5
	if discp.SugestãoFinal != n_expected {
		t.Errorf("Got: %4.1f Expected: %4.1f", discp.SugestãoFinal, n_expected)
	}
}
