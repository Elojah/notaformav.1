/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   psql_test.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/06 17:37:12 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/10 16:19:36 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package psql

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-server/models"
	"github.com/zemirco/uid"
	"testing"
)

func testFormation(t *testing.T) {
	err := Drop(`test_formation`)
	if err != nil {
		fmt.Println(err.Error())
	}

	testVarF := models.FormationModel{}
	err = CreateTable(testVarF, `test_formation`)
	if err != nil {
		t.Error("Failed creating table" + err.Error())
	}
	err = json.Unmarshal(exampleF, &testVarF)
	if err != nil {
		t.Error("Failed example unmarshal" + err.Error())
	}
	err = InsertRow(`test_formation`, testVarF)
	if err != nil {
		t.Error("Failed to insert formation" + err.Error())
	}
}

func testOrganism(t *testing.T) {
	err := Drop(`test_organism`)
	if err != nil {
		fmt.Println(err.Error())
	}

	testVarO := models.OrganismModel{}
	err = CreateTable(testVarO, `test_organism`)
	if err != nil {
		t.Error("Failed creating table" + err.Error())
	}
	err = json.Unmarshal(exampleO, &testVarO)
	if err != nil {
		t.Error("Failed example unmarshal:" + err.Error())
	}
	testVarO.Id = uid.New(11)
	err = InsertRow(`test_organism`, testVarO)
	if err != nil {
		t.Error("Failed to insert organism:" + err.Error())
	}
}

func testDouble(t *testing.T) {
	testVarO := models.OrganismModel{}
	testVarF := models.FormationModel{}
	uniqueKeyMap := map[string]string{
		`name`: `Name`,
	}
	doublon := InsertRowUnique(`test_organism`, testVarO, uniqueKeyMap)
	if doublon {
		t.Error("Double insertion worked, it definitely shouldn't...")
	}
	err := json.Unmarshal(exampleO2, &testVarO)
	if err != nil {
		t.Error("Failed example unmarshal:" + err.Error())
	}
	testVarO.Id = uid.New(11)
	doublon = InsertRowUnique(`test_organism`, testVarO, uniqueKeyMap)
	if !doublon {
		t.Error("Detected not unique insertion, but they are different")
	}
	query := models.QueryModel{}
	query.Filter = make(map[string]string)
	query.Filter[`id`] = testVarO.Id
	row, err := GetRow(`test_organism`, models.OrganismModel{}, query)
	if err != nil || row == nil {
		t.Error("Failed get row by id:" + err.Error())
	}
	testOrgaRes, err := models.RowToStruct(row, &models.OrganismModel{})
	if err != nil {
		t.Error("Failed map row into struct:" + err.Error())
	}
	spew.Dump(testOrgaRes)
	query.Filter[`id`] = testVarF.Id
	row, err = GetRow(`test_formation`, models.FormationModel{}, query)
	testFormRes, err := models.RowToStruct(row, &models.FormationModel{})
	if err != nil {
		t.Error("Failed map row into struct:" + err.Error())
	}
	spew.Dump(testFormRes)
}

func testSearch(t *testing.T) {
	queryOptions := models.QueryModel{}
	queryOptions.Page = make(map[string]string)
	queryOptions.Page["number"] = `25`
	queryOptions.Page["limit"] = `25`
	queryOptions.Page["offset"] = `0`

	rows, err := Search(`test_organism`, models.OrganismModel{}, queryOptions, `art`)
	if err != nil {
		t.Error("Failed to search organisms:" + err.Error())
	}

	defer rows.Close()
	jsonContent := models.JSONContent{}
	for rows.Next() {
		organism, err := models.RowsToStruct(rows, &models.OrganismModel{})
		if err != nil {
			fmt.Println("Error occured retrieving organism organisms: ", err)
			return
		}
		// TODO generic function to build JSONContent around attribute
		jsonContent.Data = append(jsonContent.Data, models.JSONData{
			Attributes: organism,
		})
	}
	spew.Dump(jsonContent)
}

func testIndex(t *testing.T) {
	err := Index(`test_formation`, models.FormationModel{})
	if err != nil {
		t.Error("Failed to index formations:" + err.Error())
	}
	err = Index(`test_organism`, models.OrganismModel{})
	if err != nil {
		t.Error("Failed to index organisms:" + err.Error())
	}
}

func TestAll(t *testing.T) {
	err := Open()
	if err != nil {
		t.Error("Failed example unmarshal" + err.Error())
	}

	testFormation(t)
	testOrganism(t)
	testDouble(t)
	testIndex(t)
	testSearch(t)
}

var (
	exampleO = []byte(`
	{
	  "name": "3WA''",
	  "corporatename": "3W Aca'''demy",
	  "contact": {
	    "street": "3 Quai Kléber c/o 'Regus, Tour Sébastopol",
	    "zip": "67000",
	    "locality": "Strasbourg",
	    "tel": " 0140581972",
	    "fax": "",
	    "website": "http://3wa.fr",
	    "mail": "contact@3wa.fr",
	    "name": ""
	  },
	  "siret": " 75404770200024",
	  "registrationnumber": "11754910875",
	  "fields": {
	    "Arts": [
	      "Créati'on site internet",
	      "Intégration web"
	    ],
	    "Information, communication": [
	      "Langage PH''P"
	    ]
	  }
	}`)
	exampleO2 = []byte(`
		{
		  "name": "13WA''",
		  "corporatename": "3W Aca'''demy",
		  "contact": {
		    "street": "3 Quai Kléber c/o 'Regus, Tour Sébastopol",
		    "zip": "67000",
		    "locality": "Strasbourg",
		    "tel": " 0140581972",
		    "fax": "",
		    "website": "http://3wa.fr",
		    "mail": "contact@3wa.fr",
		    "name": ""
		  },
		  "siret": " 75404770200024",
		  "registrationnumber": "11754910875",
		  "fields": {
		    "Arts": [
		      "Créati'on site internet",
		      "Intégration web"
		    ],
		    "Information, communication": [
		      "Langage PH''P"
		    ]
		  },
		  "formations" : ["4648d6", "asdasd"]
		}`)

	exampleF = []byte(`
	{
	  "name": "FPC Développeur Intégrateur web",
	  "objectives": "La 3W Academy forme des développeurs-intégrateurs web pour un accès direct à l'emploi. A lissue des 3 mois, les étudiants doivent être en mesure de travailler en équipe avec des webdesigners et des chefs de projet afin de réaliser des sites Internet modernes, fonctionnels et solides",
	  "programme": [
	    "Intégration web (112 heures), créer des sites responsive, création de 7 sites en partant de maquettes, HTML5, CSS",
	    "Développement web (288 heures), créer des sites dynamiques, création de 10 applications web, Javascript, PHP"
	  ],
	  "validation": [
	    "Titre ou diplôme homologué"
	  ],
	  "type": [
	  "Certification",
	  " TE,ST ",
	  " TE,ST ",
	  "Certification",
	  " TE,ST "
	  ],
	  "outputlevel": "niveau III (BTS, DUT)",
	  "intendedoccupations": [
	    {
	      "romcode": "E1104",
	      "entitle": "Conception de contenus multimédias"
	    },
	    {
	      "romcode": "M1805",
	      "entitle": "Études et développement informatique"
	    }
	  ],
	  "pedagogicalterms": "",
	  "duration": [
	    "400 heures en centre"
	  ],
	  "workstudyterms": [
	    "La formation ne se déroule pas en alternance"
	  ],
	  "publicservicecontract": false,
	  "financers": null,
	  "publicaccess": [
	    "Tout public"
	  ],
	  "admissionterm": null,
	  "entrylevel": "niveau IV (BP, BT, baccalauréat professionnel ou technologique)",
	  "prerequisite": [
	    "Aucune connaissance préalable de la programmation n'est requise mais une mise à niveau préalable peut être requise après entretien avec le candidat"
	  ],
	  "location": {
	    "street": "8 Rue du Faubourg de Saverne",
	    "zip": "",
	    "locality": "67000 Strasbourg",
	    "tel": "01 40 58 19 72",
	    "fax": "",
	    "website": "http://3wa.fr",
	    "mail": "contact@3wa.fr",
	    "name": ""
	  },
	  "contact": {
	    "street": "",
	    "zip": "",
	    "locality": "",
	    "tel": "01 40 58 19 72",
	    "fax": "",
	    "website": "http://3wa.fr",
	    "mail": "contact@3wa.fr",
	    "name": "Claire Viguier"
	  },
	  "eligibilityemployee": [
	    {
	      "name": "",
	      "code": "180564",
	      "startperiodvalidity": "13/06/2016",
	      "endperiodvalidity": "31/03/2017",
	      "region": "Toutes les régions",
	      "professionalbranch": [
	        {
	          "nafcode": "25.11Z"
	        },
	        {
	          "nafcode": "43.32C"
	        },
	        {
	          "nafcode": "58.12Z"
	        },
	        {
	          "nafcode": "58.21Z"
	        },
	        {
	          "nafcode": "58.29A"
	        },
	        {
	          "nafcode": "58.29B"
	        },
	        {
	          "nafcode": "58.29C"
	        },
	        {
	          "nafcode": "62.01Z"
	        },
	        {
	          "nafcode": "62.02A"
	        },
	        {
	          "nafcode": "62.02B"
	        },
	        {
	          "nafcode": "62.03Z"
	        },
	        {
	          "nafcode": "62.09Z"
	        },
	        {
	          "nafcode": "63.11Z"
	        },
	        {
	          "nafcode": "63.12Z"
	        },
	        {
	          "nafcode": "68.20B"
	        },
	        {
	          "nafcode": "68.32A"
	        },
	        {
	          "nafcode": "70.21Z"
	        },
	        {
	          "nafcode": "70.22Z"
	        },
	        {
	          "nafcode": "71.12B"
	        },
	        {
	          "nafcode": "71.20B"
	        },
	        {
	          "nafcode": "73.20Z"
	        },
	        {
	          "nafcode": "74.30Z"
	        },
	        {
	          "nafcode": "74.90B"
	        },
	        {
	          "nafcode": "78.10Z"
	        },
	        {
	          "nafcode": "78.30Z"
	        },
	        {
	          "nafcode": "82.30Z"
	        },
	        {
	          "nafcode": "90.04Z"
	        }
	      ]
	    },
	    {
	      "name": "TEST",
	      "code": "180564",
	      "startperiodvalidity": "13/06/2016",
	      "endperiodvalidity": "31/03/2017",
	      "region": "2442",
	      "professionalbranch": [
	        {
	          "nafcode": "25.11Z"
	        },
	        {
	          "nafcode": "43.32C"
	        }
	      ]
	    }
	  ],
	  "eligibilityjobseeker": null,
	  "eligibilityall": null,
	  "sessions": [
	  {
	    "startperiodvalidity": "03/10/2016",
	    "endperiodvalidity": "30/12/2016",
	    "adress": {
	      "street": "67000 Strasbourg",
	      "zip": "",
	      "locality": "3 Quai Kléber c/o Regus, Tour Sébastopol"
	    },
	    "recruitmentstate": "",
	    "terms": "Dispositif de formation en entrées et sorties permanentes"
	  },
	  {
	    "startperiodvalidity": "03/10/2014",
	    "endperiodvalidity": "30/12/2014",
	    "adress": {
	      "street": "67000 Strasbourgasdfs",
	      "zip": "",
	      "locality": "3 Quai Sébastopol"
	    },
	    "recruitmentstate": "",
	    "terms": "Dispositif den entrées et sorties permanentes"
	  }
	  ],
	  "lastupdate": "",
	  "ref": ""
	}`)
)
