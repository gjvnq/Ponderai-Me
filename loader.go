package main

import "errors"

const DisciplinaURL = "https://www.sistemas.usp.br/jupiterweb/obterDisciplina?sgldis={Código}"

func GetDisciplinaOnline(código string) (DisciplinaT, error) {
	return DisciplinaT{}, errors.New("not implemented")
}
