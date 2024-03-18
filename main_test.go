package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func httpClient(client *http.Client, url string) {

	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	s, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	if len(s) > 1024 {
		s = s[0:1024]
	}

	fmt.Println(string(s))
}

func TestEmailPassIsInvalid(t *testing.T) {
	// expected fail
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/reg?email=hello@world.com")

	// expected fail
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/reg?email=a@b.c&pass=a123'45Z")

	// expected fail
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/reg?email=hello3@world.com&pass=dsf23dsaf34")

	// expected fail
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/reg?email=olleh4@moc.dlrow&pass=fcsdf324A1")
}

func TestReg(t *testing.T) {

	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/reg?email=hello96@world.com&pass=DQdqw2'01/dzg")

	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/reg?email=hello99@world.com&pass=ewdev4tvgw^2Hz")
}

func TestAuth(t *testing.T) {
	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/auth?email=hello96@world.com&pass=DQdqw2'01/dzg")

	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/auth?email=hello99@world.com&pass=ewdev4tvgw^2Hz")
}

func TestConfirm(t *testing.T) {
	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/confirm?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvOTZAd29ybGQuY29tIiwiaXNzIjoidGVzdCIsInN1YiI6InNvbWVib2R5IiwiYXVkIjpbInNvbWVib2R5X2Vsc2UiXSwiZXhwIjoxNzExOTM4MzI0LCJuYmYiOjE3MTA3Mjg3MjQsImlhdCI6MTcxMDcyODcyNCwianRpIjoiMSJ9.ks-q5WyeVgQOKBNvdzmQAfmppXE0dScS80mAOO-RY6A")

	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/confirm?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvOTlAd29ybGQuY29tIiwiaXNzIjoidGVzdCIsInN1YiI6InNvbWVib2R5IiwiYXVkIjpbInNvbWVib2R5X2Vsc2UiXSwiZXhwIjoxNzExOTM4MzI0LCJuYmYiOjE3MTA3Mjg3MjQsImlhdCI6MTcxMDcyODcyNCwianRpIjoiMSJ9.KKiALdW5WOInu_tCb_8zQDZgel9wPCP-NCAME1qcjOU")
}

func Test(t *testing.T) {
	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/recommendation?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvOTZAd29ybGQuY29tIiwiaXNzIjoidGVzdCIsInN1YiI6InNvbWVib2R5IiwiYXVkIjpbInNvbWVib2R5X2Vsc2UiXSwiZXhwIjoxNzExOTM4OTY0LCJuYmYiOjE3MTA3MjkzNjQsImlhdCI6MTcxMDcyOTM2NCwianRpIjoiMSJ9.2V0UcOF0ozJcOx6cUZw5a_TrZmObFZqZcvnrrhTpCn4")

	// expected ok
	httpClient(&http.Client{}, "http://127.0.0.1:8888/v1/recommendation?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImhlbGxvOTlAd29ybGQuY29tIiwiaXNzIjoidGVzdCIsInN1YiI6InNvbWVib2R5IiwiYXVkIjpbInNvbWVib2R5X2Vsc2UiXSwiZXhwIjoxNzExOTM4OTY0LCJuYmYiOjE3MTA3MjkzNjQsImlhdCI6MTcxMDcyOTM2NCwianRpIjoiMSJ9.0PdU0Epiw5Hx35Uu9vDo9uT3dSbZCDCmjNkDQHeWxeo")
}
