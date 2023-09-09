package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// var endpoint = func() string {
// 	postfix := "/.netlify/functions/"
// 	envPrefix := os.Getenv("endpointprefix")
// 	if envPrefix != "" {
// 		return "https://" + envPrefix + postfix
// 	}
//
// 	return "https://" + "dotdepot.pyros.dev" + postfix
// }()

type CRUDStore[T IsEmptyer] struct {
	Endpoint   string
	Collection string
	Username   string
	Password   string
}

func (s CRUDStore[T]) ep(action string, parameters map[string]string) string {
	paramStr := ""
	for k, v := range parameters {
		paramStr += fmt.Sprintf("&%s=%s", k, v)
	}
	return fmt.Sprintf("%s/%s?collection=%s&password=%s&username=%s%s", s.Endpoint, action, s.Collection, s.Password, s.Username, paramStr)
}

func (s CRUDStore[T]) Get(q Query) (*T, error) {
	res, err := http.Get(s.ep("one", q.ToParameters()))
	if err != nil {
		return nil, err
	}
	var document T
	err = json.NewDecoder(res.Body).Decode(&document)
	return &document, err
}

func (s CRUDStore[T]) GetMany(q Query) ([]T, error) {
	res, err := http.Get(s.ep("many", q.ToParameters()))
	if err != nil {
		return nil, err
	}
	var documents []T
	err = json.NewDecoder(res.Body).Decode(&documents)
	return documents, err
}

func (s CRUDStore[T]) Put(doc T) error {
	jdoc, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	_, err = http.Post(s.ep("put", nil), "application/json", bytes.NewReader(jdoc))
	return err
}

func (s CRUDStore[T]) Update(q Query, newValue T) error {
	panic("Not implemented!")
}

func (s CRUDStore[T]) Delete(q Query) error {
	panic("Not implemented!")
}
