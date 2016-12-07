/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   login.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/13 13:21:37 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/03 17:23:39 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-server/cache"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"github.com/go-server/psql"
	"github.com/julienschmidt/httprouter"
	"github.com/zemirco/uid"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"reflect"
)

type GGLResponse struct {
	Iss            string `json:"iss"`
	Iat            string `json:"iat"`
	Exp            string `json:"exp"`
	At_hash        string `json:"at_hash"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Email_verified string `json:"email_verified"`
	Azp            string `json:"azp"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Given_name     string `json:"given_name"`
	Family_name    string `json:"family_name"`
	Locale         string `json:"locale"`
	Alg            string `json:"alg"`
	Kid            string `json:"kid"`
}

func IsLog(r *http.Request) bool {
	session, err := cache.Store.Get(r, conf.COOKIE_AUTH)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// Retrieve our struct and type-assert it
	val := session.Values[conf.COOKIE_AUTH]
	tokenString, ok := val.(string)
	if ok == false {
		fmt.Println("This is not even a cookie !")
		return false
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.SESSION_SIGN_KEY), nil
	})
	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")
		maxAgeStr, ok := token.Claims.(jwt.MapClaims)["max_age"]
		if ok == false || maxAgeStr == nil {
			return false
		}
		maxage := maxAgeStr.(float64)
		if maxage < 0 {
			return false
		}
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return false
}

func (ctrl *Controller) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session, err := cache.Store.Get(r, conf.COOKIE_AUTH)
	if err != nil {
		http.Redirect(w, r, `/`, http.StatusFound)
		return
	}
	session.Values[conf.COOKIE_AUTH] = ``
	session.Save(r, w)
	http.Redirect(w, r, `/`, http.StatusFound)
}

func (ctrl *Controller) TokenSign(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error occured parsing post: ", err)
	}
	idtoken := r.Form.Get(`idtoken`)
	ggl_url_check_token := `https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=` + idtoken
	req, err := http.NewRequest(`GET`, ggl_url_check_token, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		fmt.Println("Fail identification:", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Service error response is unidentified")
		return
	}
	defer resp.Body.Close()
	gglResp := GGLResponse{}
	err = json.Unmarshal(body, &gglResp)
	if err != nil {
		fmt.Println("Service error response is unidentified")
		return
	}
	if gglResp.Aud != conf.GGL_CLIENT_ID {
		fmt.Println("Service error response is unidentified")
		return
	}
	fmt.Println("Sign in with google sucessful")
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"sub": gglResp.Sub,
		// "expires_in": time.Now().Add(time.Second * time.Duration(gglResp.Exp)).Unix(),
	}
	tokenString, err := token.SignedString([]byte(conf.SESSION_SIGN_KEY))
	if err != nil {
		fmt.Println("Error creating auth token:", err)
		http.Redirect(w, r, "/", http.StatusFound)
	}
	session, err := cache.Store.Get(r, conf.COOKIE_AUTH)
	session.Values[conf.COOKIE_AUTH] = tokenString
	session.Save(r, w)
}

func GetTokenID(r *http.Request) (id string, err error) {
	session, err := cache.Store.Get(r, conf.COOKIE_AUTH)
	if err != nil {
		return
	}
	val := session.Values[conf.COOKIE_AUTH]
	tokenString, ok := val.(string)
	if ok == false {
		fmt.Println("This is not even a cookie !")
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.SESSION_SIGN_KEY), nil
	})
	if err != nil || !token.Valid {
		return
	}
	id, ok = token.Claims.(jwt.MapClaims)["sub"].(string)
	if ok == false {
		return "", errors.New("cant retrieve id token")
	}
	maxage, ok := token.Claims.(jwt.MapClaims)["max_age"]
	if ok == false || maxage == nil {
		return "", errors.New("Cant retrieve exp date")
	}
	maxageFloat64 := maxage.(float64)
	if maxageFloat64 < 0 {
		fmt.Println(`Token expired !`, err)
		return "", errors.New("Token has expired !")
	}
	return id, nil
}

func GetUserInfo(r *http.Request) (result models.UserInfo) {
	id, err := GetTokenID(r)
	if err != nil || len(id) == 0 {
		fmt.Println(`User is not logged:`, err)
		result.Logged = false
		return
	}
	result.Logged = true
	query := models.QueryModel{}
	query.Filter = make(map[string]string)
	query.Filter[`id`] = id
	type UserName struct {
		Email         string `json:"email" db:"email"`
		OrganismOwner string `json:"organismowner" db:"organismowner"`
	}
	// Save this in local, dont query every fuckin time
	row, err := psql.GetRow(`user_account`, UserName{}, query)
	if err != nil {
		fmt.Println(err)
		return
	}
	name, err := models.RowToStruct(row, reflect.ValueOf(&UserName{}))
	if err != nil {
		fmt.Println(err)
		return
	}
	result.Name = name.(UserName).Email
	result.OrganismOwner = name.(UserName).OrganismOwner
	return
}

func GetUserInfoWithData(r *http.Request, data *interface{}) (result models.UserInfoWithData) {
	result.UserInfo = GetUserInfo(r)
	result.Data = data
	return
}

func SendConfirmationMail(dest string, userID string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", `validation@notaforma.fr`, conf.VALIDATION_PWD, `mail.gandi.net`)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	key := uid.New(11)
	cache.KeyValidation[key] = userID
	mailInfo := models.MailInfo{
		To:      []string{dest},
		From:    `validation@notaforma.fr`,
		Subject: `Confirmation du mail d'inscription à Notaforma`,
		Body:    `Merci de suivre le lien suivant .WIP. pour confirmer votre adresse email`,
		Key:     key,
	}
	var b bytes.Buffer
	err := cache.TextTemplates.ExecuteTemplate(&b, "validation.html", mailInfo)
	if err != nil {
		fmt.Println(`Error templating validation email:`, err.Error())
	}
	err = smtp.SendMail(`mail.gandi.net:25`, auth, `validation@notaforma.fr`, mailInfo.To, b.Bytes())
	if err != nil {
		fmt.Println(`Error sending email:`, err.Error())
	}
}

func (ctrl *Controller) ValidateEmail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, `/`, http.StatusFound)
		return
	}
	keyToCheck := r.FormValue(`key`)
	id, exist := cache.KeyValidation[keyToCheck]
	if len(id) == 0 || !exist {
		http.Redirect(w, r, `/`, http.StatusFound)
		return
	}
	delete(cache.KeyValidation, keyToCheck)
	queryOptions := models.QueryModel{}
	queryOptions.Filter = make(map[string]string)
	queryOptions.Filter[`id`] = id
	toUpdate := map[string]string{
		`validate`: `true`,
	}
	psql.UpdateRow(`user_account`, queryOptions, toUpdate)
}
