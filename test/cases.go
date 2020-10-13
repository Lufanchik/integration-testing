package test

import (
	"lab.dt.multicarta.ru/tp/common/messages/cards"
	"lab.dt.multicarta.ru/tp/common/messages/carriers"
	"lab.dt.multicarta.ru/tp/common/messages/pass"
	"lab.dt.multicarta.ru/tp/common/messages/processing"
	"lab.dt.multicarta.ru/tp/common/messages/twpg"
	"time"
)

type (
	PaymentType uint8
	RequestType uint8
	AuthType    uint8
	Cases       []Case
	T           []interface{}
)

//варианты шагов, который мы можем использовать в кейсах
type (
	Case struct {
		N string
		T T
		//платежная система
		CardSystem processing.CardSystem
		//тип прохода
		PassType             pass.PassType
		CustomerId           string
		TWPGOrderId          uint64
		SkipIdempotencyCheck bool
		//Проходы по одной карте
		Card           *processing.Card
		NotDoubleCheck bool
	}

	ProcessRevisePass struct {
		req  *processing.OnlinePassEvent
		pass *pass.Pass
	}

	//генерация прохода
	Pass struct {
		//тип оплаты
		PaymentType PaymentType
		//тип карты
		PassType pass.PassType
		//тип прохода
		RequestType RequestType
		//тип авторизации
		AuthType AuthType
		//ид саб перевозчика
		SubCarrier carriers.SubCarrier
		//ид перевозчика
		Carrier carriers.Carrier
		//данные об авторизации
		Auth *processing.Auth
		//сумма, на которую мы ожидаем авторизацию
		ExpectedSum uint32
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		Parent int
		//ссылка на вход, для валидации связи
		Ingress int
		//ссылка на открывающую агрегацию
		Aggregate int
		//функция времени
		Now func() uint64
		//время отведенное на запрос в миллисекундах, работает только с онлайном
		Duration uint32
		//данные о терминале
		Terminal *processing.Terminal
		//время, которое мы ждяли перед входом
		TimeToWait time.Duration
		//является ли проход комплексным
		IsComplex bool
		//вышли ли мы за пределы таймаута комплексной поездки
		IsComplexTimeout bool
		faceId           string
		// сбрасываем emv
		EmptyEMV bool
		// проход должен быть таким же, как указанный
		Equal int
		// до авторизуем восстановленный проход
		RevisePass      int
		tapResponse     PassResponser
		passResponse    PassResponser
		id              string
		carrierID       string
		tapRequest      *processing.TapRequest
		timeRequest     uint64
		card            *processing.Card
		parent          *Pass
		equal           *Pass
		ingress         *Pass
		aggregate       *Pass
		isParent        bool
		timeToWait      time.Duration
		isCancel        bool
		isComplete      bool
		completeSum     uint32
		IsInitAggregate bool
		//сумма, на которую мы ожидаем агрегацию
		ExpectedSumAggregate uint32
	}

	//генерация прохода
	PassCheck struct {
		//тип оплаты
		PaymentType PaymentType
		//сумма, на которую мы ожидаем авторизацию
		ExpectedSum uint32
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		Parent int
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		Target int
	}

	//модификация прохода
	Updater struct {
		//функция модификатора
		f func(tap *pass.Pass)
		//ссылка на проход в нашем кейсе, которую мы считаем стартовой
		target int
	}

	//получение реестров абс
	AbsGetRegistry struct {
	}

	//Сверки по проходам
	Revise struct {
		Url      string
		Status   int
		Request  interface{}
		Response interface{}
	}

	//автосверка
	Resolve struct {
		Url      string
		Status   int
		Request  interface{}
		Response interface{}
	}

	//логин
	Login struct {
	}

	//отмена
	Cancel struct {
		Target int
		Reason processing.CancelPassRequest_CancelPassReason
	}

	//проверка парковки
	Parking struct {
		RP processing.CheckParkingResponse_Result
		R  *processing.CheckParkingRequest
	}

	//закоытие периода агрегации
	Complete struct {
		StartPass int
		Passes    []int
		Sum       int
	}

	//закрытие периода агрегации по карте
	CompleteWithCalculate struct {
		Pan string
	}

	WebAPIPasses struct {
		Passes []int
	}

	RegisterFaceId struct {
		FaceId      string
		RedirectURL string
	}

	TWPGCreateAndPayOrderStep struct {
		CustomerId string
		OrderId    uint64
	}

	TWPGOrderStatus struct {
		OrderId     uint64
		OrderStatus twpg.OrderStatus
	}

	TWPGReverseOrder struct {
		OrderId uint64
	}

	FaceIdRegistrationStatus struct {
		Id string
	}

	CommentsCRUD struct {
	}

	CardGetFull struct {
		Kind     cards.DiffExportKind
		FileType cards.FileType
	}

	ReaderConfiguration struct {
		FaceList *FaceList
		StopList *StopList
	}

	ReAuth struct {
		Id string
	}

	AuthResponse struct {
	}

	FaceList struct {
		Time     time.Time
		FileType cards.FileType
	}

	StopList struct {
		Time     time.Time
		FileType cards.FileType
	}

	CardStopList struct {
		PassId         string
		Pan            string
		ExpectedStatus cards.CardStatus
		Passes         []int
	}

	ForceReauth struct {
		PassId string
		Passes []int
	}
)

const (
	PaymentTypeFree PaymentType = iota + 1
	PaymentTypePayment
	PaymentTypeStartAggregate
	PaymentTypeAggregate
	PaymentTypePrepayed
)

const (
	RequestTypeNone RequestType = iota
	RequestTypeOffline
	RequestTypeOnline
)

const (
	AuthTypeCorrect AuthType = iota
	AuthTypeIncorrect
	AuthTypeRefund
	AuthTypeMCTokenCorrect
	AuthTypeMCTokenIncorrect
	AuthTypeMKEmulatorSuccess
	AuthTypeMKEmulatorUnsuccess
	AuthTypeMKEmulatorRandom
	AuthTypeUnsuccessWithReauth
	AuthTypeUnsuccessWithoutReauth
)

var Now func() uint64
var NowBackup = func() uint64 {
	return uint64(time.Now().UnixNano())
}

var NowCustom = func(hour, min int) func() uint64 {
	now := time.Now()
	return func() uint64 {
		return uint64(time.Date(
			now.Year(), now.Month(), now.Day(), hour, min, 0, 0, time.UTC).UnixNano())
	}
}

var NowFullDate = func(year, month, day, hour, min, sec int) func() uint64 {
	return func() uint64 {
		return uint64(time.Date(
			year, time.Month(month), day, hour, min, sec, 0, time.UTC).UnixNano())
	}
}
