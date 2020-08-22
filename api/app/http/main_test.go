package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/rssh-jp/test-mng/api/domain"
)

func TestLogin(t *testing.T) {
	ts := httptest.NewServer(newRouter(true))
	defer ts.Close()

	t.Run("success", func(t *testing.T) {
		sendData := domain.RecvLogin{
			ID:       "test",
			Password: "test",
		}

		var actual domain.SendLogin

		err := post(ts.URL+"/login", sendData, &actual)
		if err != nil {
			t.Fatal(err)
		}

		expect := domain.SendLogin{
			Message: "OK",
			Token: domain.Token{
				ID:    "test",
				Token: "BpLnfgDsc3WD9F3qNfHK6a95jjJkwzDk",
			},
		}

		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("Not match response\nexpect: %+v\nactual: %+v", expect, actual)
		}
	})
}
func TestUsersFetch(t *testing.T) {
	ts := httptest.NewServer(newRouter(true))
	defer ts.Close()

	var token string

	t.Run("login", func(t *testing.T) {
		sendData := domain.RecvLogin{
			ID:       "test",
			Password: "test",
		}

		var actual domain.SendLogin

		err := post(ts.URL+"/login", sendData, &actual)
		if err != nil {
			t.Fatal(err)
		}

		token = actual.Token.Token
	})

	t.Run("usersFetch", func(t *testing.T) {
		sendData := domain.RecvUsersFetch{
			Token: token,
		}

		var actual domain.SendUsersFetch

		err := post(ts.URL+"/users/fetch", sendData, &actual)
		if err != nil {
			t.Fatal(err)
		}

		if len(actual.Users) != 2 {
			t.Error("Could not match users length", len(actual.Users))
		}

		t.Log(actual.Users)
	})

	t.Run("usersGetOwn", func(t *testing.T) {
		sendData := domain.RecvUsersGetOwn{
			Token: token,
		}

		var actual domain.SendUsersGetOwn

		err := post(ts.URL+"/users/getown", sendData, &actual)
		if err != nil {
			t.Fatal(err)
		}

		expect := domain.SendUsersGetOwn{
			Message: "OK",
			User: domain.User{
				ID:   "test",
				Name: "test-name",
				Age:  32,
			},
		}

		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("Not match response\nexpect: %+v\nactual: %+v", expect, actual)
		}
	})

	t.Run("usersUpdate", func(t *testing.T) {
		t.Run("update", func(t *testing.T) {
			sendData := domain.RecvUsersUpdate{
				Token: token,
				User: domain.User{
					ID:   "test",
					Name: "modify-test-name",
					Age:  100,
				},
			}

			var actual domain.SendUsersUpdate

			err := post(ts.URL+"/users/update", sendData, &actual)
			if err != nil {
				t.Fatal(err)
			}

			expect := domain.SendUsersUpdate{
				Message: "OK",
			}

			if !reflect.DeepEqual(expect, actual) {
				t.Errorf("Not match response\nexpect: %+v\nactual: %+v", expect, actual)
			}
		})

		t.Run("confirm", func(t *testing.T) {
			sendData := domain.RecvUsersGetOwn{
				Token: token,
			}

			var actual domain.SendUsersGetOwn

			err := post(ts.URL+"/users/getown", sendData, &actual)
			if err != nil {
				t.Fatal(err)
			}

			expect := domain.SendUsersGetOwn{
				Message: "OK",
				User: domain.User{
					ID:   "test",
					Name: "modify-test-name",
					Age:  100,
				},
			}

			if !reflect.DeepEqual(expect, actual) {
				t.Errorf("Not match response\nexpect: %+v\nactual: %+v", expect, actual)
			}
		})
	})
}

func post(url string, sendData interface{}, response interface{}) error {
	sendDataB, err := json.Marshal(sendData)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(sendDataB)

	r, err := http.Post(url, "application/json", reader)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return err
	}

	return nil
}
