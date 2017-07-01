window.HistóricoEscolar = {"disciplinas": []};

function IsNumeric(input) {
    return (input - 0) == input && (''+input).trim().length > 0;
}

function hide_all() {
	"use strict";
	$("#histórico").hide();
	$("#nova_disciplina").hide();
	$("#sugestões_detalhadas").hide();
}

function nova_disciplina_show() {
	"use strict";
	hide_all();
	$("#nova_disciplina").show();
}

function update_histórico() {
	"use strict";
	$("#histórico_tbody").html("");
	var tmp = "";
	var ponderada = 0;
	var créditos = 0;

	for (var i=0; i < window.HistóricoEscolar.disciplinas.length; i++) {
		var discp = window.HistóricoEscolar.disciplinas[i];
		tmp += "<tr>";
		tmp += "<td>"+discp.período+"</td>";
		tmp += "<td>"+discp.código+"</td>";
		tmp += "<td>"+discp.nome+"</td>";
		tmp += "<td>"+discp.créditos+"</td>";
		tmp += "<td>"+discp.nota+"</td>";
		if (discp.sugestão === undefined) {
			discp.sugestão = "";
		}
		tmp += "<td>"+discp.sugestão+"</td>";
		tmp += "<td><button type=\"button\" class=\"btn btn-link\" onclick=\"editar_disciplina("+i+")\">Editar</button></td>";
		tmp += "</tr>";
		if (IsNumeric(discp.nota)) {
			ponderada += discp.nota*discp.créditos;
			créditos += discp.créditos;
		}
	}

	ponderada = ponderada/créditos;
	ponderada = Math.round(ponderada*100)/100;
	$("#histórico_tbody").html(tmp);
	$("#ponderadaAtual").html(ponderada);
}

function update_sugestões() {
	"use strict";
	$("#sugestões_tbody").html("");
	var tmp = "";
	for (var i=0; i < window.HistóricoEscolar.disciplinas.length; i++) {
		var discp = window.HistóricoEscolar.disciplinas[i];
		if (discp.sugestões === undefined) {
			continue;
		}
		console.log(i, discp.nome, discp.sugestões);
		var keys = Object.keys(discp.sugestões);
		for (var j=0; j < keys.length; j++) {
			if (discp.sugestões[keys[j]] === undefined) {
				console.log("discp.sugestões[keys[j]]", discp.sugestões[keys[j]])
				continue;
			}
			tmp += "<tr>";
			tmp += "<td>"+discp.período+"</td>";
			tmp += "<td>"+discp.código+"</td>";
			tmp += "<td>"+discp.nome+"</td>";
			tmp += "<td>"+keys[j]+"</td>";
			tmp += "<td>"+discp.sugestões[keys[j]]+"</td>";
			tmp += "</tr>";
			console.log(discp.nome, keys[j], discp.sugestões[keys[j]]);
		}
	}
	$("#sugestões_tbody").html(tmp);
}

function saveToLocalStorage() {
	"use strict";
	localStorage.setItem("histórico", JSON.stringify(window.HistóricoEscolar));
}

function loadFromLocalStorage() {
	"use strict";
	var retrievedObject = localStorage.getItem("histórico");
	if (retrievedObject == undefined) {
		window.HistóricoEscolar = HistóricoEscolarPadrão;
	} else {
		window.HistóricoEscolar = JSON.parse(retrievedObject);
	}

	try {
		// Limpe as sugestões
		for (var i=0; i < window.HistóricoEscolar.disciplinas.length; i++) {
			delete window.HistóricoEscolar.disciplinas[i].sugestão;
		}
		
		update_histórico();
		calcular_range();
		$("#sugestões_tbody").html("");
	} catch (e) {
		window.HistóricoEscolar = HistóricoEscolarPadrão;
		update_histórico();
		calcular_range();
		$("#sugestões_tbody").html("");
	}

}

function calcular_range() {
	"use strict";
	var pior = simula_passo(0);
	var melhor = simula_passo(999);
	$("#ponderadaRangePossível").html(pior+"-"+melhor);
}

function calcular_tudo() {
	"use strict";

	var meta = $("#inMeta").val();
	for (var i = 0; i < 100; i++) {
		var res = simula_passo(i/10);
		if (res >= meta) {
			update_sugestões();
			update_histórico();
			return;
		}
	}
	saveToLocalStorage();
	alert("Não foi possível chegar em uma ponderada de: "+meta);
	update_sugestões();
	update_histórico();
}

function run_script(index, free_grade) {
	"use strict";
	var discp = window.HistóricoEscolar.disciplinas[index];
	var vm = new Interpreter(discp.script);
	vm.setProperty(vm.global, "nota_final", vm.createPrimitive(0));	
	for (var i = 0; i < discp.variáveis.length; i++) {
		var v = discp.variáveis[i];
		var val = Math.min(free_grade, 10);
		var flag = true;
		if (discp.notas !== undefined && v in discp.notas) {
			val = fitValue(discp.notas[v], free_grade);
			flag = false;
		}
		if (discp.notas_máximas !== undefined && v in discp.notas_máximas) {
			val = Math.min(free_grade, discp.notas_máximas[v]);
		}
		if (discp.sugestões === undefined) {
			discp.sugestões = {};
		}
		if (flag) {
			window.HistóricoEscolar.disciplinas[index].sugestões[v] = val;
		}
		vm.setProperty(vm.global, v, vm.createPrimitive(val));
	}
	vm.run();
	for (var i = 0; i < discp.variáveis.length; i++) {
		var v = discp.variáveis[i];
		var value = vm.getProperty(vm.global, v);
		console.log(v, value);
		if (value.type == "undefined") {
			console.log("del", v, window.HistóricoEscolar.disciplinas[index].sugestões[v]);
			window.HistóricoEscolar.disciplinas[index].sugestões[v] = undefined;
			console.log("=", v, window.HistóricoEscolar.disciplinas[index].sugestões[v]);
		}
	}
	return Math.round(100*vm.getProperty(vm.global, "nota_final").data)/100;
}

function simula_passo(free_grade) {
	"use strict";
	var ponderada = 0;
	var créditos = 0;
	for (var i=0; i < window.HistóricoEscolar.disciplinas.length; i++) {
		var discp = window.HistóricoEscolar.disciplinas[i];
		var val_to_use = fitValue(discp.nota, free_grade);
		if (discp.script != "" && discp.script !== undefined && discp.nota == "?") {
			try {
				val_to_use = run_script(i, free_grade);
			} catch (e) {
				console.log(e);
				val_to_use = 0;
			}
		}
		ponderada += val_to_use*discp.créditos;
		window.HistóricoEscolar.disciplinas[i].sugestão = val_to_use;
		créditos += discp.créditos;
	}

	ponderada = ponderada/créditos;
	return Math.round(ponderada*100)/100;
}

function fitValue(nota, tentativa) {
	"use strict";
	if (IsNumeric(nota)) {
		return nota;
	}
	tentativa = Math.min(tentativa, 10);
	var tmp = nota.split("-");
	if (tmp.length == 2 && IsNumeric(tmp[0]) && IsNumeric(tmp[1])) {
		return Math.min(Math.max(tentativa, tmp[0]), tmp[1]);
	}
	return tentativa;
}

function download_json() {
	"use strict";
	saveToLocalStorage();
	var element = document.createElement('a');
	element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(JSON.stringify(window.HistóricoEscolar)));
	element.setAttribute('download', "histórico_escolar.json.txt");
	element.style.display = 'none';
	document.body.appendChild(element);
	element.click();
	document.body.removeChild(element);
}

hide_all();
$("#histórico").show();
$("#sugestões_detalhadas").show();
$(function() {
	loadFromLocalStorage();
});