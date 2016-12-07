/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   auth.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/14 21:54:12 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/15 15:40:09 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package apicpa

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

/*
 POST
/api/v1/oauth/token

Paramètres d'entrée

    client_id
    (clef publique)
    client_secret
    (clef secrète)

Retour

{
  "data": {
    "id": ID_TOKEN,
    "type": "tokens",
    "attributes": {
      "access_token": STRING, // Le jeton de connexion
    }
  }
}
*/
func getAuthRequest() (result *http.Request, err error) {
	authUrl := conf.CPA_API_URI + conf.CPA_AUTH_URL
	data := url.Values{}
	data.Set("client_id", conf.ID_PUBLIC_SERVICE)
	data.Add("client_secret", conf.ID_PRIVATE_SERVICE)

	if result, err = http.NewRequest("POST", authUrl, bytes.NewBufferString(data.Encode())); err == nil {
		result.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		result.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	}
	return
}

func authService() (result *http.Response, err error) {

	req, err := getAuthRequest()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	return client.Do(req)
}

func Authenticate() (result string, err error) {
	type CPAToken struct {
		Access_token string `json:"access_token"`
	}

	resp, err := authService()
	if err != nil {
		return "", err
	} else if resp.StatusCode != 200 {
		return "", getAPIErrors(resp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var receiver models.JSONContentSingleData
	err = json.Unmarshal(body, &receiver)
	if err != nil {
		return "", errors.New("Service response is unidentified")
	}
	var token CPAToken
	err = mapstructure.Decode(receiver.Data.Attributes, &token)
	if err != nil {
		return "", errors.New("Service response is unidentified")
	}
	return token.Access_token, nil
}
