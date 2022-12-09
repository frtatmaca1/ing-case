package proxy

import (
	"net/http"
	"testing"

	"github.com/frtatmaca/case/model/request"
	"github.com/frtatmaca/case/test"
)

func Test_CreatePerson(t *testing.T) {
	type args struct {
		personGreeting request.PersonGreeting
	}
	tests := []struct {
		name       string
		statusCode int
		respBody   string
		args       args
		wantErr    bool
	}{
		{
			name:       "it should do a http req",
			statusCode: http.StatusOK,
			respBody:   "{'status':'200'}",
			args: args{
				personGreeting: request.PersonGreeting{
					Person:   request.Person{Name: "William", Age: 42},
					Greeting: "Hello William (42)",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.MockServer(tt.statusCode, []byte(tt.respBody))
			defer server.Close()

			c := NewClient(server.URL, http.DefaultClient)
			_, err := c.CreatePerson(tt.args.personGreeting)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetPerson(t *testing.T) {
	type args struct {
		person request.Person
	}
	tests := []struct {
		name       string
		statusCode int
		respBody   string
		args       args
		wantErr    bool
	}{
		{
			name:       "it should do a http req",
			statusCode: http.StatusOK,
			respBody:   "{'status':'200', 'body': '{'name': 'William', 'age':42}'}",
			args: args{
				person: request.Person{Name: "William", Age: 42},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := test.MockServer(tt.statusCode, []byte(tt.respBody))
			defer server.Close()

			c := NewClient(server.URL, http.DefaultClient)
			_, err := c.GetPerson()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_GetBaseUrl(t *testing.T) {
	t.Run("it should return base url", func(t *testing.T) {
		server := test.MockServer(http.StatusOK, []byte(`{'status':'200'}`))
		defer server.Close()

		c := NewClient(server.URL, http.DefaultClient)
		baseUrl := c.GetBaseUrl()

		if baseUrl != server.URL {
			t.Errorf("GetBaseUrl() error")
			return
		}
	})
}
