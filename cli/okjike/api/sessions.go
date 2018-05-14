package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/url"
	"github.com/skip2/go-qrcode"
)

type Session struct {
	UUID string `json:"uuid"`
}

func SessionsCreate() (*Session, error) {
	resp, err := http.Get(ApiURL("sessions.create"))
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s := Session{}

	err = json.Unmarshal(data, &s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func CreateLoginQrCode(session *Session) ([]byte) {
	loginUrl := fmt.Sprintf("https://ruguoapp.com/account/scan?uuid=%s", session.UUID)
	encoded := url.QueryEscape(loginUrl)
	qrData := fmt.Sprintf("jike://page.jk/web?url=%s&displayHeader=false&displayFooter=false", encoded)
	pngData, err := qrcode.Encode(qrData, qrcode.Medium, 256)
	if err != nil {
		panic(err)
	}

	return pngData
}
