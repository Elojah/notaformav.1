/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/11/18 16:13:49 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/03 17:13:07 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	"fmt"
	"github.com/go-server/models"
	"github.com/go-server/psql"
	"os"
	"strings"
)

func launchCmd(cmd string) {
	var err error
	switch cmd {
	case `comment_formation`:
		psql.Drop(`comment_formation`)
		err = psql.CreateTable(models.CommentFormationModel{}, `comment_formation`)
		break
	case `comment_organism`:
		psql.Drop(`comment_organism`)
		err = psql.CreateTable(models.CommentOrganismModel{}, `comment_organism`)
		break
	case `organism_vote_user`:
		psql.Drop(`organism_vote_user`)
		err = psql.CreateTable(models.OrganismVoteUser{}, `organism_vote_user`)
		break
	case `comment_organism_vote_user`:
		psql.Drop(`comment_organism_vote_user`)
		err = psql.CreateTable(models.CommentOrganismVoteUser{}, `comment_organism_vote_user`)
		break
	case `formation_vote_user`:
		psql.Drop(`formation_vote_user`)
		err = psql.CreateTable(models.FormationVoteUser{}, `formation_vote_user`)
		break
	case `comment_formation_vote_user`:
		psql.Drop(`comment_formation_vote_user`)
		err = psql.CreateTable(models.CommentFormationVoteUser{}, `comment_formation_vote_user`)
		break
	case `user_account`:
		psql.Drop(`user_account`)
		err = psql.CreateTable(models.User{}, `user_account`)
		break
	default:
		fmt.Println(`Unrecognized table`)
	}
	if err != nil {
		fmt.Println(`Error occured creating table: ` + err.Error())
	} else {
		fmt.Println(`OK`)
	}
}

func writeLoop() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("$>")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n\t ")
		if text == `exit` {
			panic("Bye")
		}
		launchCmd(text)
	}
}

func main() {
	err := psql.Open()
	if err != nil {
		fmt.Println(`Error opening table: ` + err.Error())
		return
	}
	writeLoop()
}
