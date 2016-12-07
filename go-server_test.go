/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   go-server_test.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/25 15:18:03 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/25 16:52:16 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	// "github.com/go-server/models"
	"io/ioutil"
	"net/http"
	"testing"
)

type Adress struct {
	Zip string `json:"zip" db:"zip"`
}

type Location struct {
	Adress Adress `json:"adress" db:"adress"`
}

type Formation struct {
	Name     string   `json:"name" db:"name"`
	RomeCode []string `json:"romecode" db:"romecode"`
	Location Location `json:"location" db:"location"`
}

type JSONMetaPagination struct {
	Total       int `json:"total"`
	Total_pages int `json:"total_pages"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
	Count       int `json:"count"`
}

type JSONData struct {
	Id         string    `json:"id"`
	Type       string    `json:"type"`
	Attributes Formation `json:"attributes"`
	Links      struct {
		Self  string `json:"self"`
		First string `json:"first"`
	} `json:"links"`
	Meta struct {
		Creation     string `json:"creation"`
		Modification string `json:"modification"`
		Version      int    `json:"version"`
	}
}

type JSONContent struct {
	Meta JSONMetaPagination `json:"meta"`
	Data []JSONData         `json:"data"`
}

func doRequest(reqUrl string, reqType string) (result *http.Response, err error) {
	req, err := http.NewRequest(reqType, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("romecode[]", `H1203`)
	q.Add("romecode[]", `H1210`)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	return client.Do(req)
}

func get(reqUrl string) (data []byte, err error) {
	resp, err := doRequest(reqUrl, "GET")
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 && resp.StatusCode != 206 {
		fmt.Println("URL:", reqUrl)
		return nil, errors.New("Service response is not ok")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Service response is unidentified")
	}
	defer resp.Body.Close()
	return body, nil
}

func TestMatch(t *testing.T) {
	data, err := get(`http://www.notaforma.fr/search/match`)
	model := JSONContent{}
	json.Unmarshal(data, &model)
	if err != nil {
		fmt.Println(err.Error())
	}
	spew.Dump(model)
	for _, val := range model.Data {
		spew.Dump(val.Attributes)
	}
	// fmt.Println(string(data))
	// fmt.Println(model.Meta.Count)
}
