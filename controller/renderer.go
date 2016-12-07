/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   renderer.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/13 13:15:59 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/03 17:27:35 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package controller

import (
	"github.com/go-server/cache"
	// "github.com/go-server/models"
	// "github.com/go-server/psql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// HTML rendering

func (ctrl *Controller) RenderHome(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "home.html", GetUserInfo(r))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func (ctrl *Controller) RenderDashboard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	err := cache.Templates.ExecuteTemplate(w, "dashboard.html", GetUserInfoWithData(r))
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func (ctrl *Controller) RenderFormations(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "all-formation.html", GetUserInfo(r))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (ctrl *Controller) RenderOrganisms(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "all-organisms.html", GetUserInfo(r))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) RenderLogin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) RenderSubscribe(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "subscribe.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ctrl *Controller) RenderCriteria(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := cache.Templates.ExecuteTemplate(w, "criteria.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
