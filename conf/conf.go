/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   conf.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/12 19:12:30 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 21:48:50 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package conf

const (
	// Public
	CPA_API_URI                = "https://api-cpa.herokuapp.com/api/v1/"
	CPA_AUTH_URL               = "oauth/token/"
	CPA_SERVICE_URL            = "services/"
	CPA_COLLECTION_URL         = "collections/"
	CPA_COLLECTION_DATA_URL    = "relationships/donnees/"
	CPA_COLLECTION_SERVICE_URL = "relationships/collections/"
	POST_ACCESS_TOKEN          = "?access_token="
	FC_URL                     = "https://fcp.integ01.dev-franceconnect.fr/api/v1/"
	FC_COOKIE_NAME             = "fc_session"
	COOKIE_AUTH                = "auth"
	MY_HOME_URL                = "http://localhost:4242/"
	// Private
	DATABASE_URL       = ""
	ID_PUBLIC_SERVICE  = ""
	ID_PRIVATE_SERVICE = ""
	FC_KEY             = ""
	FC_SECRET          = ""
	SESSION_SIGN_KEY   = ""
	SESSION_STORE_KEY  = ""
	FC_FS_STATE        = ""
	FC_FS_NONCE        = ""
	GGL_CLIENT_ID      = ""
	GGL_SECRET_CODE    = ""
	VALIDATION_PWD     = ""
)
