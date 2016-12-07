/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   controller_test.go                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/09 19:00:00 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/25 15:16:49 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func doRequest(reqUrl string, reqType string) (result *http.Response, err error) {
	req, err := http.NewRequest(reqType, reqUrl, nil)
	if err != nil {
		return nil, err
	}
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
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}
