package db

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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
	var passwordQuery string
	if s.Password != "" {
		passwordQuery = fmt.Sprintf("&password=%s", s.Password)
	} else {
		passwordQuery = ""
	}
	return fmt.Sprintf("%s/%s?collection=%s%s&username=%s%s", s.Endpoint, action, s.Collection, passwordQuery, s.Username, paramStr)
}

func (s CRUDStore[T]) Get(q Query) (*T, error) {
	jquery, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, s.ep("one", nil), strings.NewReader(string(jquery)))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var document T
	err = json.NewDecoder(res.Body).Decode(&document)
	return &document, err
}

func (s CRUDStore[T]) GetMany(q Query) ([]T, error) {
	jquery, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, s.ep("many", nil), strings.NewReader(string(jquery)))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var documents []T
	err = json.NewDecoder(res.Body).Decode(&documents)
	return documents, err
}

func (s CRUDStore[T]) Put(doc T) error {
	panic("Not implemented")
	// jdoc, err := json.Marshal(doc)
	// if err != nil {
	// 	return err
	// }
	// _, err = http.Post(s.ep("put", nil), "application/json", bytes.NewReader(jdoc))
	// return err
}

func (s CRUDStore[T]) Update(q Query, newValue T) error {
	panic("Not implemented!")
}

func (s CRUDStore[T]) Delete(q Query) error {
	panic("Not implemented!")
}
