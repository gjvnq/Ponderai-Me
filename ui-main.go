package main

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/go-ansi"
	. "github.com/logrusorgru/aurora"
	"io/ioutil"
	"encoding/json"
	"os"
	"strings"
)

func simpleReadLine() (string, error) {
	stdioScanner := bufio.NewScanner(os.Stdin)
	if !stdioScanner.Scan() {
		return "", stdioScanner.Err()
	}
	return stdioScanner.Text(), nil
}

func help() {
	fmt.Println("Comandos disponíveis:")
	cmd := []string{
		"abrir [arquivo]",
		"salvar [arquivo]",
		"sair",
		"ajuda",
		"hist",
		"add",
		"edt [número no histórico]",
		"del [número no histórico]",
		"meta [média]",
		"otimista",
		"pessimista"}
	cmd_help := []string{
		"Abre um arquivo de histórico escolar",
		"Salva o arquivo de histórico escolar",
		"Saí sem salvar",
		"Mostra esse texto",
		"Mostra o histórico escolar em uso e a ponderada",
		"Mostra o histórico escolar em uso e a ponderada",
		"Adiciona uma disciplina",
		"Edita uma disciplina",
		"Remove uma disciplina",
		"Calcula as notas necessárias para obter a ponderada desejada",
		"Calcula a maior média ponderada possível",
		"Calcula a menor média ponderada possível"}
	for i := 0; i < len(cmd); i++ {
		ansi.Printf("  %s", Green(cmd[i]))
		ansi.CursorHorizontalAbsolute(30)
		ansi.Printf("%s\n", cmd_help[i])
	}
}

func histórico(cmd []string) {
	fmt.Println(cmd)
}

func salvar(cmd []string) {
	filename := Filename
	if len(cmd) > 0 {
		filename = cmd[0]
		if !strings.HasSuffix(filename, ".json") {
			filename += ".json"
		}
	}
	if len(cmd) == 0 && Filename == "" {
		ansi.Println(Brown("Por favor, informa o nome do arquivo. Ex: salvar meu_arquivo"))
		return
	}

	json_dat, err := json.Marshal(Hist)
	if err != nil {
		ansi.Println(Bold(Red("Falha ao codificar o histórico em json: "+err.Error())))
		return
	}
	err = ioutil.WriteFile(filename, json_dat, 0644)
	if err != nil {
		ansi.Println(Bold(Red("Falha ao salvar o histórico: "+err.Error())))
		return
	}
}

func abrir(cmd []string) {
	if len(cmd) == 0 {
		ansi.Println(Brown("Por favor, informa o nome do arquivo. Ex: abrir meu_arquivo.json"))
		return
	}

	json_dat, err := ioutil.ReadFile(cmd[0])
	if err != nil {
		ansi.Println(Bold(Red("Falha ao abrir o arquivo: "+err.Error())))
		return
	}
	err = json.Unmarshal(json_dat, Hist)
	if err != nil {
		ansi.Println(Bold(Red("Falha ao decodificar o histórico em json: "+err.Error())))
		return
	}
}

func add(cmd []string) {
	raw_input := ""
	ano := 0
	n := 0

	// Obtenha infromações
	fmt.Printf("Ano e semestre da disciplina (ex: 2017 1 ou 2016 2): ")
	fmt.Scanf("%d %d", &ano, %n)
	fmt.Printf("Código da disciplina (ex: 2017 1 ou 2016 2): ")
}

var Hist = NovoHistórico("", "")
var Filename = ""

func main() {
	for {
		ansi.Printf(Bold(Green("--> ")).String())
		raw_input, _ := simpleReadLine()
		raw_input = strings.TrimSpace(raw_input)
		cmd := strings.Split(raw_input, " ")
		switch cmd[0] {
		case "ajuda":
			help()
		case "hist":
			histórico(cmd[1:])
		case "salvar":
			salvar(cmd[1:])
		case "abrir":
			abrir(cmd[1:])
		case "add":
			add(cmd[1:])
		case "sair":
			os.Exit(0)
		default:
			ansi.Printf("Que tal o comando %s?\n", Green("ajuda"))
		}
	}
}
