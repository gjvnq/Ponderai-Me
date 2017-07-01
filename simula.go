package main

import "errors"

var NoWayErr = errors.New("no way to get that grade")

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func MaiorNotaPossível(hist *HistóricoT) float64 {
	maior_nota := 10.0
	for i := 0; i < len(hist.Disciplinas); i++ {
		discp := hist.Disciplinas[i]
		if discp.NotasMáximas != nil {
			for _, n_max := range discp.NotasMáximas {
				maior_nota = max(maior_nota, n_max)
			}
		}
	}
	return maior_nota
}

func MelhorCaso(hist *HistóricoT) error {
	return SimulaPasso(hist, 999, false)
}

func PiorCaso(hist *HistóricoT) error {
	return SimulaPasso(hist, 0, true)
}

func Meta(meta float64, hist *HistóricoT, use_min_grades bool) error {
	dx := 0.1
	max_nota := MaiorNotaPossível(hist)
	for x := 0.0; x < max_nota; x += dx {
		SimulaPasso(hist, x, use_min_grades)
		if hist.Média >= meta {
			// Encontramos as notas necessárias
			return nil
		}
	}
	return NoWayErr
}

func SimulaPasso(hist *HistóricoT, free_grade float64, use_min_grades bool) error {
	accumulator_nom := 0.0
	accumulator_den := 0
	hist.Média = 0.0
	for i := 0; i < len(hist.Disciplinas); i++ {
		discp := hist.Disciplinas[i]
		// A disciplina já está "fechada"?
		if discp.NotaFinal == nil {
			err := discp.RunScript(free_grade, use_min_grades)
			if err != nil {
				// do something?
				discp.SugestãoFinal = 0.0
			}
		} else {
			discp.SugestãoFinal = discp.NotaFinal.Try(min(10, free_grade))
		}
		discp.SugestãoFinal = min(discp.SugestãoFinal, 10)
		accumulator_nom += float64(discp.Créditos) * discp.SugestãoFinal
		accumulator_den += discp.Créditos
	}
	hist.Média = accumulator_nom / float64(accumulator_den)
	return nil
}
