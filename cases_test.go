package integration_testing

import (
	"lab.siroccotechnology.ru/tp/integration-testing/passes/mck"
)

func init() {
	//последовательный запуск
	//Add(passes.CasesCancel)
	Add(mck.CasesComplexTimeMCK)

	////параллельный запуск
	////общие кейсы проходов
	////AddP(passes.CasesCancel)
	//AddP(passes.CasesSimpleComplexPass)
	//AddP(passes.CasesSimplePass)
	//
	////МТППК
	//AddP(mtppk.CasesMTPPKPasses)
	//
	////ММТС
	//AddP(mmts.CasesComplexPassMMTS1)
	//AddP(mmts.CasesComplexPassMMTS2)
	//AddP(mmts.CasesComplexPassMMTS3)
	//AddP(mmts.CasesComplexPassMMTS4)
	//AddP(mmts.CasesComplexPassMMTS5)
	////AddP(mmts.CasesOfflineMetroComplexMMTS)
	//
	////MM
	//AddP(mm.CasesMetroComplexMM1)
	//AddP(mm.CasesMetroComplexMM2)
	//AddP(mm.CasesMetroComplexMM3)
	//AddP(mm.CasesMetroComplexMM4)
	////AddP(mm.CasesOfflineMetroComplexMM)
	//
	////МЦК
	////AddP(mck.CasesMetroComplexMCK)
	//AddP(mck.CasesMetroComplexMCK1)
	//AddP(mck.CasesMetroComplexMCK2)
	//AddP(mck.CasesMetroComplexMCK3)
	//AddP(mck.CasesMetroComplexMCK4)
	////AddP(mck.CasesOfflineMetroComplexMCK)
	//
	////МЦД
	//AddP(mcd.CasesComplexMCD)
	//AddP(mcd.CasesComplexMCD2)
	//AddP(mcd.CasesMetroComplexMCDMSK1)
	//AddP(mcd.CasesMetroComplexMCDMSK2)
	//AddP(mcd.CasesMetroComplexMCDMSK3)
	//AddP(mcd.CasesMetroComplexMCDMSK4)
	//AddP(mcd.CasesMetroComplexMCDMSK5)
	//AddP(mcd.CasesComplexTimeMCD)
	//
	////APM
	//AddP(apm.CasesApmGateway)
	//
	////Parking
	////AddP(parking.CasesParkingPass)
	//
	////черновики, которые пока что не запускаются
	////AddP(passes.CasesWrongTimeComplexPass)
	////AddP(passes.CasesScopeCheckPass)
}
