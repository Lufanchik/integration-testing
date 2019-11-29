package integration_testing

import (
	"lab.siroccotechnology.ru/tp/integration-testing/passes"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mm"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mmts"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mtppk"
)

func init() {
	//последовательный запуск
	//Add(passes.CasesCancel)

	//параллельный запуск - общие кейсы
	AddP(passes.CasesCancel)
	//AddP(passes.CasesWrongTimeComplexPass)
	//AddP(passes.CasesScopeCheckPass)
	AddP(passes.CasesSimpleComplexPass)
	AddP(passes.CasesSimplePass)

	//МТППК
	AddP(mtppk.CasesMTPPKPasses)

	//ММТС
	AddP(mmts.CasesComplexPassMMTS1)
	AddP(mmts.CasesComplexPassMMTS2)
	AddP(mmts.CasesComplexPassMMTS3)
	AddP(mmts.CasesComplexPassMMTS4)
	AddP(mmts.CasesComplexPassMMTS5)
	AddP(mmts.CasesOfflineMetroComplexMMTS)

	//MM
	AddP(mm.CasesMetroComplexMM1)
	AddP(mm.CasesMetroComplexMM2)
	AddP(mm.CasesMetroComplexMM3)
	AddP(mm.CasesMetroComplexMM4)
	AddP(mm.CasesOfflineMetroComplexMM)
}
