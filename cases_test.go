package integration_testing

import (
	"lab.siroccotechnology.ru/tp/integration-testing/parking"
	"lab.siroccotechnology.ru/tp/integration-testing/passes"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/face"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mcd"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mck"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mgt"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mm"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mmts"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mt"
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mtppk"
	"lab.siroccotechnology.ru/tp/integration-testing/registry"
	"lab.siroccotechnology.ru/tp/integration-testing/webapi"
)

func init() {
	//параллельный запуск
	//общие кейсы проходов
	//CasesCancel 25 26 27 пока на доработке
	AddP(passes.CasesCancel)

	//МТППК
	AddP(mtppk.CasesMTPPKPasses)
	AddP(mtppk.CasesMTPPK_MGT)
	AddP(mtppk.CasesMTPPK_MCD_MO)
	AddP(mtppk.CasesMTPPK_MCK)
	AddP(mtppk.CasesMTPPK_MCD_MSK)
	AddP(mtppk.CasesMTPPK_single)

	//ММТС
	AddP(mmts.CasesComplexPassMMTS1)
	AddP(mmts.CasesComplexPassMMTS2)
	AddP(mmts.CasesComplexPassMMTS3)
	AddP(mmts.CasesComplexPassMMTS4)
	AddP(mmts.CasesComplexPassMMTS5)

	//MM
	AddP(mm.CasesMetroComplexMM1)
	AddP(mm.CasesMetroComplexMM2)
	AddP(mm.CasesMetroComplexMM3)
	AddP(mm.CasesMetroComplexMM4)

	//МЦК
	AddP(mck.CasesMetroComplexMCK1)
	AddP(mck.CasesMetroComplexMCK2)
	AddP(mck.CasesMetroComplexMCK3)
	AddP(mck.CasesMetroComplexMCK4)

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

	//Мобильная тройка
	AddP(mt.CasesMetroComplexMT)

	//APM
	//AddP(apm.CasesApmGateway)

	//МГТ
	AddP(mgt.CasesMGT)
	AddP(mgt.CasesMGT_1)
	AddP(mgt.CasesMGT_2)
	AddP(mgt.CasesMGT_3)

	//Parking
	AddP(parking.CasesParkingPass)

	//WebAPI
	AddP(webapi.CasesWEBAPI)

	//Revise
	AddAR(registry.CasesReviseGetTaskList)
	//Resolve
	AddAR(registry.CasesResolveGetTaskList)
	//Face
	AddP(face.CasesAuthWithFace)

	//черновики, которые пока что не запускаются
	//AddP(passes.CasesWrongTimeComplexPass)
	//AddP(passes.CasesScopeCheckPass)

	//тесты с нестандартным временем - надо придумать как с ними работать
	//AddP(mm.CasesComplexTimeMM)
	//AddP(mcd.CasesComplexTimeMCD)
	//Add(mmts.CasesComplexTimeMMTS)
	//AddP(mcd.CasesComplexTimeMCD)
}
