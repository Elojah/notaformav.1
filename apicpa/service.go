/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   service.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/14 21:54:14 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/14 18:02:38 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package apicpa

import (
	"encoding/json"
	"errors"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"github.com/mitchellh/mapstructure"
)

/* GET
/api/v1/services
Liste des services référencés sur l'API

Paramètres d'entrée

    access_token

Retour

{
  "data": [{
    "id": ID_SERVICE,
    "type": "services",
    "attributes": {
      "nom": STRING,
      "description": STRING,
      "site_internet": STRING
    }
  }]
}

*/
func GetServiceID(serviceName []string, accessToken string) (servicesID map[string]string, err error) {

	type CPAServiceModel struct {
		Nom           string `json:"nom"`
		Description   string `json:"description"`
		Site_internet string `json:"site_internet"`
	}

	serviceInfoUrl := conf.CPA_API_URI + conf.CPA_SERVICE_URL + conf.POST_ACCESS_TOKEN + accessToken
	data, err := get(serviceInfoUrl)
	if err != nil {
		return nil, err
	}
	var content models.JSONContent
	err = json.Unmarshal(data, &content)
	if err != nil {
		return nil, errors.New("Service response is unidentified")
	}
	servicesID = make(map[string]string)
	for _, value := range content.Data {

		id := value.Id
		var service CPAServiceModel
		err := mapstructure.Decode(value.Attributes, &service)
		if err != nil {
			return nil, errors.New("Service response is unidentified")
		}
		name := service.Nom
		if stringInSlice(name, serviceName) {
			servicesID[name] = id
		}
	}
	return
}
