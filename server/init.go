/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/11 20:53:48 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 18:55:21 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package server

import (
	"fmt"
	"github.com/go-server/cache"
	c "github.com/go-server/controller"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func linkHandler(r *httprouter.Router, ctrl *c.Controller) {
	// Login
	r.GET("/login", ctrl.RenderLogin)
	r.GET("/logout", ctrl.Logout)
	r.GET("/subscribe", ctrl.RenderSubscribe)
	r.GET("/valid", ctrl.ValidateEmail)
	// r.GET("/fc/login", ctrl.CallbackFC)
	// r.GET("/logout", ctrl.LogoutFC)
	// r.GET("/fc/logout", ctrl.RenderHome)

	// HTML rendering
	r.GET("/", ctrl.RenderHome)
	r.GET("/search", ctrl.RenderSearch)
	r.GET("/formation/:id", ctrl.RenderFormation)
	r.GET("/organism/:id", ctrl.RenderOrganism)
	r.GET("/organisms", ctrl.RenderOrganisms)
	r.GET("/formations", ctrl.RenderFormations)
	r.GET("/dashboard", ctrl.RenderDashboard)
	r.GET("/criteres", ctrl.RenderCriteria)
	// r.GET("/formations/new", ctrl.RenderNewFormation)
	// r.GET("/establishments/new", ctrl.RenderNewEstablishment)

	// API
	r.GET("/api/organism/:id/comments", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, err := c.GetTokenID(r)
		if err != nil {
			ctrl.CommentOrganismRequest.WriteAsJSON(w, r, p)
		} else {
			ctrl.CommentOrganismRequestWithVote(id).WriteAsJSON(w, r, p)
		}
	})
	r.GET("/api/formation/:id/comments", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id, err := c.GetTokenID(r)
		if err != nil {
			ctrl.CommentFormationRequest.WriteAsJSON(w, r, p)
		} else {
			ctrl.CommentFormationRequestWithVote(id).WriteAsJSON(w, r, p)
		}
	})
	r.GET("/api/formation", ctrl.FormationsRequest.WriteAsJSON)
	r.GET("/api/organism", ctrl.OrganismsRequest.WriteAsJSON)
	r.GET("/api/search/formation", ctrl.SearchFormationMinRequest.WriteAsJSON)
	r.GET("/api/search/organism", ctrl.SearchOrganismMinRequest.WriteAsJSON)
	r.GET("/api/search/formation/match", ctrl.SearchFormationMatchRequest.WriteAsJSON)
	r.GET("/api/formations/all", ctrl.FormationsMinRequest.WriteAsJSON)
	r.GET("/api/organisms/all", ctrl.OrganismsMinRequest.WriteAsJSON)

	// // Data post
	r.POST("/organism/:id/comments", ctrl.PostCommentOrganism)
	r.POST("/formation/:id/comments", ctrl.PostCommentFormation)
	r.POST("/tokensignin", ctrl.TokenSign)
	r.POST("/login", ctrl.LoginUser)
	r.POST("/subscribe", ctrl.SubscribeUser)
	r.POST("/organism/:id/comment/:commentid", ctrl.VoteCommentOrganism)
	r.POST("/formation/:id/comment/:commentid", ctrl.VoteCommentFormation)
	r.POST("/dashboard/edit", ctrl.EditDashboardInfo)
	r.POST("/dashboard/formation/edit", ctrl.EditDashboardFormationInfo)
	// r.POST("/establishments/new", ctrl.PostEstablishment)
	// r.POST("/formations/new", ctrl.PostFormation)
}

func launchServer(r *httprouter.Router) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4242"
	}
	fmt.Println("Server listening on port: " + port)
	http.ListenAndServe(":"+port, context.ClearHandler(r))
}

func Init() {
	r := httprouter.New()
	var ctrl c.Controller
	err := ctrl.Init()
	if err != nil {
		panic(err)
	}
	err = ctrl.InitDB()
	// apifc.RegisterTypes() // Probably keep function but move in controller package
	if err != nil {
		panic(err)
	}
	linkHandler(r, &ctrl)
	r.ServeFiles("/static/*filepath", http.Dir("static"))
	cache.Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	launchServer(r)
}
