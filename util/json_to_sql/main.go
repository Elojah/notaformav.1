/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/10 16:33:52 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 14:05:22 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-server/models"
	"github.com/go-server/psql"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile(`../scrap_my_intercarif/log_organism_idf.json`)
	if err != nil {
		fmt.Println(`Error occured opening file: ` + err.Error())
		return
	}
	var bigJsonArrayOrganism []models.OrganismModel
	err = json.Unmarshal(file, &bigJsonArrayOrganism)
	if err != nil {
		fmt.Println(`Error occured unmarshaling organism data: ` + err.Error())
		return
	}
	file, err = ioutil.ReadFile(`../scrap_my_intercarif/log_formation_idf.json`)
	if err != nil {
		fmt.Println(`Error occured opening file: ` + err.Error())
		return
	}
	var bigJsonArrayFormation []models.FormationModel
	err = json.Unmarshal(file, &bigJsonArrayFormation)
	if err != nil {
		fmt.Println(`Error occured unmarshaling formation data: ` + err.Error())
		return
	}
	psql.Open()
	psql.Drop(`organism`)
	err = psql.CreateTable(models.OrganismModel{}, `organism`)
	if err != nil {
		fmt.Println(`Error occured creating organism table: ` + err.Error())
		return
	}
	psql.Drop(`formation`)
	err = psql.CreateTable(models.FormationModel{}, `formation`)
	if err != nil {
		fmt.Println(`Error occured creating formation table: ` + err.Error())
		return
	}
	for _, val := range bigJsonArrayOrganism {
		err = psql.InsertRow(`organism`, val)
		if err != nil {
			fmt.Println(`Error occured inserting organism data: ` + err.Error())
			return
		}
	}
	for _, val := range bigJsonArrayFormation {
		err = psql.InsertRow(`formation`, val)
		if err != nil {
			fmt.Println(`Error occured inserting formation data: ` + err.Error())
			return
		}
	}
	err = psql.Index(`organism`, models.OrganismModel{})
	if err != nil {
		fmt.Println(`Error occured indexing organism data: ` + err.Error())
		return
	}
	err = psql.Index(`formation`, models.FormationModel{})
	if err != nil {
		fmt.Println(`Error occured indexing formation data: ` + err.Error())
		return
	}
}
