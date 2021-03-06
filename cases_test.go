package integration_testing

import (
	mk_emulator "lab.dt.multicarta.ru/tp/integration-testing/mk-emulator"
	"lab.dt.multicarta.ru/tp/integration-testing/passes/mcd"
	"lab.dt.multicarta.ru/tp/integration-testing/passes/mck"
	"lab.dt.multicarta.ru/tp/integration-testing/passes/mm"
	"lab.dt.multicarta.ru/tp/integration-testing/passes/mmts"
)

func init() {
	//параллельный запуск
	//общие кейсы проходов
	//CasesCancel 25 26 27 пока на доработке
	//AddP(passes.CasesCancel)
	//AddP(aggregate.MetroAggregate)

	//МТППК
	//AddP(mtppk.CasesMTPPKPasses)
	//AddP(mtppk.CasesMTPPK_MGT)
	//AddP(mtppk.CasesMTPPK_MCD_MO)
	//AddP(mtppk.CasesMTPPK_MCK)
	//AddP(mtppk.CasesMTPPK_MCD_MSK)
	//AddP(mtppk.CasesMTPPK_single)

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
	//AddP(mt.CasesMetroComplexMT)

	//APM
	//AddP(apm.CasesApmGateway)

	//МГТ
	//AddP(mgt.CasesMGT_1)
	//AddP(mgt.CasesMGT_2)
	//AddP(mgt.CasesMGT_3)
	//AddP(mgt.CasesMGTWithEmptyEMV)

	//Parking
	//AddP(parking.CasesParkingPass)

	//WebAPI
	//AddP(webapi.CasesWEBAPI)

	//Revise
	//AddAR(registry.CasesReviseGetTaskList)
	//Resolve
	//AddAR(registry.CasesResolveGetTaskList)
	//Face
	//AddP(face.MetroComplexFaceID)

	//Comments CRUDs
	//AddAR(comments.CasesCommentsCRUD)

	//Отмена оплаты в TWPG
	//AddP(twpg.CaseTWPGReverseOrder)

	////скачивание полного вайт листа на ридер
	//AddP(cards.FaceList)
	////запрос от ридера, включая скачивание полного вайт листа
	//AddP(webapi.ReaderCase)
	////запрос на восстановление прохода, если недостаточно данных для авторизации
	//AddP(resolve.CasesResolve)

	//Токен Мастеркард
	//AddP(mastercard.CasesTokensMC)

	//MK-emulator
	AddP(mk_emulator.CasesEmulatorMK)

	//Stop list add and remove card
	//AddP(cards.CardsStopList)

	//черновики, которые пока что не запускаются
	//AddP(passes.CasesWrongTimeComplexPass)
	//AddP(passes.CasesScopeCheckPass)

	//тесты с нестандартным временем - надо придумать как с ними работать
	//AddP(mm.CasesComplexTimeMM)
	//AddP(mcd.CasesComplexTimeMCD)
	//Add(mmts.CasesComplexTimeMMTS)
	//AddP(mcd.CasesComplexTimeMCD)
}
