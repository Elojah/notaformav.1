/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/13 13:35:40 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 17:34:51 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
	// // "fmt"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"github.com/go-server/psql"
	"github.com/jmoiron/sqlx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"reflect"
)

type SQLRequest struct {
	Model reflect.Value
	Query func(query models.QueryModel) (*sqlx.Rows, error)
}

type Controller struct {
	ServicesID                  map[string]string
	CollectionsID               map[string]string
	AccessToken                 string
	Conf                        *oauth2.Config
	SearchOrganismRequest       SQLRequest
	SearchOrganismMinRequest    SQLRequest
	SearchFormationRequest      SQLRequest
	SearchFormationMinRequest   SQLRequest
	SearchFormationMatchRequest SQLRequest
	FormationsRequest           SQLRequest
	FormationsMinRequest        SQLRequest
	OrganismsRequest            SQLRequest
	OrganismsMinRequest         SQLRequest
	CommentFormationRequest     SQLRequest
	CommentOrganismRequest      SQLRequest
}

func (ctrl *Controller) InitDB() (err error) {
	err = psql.Open()
	if err != nil {
		return
	}

	// TEST
	// data, err := dbheroku.GetCollectionDataID("establishments")
	// if err != nil {
	// 	return err
	// }
	// for _, val := range data {
	// 	row, err := dbheroku.GetData(val, "establishments")
	// 	var establishment models.EstablishmentModel
	// 	err = row.StructScan(&establishment)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println(establishment)
	// }
	// var queryOptions models.QueryModel
	// queryOptions.Page = make(map[string]string)
	// queryOptions.Filter = make(map[string]string)
	// rows, err := dbheroku.Search(queryOptions, "lyon")
	// if err != nil {
	// 	return
	// }
	// defer rows.Close()
	// var establishments []models.EstablishmentModel
	// for rows.Next() {
	// 	var establishment models.EstablishmentModel
	// 	err = rows.StructScan(&establishment)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	establishments = append(establishments, establishment)
	// }
	// for _, val := range establishments {
	// 	fmt.Println(val)
	// }
	return
}

/*
SearchOrganismRequest       SQLRequest
SearchOrganismMinRequest       SQLRequest
SearchFormationRequest      SQLRequest
SearchFormationMinRequest      SQLRequest
SearchFormationMatchRequest SQLRequest
FormationsRequest           SQLRequest
FormationsMinRequest        SQLRequest
OrganismsRequest            SQLRequest
OrganismsMinRequest            SQLRequest
CommentFormationRequest     SQLRequest
CommentOrganismRequest      SQLRequest
*/

func (ctrl *Controller) Init() (err error) {
	organismSpec := func(input string) string {
		return input
	}
	formationSpec := func(string) string {
		return ` id, name, parentid, programme, globalrating `
	}
	formationSpecRC := func(string) string {
		return ` id, name, romecode, location_adress_zip `
	}
	ctrl.SearchOrganismRequest = SQLRequest{
		Model: reflect.ValueOf(&models.OrganismModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.Search(`organism`, models.OrganismModel{}, query)
		},
	}
	ctrl.SearchOrganismMinRequest = SQLRequest{
		Model: reflect.ValueOf(&models.OrganismModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.SearchSpec(`organism`, models.OrganismModel{}, query, organismSpec)
		},
	}
	ctrl.SearchFormationRequest = SQLRequest{
		Model: reflect.ValueOf(&models.FormationModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.Search(`formation`, models.FormationModel{}, query)
		},
	}
	ctrl.SearchFormationMinRequest = SQLRequest{
		Model: reflect.ValueOf(&models.FormationModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.SearchSpec(`formation`, models.FormationModel{}, query, formationSpec)
		},
	}
	ctrl.SearchFormationMatchRequest = SQLRequest{
		Model: reflect.ValueOf(&models.FormationModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			if len(query.Filter[`romecode`]) == 0 {
				return nil, nil
			}
			return psql.SearchMatchArray(`formation`, models.FormationModel{}, query, formationSpecRC)
		},
	}
	ctrl.FormationsRequest = SQLRequest{
		Model: reflect.ValueOf(&models.FormationModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRows(`formation`, models.FormationModel{}, query)
		},
	}
	ctrl.FormationsMinRequest = SQLRequest{
		Model: reflect.ValueOf(&models.FormationModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRowsSpec(`formation`, models.FormationModel{}, query, formationSpec)
		},
	}
	ctrl.OrganismsRequest = SQLRequest{
		Model: reflect.ValueOf(&models.OrganismModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRows(`organism`, models.OrganismModel{}, query)
		},
	}
	ctrl.OrganismsMinRequest = SQLRequest{
		Model: reflect.ValueOf(&models.OrganismModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRowsSpec(`organism`, models.OrganismModel{}, query, organismSpec)
		},
	}
	ctrl.CommentFormationRequest = SQLRequest{
		Model: reflect.ValueOf(&models.CommentFormationModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRows(`comment_formation`, models.CommentFormationModel{}, query)
		},
	}
	ctrl.CommentOrganismRequest = SQLRequest{
		Model: reflect.ValueOf(&models.CommentOrganismModel{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRows(`comment_organism`, models.CommentOrganismModel{}, query)
		},
	}
	// OAUTH2 Conf for ggl login
	ctrl.Conf = &oauth2.Config{
		ClientID:     conf.GGL_CLIENT_ID,
		ClientSecret: conf.GGL_SECRET_CODE,
		RedirectURL:  conf.MY_HOME_URL + `auth`,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
			"https://www.googleapis.com/auth/plus.me",        //	Know who you are on Google
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return
}

func (ctrl *Controller) CommentFormationRequestWithVote(userid string) SQLRequest {
	return SQLRequest{
		Model: reflect.ValueOf(&models.CommentFormationModelWithVote{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRowsWithSubRequest(`comment_formation`, models.CommentFormationModel{}, query,
				`(SELECT vote FROM comment_formation_vote_user
				WHERE id=comment_formation.id
				AND formationid=comment_formation.parentid
				AND userid='`+userid+`') AS vote`)
		},
	}
}

func (ctrl *Controller) CommentOrganismRequestWithVote(userid string) SQLRequest {
	return SQLRequest{
		Model: reflect.ValueOf(&models.CommentOrganismModelWithVote{}),
		Query: func(query models.QueryModel) (*sqlx.Rows, error) {
			return psql.GetRowsWithSubRequest(`comment_organism`, models.CommentOrganismModel{}, query,
				`(SELECT vote FROM comment_organism_vote_user
				WHERE id=comment_organism.id
				AND organismid=comment_organism.parentid
				AND userid='`+userid+`') AS vote`)
		},
	}
}
