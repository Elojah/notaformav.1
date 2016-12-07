/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   post_data.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/13 13:19:38 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 19:38:04 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
	"encoding/json"
	"errors"
	// "github.com/go-server/apicpa"
	// "github.com/go-server/apifc"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-server/cache"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"github.com/go-server/psql"
	"github.com/julienschmidt/httprouter"
	"github.com/zemirco/uid"
	"time"
	// "github.com/mitchellh/mapstructure"
	"net/http"
)

func marshalledErr(message error) (result []byte) {
	jsonErr := models.JSONError{
		Message: func(check error) string {
			if check != nil {
				return check.Error()
			}
			return `Unknown`
		}(message),
	}
	result, err := json.Marshal(jsonErr)
	if err != nil {
		fmt.Print(`Error occured marshaling an error`)
	}
	return
}

func marshalledSuccess() (result []byte) {
	jsonErr := models.JSONError{
		Message: `Success`,
	}
	result, err := json.Marshal(jsonErr)
	if err != nil {
		fmt.Print(`Error occured marshaling an error`)
	}
	return
}

// POST requests
func (ctrl *Controller) PostCommentFormation(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	formationId := p.ByName("id")
	err := r.ParseForm()
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	authorid, err := GetTokenID(r)
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	commentInterface := models.FormToStruct(&models.CommentFormationModel{}, &r.Form)
	comment := commentInterface.(models.CommentFormationModel)
	comment.Id = uid.New(11)
	comment.Author = authorid
	comment.ParentID = formationId
	comment.UpVote = 0
	comment.DownVote = 0
	comment.PostDate = time.Now()
	if err != nil {
		fmt.Print(`This form is corrupted`)
		w.Write(marshalledErr(err))
		return
	}
	unique := psql.InsertRowUnique(`comment_formation`, comment, []string{`Author`, `ParentID`})
	if unique == false {
		w.Write(marshalledErr(errors.New(`You already commented this page`)))
		return
	}
	psql.UpdateAverage(`formation`, `globalrating`, comment.ParentID)
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = comment.ParentID
	toUpdate := map[string]string{
		`ncomments`: `ncomments + 1`,
	}
	ok := psql.UpdateRowNoQuote(`formation`, queryOptions, toUpdate)
	if ok == false {
		w.Write(marshalledErr(errors.New(`Error while adding your vote`)))
	}
	w.Write(marshalledSuccess())
}

func (ctrl *Controller) PostCommentOrganism(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	organismId := p.ByName("id")
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error occured parsing post: ", err)
		return
	}
	authorid, err := GetTokenID(r)
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	commentInterface := models.FormToStruct(&models.CommentOrganismModel{}, &r.Form)
	comment := commentInterface.(models.CommentOrganismModel)
	comment.Id = uid.New(11)
	comment.Author = authorid
	comment.ParentID = organismId
	comment.UpVote = 0
	comment.DownVote = 0
	comment.PostDate = time.Now()
	if err != nil {
		fmt.Print(`This form is corrupted`)
		w.Write(marshalledErr(err))
		return
	}
	unique := psql.InsertRowUnique(`comment_organism`, comment, []string{`Author`, `ParentID`})
	if unique == false {
		w.Write(marshalledErr(errors.New(`You already commented this page`)))
		return
	}
	psql.UpdateAverage(`organism`, `globalrating`, comment.ParentID)
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = comment.ParentID
	toUpdate := map[string]string{
		`ncomments`: `ncomments + 1`,
	}
	ok := psql.UpdateRowNoQuote(`organism`, queryOptions, toUpdate)
	if ok == false {
		w.Write(marshalledErr(errors.New(`Error while adding your vote`)))
	}
	w.Write(marshalledSuccess())
}

func (ctrl *Controller) VoteCommentFormation(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	formationId := p.ByName("id")
	commentId := p.ByName("commentid")
	certVote := func(s string) bool {
		fmt.Println(`GOT: `, s)
		if s == `1` {
			return true
		} else {
			return false
		}
	}(r.Form.Get(`vote`))
	authorid, err := GetTokenID(r)
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	userVote := models.CommentFormationVoteUser{
		FormationVoteUser: models.FormationVoteUser{
			UserID:      authorid,
			FormationID: formationId,
		},
		CommentVotes: models.CommentVotes{
			CommentID: commentId,
		},
		Vote: certVote,
	}
	unique := psql.InsertRowUnique(`comment_formation_vote_user`, userVote, []string{`UserID`, `FormationID`, `CommentID`})
	if unique == false {
		w.Write(marshalledErr(errors.New(`You already voted for this comment`)))
		return
	}
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = userVote.CommentID
	queryOptions.Filter[`parentid`] = userVote.FormationID
	voteStr := func(b bool) string {
		if b {
			return `upvote`
		} else {
			return `downvote`
		}
	}(userVote.Vote)
	toUpdate := map[string]string{
		voteStr: voteStr + ` + 1`,
	}
	ok := psql.UpdateRowNoQuote(`comment_formation`, queryOptions, toUpdate)
	if ok == false {
		w.Write(marshalledErr(errors.New(`Error while adding your vote`)))
	}
	w.Write(marshalledSuccess())
}

func (ctrl *Controller) VoteCommentOrganism(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	organismId := p.ByName("id")
	commentId := p.ByName("commentid")
	certVote := func(s string) bool {
		if s == `1` {
			return true
		} else {
			return false
		}
	}(r.Form.Get(`vote`))
	authorid, err := GetTokenID(r)
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	userVote := models.CommentOrganismVoteUser{
		OrganismVoteUser: models.OrganismVoteUser{
			UserID:     authorid,
			OrganismID: organismId,
		},
		CommentVotes: models.CommentVotes{
			CommentID: commentId,
		},
		Vote: certVote,
	}
	unique := psql.InsertRowUnique(`comment_organism_vote_user`, userVote, []string{`UserID`, `OrganismID`, `CommentID`})
	if unique == false {
		w.Write(marshalledErr(errors.New(`You already voted for this comment`)))
		return
	}
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = userVote.CommentID
	queryOptions.Filter[`parentid`] = userVote.OrganismID
	voteStr := func(b bool) string {
		if b {
			return `upvote`
		} else {
			return `downvote`
		}
	}(userVote.Vote)
	toUpdate := map[string]string{
		voteStr: voteStr + ` + 1`,
	}
	ok := psql.UpdateRowNoQuote(`comment_organism`, queryOptions, toUpdate)
	if ok == false {
		w.Write(marshalledErr(errors.New(`Error while adding your vote`)))
	}
	w.Write(marshalledSuccess())
}

/*
	CHECK EMAILS AND PASSWORD BACKEND FOR LOGIN AND SUBSCRIBE
	TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
*/

func (ctrl *Controller) LoginUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error occured parsing post: ", err)
		http.Redirect(w, r, `/login`, http.StatusForbidden)
		return
	}
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`email`] = r.Form.Get(`email`)
	queryOptions.Filter[`password`] = r.Form.Get(`password`)
	row, err := psql.GetRow(`user_account`, models.User{}, queryOptions)
	if err != nil || row == nil {
		fmt.Println(`Wrong email/password`)
		ctrl.RenderLogin(w, r, nil)
		return
	}
	user := models.User{}
	err = row.StructScan(&user)
	if err != nil {
		fmt.Println(`Account is corrupted:`, err)
		ctrl.RenderLogin(w, r, nil)
		return
	}
	if !user.Validate {
		fmt.Println(`You must validate your account`)
		ctrl.RenderLogin(w, r, nil)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"sub":     user.ID,
		"max_age": time.Now().Add(time.Second * time.Duration(3600*6)).Unix(),
	}
	tokenString, err := token.SignedString([]byte(conf.SESSION_SIGN_KEY))
	if err != nil {
		fmt.Println("Error creating auth token:", err)
		ctrl.RenderLogin(w, r, nil)
		return
	}
	session, err := cache.Store.Get(r, conf.COOKIE_AUTH)
	session.Values[conf.COOKIE_AUTH] = tokenString
	session.Save(r, w)
	http.Redirect(w, r, `/`, http.StatusFound)
}

func (ctrl *Controller) SubscribeUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error occured parsing post: ", err)
		ctrl.RenderSubscribe(w, r, nil)
		return
	}
	user := models.User{
		ID:       uid.New(11),
		Email:    r.Form.Get(`email`),
		Password: r.Form.Get(`password`),
		Validate: false,
	}
	err = psql.InsertRow(`user_account`, user)
	if err != nil {
		fmt.Println("Error occured creating user: ", err)
		ctrl.RenderSubscribe(w, r, nil)
		return
	}
	// Send email
	SendConfirmationMail(user.Email, user.ID)
	r.Method = "GET"
	http.Redirect(w, r, `/`, http.StatusFound)
}

func updateOrganism(field string, val []string, w http.ResponseWriter, organismId string) {
	if len(val) != 1 {
		w.Write(marshalledErr(errors.New(`You can't assign multiple values to this field`)))
		return
	}
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = organismId
	toUpdate := map[string]string{
		field: val[0],
	}
	psql.UpdateRow(`organism`, queryOptions, toUpdate)
	w.Write(marshalledSuccess())
	return
}

func (ctrl *Controller) EditDashboardInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var userInfoData models.UserInfoWithData
	userInfoData.UserInfo = GetUserInfo(r)
	if userInfoData.Logged == false || len(userInfoData.OrganismOwner) == 0 {
		w.Write(marshalledErr(errors.New(`Error while editing field, you're not connected anymore`)))
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	if len(r.Form) > 1 {
		w.Write(marshalledErr(errors.New(`You can't edit multiple values at the same time`)))
		return
	}
	for key, val := range r.Form {
		switch key {
		case `name`, `corporatename`, `siret`:
			updateOrganism(key, val, w, userInfoData.UserInfo.OrganismOwner)
			return
		case `street`, `zip`, `locality`:
			updateOrganism(`contact_adress_`+key, val, w, userInfoData.UserInfo.OrganismOwner)
			return
		case `tel`, `fax`, `website`:
			updateOrganism(`contact_`+key, val, w, userInfoData.UserInfo.OrganismOwner)
			return
		default:
			w.Write(marshalledErr(errors.New(`Unrecognized field name`)))
			return
		}
	}
	w.Write(marshalledErr(errors.New(`Data is corrupted`)))
}

func updateFormation(field string, val []string, w http.ResponseWriter, formationId string) {
	if len(val) != 1 {
		w.Write(marshalledErr(errors.New(`You can't assign multiple values to this field`)))
		return
	}
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = formationId
	toUpdate := map[string]string{
		field: val[0],
	}
	psql.UpdateRow(`formation`, queryOptions, toUpdate)
	w.Write(marshalledSuccess())
	return
}

func updateFormationArray(field string, val []string, w http.ResponseWriter, formationId string) {
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = formationId
	toUpdate := map[string]string{
		field: `ARRAY[` + models.SerializeStringArraySQL(val) + `]`,
	}
	psql.UpdateRow(`formation`, queryOptions, toUpdate)
	w.Write(marshalledSuccess())
	return
}

func (ctrl *Controller) EditDashboardFormationInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var userInfoData models.UserInfoWithData
	userInfoData.UserInfo = GetUserInfo(r)
	if userInfoData.Logged == false || len(userInfoData.OrganismOwner) == 0 {
		w.Write(marshalledErr(errors.New(`Error while editing field, you're not connected anymore`)))
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.Write(marshalledErr(err))
		return
	}
	if len(r.Form) > 1 {
		w.Write(marshalledErr(errors.New(`You can't edit multiple values at the same time`)))
		return
	}
	// fmt.Println(r.Form)
	for key, val := range r.Form {
		if len(val) < 2 || len(key) < 3 {
			w.Write(marshalledErr(errors.New(`Couldn't retrieve formation id`)))
			return
		}
		formationId := val[0]
		val = val[1:]
		key = key[:len(key)-2]
		switch key {
		case `name`, `objectives`:
			updateFormation(key, val, w, formationId)
			return
		case `street`, `zip`, `locality`:
			updateFormation(`contact_adress_`+key, val, w, formationId)
			return
		case `programme`:
			updateFormationArray(key, val, w, formationId)
			return
		case `tel`, `mail`, `fax`, `website`, `contact-name`:
			if key == `contact-name` {
				key = `name`
			}
			updateFormation(`contact_`+key, val, w, formationId)
			return
		default:
			w.Write(marshalledErr(errors.New(`Unrecognized field name`)))
			return
		}
	}
	w.Write(marshalledErr(errors.New(`Data is corrupted`)))
}
