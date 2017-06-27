package main

import (
	"github.com/robertkrimen/otto"
	"time"
	"errors"
)

var halt = errors.New("Stahp")

func (discp *DisciplinaT) RunScript(free_grade float64, use_min_grades bool) (ret_err error) {
	vm := otto.New()
	discp.SugestãoFinal = 0
	defer func() {
        if caught := recover(); caught != nil {
            if caught == halt {
                ret_err = errors.New("hi")
            } else {
            	panic(caught) // Something else happened, repanic!
            }
        }
    }()

    // Prepare a vm
	vm.Interrupt = make(chan func(), 1)
	discp.Sugestões = make(map[string]float64)
	for _, varname := range discp.Vars {
		nota, ok := discp.NotasAtéAgora[varname]
		nota_val := nota.Val
		if !ok {
			nota_val = free_grade
		} else if nota.IsRange && use_min_grades {
			if nota_val < nota.Min {
				nota_val = nota.Min
			}
		}
		vm.Set(varname, nota_val)
		discp.Sugestões[varname] = nota_val
	}

	go func() {
        time.Sleep(time.Second) // Stop after one seconds
        vm.Interrupt <- func() {
            panic(halt)
        }
    }()
	_, err := vm.Run(discp.JSScript)
	if err != nil {
		ret_err = errors.New("falha ao rodar script")
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