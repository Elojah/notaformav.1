/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   collection.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/14 21:54:18 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/14 18:00:40 by hdezier          ###   ########.fr       */
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
/api/v1/services/ID_SERVICE
/relationships/collections
Toutes les collections du service

Paramètres d'entrée

    access_token

Retour

{
  "data": [{
    "id": ID_COLLECTION,
    "type": "collections",
    "attributes": {
      "nom": STRING,
      "description": STRING,
      "tableau_de_donnees": BOOLEAN, // Vrai si la donnée est un tableau de champs
      "jeton_fc_lecture_ecriture": BOOLEAN, // Vrai si jeton FranceConnect requis en lecture et écriture
      "jeton_fc_lecture_seulement": BOOLEAN // Vrai si jeton FranceConnect requis en lecture seulement
    }
  }]
}

Status

    200 - Collections trouvées
    400 - Paramètres d'entrée incorrect

*/
func GetCollectionID(collectionName []string, accessToken string) (collectionsID map[string]string, err error) {

	type CPACollectionModel struct {
		Nom                        string `json:"nom"`
		Description                string `json:"description"`
		Tableau_de_donnees         bool   `json:"tableau_de_donnees"`
		Jeton_fc_lecture_ecriture  bool   `json:"jeton_fc_lecture_ecriture"`
		Jeton_fc_lecture_seulement bool   `json:"jeton_fc_lecture_seulement"`
	}

	collectionInfoUrl := conf.CPA_API_URI + conf.CPA_COLLECTION_URL + conf.POST_ACCESS_TOKEN + accessToken
	data, err := get(collectionInfoUrl)
	if err != nil {
		return nil, err
	}
	var receiver models.JSONContent
	err = json.Unmarshal(data, &receiver)
	if err != nil {
		return nil, errors.New("Service global response is unidentified")
	}
	collectionsID = make(map[string]string)
	for _, value := range receiver.Data {
		id := value.Id
		var collection CPACollectionModel
		err := mapstructure.Decode(value.Attributes, &collection)
		if err != nil {
			return nil, errors.New("Service detail response is unidentified")
		}
		name := collection.Nom
		if stringInSlice(name, collectionName) {
			collectionsID[name] = id
		}
	}
	return
}
