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
	NotaFinal     *NotaT // nil não significa nada
	SugestãoFinal float64
	Sugestões     map[string]float64
	Vars          []string // Lista de variáveis usadas pelo script
	NotasAtéAgora map[string]*NotaT
	NotasMáximas  map[string]float64 // Para aquelas maravilhas que são trabalhos valendo 12
	JSScript      string
	Semestre      *SemestreT
}

type SemestreT struct {
	Ano int
	N   int // 1 ou 2
}

type HistóricoT struct {
	Aluno        string
	Universidade string
	Disciplinas  []*DisciplinaT
	Média        float64
}

func (discp *DisciplinaT) Init() {
	if discp.Sugestões == nil {
		discp.Sugestões = make(map[string]float64)
	}
	if discp.NotasAtéAgora == nil {
		discp.NotasAtéAgora = make(map[string]*NotaT)
	}
	if discp.Vars == nil {
		discp.Vars = make([]string, 0)
	}
}

func (hist *HistóricoT) Init() {
	if hist.Disciplinas == nil {
		hist.Disciplinas = make([]*DisciplinaT, 0)
	}
}

func (sem *SemestreT) Init() {
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

func (n NotaT) Try(attempt float64) float64 {
	if n.IsFixed {
		return n.Val
	}
	if n.IsRange {
		if attempt < n.Min {
			return n.Min
		} else if attempt > n.Max {
			return n.Max
		} else {
			return attempt
		}
	}
	panic("NotaT.IsFixed and NotaT.IsRange both false")
	return 0
}

func (n NotaT) TryNoMin(attempt float64) float64 {
	if n.IsFixed {
		return n.Val
	}
	if n.IsRange {
		if attempt > n.Max {
			return n.Max
		} else {
			return attempt
		}
	}
	panic("NotaT.IsFixed and NotaT.IsRange both false")
	return 0
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

func NovoHistórico(aluno, universidade string) *HistóricoT {
	h := HistóricoT{}
	h.Init()
	h.Aluno = aluno
	h.Universidade = universidade
	return &h
}

func NovoSemestre(ano, n int) *SemestreT {
	s := SemestreT{}
	s.Init()
	s.Ano = ano
	s.N = n
	return &s
}

func NovaDisciplina(código, nome string, créditos int) *DisciplinaT {
	d := DisciplinaT{}
	d.Init()
	d.Código = código
	d.Nome = nome
	d.Créditos = créditos
	d.Sugestões = make(map[string]float64)
	d.Vars = make([]string, 0)
	d.NotasAtéAgora = make(map[string]*NotaT)
	d.NotasMáximas = make(map[string]float64)
	return &d
}

func NovaNotaFixa(val float64) *NotaT {
	n := NotaT{}
	n.IsFixed = true
	n.Val = val
	return &n
}

func NovaNotaRange(min, max float64) *NotaT {
	n := NotaT{}
	n.IsRange = true
	if min < max {
		n.Min = min
		n.Max = max
	} else if min > max {
		n.Min = max
		n.Max = min
	} else if min == max {
		return NovaNotaFixa(min)
	}
	return &n
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
