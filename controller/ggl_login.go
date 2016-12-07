/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ggl_login.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/17 14:23:24 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/17 14:25:18 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
// "golang.org/x/oauth2"
// "golang.org/x/oauth2/google"
)

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

type AuthToken struct {
	State string `json:"state"`
}

// func (ctrl *Controller) AuthHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	// state = randToken()
// 	// session, err := cache.Store.Get(r, conf.FC_COOKIE_NAME)
// 	// session.Values["state"] = state
// 	// session.Save(r, w)
// 	// session.Save()
// 	// c.Writer.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
// }
