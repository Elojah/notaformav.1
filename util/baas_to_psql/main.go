/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/14 15:06:57 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/17 16:30:47 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"github.com/go-server/apicpa"
	"github.com/go-server/dbheroku"
	"github.com/go-server/models"
)

func resetDB() (err error) {

	accessToken, err := apicpa.Authenticate()
	if err != nil {
		return
	}
	fmt.Println("AccessToken OK: ", accessToken)

	// collectionsName := []string{"Establishments"}
	// collectionsID, err := apicpa.GetCollectionID(collectionsName, accessToken)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("CollectionsID OK")

	err = dbheroku.Open()
	// if err != nil {
	// 	return
	// }
	// _, err = dbheroku.Exec("sql/drop_establishments.sql")
	// if err != nil {
	// 	return
	// }
	// _, err = dbheroku.Exec("sql/create_establishments.sql")
	// if err != nil {
	// 	return
	// }
	// err = dbheroku.InsertFromAPI(collectionsID["Establishments"], accessToken)
	// if err != nil {
	// 	return
	// }
	err = dbheroku.Index()
	if err != nil {
		return
	}
	var queryOptions models.QueryModel
	queryOptions.Page = make(map[string]string)
	queryOptions.Filter = make(map[string]string)
	results, n, err := dbheroku.Search(queryOptions, "lyon")
	fmt.Println(n, " results:", results)
	if err != nil {
		return
	}
	return
}

func main() {
	err := resetDB()
	fmt.Println(err)
}
