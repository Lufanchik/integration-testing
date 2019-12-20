package integration_testing

import (
	"lab.siroccotechnology.ru/tp/integration-testing/apm"
	"lab.siroccotechnology.ru/tp/integration-testing/parking"
	"lab.siroccotechnology.ru/tp/integration-testing/passes"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mcd"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mck"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mgt"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mm"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mmts"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mmts_mtppk"
	"lab.siroccotechnology.ru/tp/integration-testing/webapi"
)

func init() {
	//параллельный запуск
	//общие кейсы проходов
	AddP(passes.CasesCancel)
	AddP(passes.CasesSimpleComplexPass)

	//МТППК
	//AddP(mtppk.CasesMTPPKPasses)

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

	//МЦК
	AddP(mck.CasesMetroComplexMCK1)
	AddP(mck.CasesMetroComplexMCK2)
	AddP(mck.CasesMetroComplexMCK3)
	AddP(mck.CasesMetroComplexMCK4)
	AddP(mck.CasesOfflineMetroComplexMCK)

	//МЦД
	AddP(mcd.CasesMetroComplexMCDMSK1)
	AddP(mcd.CasesMetroComplexMCDMSK2)
	AddP(mcd.CasesMetroComplexMCDMSK3)
	AddP(mcd.CasesMetroComplexMCDMSK4)
	AddP(mcd.CasesMetroComplexMCDMSK5)
	AddP(mcd.CasesComplexMCDMOPartOne)
	AddP(mcd.CasesComplexMCDMOPartTwo)
	AddP(mcd.CasesComplexMCDMOPartThree)
	AddP(mcd.CasesComplexMCDMOPartFour)
	AddP(mcd.CasesComplexMCDMOPartFife)
	AddP(mcd.CasesOfflineMetroComplexMCD)
	AddP(mcd.CasesOfflineMetroComplexMCDMO)


	//MMTS_MTPPK
	Add(mmts_mtppk.CasesComplexPassMMTS1MTPPK)

	//APM
	AddP(apm.CasesApmGateway)

	//МГТ
	AddP(mgt.CasesMGT2)

	//Parking
	AddP(parking.CasesParkingPass)

	//WebAPI
	AddP(webapi.CasesWEBAPI)

	//черновики, которые пока что не запускаются
	//AddP(passes.CasesWrongTimeComplexPass)
	//AddP(passes.CasesScopeCheckPass)

	//тесты с нестандартным временем - надо придумать как с ними работать
	//AddP(mm.CasesComplexTimeMM)
	//AddP(mcd.CasesComplexTimeMCD)
	//Add(mmts.CasesComplexTimeMMTS)
	//AddP(mcd.CasesComplexTimeMCD)
}
