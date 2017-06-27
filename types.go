package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	UIState_Welcome = iota
	UIState_FileSelect
	UIState_MainView
)

type NotaT struct {
	IsFixed bool
	Val     float64
	IsRange bool
	Min     float64
	Max     float64
}

type DisciplinaT struct {
	Nome          string
	Código        string
	Créditos      int
	NotaFinal     NotaT
	SugestãoFinal float64
	Sugestões     map[string]float64
	Vars          []string // Lista de variáveis usadas pelo script
	NotasAtéAgora map[string]NotaT
	JSScript      string
}

type SemestreT struct {
	Ano         int
	N           int // 1 ou 2
	Disciplinas []DisciplinaT
}

type HistóricoT struct {
	Aluno        string
	Universidade string
	Semestres    []SemestreT
	Medias       []float64
}

func (discp *DisciplinaT) Init() {
	if discp.Sugestões == nil {
		discp.Sugestões = make(map[string]float64)
	}
	if discp.NotasAtéAgora == nil {
		discp.NotasAtéAgora = make(map[string]NotaT)
	}
	if discp.Vars == nil {
		discp.Vars = make([]string, 0)
	}
}

func (n1 NotaT) Equals(n2 NotaT) bool {
	if n1.IsRange != n2.IsRange || n1.IsFixed != n2.IsFixed {
		return false
	}
	if n1.IsFixed {
		return n1.Val == n2.Val
	}
	if n1.IsRange {
		return n1.Min == n2.Min && n1.Max == n2.Max
	}
	panic("unforseen case")
}

func (n NotaT) String() string {
	if n.IsFixed {
		return fmt.Sprintf("%4.1f", n.Val)
	}
	if n.IsRange {
		return fmt.Sprintf("[%4.1f %4.1f]", n.Min, n.Max)
	}
	return "?"
}

func NewNotaFixed(val float64) NotaT {
	n := NotaT{}
	n.IsFixed = true
	n.Val = val
	return n
}

func NewNotaRange(min, max float64) NotaT {
	n := NotaT{}
	n.IsRange = true
	n.Min = min
	n.Max = max
	return n
}

func HistóricoFromJSONFile(filename string) (HistóricoT, error) {
	raw, err := ioutil.ReadFile("./pages.json")
	if err != nil {
		return HistóricoT{}, errors.New("falha ao ler o arquivo")
	}
	return HistóricoFromJSON(raw)
}

func HistóricoFromJSON(raw []byte) (HistóricoT, error) {
	var h HistóricoT
	json.Unmarshal(raw, &h)
	if json.Unmarshal(raw, &h) != nil {
		return HistóricoT{}, errors.New("falha ao entender o arquivo")
	}
	return h, nil
}
