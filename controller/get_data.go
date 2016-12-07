/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   get_data.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/13 13:19:33 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 17:44:01 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"bytes"
	"github.com/go-server/cache"
	"github.com/go-server/models"
	"github.com/go-server/psql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// Handler is useless FTM but interesting, maybe one day...
type Handler func(...interface{}) (interface{}, error)

var EmptyHandler = func(...interface{}) (interface{}, error) { return nil, nil }

func MakeHandler(fn interface{}) Handler {
	v_fn := reflect.ValueOf(fn)
	if v_fn.Type().Kind() != reflect.Func {
		return EmptyHandler
	}
	return func(args ...interface{}) (interface{}, error) {
		vargs := make([]reflect.Value, len(args))
		for i, val := range args {
			vargs[i] = reflect.ValueOf(val)
		}
		results := v_fn.Call(vargs)
		return results[0].Interface(), results[1].Interface().(error)
	}
}

// Helpers
func setMeta(queryOptions *models.QueryModel, model *models.JSONContent) (err error) {
	if len(model.Data) > 0 {
		model.Meta.Total = model.Data[0].Attributes.(models.Countable).GetCount()
	} else {
		model.Meta.Total = 0
	}
	model.Meta.Limit, err = strconv.Atoi(queryOptions.Page["limit"])
	if err != nil {
		model.Meta.Limit = 25 // Default value
	}
	model.Meta.Offset, err = strconv.Atoi(queryOptions.Page["offset"])
	if err != nil {
		model.Meta.Offset = 0 // Default value
	}
	model.Meta.Total_pages = 1
	if model.Meta.Limit != 0 {
		model.Meta.Total_pages += (model.Meta.Total / model.Meta.Limit)
	}
	model.Meta.Count = model.Meta.Total - model.Meta.Offset
	return
}

func reqToQueryOptions(r *http.Request, p httprouter.Params) (queryOptions models.QueryModel, err error) {
	r.ParseForm()
	queryOptions.Page = make(map[string]string)
	queryOptions.Filter = make(map[string]string)
	queryOptions.Page["number"] = r.FormValue("page")
	queryOptions.Page["limit"] = r.FormValue("limit")

	number, err := strconv.Atoi(r.FormValue("page"))
	limit, err := strconv.Atoi(r.FormValue("limit"))

	queryOptions.Page["offset"] = strconv.Itoa((number - 1) * limit)
	// Unsafe then ?
	queryOptions.Search = r.FormValue(`q`)

	id := r.FormValue(`id`)
	if len(id) != 0 {
		queryOptions.Filter[`id`] = id
	}

	parentid := p.ByName(`id`)
	if len(parentid) != 0 {
		queryOptions.Filter[`parentid`] = parentid
	}

	if len(r.FormValue("romecode[]")) != 0 {
		for key, val := range r.Form {
			if key != `romecode[]` {
				continue
			}
			var b bytes.Buffer
			for _, value := range val {
				b.WriteString(models.SerializeSQL(value))
				b.WriteString(`,`)
			}
			queryOptions.Filter["romecode"] = b.String()[:b.Len()-1] // tmp
			break
		}
	}

	fmt.Println(queryOptions.Filter["romecode"])
	return
}

func (req SQLRequest) WriteAsJSON(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	queryOptions, err := reqToQueryOptions(r, p)
	rows, err := req.Query(queryOptions)
	jsonContent := models.JSONContent{}
	if err != nil || rows == nil {
		fmt.Println("Error occured during data retrieving: ", err)
		setMeta(&queryOptions, &jsonContent)
		data, _ := json.Marshal(jsonContent)
		w.Write(data)
		return
	}
	defer rows.Close()
	start := time.Now()
	for rows.Next() {
		rowStruct, err := models.RowsToStruct(rows, req.Model)
		if err != nil {
			fmt.Println("Error occured retrieving rowStruct: ", err)
			return
		}
		// TODO generic function to build JSONContent around attribute
		jsonContent.Data = append(jsonContent.Data, models.JSONData{
			Attributes: rowStruct,
		})
	}
	// Meta setting
	err = setMeta(&queryOptions, &jsonContent)
	data, err := json.Marshal(jsonContent)
	if err != nil {
		fmt.Println("Error occured retrieving rowStruct: ", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Println("Treatment took ", elapsed)
	start = time.Now()
	w.Write(data)
	elapsed = time.Since(start)
	fmt.Println("Writing data took ", elapsed)
}

// !Helpers

func (ctrl *Controller) RenderSearch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "search.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) RenderFormation(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	query := models.QueryModel{}
	query.Filter = make(map[string]string)
	// TODO Check p.ByName(id) is valid !!!
	query.Filter[`id`] = p.ByName("id")

	// TODO remove test_DB
	row, err := psql.GetRow(`formation`, models.FormationModel{}, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	formation, err := models.RowToStruct(row, reflect.ValueOf(&models.FormationModel{}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userInfo := GetUserInfoWithData(r, &formation)
	if userInfo.Logged {
		query.Filter = make(map[string]string)
		query.Filter[`author`], err = GetTokenID(r)
		query.Filter[`parentid`] = p.ByName("id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userInfo.AlreadyCommented = psql.RowExist(`comment_formation`, models.CommentFormationModel{}, query)
	}
	err = cache.Templates.ExecuteTemplate(w, "formation.html", userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) RenderOrganism(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	query := models.QueryModel{}
	query.Filter = make(map[string]string)
	// TODO Check p.ByName(id) is valid !!!
	query.Filter[`id`] = p.ByName("id")

	// TODO remove test_DB
	row, err := psql.GetRow(`organism`, models.OrganismModel{}, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	organism, err := models.RowToStruct(row, reflect.ValueOf(&models.OrganismModel{}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userInfo := GetUserInfoWithData(r, &organism)
	if userInfo.Logged {
		query.Filter = make(map[string]string)
		query.Filter[`author`], err = GetTokenID(r)
		query.Filter[`parentid`] = p.ByName("id")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userInfo.AlreadyCommented = psql.RowExist(`comment_organism`, models.CommentOrganismModel{}, query)
	}
	err = cache.Templates.ExecuteTemplate(w, "organism.html", userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) RenderDashboard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var userInfoData models.UserInfoWithData
	userInfoData.UserInfo = GetUserInfo(r)
	if userInfoData.Logged == false || len(userInfoData.OrganismOwner) == 0 {
		fmt.Println(`You don't own any organism`)
		http.Error(w, `You don't own any organism`, http.StatusInternalServerError)
		return
	}
	query := models.QueryModel{}
	query.Filter = make(map[string]string)
	query.Filter[`id`] = userInfoData.OrganismOwner

	row, err := psql.GetRow(`organism`, models.OrganismModel{}, query)
	if err != nil {
		fmt.Println(`You don't own any organism`)
		http.Error(w, `You don't own any organism`, http.StatusInternalServerError)
		return
	}
	organismInfo := models.OrganismWithFormationNames{}
	organismInterface, err := models.RowToStruct(row, reflect.ValueOf(&models.OrganismModel{}))
	organismInfo.OrganismModel = organismInterface.(models.OrganismModel)
	if err != nil {
		fmt.Println(`You don't own any organism`)
		http.Error(w, `You don't own any organism`, http.StatusInternalServerError)
		return
	}
	query.Filter = make(map[string]string)
	query.Filter[`parentid`] = userInfoData.OrganismOwner
	rows, err := psql.GetRowsSpec(`formation`, models.FormationModel{}, query, func(string) string {
		return ` id, name `
	})
	if err != nil {
		http.Error(w, `Error retrieving formation data`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		rowStruct, err := models.RowsToStruct(rows, reflect.ValueOf(&models.FormationModel{}))
		if err != nil {
			fmt.Println("Error occured retrieving rowStruct: ", err)
			return
		}
		organismInfo.FormationsInfo = append(organismInfo.FormationsInfo, rowStruct.(models.FormationModel))
	}
	userInfoData.Data = organismInfo
	err = cache.Templates.ExecuteTemplate(w, "dashboard.html", userInfoData)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
