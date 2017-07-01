HistóricoEscolar = {"disciplinas": []};

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

	for (var i=0; i < HistóricoEscolar.disciplinas.length; i++) {
		var discp = HistóricoEscolar.disciplinas[i];
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
	for (var i=0; i < HistóricoEscolar.disciplinas.length; i++) {
		var discp = HistóricoEscolar.disciplinas[i];
		if (discp.sugestões === undefined) {
			continue;
		}
		var keys = Object.keys(discp.sugestões);
		for (var j=0; j < keys.length; j++) {
			if (discp.sugestões[keys[j]] === undefined) {
				continue;
			}
			tmp += "<tr>";
			tmp += "<td>"+discp.período+"</td>";
			tmp += "<td>"+discp.código+"</td>";
			tmp += "<td>"+discp.nome+"</td>";
			tmp += "<td>"+keys[j]+"</td>";
			tmp += "<td>"+discp.sugestões[keys[j]]+"</td>";
			tmp += "</tr>";
		}
	}
	$("#sugestões_tbody").html(tmp);
}

function saveToLocalStorage() {
	"use strict";
	localStorage.setItem("histórico", JSON.stringify(HistóricoEscolar));
}

function loadFromLocalStorage() {
	"use strict";
	var retrievedObject = localStorage.getItem("histórico");
	if (retrievedObject == undefined) {
		HistóricoEscolar = HistóricoEscolarPadrão;
	} else {
		HistóricoEscolar = JSON.parse(retrievedObject);
	}

	try {
		// Limpe as sugestões
		for (var i=0; i < HistóricoEscolar.disciplinas.length; i++) {
			delete HistóricoEscolar.disciplinas[i].sugestão;
		}
		// Ordene
		function compare(a,b) {
		  if (a.período < b.período)
		    return -1;
		  if (a.período > b.período)
		    return 1;
		  if (a.código < b.código)
		    return -1;
		  if (a.código > b.código)
		    return 1;
		  return 0;
		}
		HistóricoEscolar.disciplinas.sort(compare);
		
		update_histórico();
		calcular_range();
		$("#sugestões_tbody").html("");
	} catch (e) {
		HistóricoEscolar = HistóricoEscolarPadrão;
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
	saveToLocalStorage();
}

function calcular_tudo() {
	"use strict";

	var meta = $("#inMeta").val();
	for (var i = 0; i <= 300; i++) {
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
	var discp = HistóricoEscolar.disciplinas[index];
	var vm = new Interpreter(discp.script);
	vm.setProperty(vm.global, "nota_final", vm.createPrimitive(0));	
	discp.sugestões = {};
	for (var i = 0; i < discp.variáveis.length; i++) {
		var v = discp.variáveis[i];
		var val = Math.min(free_grade, 10);
		var flag = true;
		if (discp.notas !== undefined && v in discp.notas) {
			flag = !IsNumeric(discp.notas[v]);
			val = fitValue(discp.notas[v], free_grade);
		}
		if (discp.notasMáximas !== undefined && v in discp.notasMáximas) {
			val = Math.min(free_grade, discp.notasMáximas[v]);
		}
		if (flag) {
			HistóricoEscolar.disciplinas[index].sugestões[v] = val;
			console.log(discp.nome, v, val);
		}
		vm.setProperty(vm.global, v, vm.createPrimitive(val));
	}
	vm.run();
	for (var i = 0; i < discp.variáveis.length; i++) {
		var v = discp.variáveis[i];
		var value = vm.getProperty(vm.global, v);
		if (value.type == "undefined") {
			HistóricoEscolar.disciplinas[index].sugestões[v] = undefined;
		}
	}
	return Math.round(100*vm.getProperty(vm.global, "nota_final").data)/100;
}

function simula_passo(free_grade) {
	"use strict";
	var ponderada = 0;
	var créditos = 0;
	for (var i=0; i < HistóricoEscolar.disciplinas.length; i++) {
		var discp = HistóricoEscolar.disciplinas[i];
		var val_to_use = fitValue(discp.nota, free_grade);
		if (discp.script != "" && discp.script !== undefined && discp.nota == "?") {
			try {
				val_to_use = run_script(i, free_grade);
			} catch (e) {
				console.log(e);
				val_to_use = 0;
			}
		}
		val_to_use = Math.min(val_to_use, 10);
		ponderada += val_to_use*discp.créditos;
		HistóricoEscolar.disciplinas[i].sugestão = val_to_use;
		créditos += discp.créditos;
	}

	ponderada = ponderada/créditos;
	return Math.round(ponderada*100)/100;
}

function fitValue(nota, tentativa) {
	"use strict";
	if (IsNumeric(nota)) {
		return Number(nota);
	}
	tentativa = Math.min(tentativa, 10);
	var tmp = nota.split("-");
	if (tmp.length == 2 && IsNumeric(tmp[0]) && IsNumeric(tmp[1])) {
		return Math.min(Math.max(tentativa, tmp[0]), tmp[1]);
	}
	return Number(tentativa);
}

function download_json() {
	"use strict";
	saveToLocalStorage();
	var element = document.createElement('a');
	element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(JSON.stringify(HistóricoEscolar)));
	element.setAttribute('download', "histórico_escolar.json.txt");
	element.style.display = 'none';
	document.body.appendChild(element);
	element.click();
	document.body.removeChild(element);
}

function enter_json_file() {
	$("#inFile").click();
}

function enter_json_file_2() {
	if (typeof window.FileReader !== 'function') {
        alert("O seu navegador não tem a file API, a qual é necessária para abrir arquivos.");
        return;
    }

	function receivedText() {
        var txt = fr.result;
        try {
			HistóricoEscolar = JSON.parse(txt);
			saveToLocalStorage();
			loadFromLocalStorage();
			alert("Arquivo carregado!");
		} catch (e) {
			alert("Falha ao abrir arquivo: "+e);
		}
        
    }

	try {
		file = $("#inFile")[0].files[0];
        fr = new FileReader();
        fr.onload = receivedText;
        fr.readAsText(file);
	} catch (e) {
		alert("Falha ao abrir arquivo: "+e);
	}
}

function nova_disciplina() {
	HistóricoEscolar.disciplinas.push({});
	editar_disciplina(HistóricoEscolar.disciplinas.length-1);
}

function editar_disciplina(index) {
	"use strict";
	location.hash = "nova_disciplina_"+index;
	var discp = HistóricoEscolar.disciplinas[index];
	var tmp = "";
	var tmp2 = "";
	var notas = discp.notas || {};
	for (var key in notas) {
		tmp += " "+key+"="+notas[key];
	}
	var notas2 = discp.notasMáximas || {};
	for (var key in notas2) {
		tmp2 += " "+key+"="+notas2[key];
	}

	$("#inDiscpCódigo").val(discp.código);
	$("#inDiscpNome").val(discp.nome);
	$("#inDiscpPeríodo").val(discp.período);
	$("#inDiscpMédia").val(discp.nota);
	$("#inDiscpCréditos").val(discp.créditos);
	var vars = discp.variáveis || [];
	$("#inDiscpVariáveis").val(vars.join(" "));
	$("#inDiscpNotas").val(tmp);
	$("#inDiscpNotasMáximas").val(tmp2);
	var src = discp.script || "";
	$("#inDiscpScript").val(src);

	$("#salvar_disciplina").unbind("click")
	$("#salvar_disciplina").click(function () {
		salvar_disciplina(index);
	});
	$("#remover_disciplina").unbind("click")
	$("#remover_disciplina").click(function () {
		remover_disciplina(index);
	});
}

function remover_disciplina(index) {
	HistóricoEscolar.disciplinas.splice(index, 1);
	location.hash = "";
	saveToLocalStorage();
	loadFromLocalStorage();
}

function salvar_disciplina(index) {
	HistóricoEscolar.disciplinas[index].notas = HistóricoEscolar.disciplinas[index].notas || {};
	HistóricoEscolar.disciplinas[index].notasMáximas = HistóricoEscolar.disciplinas[index].notasMáximas || {};

	HistóricoEscolar.disciplinas[index].código = $("#inDiscpCódigo").val();
	HistóricoEscolar.disciplinas[index].nome = $("#inDiscpNome").val();
	HistóricoEscolar.disciplinas[index].créditos = Number($("#inDiscpCréditos").val());
	HistóricoEscolar.disciplinas[index].período = $("#inDiscpPeríodo").val();
	HistóricoEscolar.disciplinas[index].nota = $("#inDiscpMédia").val();
	HistóricoEscolar.disciplinas[index].script = $("#inDiscpScript").val();
	HistóricoEscolar.disciplinas[index].variáveis = $("#inDiscpVariáveis").val().split(" ");
	HistóricoEscolar.disciplinas[index].notas = {};
	var tmp = $("#inDiscpNotas").val().split(" ");
	for (var i=0; i < tmp.length; i++) {
		try {
			var tmp2 = tmp[i].split("=");
			if (tmp2.length == 2) {
				HistóricoEscolar.disciplinas[index].notas[tmp2[0]] = tmp2[1];
			}
		} catch (e) {
			console.log(e);
		}
	}
	console.log(HistóricoEscolar.disciplinas[index].notas);
	tmp = $("#inDiscpNotasMáximas").val().split(" ");
	for (var i=0; i < tmp.length; i++) {
		try {
			var tmp2 = tmp[i].split("=");
			if (tmp2.length == 2) {
				HistóricoEscolar.disciplinas[index].notasMáximas[tmp2[0]] = tmp2[1];
			}
		} catch (e) {
			console.log(e);
		}
	}
	console.log(HistóricoEscolar.disciplinas[index].notasMáximas);
	location.hash = "";
	saveToLocalStorage();
	loadFromLocalStorage();
}

hide_all();
$("#histórico").show();
$("#sugestões_detalhadas").show();
$(function() {
	loadFromLocalStorage();
	location.hash = "";
	window.onhashchange = function() {
		if (location.hash.startsWith("#nova_disciplina")) {
			hide_all();
			$("#nova_disciplina").show();
			return;
		}
		hide_all();
		$("#histórico").show();
		$("#sugestões_detalhadas").show();
	}
});