/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/26 17:45:10 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/18 15:55:00 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package apifc

/*import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-server/cache"
	"github.com/go-server/conf"
	"github.com/go-server/models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Code by Marion
// So be nice plz...

func RegisterTypes() {
	gob.Register(&models.FCToken{})
	// Nothing ftm
}
*/
/*
func GetTokenID(r *http.Request) (id string, err error) {
	session, err := cache.Store.Get(r, conf.FC_COOKIE_NAME)
	if err != nil {
		return
	}
	val := session.Values[conf.FC_COOKIE_NAME]
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
	id, ok = token.Claims.(jwt.MapClaims)["id_token"].(string)
	if ok == false {
		return "", errors.New("cant retrieve id token")
	}
	idSeparator := strings.LastIndex(id, ".")
	id = id[:idSeparator]
	return id, nil
}

func IsLog(r *http.Request) bool {
	session, err := cache.Store.Get(r, conf.FC_COOKIE_NAME)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	// Retrieve our struct and type-assert it
	val := session.Values[conf.FC_COOKIE_NAME]
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

func Authenticate(w http.ResponseWriter, r *http.Request) {

	redir := conf.FC_URL + "authorize?"
	redir += "response_type=code"
	redir += "&client_id=" + conf.FC_KEY
	redir += "&redirect_uri=" + conf.MY_HOME_URL + "fc/login"
	redir += "&scope=openid%20birth%20profile"
	redir += "&state=" + conf.FC_FS_STATE
	redir += "&nonce=" + conf.FC_FS_NONCE
	http.Redirect(w, r, redir, 302)
}

func Callback(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Callback FranceConnect...")
	session, err := cache.Store.Get(r, conf.FC_COOKIE_NAME)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
	}
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	session.Values["code"] = code
	session.Values["state"] = state
	session.Save(r, w)

	url := conf.FC_URL + "token"
	postData := models.FCPostCode{
		Grant_type:    "authorization_code",
		Redirect_uri:  conf.MY_HOME_URL + "fc/login",
		Client_id:     conf.FC_KEY,
		Client_secret: conf.FC_SECRET,
		Code:          code,
	}
	b, err := json.Marshal(postData)
	resp, err := http.Post(url, "application/json;charset=utf-8", bytes.NewReader(b))

	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err, resp.StatusCode)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var tokenResp models.FCToken
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		fmt.Println("Callback FranceConnect Error")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"access_token": tokenResp.Access_token,
		"id_token":     tokenResp.Id_token,
		"expires_in":   time.Now().Add(time.Second * time.Duration(tokenResp.Expires_in)).Unix(),
	}
	tokenString, err := token.SignedString([]byte(conf.SESSION_SIGN_KEY))
	if err != nil {
		fmt.Println("Error creating session:", err)
		http.Redirect(w, r, "/", http.StatusFound)
	}

	session.Values[conf.FC_COOKIE_NAME] = tokenString

	session.Save(r, w)
	spew.Dump(tokenResp)
	fmt.Println("Callback FranceConnect Success")
	http.Redirect(w, r, "/", http.StatusFound)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenId, err := GetTokenID(r)
	if err != nil {
		fmt.Println("Failed to hint ID")
	}

	redirUrl := conf.FC_URL + "logout"
	redirUrl += `?id_token_hint=` + tokenId
	redirUrl += `&state=` + conf.FC_FS_STATE
	redirUrl += `&post_logout_redirect_uri=` + conf.MY_HOME_URL + `fc/logout`
	http.Redirect(w, r, redirUrl, 302)
}
*/
