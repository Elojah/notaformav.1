/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   post.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/14 22:54:29 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/09 19:29:26 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package apicpa

import (
// "bytes"
// "encoding/csv"
// "encoding/json"
// "errors"
// "fmt"
// // "github.com/davecgh/go-spew/spew"
// "github.com/go-server/conf"
// "github.com/go-server/models"
// "github.com/zemirco/uid"
// "net/http"
// "net/url"
// "os"
)

/* POST
/api/v1/services/ID_SERVICE
/relationships/collections/ID_COLLECTION
/relationships/donnees
Ajouter ou modifier une donnée

Entrée
tableau_de_donnees == false

{
  "access_token": "JETON_DE_CONNEXION",
  "fc_token": "JETON_DE_CONNEXION_FRANCECONNECT", (requis si jeton de connexion FranceConnect requis)
  "data": {
    "id": ID_DONNEE, (omis si jeton de connexion FranceConnect requis)
    "type": "donnees",
    "attributes": {
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }
  }
}

tableau_de_donnees == true

{
  "access_token": "JETON_DE_CONNEXION",
  "fc_token": "JETON_DE_CONNEXION_FRANCECONNECT", (requis si jeton de connexion FranceConnect requis)
  "data": {
    "id": ID_DONNEE, (omis si jeton de connexion FranceConnect requis)
    "type": "donnees",
    "attributes": [{
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }]
  }
}

Retour
tableau_de_donnees == false

{
  "data": {
    "id": ID_DONNEE,
    "type": "donnees",
    "attributes": {
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }
  }
}

tableau_de_donnees == true

{
  "data": {
    "id": ID_DONNEE,
    "type": "donnees",
    "attributes": [{
      CHAMPS_DEFINIS_PAR_LE_SERVICE
    }]
  }
}

Status

    201 - Donnée créée
    200 - Donnée modifiée
    400 - Paramètres d'entrée incorrect
*/

// func postCollectionData(postEntry interface{}, collectionID string) (err error) {
// 	postUrl := conf.CPA_API_URI + conf.CPA_COLLECTION_URL + collectionID + "/" + conf.CPA_COLLECTION_DATA_URL

// 	b, err := json.Marshal(postEntry)
// 	// spew.Dump(b)
// 	resp, err := http.Post(postUrl, "application/json;charset=utf-8", bytes.NewReader(b))
// 	if err != nil {
// 		return err
// 	}
// 	if resp.StatusCode == 400 {
// 		fmt.Println("Post to collections failed:" + getAPIErrors(resp).Error())
// 		return errors.New("Failed to post")
// 	}
// 	fmt.Println("Post request DONE")
// 	return
// }

// func PostCollectionFromCSV(collectionID string, accessToken string, filename string, model interface{}) (err error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	reader.Comma = ';'
// 	lines, err := reader.ReadAll()
// 	if err != nil {
// 		return err
// 	}
// 	postEntry := models.CPAPostModel{Data: models.JSONData{Attributes: model}}
// 	postEntry.Access_token = accessToken
// 	postEntry.Fc_token = ""
// 	postEntry.Data.Type = "donnees"
// 	for _, line := range lines {
// 		postEntry.Data.Id = uid.New(11)
// 		err = postEntry.Data.Attributes.(models.Settable).SetFromCsv(line)
// 		if err != nil {
// 			break
// 		}
// 		err = postCollectionData(postEntry, collectionID)
// 		if err != nil {
// 			break
// 		}
// 	}
// 	return
// }

// func PostCollectionFromPost(collectionID string, accessToken string, postData *url.Values, model interface{}) (id string, err error) {
// 	postEntry := models.CPAPostModel{Data: models.JSONData{Attributes: model}}
// 	postEntry.Access_token = accessToken
// 	postEntry.Fc_token = ""
// 	postEntry.Data.Id = uid.New(11)
// 	postEntry.Data.Type = "donnees"
// 	postEntry.Data.Attributes.(models.Settable).SetFromUrl(postData)
// 	fmt.Println("struct data to send:", postEntry.Data.Attributes)
// 	return postEntry.Data.Id, postCollectionData(postEntry, collectionID)
// }
