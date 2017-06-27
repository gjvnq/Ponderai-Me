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
	Nota     NotaT
}

type SemestreT struct {
	Ano         int
	N           int // 1 ou 2
	Disciplinas []DisciplinaT
}

type HistoricoT struct {
	Aluno        string
	Universidade string
	Semestres    []SemestreT
	Medias       []int
}

func HistoricoFromJSONFile(filename string) (HistoricoT, error) {
	raw, err := ioutil.ReadFile("./pages.json")
	if err != nil {
		return HistoricoT{}, errors.New("falha ao ler o arquivo")
	}
	return HistoricoFromJSON(raw)
}

func HistoricoFromJSON(raw []byte) (HistoricoT, error) {
	var h HistoricoT
	json.Unmarshal(raw, &h)
	if json.Unmarshal(raw, &h) != nil {
		return HistoricoT{}, errors.New("falha ao entender o arquivo")
	}
	return h, nil
}
