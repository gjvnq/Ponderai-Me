package main

import (
	"errors"
	"github.com/robertkrimen/otto"
	"sort"
	"time"
)

var HaltErr = errors.New("script took too long")

func (discp *DisciplinaT) RunScript(free_grade float64, use_min_grades bool) (ret_err error) {
	vm := otto.New()
	discp.SugestãoFinal = 0
	defer func() {
		if caught := recover(); caught != nil {
			if caught == HaltErr {
				ret_err = errors.New("hi")
			} else {
				panic(caught) // Something else happened, repanic!
			}
		}
	}()

	// Prepare a vm
	vm.Interrupt = make(chan func(), 1)
	discp.Sugestões = make(map[string]float64)
	// Coloque as notas
	for _, varname := range discp.Vars {
		nota, ok := discp.NotasAtéAgora[varname]
		nota_val := 0.0
		if ok {
			nota_val = nota.Val
			if nota.IsRange && use_min_grades {
				if nota_val < nota.Min {
					nota_val = nota.Min
				}
			}
		} else {
			nota_val = free_grade
		}

		nota_max, ok := discp.NotasMáximas[varname]
		if ok && nota_val > nota_max {
			nota_val = nota_max
		}
		vm.Set(varname, nota_val)
		discp.Sugestões[varname] = nota_val
	}
	vm.Set("nota_final", 0.0)

	// Instale algumas funções
	vm.Set("max", func(call otto.FunctionCall) otto.Value {
		l := make([]float64, 0)
		for _, arg := range call.ArgumentList {
			num, err := arg.ToFloat()
			if err == nil {
				l = append(l, num)
			}
		}
		sort.Float64s(l)
		result, _ := call.Otto.ToValue(l[len(l)-1])
		return result
	})

	go func() {
		time.Sleep(time.Second) // Stop after one seconds
		vm.Interrupt <- func() {
			panic(HaltErr)
		}
	}()
	_, err := vm.Run(discp.JSScript)
	if err != nil {
		ret_err = errors.New("falha ao rodar script: " + err.Error())
		return
	}
	nota_final, err := vm.Get("nota_final")
	if err != nil {
		ret_err = errors.New("falha ao extarir nota do script")
		return
	}
	discp.SugestãoFinal, err = nota_final.ToFloat()
	if err != nil {
		ret_err = errors.New("falha ao extarir nota do script")
		return
	}
	return
}
