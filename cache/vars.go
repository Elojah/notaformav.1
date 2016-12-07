/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   vars.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/15 18:01:32 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/27 16:12:58 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package cache

import (
	"github.com/go-server/conf"
	"github.com/gorilla/sessions"
	"html/template"
	"os"
	"path/filepath"
	ttpl "text/template"
)

// Template caching
var (
	Cwd, _       = os.Getwd() // Useless ftm
	TemplatePath = "./templates/"
)

var Templates = template.Must(template.ParseFiles(
	filepath.Join(TemplatePath, "common/header.html"),
	filepath.Join(TemplatePath, "common/footer.html"),
	filepath.Join(TemplatePath, "common/scripts.html"),
	filepath.Join(TemplatePath, "home.html"),
	filepath.Join(TemplatePath, "dashboard.html"),
	filepath.Join(TemplatePath, "criteria.html"),
	filepath.Join(TemplatePath, "login.html"),
	filepath.Join(TemplatePath, "subscribe.html"),
	filepath.Join(TemplatePath, "search.html"),
	filepath.Join(TemplatePath, "new-comment-formation.html"),
	filepath.Join(TemplatePath, "new-comment-organism.html"),
	filepath.Join(TemplatePath, "comments.html"),
	filepath.Join(TemplatePath, "formation.html"),
	filepath.Join(TemplatePath, "validation.html"),
	filepath.Join(TemplatePath, "organism.html"),
	filepath.Join(TemplatePath, "all-formation.html"),
	filepath.Join(TemplatePath, "formation/new-formation.html"),
	filepath.Join(TemplatePath, "all-organisms.html"),
	filepath.Join(TemplatePath, "list-organisms.html"),
	filepath.Join(TemplatePath, "list-formations.html"),
	filepath.Join(TemplatePath, "establishment/new-establishment.html"),
))

var TextTemplates = ttpl.Must(ttpl.ParseFiles(
	filepath.Join(TemplatePath, "validation.html"),
))

var Store = sessions.NewCookieStore([]byte(conf.SESSION_STORE_KEY))

var KeyValidation = map[string]string{}
