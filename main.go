package main

// import ui "github.com/SAPessi/termui"

// var H Historico
// var UIState int = UIState_MainView
// var cmd_list_p ui.Par

func main() {
	// err := ui.Init()
	// ui.Body.BgColor = ui.ColorBlack
	// if err != nil {
	// 	panic(err)
	// }
	// defer ui.Close()

	// // Lista de comandos de teclado
	// cmd_list_p = ui.NewPar("[ESQ] Sair/Voltar [Ctrl+O] Abrir arquivo [Ctrl+S] Salvar arquivo [Ctrl+Shift+S] Salvar arquivo como [N] Adicionar disciplina [R] Remover disciplina [M] Definir meta")
	// cmd_list_p.Float = ui.AlignCenter
	// cmd_list_p.Height = 3
	// cmd_list_p.TextFgColor = ui.ColorWhite
	// cmd_list_p.BorderFg = ui.ColorCyan

	// drawUiState()
	// ui.Handle("/sys/kbd/C-c", func(ui.Event) {
	// 	ui.StopLoop()
	// })
	// ui.Handle("/sys/kbd/C-o", func(ui.Event) {
	// })
	// ui.Handle("/sys/wnd/resize", func(e ui.Event) {
	// 	ui.Body.Width = ui.TermWidth()
	// 	ui.Body.Align()
	// 	ui.Clear()
	// 	ui.Render(ui.Body)
	// })
	// ui.Loop()
}

func drawMain() {
	// // Lista de disciplinas
	// discp_tbl := ui.NewTable()
	// discp_tbl.Rows = H.TableRows()
	// discp_tbl.FgColor = ui.ColorWhite
	// discp_tbl.BgColor = ui.ColorDefault
	// discp_tbl.Y = 0
	// discp_tbl.X = 0
	// discp_tbl.Width = 62
	// discp_tbl.Height = 7

	// ui.Body.AddRows(
	// 	ui.NewRow(
	//     	ui.NewCol(10, 0, discp_tbl)),
	//     ui.NewRow(
	//         ui.NewCol(12, 0, cmd_list_p)))

	//    ui.Body.Align()
	// ui.Render(ui.Body)
}

func drawUiState() {
	// switch UIState {
	// 	case UIState_MainView:
	// 	drawMain()
	// }
}
