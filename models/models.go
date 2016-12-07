/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   models.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/14 22:55:25 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 16:17:48 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package models

import (
	"time"
)

type CPAPostModel struct {
	Access_token string   `json:"access_token"`
	Fc_token     string   `json:"fc_token"`
	Data         JSONData `json:"data"`
}

type Converter func(string) string

// JSON API
/*
id: a unique identifier for this particular occurrence of the problem.
links: a links object containing the following members:
    about: a link that leads to further details about this particular occurrence of the problem.
status: the HTTP status code applicable to this problem, expressed as a string value.
code: an application-specific error code, expressed as a string value.
title: a short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
detail: a human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized.
source: an object containing references to the source of the error, optionally including any of the following members:
    pointer: a JSON Pointer [RFC6901] to the associated entity in the request document [e.g. "/data" for a primary data object, or "/data/attributes/title" for a specific attribute].
    parameter: a string indicating which URI query parameter caused the error.
meta: a meta object containing non-standard meta-information about the error.
*/
type JSONError struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Links   struct {
		About string `json:"about"`
	} `json:"links"`
	Status string `json:"status"`
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Source struct {
		Pointer   string `json:"pointer"`
		Parameter string `json:"parameter"`
	} `json:"source"`
	Meta JSONMetaPagination `json:"meta"`
}

type JSONData struct {
	Id         string      `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
	Links      struct {
		Self  string `json:"self"`
		First string `json:"first"`
	} `json:"links"`
	Meta struct {
		Creation     string `json:"creation"`
		Modification string `json:"modification"`
		Version      int    `json:"version"`
	}
}

type JSONMetaPagination struct {
	Total       int `json:"total"`
	Total_pages int `json:"total_pages"`
	Offset      int `json:"offset"`
	Limit       int `json:"limit"`
	Count       int `json:"count"`
}

type JSONContent struct {
	Meta JSONMetaPagination `json:"meta"`
	Data []JSONData         `json:"data"`
}

type JSONContentSingleData struct {
	Meta JSONMetaPagination `json:"meta"`
	Data JSONData           `json:"data"`
}

// Custom type

/*
Structs definitions
*/
type Adress struct {
	Street   string `json:"street" db:"street"`
	Zip      string `json:"zip" db:"zip"`
	Locality string `json:"locality" db:"locality" search:"B"`
}

type Contact struct {
	Adress  `json:"adress" db:"adress"`
	Tel     string `json:"tel" db:"tel"`
	Fax     string `json:"fax" db:"fax"`
	Website string `json:"website" db:"website" search:"B"`
	Mail    string `json:"mail" db:"mail" search:"B"`
	Name    string `json:"name" db:"name" search:"D"`
}

type ProfessionalBranch struct {
	NafCode string `json:"nafcode" db:"nafcode"`
}

type Session struct {
	StartPeriodValidity string `json:"startperiodvalidity" db:"startperiodvalidity"`
	EndPeriodValidity   string `json:"endperiodvalidity" db:"endperiodvalidity"`
	Adress              Adress `json:"adress" db:"adress"`
	RecruitmentState    string `json:"recruitmentstate" db:"recruitmentstate"`
	Terms               string `json:"terms" db:"terms"`
}

type CPFEligibility struct {
	Name                string               `json:"name" db:"name"`
	Code                string               `json:"code" db:"code"`
	StartPeriodValidity string               `json:"startperiodvalidity" db:"startperiodvalidity"`
	EndPeriodValidity   string               `json:"endperiodvalidity" db:"endperiodvalidity"`
	Region              string               `json:"region" db:"region"`
	Branchs             []ProfessionalBranch `json:"professionalbranch" db:"professionalbranch"`
}

type Formation struct {
	ParentId              string           `json:"parentid" db:"parentid"`
	Name                  string           `json:"name" db:"name" search:"A"`
	Objectives            string           `json:"objectives" db:"objectives" search:"A"`
	Programme             []string         `json:"programme" db:"programme" search:"A"`
	Validation            []string         `json:"validation" db:"validation"`
	Type                  []string         `json:"type" db:"type"`
	OutputLevel           string           `json:"outputlevel" db:"outputlevel"`
	RomeCode              []string         `json:"romecode" db:"romecode" search:"A"`
	PedagogicalTerms      string           `json:"pedagogicalterms" db:"pedagogicalterms"`
	Duration              []string         `json:"duration" db:"duration"`
	WorkStudyTerms        []string         `json:"workstudyterms" db:"workstudyterms"`
	PublicServiceContract bool             `json:"publicservicecontract" db:"publicservicecontract"`
	Financers             []string         `json:"financers" db:"financers"`
	PublicAccess          []string         `json:"publicaccess" db:"publicaccess"`
	AdmissionTerm         []string         `json:"admissionterm" db:"admissionterm"`
	EntryLevel            string           `json:"entrylevel" db:"entrylevel"`
	Prerequisite          []string         `json:"prerequisite" db:"prerequisite"`
	Location              Contact          `json:"location" db:"location"`
	Contact               Contact          `json:"contact" db:"contact"`
	EligibilityEmployee   []CPFEligibility `json:"eligibilityemployee" db:"eligibilityemployee"`
	EligibilityJobSeeker  []CPFEligibility `json:"eligibilityjobseeker" db:"eligibilityjobseeker"`
	EligibilityAll        []CPFEligibility `json:"eligibilityall" db:"eligibilityall"`
	Sessions              []Session        `json:"sessions" db:"sessions"`
	LastUpdate            string           `json:"lastupdate" db:"lastupdate"`
	Ref                   string           `json:"ref" db:"ref"`
}

type Organism struct {
	Name               string              `json:"name" db:"name" search:"A"`
	CorporateName      string              `json:"corporatename" db:"corporatename" search:"A"`
	Contact            Contact             `json:"contact" db:"contact"`
	Siret              string              `json:"siret" db:"siret"`
	RegistrationNumber string              `json:"registrationnumber" db:"registrationnumber"`
	Fields             map[string][]string `json:"fields" db:"fields" search:"A"`
	Formations         []string            `json:"formations" db:"formations"`
}

type QueryModel struct {
	Filter map[string]string `json:"filter"`
	Page   map[string]string `json:"page"`
	Search string            `json:"search"`
}

type Countable interface {
	GetCount() int
}

type DBCountable struct {
	Total_count int    `json:"-" db:"totalcount"`
	Id          string `json:"id" db:"id" modifiers:"primary key not null"`
}

func (this DBCountable) GetCount() int {
	return this.Total_count
}

type Commentable struct {
	GlobalRating int `json:"globalrating" db:"globalrating"`
	NComments    int `json:"ncomments" db:"ncomments"`
}

type CommentModel struct {
	DBCountable
	Author       string    `json:"author" db:"author"`
	Role         string    `json:"role" db:"role"`
	Content      string    `json:"content" db:"content"`
	ParentID     string    `json:"parentid" db:"parentid"`
	GlobalRating int       `json:"globalrating" db:"globalrating"`
	UpVote       int       `json:"upvote" db:"upvote"`
	DownVote     int       `json:"downvote" db:"downvote"`
	PostDate     time.Time `json:"postdate" db:"postdate"`
}

type CriteriasFormation struct {
	// Rating criterias
	QualityRating             int `json:"qualityrating" db:"qualityrating"`
	TeachersRating            int `json:"teachersrating" db:"teachersrating"`
	AffordabilityRating       int `json:"affordabilityrating" db:"affordabilityrating"`
	HeadCountRating           int `json:"headcountrating" db:"headcountrating"`
	MonitoringRating          int `json:"monitoringrating" db:"monitoringrating"`
	EquipmentRating           int `json:"equipmentrating" db:"equipmentrating"`
	ExternalRating            int `json:"externalrating" db:"externalrating"`
	ProfessionalisationRating int `json:"professionalisationrating" db:"professionalisationrating"`
	SalaryRating              int `json:"salaryrating" db:"salaryrating"`
	RecognitionRating         int `json:"recognitionrating" db:"recognitionrating"`
	AmbianceRating            int `json:"ambiancerating" db:"ambiancerating"`
	ExtraActivityRating       int `json:"extraactivityrating" db:"extraactivityrating"`
}

type CriteriasOrganism struct {
	// Rating criterias
	HygieneRating       int `json:"hygienerating" db:"hygienerating"`
	SizeRating          int `json:"sizerating" db:"sizerating"`
	AdminRating         int `json:"adminrating" db:"adminrating"`
	AccessibilityRating int `json:"accessibilityrating" db:"accessibilityrating"`
	EnvironmentalRating int `json:"environmentalrating" db:"environmentalrating"`
	StuffRating         int `json:"stuffrating" db:"stuffrating"`
}

type FormationModel struct {
	Formation
	DBCountable
	CriteriasFormation
	Commentable
}

type OrganismModel struct {
	Organism
	DBCountable
	CriteriasOrganism
	Commentable
}

type CommentFormationModel struct {
	CommentModel
	CriteriasFormation
}

type CommentFormationModelWithVote struct {
	CommentFormationModel
	Vote int `json:"vote" db:"vote"` // 0 for no vote, -1 for downvote, 1 for upvote
}

type CommentOrganismModel struct {
	CommentModel
	CriteriasOrganism
}

type CommentOrganismModelWithVote struct {
	CommentOrganismModel
	Vote int `json:"vote" db:"vote"` // 0 for no vote, -1 for downvote, 1 for upvote
}

type CommentVotes struct {
	CommentID string `json:"commentid" db:"commentid"`
}

type OrganismVoteUser struct {
	UserID     string `json:"userid" db:"userid"`
	OrganismID string `json:"organismid" db:"organismid"`
}

type CommentOrganismVoteUser struct {
	OrganismVoteUser
	CommentVotes
	Vote bool `json:"vote" db:"vote"`
}

type FormationVoteUser struct {
	UserID      string `json:"userid" db:"userid"`
	FormationID string `json:"formationid" db:"formationid"`
}

type CommentFormationVoteUser struct {
	FormationVoteUser
	CommentVotes
	Vote bool `json:"vote" db:"vote"`
}

type OrganismWithFormationNames struct {
	OrganismModel
	FormationsInfo []FormationModel `json:"formationsinfo" db:"formationsinfo"`
}

type User struct {
	ID            string `json:"id" db:"id"`
	Email         string `json:"email" db:"email"`
	Password      string `json:"password" db:"password"`
	Validate      bool   `json:"validate" db:"validate"`
	FirstName     string `json:"firstname" db:"firstname"`
	LastName      string `json:"lastname" db:"lastname"`
	OrganismOwner string `json:"organismowner" db:"organismowner"`
}

type UserInfo struct {
	Name             string `json:"name" db:"name"`
	Logged           bool   `json:"logged" db:"logged"`
	AlreadyCommented bool   `json:"alreadycommented" db:"alreadycommented"`
	OrganismOwner    string `json:"organismowner" db:"organismowner"`
}

type UserInfoWithData struct {
	UserInfo
	Data interface{}
}

type MailInfo struct {
	To      []string
	From    string
	Subject string
	Body    string
	Key     string
}
