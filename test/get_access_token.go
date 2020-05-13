package test

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"lab.dt.multicarta.ru/tp/common/messages/user"
	"net/http"
	"testing"
)

func GetAccessToken(t *testing.T) (AT string, err error) {
	httpClient := http.DefaultClient

	rq := &user.LoginRequest{
		Email:    "test@1234",
		Password: "test1234",
	}

	usrResp := &user.JWTResponse{}

	data, _ := json.Marshal(rq)
	loginData := bytes.NewReader(data)

	//пробуем залогиниться
	resp, err := httpClient.Post(ApmApiUrl+"/twirp/proto.ApmAPIGatewayPublic/Login", "application/json", loginData)
	//неизвестная ошибка
	if resp == nil {
		return "", errors.New("Login Resp is nil")
	}
	if err != nil {
		log.Error(err)
		return
	}

	//логин прошел успешно
	if resp.StatusCode == http.StatusOK {
		responseData, err := ioutil.ReadAll(resp.Body)
		if err = json.Unmarshal(responseData, &usrResp); err != nil {
			log.Error(err)
			return "", errors.Errorf("Login resp unmarshal error: %s", err)
		}
		return usrResp.AccessToken, nil
	}

	//логин неуспешен
	if resp.StatusCode != http.StatusNotFound {
		//логин неуспешен, неизвестный ответ сервера/ошибка сервера
		log.Error(err)
		return "", errors.Errorf("Unknown Login api server status: %d", resp.StatusCode)
	}

	//регистрация
	rqReg := &user.RegisterRequest{
		Email:     "test@1234",
		FirstName: "Tester",
		LastName:  "Test",
	}
	data, _ = json.Marshal(rqReg)
	registerData := bytes.NewReader(data)

	resp, err = httpClient.Post(ApmApiUrl+"/twirp/proto.ApmAPIGatewayPublic/Register", "application/json", registerData)
	//ошибка регистрации
	if resp == nil {
		return "", errors.New("Error Register resp is nil")
	}

	if err != nil {
		log.Error(err)
		return
	}

	//успешная регистрация
	responseData, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(responseData, &usrResp); err != nil {
		log.Error(err)
		return "", errors.Errorf("Register resp unmarshal error: %s", err)
	}

	return usrResp.AccessToken, nil
}
