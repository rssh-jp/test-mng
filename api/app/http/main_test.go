package main

import (
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
		r, err := http.Get(ts.URL + "/login?id=test&password=test")
		if err != nil {
			t.Fatal(err)
		}

		defer r.Body.Close()

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		var actual domain.Token
		err = json.Unmarshal(data, &actual)
		if err != nil {
			t.Fatal(err)
		}

		expect := domain.Token{
			ID:    "test",
			Token: "J+!N>ip\"asYzQ%Wk#t_upS\\mt#V|w>{i",
		}

		if !reflect.DeepEqual(expect, actual) {
			t.Errorf("Not match response\nexpect: %+v\nactual: %+v", expect, actual)
		}
	})
}
