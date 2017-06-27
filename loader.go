package main

import (
	"bytes"
	"errors"
	"github.com/bjarneh/latinx"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const DisciplinaURL = "https://uspdigital.usp.br/jupiterweb/obterDisciplina?sgldis={Código}&verdis=1"

var USPTítulo = regexp.MustCompile("(?:Disciplina:\\s*)([A-Z0-9]*)(?:\\s*-\\s*)(.*?)(?:<\\/b>)")
var USPCréditoAulaRegEx = regexp.MustCompile("(?:Cr(?:é|e|&eacute;)ditos\\s*?Aula:(?:.|\\n)*?)([0-9])(?:\\s*?<\\/span>)")
var USPCréditoTabalhoRegEx = regexp.MustCompile("(?:Cr(?:é|e|&eacute;)ditos\\s*?Trabalho:(?:.|\\n)*?)([0-9])(?:\\s*?<\\/span>)")

func GetDisciplinaOnline(código string) (discp DisciplinaT, ret_err error) {
	defer func() {
		if r := recover(); r != nil {
			discp = DisciplinaT{}
			ret_err = errors.New("Falha ao analisar página da disciplina")
		}
	}()

	url := DisciplinaURL
	url = strings.Replace(url, "{Código}", código, 1)
	txt, err := webPageText(url)
	if err != nil {
		ret_err = err
		return
	}

	matches := USPCréditoAulaRegEx.FindAllStringSubmatch(txt, 10)
	cr_aula, err := strconv.Atoi(matches[0][1])

	matches = USPCréditoTabalhoRegEx.FindAllStringSubmatch(txt, 10)
	cr_trab, err := strconv.Atoi(matches[0][1])

	matches = USPTítulo.FindAllStringSubmatch(txt, 10)
	discp.Código = matches[0][1]
	discp.Nome = matches[0][2]

	discp.Créditos = cr_aula + cr_trab

	ret_err = nil
	return
}

func webPageText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("falha ao obter página")
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("falha ao obter página")
	}
	resp.Body.Close()
	str := string(buf)
	i := bytes.Index(buf, []byte("charset=iso-8859-1"))
	if i != 0 && i < 100000 {
		converter := latinx.Get(latinx.ISO_8859_1)
		utf8bytes, err := converter.Decode(buf)
		if err == nil {
			str = string(utf8bytes)
		}
	}
	return str, nil
}
