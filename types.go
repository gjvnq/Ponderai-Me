package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

const (
	UIState_Welcome = iota
	UIState_FileSelect
	UIState_MainView
)

type NotaT struct {
	IsFixed bool
	Val     int
	IsRange bool
	Min     int
	Max     int
}

type DisciplinaT struct {
	Nome     string
	Código   string
	Créditos int
	NotaFinal     NotaT
	VarEntrada map[string]NotaT
	VarSaídaComMin map[string]NotaT
	VarSaídaSemMin map[string]NotaT
	LuaScript string
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
	Medias       []int
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

func (N NotaT) RunScript(free_grade int, use_min_grades bool) error {
	return nil
}