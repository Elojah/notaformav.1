/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/14 14:44:53 by hdezier           #+#    #+#             */
/*   Updated: 2016/10/19 19:25:43 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"github.com/go-server/apicpa"
	"github.com/go-server/models"
)

func Do(filename string) (err error) {
	AccessToken, err := apicpa.Authenticate()
	if err != nil {
		return
	}
	fmt.Println("AccessToken OK: ", AccessToken)
	// Services ID
	servicesName := []string{"FormAdvisor"}
	ServicesID, err := apicpa.GetServiceID(servicesName, AccessToken)
	if err != nil {
		return
	}
	fmt.Println("ServicesID OK: ")
	for key, val := range ServicesID {
		fmt.Println(key, " -> ", val)
	}
	// Collections ID
	collectionsName := []string{"Establishments", "Formations", "Comments_Formation", "Comments_Establishment"}
	CollectionsID, err := apicpa.GetCollectionID(collectionsName, AccessToken)
	if err != nil {
		return
	}
	fmt.Println("CollectionsID OK")
	for key, val := range CollectionsID {
		fmt.Println(key, " -> ", val)
	}
	// Fill DB
	err = apicpa.PostCollectionFromCSV(CollectionsID["Establishments"],
		AccessToken,
		filename,
		&models.EstablishmentModel{})
	if err != nil {
		fmt.Println(err)
		return
	}
	// Collection data ID
	collectionsDataID, err := apicpa.GetCollectionDataID(CollectionsID["Establishments"], AccessToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(collectionsDataID)
	return
}

func main() {
	csv_filename := "test/fr-esr-principaux-etablissements-enseignement-superieur.csv"
	Do(csv_filename)
}
