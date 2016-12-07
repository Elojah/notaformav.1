/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   models_test.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/05 19:18:57 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/06 19:57:23 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestModels(t *testing.T) {

	testVarF := FormationModel{}
	err := json.Unmarshal(exampleF, &testVarF)
	if err != nil {
		fmt.Println("Failed example unmarshal")
		t.Error("Failed example unmarshal" + err.Error())
	}

	fmt.Println(SerializeFieldsNameSQL(testVarF))
	fmt.Println(SerializeFieldsNameSQLCreation(testVarF))
	testVarO := OrganismModel{}
	json.Unmarshal(exampleO, &testVarO)

	// testVarCF := CommentFormationModel{}

	// testVarCO := CommentOrganismModel{}

	// fmt.Println(`testVarF: `, testVarF.SerializeSQL())
	// if err != nil {
	// 	fmt.Println("Failed inserting formation")
	// 	t.Error("Failed example unmarshal" + err.Error())
	// }
	// fmt.Println(`testVarO: `, testVarO.SerializeSQL())
	// fmt.Println(`testVarCF: `, testVarCF.SerializeSQL())
	// fmt.Println(`testVarCO: `, testVarCO.SerializeSQL())
}

var (
	exampleO = []byte(`
{
  "name": "3WA",
  "corporatename": "3W Academy",
  "contact": {
    "street": "3 Quai Kléber c/o Regus, Tour Sébastopol",
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
      "Création site internet",
      "Intégration web"
    ],
    "Information, communication": [
      "Langage PHP"
    ]
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
	    "Certification"
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
	    }
	  ],
	  "lastupdate": "",
	  "ref": ""
	}`)
)
