var HistóricoEscolarPadrão = {
	"disciplinas": [
		{"código": "FCI0711", "nome": "Panorama da Biotecnologia Contemporânea", "nota": 10, "créditos": 2, "período": "2016-1"},
		{"código": "FCI0751", "nome": "Microbiologia", "nota": 6.5, "créditos": 5, "período": "2016-1"},
		{"código": "FCM0501", "nome": "Física I", "nota": 7.2, "créditos": 8, "período": "2016-1"},
		{"código": "SMA0301", "nome": "Cálculo I", "nota": 6.3, "créditos": 6, "período": "2016-1"},
		{"código": "SMA0330", "nome": "Complementos de Geometria e Vetores", "nota": 6.0, "créditos": 4, "período": "2016-1"},
		{"código": "SQM0406", "nome": "Fundamentos de Química Estrutural", "nota": 6.4, "créditos": 4, "período": "2016-1"},
		{"código": "FCI0730", "nome": "Biologia Molecular e Celular I", "nota": 8.5, "créditos": 4, "período": "2016-2"},
		{"código": "FCM0103", "nome": "Laboratório de Física I", "nota": 6.7, "créditos": 2, "período": "2016-2"},
		{"código": "FCM0502", "nome": "Física II", "nota": 6.7, "créditos": 8, "período": "2016-2"},
		{"código": "SCC0172", "nome": "Introdução à Programação para Biologia Molecular", "nota": 9.5, "créditos": 4, "período": "2016-2"},
		{"código": "SMA0332", "nome": "Cálculo II", "nota": 5.0, "créditos": 6, "período": "2016-2"},
		{"código": "SQF0373", "nome": "Química Geral para CFBio", "nota": 6.6, "créditos": 4, "período": "2016-2"},
		{
			"código": "7600007", "nome": "Física III", "nota": "?", "créditos": 4, "período": "2017-1",
			"variáveis": ["P1", "P2", "P3", "P4", "T1", "T2", "T3", "T4", "REC"],
			"notas": {"P1": 2.8, "P2": 3, "T1": 0, "T2": 9, "T3":"3-7"},
			"script": "nota_final = 0.8*(P1+P2+P3+P4)/4 + 0.2*(T1+T2+T3+T4)/4;\nif (nota_final < 5) {\n  nota_final = (nota_final+REC)/2;\n} else {\n  REC = undefined;\n}"
		},
		{
			"código": "7600014", "nome": "Laboratório de Física II", "nota": "?", "créditos": 3, "período": "2017-1",
			"variáveis": ["P1", "P2", "M_adi", "R1", "R2", "R3", "R4", "R5", "R6", "S1", "S2"],
			"script": "mr = (R1+R2+R3+R4+R5+R6)/6;\nms = (S1+S2)/2;\nmp = (P1+P2)/2;\nnota_final = 0.2*M_adi + 0.3*mr + 0.2*ms + 0.3*mp;\nif (mr < 5) {\n  nota_final = mr;\n}\nif (ms < 5) {\n  nota_final = ms;\n}\nif (mp < 5) {\n  nota_final = mp;\n}\nif (M_adi < 5) {\n  nota_final = M_adi;\n}\n"
		},
		{
			"código": "7600017", "nome": "Introdução à Física Computacional", "nota": "?", "créditos": 4, "período": "2017-1",
			"notasMáximas": {"P3": 12},
			"variáveis": ["P1", "P2", "P3", "P4", "P5", "P6"],
			"script": "nota_final = (P1+P2+P3+P4+P5+P6)/6;"
		},
		{
			"código": "SMA0356", "nome": "Cálculo IV", "nota": "?", "créditos": 4, "período": "2017-1",
			"variáveis": ["P1", "P2"],
			"script": "nota_final = P1;\nif (P2 > P1) {\n  nota_final = P2;\n}"
		},
		{
			"código": "SQM0485", "nome": "Princípios de Química Orgânica e Bioquímica de Macromoléculas ", "nota": "?", "créditos": 4, "período": "2017-1"
		},
		{
			"código": "SQM0486", "nome": "Laboratório de Bioquímica para CFBio", "nota": "?", "créditos": 2, "período": "2017-1"
		}
	]
}