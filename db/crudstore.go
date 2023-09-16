package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handleHTTPError(res *http.Response) error {
	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		return fmt.Errorf("[error %v] %v", res.StatusCode, string(body))
	}

	return nil
}

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
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	err = handleHTTPError(res)
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
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	err = handleHTTPError(res)
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
	req, err := http.NewRequest(http.MethodPost, s.ep("put", nil), strings.NewReader(string(jdoc)))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	err = handleHTTPError(res)
	if err != nil {
		return err
	}
	return nil
}

func (s CRUDStore[T]) Update(q Query, newValue T) error {
	jbody, err := json.Marshal(Query{
		"query": q,
		"doc":   newValue,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, s.ep("update", nil), strings.NewReader(string(jbody)))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	err = handleHTTPError(res)
	if err != nil {
		return err
	}
	return nil
}

func (s CRUDStore[T]) Delete(q Query) error {
	panic("Not implemented!")
}
