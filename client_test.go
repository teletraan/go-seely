package seely

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/openlyinc/pointy"
	"github.com/syfun/go-graphql"
)

type graphqlHandler func(*graphql.Request) *graphql.Response

type graphqlMux struct {
	mu sync.RWMutex
	m  map[string]graphqlHandler
}

func (mux *graphqlMux) handlerFunc(operationName string, f graphqlHandler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if mux.m == nil {
		mux.m = make(map[string]graphqlHandler)
	}

	mux.m[operationName] = f
}

func (mux *graphqlMux) want(operationName string, data graphql.JSON) {
	mux.handlerFunc(operationName, func(req *graphql.Request) *graphql.Response {
		return &graphql.Response{Data: graphql.JSON{operationName: data}}
	})
}

func setup() (client *Client, mux *graphqlMux, teardown func()) {
	mux = new(graphqlMux)

	mux.want("login", graphql.JSON{"token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InlAdC5pbyIsImV4cCI6MTU4MTA1ODQxNSwiZW1haWwiOiJ5QHQuaW8iLCJhY2NvdW50IjoieUB0LmlvIn0.D7HSd50jspa_YrNuToX5scGzvD_IrNxXhyBSd1OnEgA"})

	apiHandler := http.NewServeMux()
	apiHandler.HandleFunc("/graphql/", func(w http.ResponseWriter, r *http.Request) {
		var req graphql.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Fprintln(w, err.Error())
			http.Error(w, "Malformed request.", http.StatusBadRequest)
			return
		}
		handler, ok := mux.m[req.OperationName]
		if !ok {
			fmt.Fprintf(w, "unspported operation name '%v'\n", req.OperationName)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		resp := handler(&req)
		b, _ := json.Marshal(resp)
		w.Write(b)
	})
	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)
	client = New(server.URL+"/graphql/", "", "")
	return client, mux, server.Close
}

var page = &Page{NoPage: pointy.Bool(true)}
