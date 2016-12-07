/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/19 18:43:11 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/22 15:47:27 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"github.com/go-server/dbheroku"
	"github.com/go-server/models"
)

func resetAndCreate() (err error) {

	// collectionsName := []string{"Establishments"}
	// collectionsID, err := apicpa.GetCollectionID(collectionsName, accessToken)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("CollectionsID OK")

	err = dbheroku.Open()
	if err != nil {
		return
	}
	err = dbheroku.Drop("establishment")
	err = dbheroku.Drop("formation")
	err = dbheroku.Drop("comment_establishment")
	err = dbheroku.Drop("comment_formation")
	if err != nil {
		return
	}
	_, err = dbheroku.Exec("../sql/create_establishment.sql")
	if err != nil {
		return
	}
	_, err = dbheroku.Exec("../sql/create_formation.sql")
	if err != nil {
		return
	}
	_, err = dbheroku.Exec("../sql/create_comment_establishment.sql")
	if err != nil {
		return
	}
	_, err = dbheroku.Exec("../sql/create_comment_formation.sql")
	if err != nil {
		return
	}
	csv_filename := "../test/fr-esr-principaux-etablissements-enseignement-superieur.csv"
	err = dbheroku.InsertEstablishmentFromCSV(csv_filename)
	if err != nil {
		return
	}
	err = dbheroku.Index()
	if err != nil {
		return
	}
	var queryOptions models.QueryModel
	queryOptions.Page = make(map[string]string)
	queryOptions.Filter = make(map[string]string)
	results, err := dbheroku.Search(queryOptions, "lyon")
	fmt.Println(results)
	if err != nil {
		return
	}
	return
}

func main() {
	err := resetAndCreate()
	fmt.Println(err)
}
