/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/27 17:30:24 by hdezier           #+#    #+#             */
/*   Updated: 2016/11/19 16:08:16 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-server/models"
	"github.com/zemirco/uid"
	"os"
	"regexp"
	"strings"
)

const (
	root = "http://www.intercariforef.org/formations/"
	logF = "log_formation_idf.json"
	logO = "log_organism_idf.json"
)

// func decodeCaptcha(url string) (code string) {
// 	response, e := http.Get(url)
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// 	defer response.Body.Close()
// 	//open a file for writing
// 	filePNG, err := os.Create("captcha.png")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// Use io.Copy to just dump the response body to the filePNG. This supports huge files
// 	_, err = io.Copy(filePNG, response.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	filePNG.Close()

// 	fmt.Println("PNG has been copied")

// 	client, _ := gosseract.NewClient()
// 	code, err = client.Src("captcha.png").Out()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("PNG has been analyzed: ", code)
// 	return
// }

func throughCaptcha(url string) (doc *goquery.Document) {
	for {
		doc, err := goquery.NewDocument(url)
		if err != nil || strings.HasPrefix(doc.Text(), "\n\nVous avez dépassé le nombre de visites autorisées par jour.") {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter to continue/ quit to exit")
			text, _ := reader.ReadString('\n')
			if text == `quit` {
				panic("Bye")
			}
			continue
		}
		return doc
	}
}

func writeLog(filename string, text string) {
	fmt.Println(text)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(text + `,`); err != nil {
		panic(err)
	}
}

func mainTestCaptcha() {
	url := `http://www.intercariforef.org/formations/recherche-organismes.html`
	throughCaptcha(url)
}

func mainTestOrga() {
	f, e := os.Open("test/orga.html")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		fmt.Println(e)
	}
	parseOrganism(doc)
}

func mainTestForma() {
	f, e := os.Open("test/formation.html")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		fmt.Println(e)
	}
	formation := parseFormation(doc, `TEST`)
	spew.Dump(formation)
}

func mainTestBranch() {
	f, e := os.Open("test/branches.html")
	if e != nil {
		fmt.Println(e)
	}
	defer f.Close()
	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		fmt.Println(e)
	}
	branches := parseBranchs(doc)
	spew.Dump(branches)
}

func mainScrap() {
	mainUrl := root + "annuaire-formation.html"
	doc := throughCaptcha(mainUrl)
	regionRgx := regexp.MustCompile("^liste-departement-")
	doc.Find("#content_onglet").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Found first content")
		s.Children().First().Children().FilterFunction(func(i int, s *goquery.Selection) bool {
			href, hasHref := s.Attr("href")
			return hasHref && regionRgx.Match([]byte(href))
		}).Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			// Conditional region
			if href != `liste-departement-11.html` {
				return
			}
			parseDepartements(throughCaptcha(root + href))
		})
	})
}

func main() {
	_, _ = os.OpenFile(logF, os.O_CREATE, 0600)
	_, _ = os.OpenFile(logO, os.O_CREATE, 0600)
	// mainTestBranch()
	// mainTestForma()
	// mainTestOrga()
	// mainTestCaptcha()
	mainScrap()
}

func parseDepartements(doc *goquery.Document) {
	fmt.Println("Found departement")
	departementRgx := regexp.MustCompile("^liste-organisme-")
	doc.Find("#content_onglet").Each(func(i int, s *goquery.Selection) {
		s.Children().First().Children().FilterFunction(func(i int, s *goquery.Selection) bool {
			href, hasHref := s.Attr("href")
			return hasHref && departementRgx.Match([]byte(href))
		}).Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			parseOrganisms(throughCaptcha(root + href))
		})
	})
}

func parseOrganisms(doc *goquery.Document) {
	fmt.Println("Found organisms")
	organismRgx := regexp.MustCompile(`^http\:\/\/www.intercariforef.org\/formations\/`)
	doc.Find("#content_onglet").Each(func(i int, s *goquery.Selection) {
		s.Children().First().
			ChildrenFiltered("table").
			ChildrenFiltered("tbody").
			ChildrenFiltered("tr").
			ChildrenFiltered("td").
			Children().
			FilterFunction(func(i int, s *goquery.Selection) bool {
				href, hasHref := s.Attr("href")
				return hasHref && organismRgx.Match([]byte(href))
			}).Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			parseOrganism(throughCaptcha(href))
		})
	})
}

func parseOrganism(doc *goquery.Document) {
	organism := models.OrganismModel{}
	organism.Id = uid.New(11)
	doc.Find("#content_onglet").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(i int, s *goquery.Selection) {
			classes, hasClass := s.Attr("class")
			if !hasClass {
				return
			}
			switch classes {
			case "titreFormationContainer":
				organism.Name = parseMainName(s)
				break
			case "onglet_bloc":
				organism.CorporateName,
					organism.Contact.Adress,
					organism.Contact.Tel,
					organism.Contact.Mail,
					organism.Contact.Website,
					organism.Siret,
					organism.RegistrationNumber = parseOrganismInfo(s)
				break
			case "blocAfficheCache":
				id, hasId := s.Attr("id")
				if !hasId {
					return
				}
				switch id {
				case "ACB_DomainesFormation":
					organism.Fields = parseOrganismFields(s)
					break
				case "ACB_DetailFormations":
					organism.Formations = parseFormationsLink(s, organism.Id)
					break
				}
				break
			}
		})
	})
	jsonOrganism, err := json.Marshal(organism)
	if err != nil {
		fmt.Println("Error occured marshaling organism")
	}
	writeLog(logO, string(jsonOrganism))
}

func parseMainName(s *goquery.Selection) string {
	return s.Children().First().Text()
}

func parseOrganismInfo(s *goquery.Selection) (corporateName string, adress models.Adress, tel string, mail string, website string, siret string, registrationNumber string) {
	innerDiv := s.Children().First()
	innerTxt := innerDiv.Text()
	innerTxt = innerTxt[1:]
	corporateName = innerTxt[17:strings.Index(innerTxt, "\n")]
	innerTxt = innerTxt[strings.Index(innerTxt, "\n")+2:]

	adress.Street = innerTxt[:strings.Index(innerTxt, "\n")]
	innerTxt = innerTxt[strings.Index(innerTxt, "\n")+1:]
	adress.Zip = innerTxt[:strings.Index(innerTxt, "\n")]
	innerTxt = innerTxt[strings.Index(innerTxt, "\n")+1:]
	adress.Locality = innerTxt[:strings.Index(innerTxt, "\n")]
	innerTxt = innerTxt[strings.Index(innerTxt, "\n")+2:]

	lines := strings.Split(innerTxt, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, `Site Internet : `) {
			website = line[16:]
		} else if strings.HasPrefix(line, `Tél : `) {
			tel = line[6:]
		} else if strings.HasPrefix(line, `Siret :`) {
			siret = line[7:]
		} else if strings.HasPrefix(line, `Enregistrée sous le numéro : `) {
			registrationNumber = line[31:]
		} else if strings.HasPrefix(line, `m_a_i_l_trompe_robot`) {
			mail = parseMailFromJS(line)
		}
	}
	return
}

func parseOrganismFields(s *goquery.Selection) (result map[string][]string) {
	result = make(map[string][]string)
	fieldsStr := s.Text()
	var currentKey string
	for ; ; fieldsStr = fieldsStr[strings.Index(fieldsStr, "\n")+1:] {
		endLine := strings.Index(fieldsStr, "\n")
		if endLine == -1 {
			return
		}
		line := fieldsStr[:strings.Index(fieldsStr, "\n")]
		line = strings.Trim(line, ` `)
		if line == "" {
			continue
		}
		if line[len(line)-1] == ':' {
			currentKey = line[:len(line)-1]
		} else {
			skills := strings.Split(line, `, `)
			for _, val := range skills {
				result[currentKey] = append(result[currentKey], strings.Trim(val, ` `))
			}
		}
	}
	return
}

func parseMailFromJS(content string) (result string) {
	if content == "" {
		return ""
	}
	splittedStr := strings.Split(content, `"`)
	if len(splittedStr) < 4 {
		return ""
	}
	return splittedStr[1] + `@` + splittedStr[3]
}

func parseFormationsLink(s *goquery.Selection, id string) (result []string) {
	fmt.Println("Found formations")
	s.Children().First().ChildrenFiltered("ul").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered("li").Each(func(i int, s *goquery.Selection) {
			href, hasHref := s.Children().First().Attr("href")
			if !hasHref {
				return
			}
			result = append(result, parseFormation(throughCaptcha(href[:len(href)-6]), id))
		})
	})
	return
}

/*
type Formation struct {
	Name                  string           `json:"name"`
	Objectives            string           `json:"objectives"`
	Programme             []string           `json:"programme"`
	Validation            []string           `json:"validation"`
	Type                  []string           `json:"type"`
	OutputLevel           string           `json:"outputlevel"`
	RomeCode   []string            			`json:"RomeCode"`
	PedagogicalTerms      string           `json:"pedagogicalterms"`
	Duration              []string           `json:"duration"`
	WorkStudyTerms        string           `json:"workstudyterms"`
	PublicServiceContract bool             `json:"publicservicecontract"`
	Financers              []string         `json:"financers"`
	PublicAccess          []string         `json:"publicaccess"`
	AdmissionTerm         []string         `json:"admissionterm"`
	EntryLevel            []string         `json:"entrylevel"`
	Prerequisite          []string           `json:"prerequisite"`
	Location              string           `json:"location"`
	Contact               Contact          `json:"contact"`
	EligibilityEmployee   []CPFEligibility `json:"eligibilityemployee"`
	EligibilityJobSeeker  []CPFEligibility `json:"eligibilityjobseeker"`
	EligibilityAll        []CPFEligibility `json:"eligibilityall"`
}
*/
func parseFormation(doc *goquery.Document, parentId string) (id string) {
	formation := models.FormationModel{}
	id = uid.New(11)
	formation.Id = id
	doc.Find("#content_onglet").Each(func(i int, s *goquery.Selection) {
		s.Children().Each(func(i int, s *goquery.Selection) {
			formation.ParentId = parentId
			classes, hasClass := s.Attr("class")
			if !hasClass {
				return
			}
			switch classes {
			case "titreFormationContainer":
				formation.Name = parseMainName(s)
				break
			case "blocAfficheCache":
				id, hasId := s.Attr("id")
				if !hasId {
					return
				}
				switch id {
				case "ACB_ObjectifProgramme":
					formation.Objectives,
						formation.Programme,
						formation.Validation,
						formation.Type,
						formation.OutputLevel = parseFormationInfo(s)
					break
				case "ACB_metiers":
					formation.RomeCode = parseJobs(s)
					break
				case "ACB_DureeRythme":
					formation.PedagogicalTerms,
						formation.Duration,
						formation.WorkStudyTerms,
						formation.PublicServiceContract,
						formation.Financers = parseDurationAndTerms(s)
					break
				case "ACB_condAccess":
					formation.PublicAccess,
						formation.AdmissionTerm,
						formation.EntryLevel,
						formation.Prerequisite = parseConditions(s)
					break
				case "ACB_lieuDeFormation":
					formation.Location = parseFormationLocation(s)
					break
				case "ACB_Inscription":
					formation.Contact = parseFormationContact(s)
					break
				case "ACB_eligibleCPF1":
					formation.EligibilityEmployee = parseEligibility(s)
					break
				case "ACB_eligibleCPF2":
					formation.EligibilityJobSeeker = parseEligibility(s)
					break
				case "ACB_eligibleCPF3":
					formation.EligibilityAll = parseEligibility(s)
					break
				case "ACB_Sessions":
					formation.Sessions = parseSessions(s)
					break
				}
				break
			}
		})
	})
	jsonFormation, err := json.Marshal(formation)
	if err != nil {
		fmt.Println("Error occured marshaling formation")
	}
	writeLog(logF, string(jsonFormation))
	return
}

const (
	trimStr           = " -,•.;:\n*<>/0123456789"
	trimStrWithNumber = " -,•.;:\n*<>/"
)

func HTMLToText(html string) (text string) {
	parsedDiv, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	return parsedDiv.Text()
}

func splitHTMLNewLineComma(html string) (result []string) {
	splittedBr := strings.Split(html, `<br/>`)
	for _, val := range splittedBr {
		parsedDiv, err := goquery.NewDocumentFromReader(strings.NewReader(val))
		if err != nil || len(parsedDiv.Text()) == 0 {
			continue
		}
		if len(splittedBr) != 1 {
			result = append(result, strings.Trim(parsedDiv.Text(), trimStr))
			continue
		}
		// If only one line we resplit by comma
		comaSplitted := strings.Split(parsedDiv.Text(), `;`)
		if len(comaSplitted) == 1 {
			comaSplitted = strings.Split(parsedDiv.Text(), `,`)
		}
		for _, val := range comaSplitted {
			result = append(result, HTMLToText(strings.Trim(val, trimStr)))
		}
	}
	return
}

/*
Objectives            string           `json:"objectives"`
Programme             []string           `json:"programme"`
Validation            []string           `json:"validation"`
Type                  []string           `json:"type"`
OutputLevel           string           `json:"outputlevel"`
*/
func parseFormationInfo(s *goquery.Selection) (objectives string, programme []string, validation []string, types []string, outputLevel string) {
	outerHtml, err := goquery.OuterHtml(s)
	if err != nil {
		return
	}
	lines := strings.Split(outerHtml, "\n")
	for _, line := range lines {
		if len(line) <= 24 {
			continue
		}
		if strings.HasPrefix(line[24:], `Objectifs : `) {
			parsedDiv, err := goquery.NewDocumentFromReader(strings.NewReader(line[24+12:]))
			if err != nil {
				continue
			}
			objectives = strings.Trim(parsedDiv.Text(), trimStr)
		} else if strings.HasPrefix(line[24:], `Programme de la formation : `) {
			programme = splitHTMLNewLineComma(line[24+28:])
		} else if strings.HasPrefix(line[24:], `Validation et sanction : `) {
			validation = splitHTMLNewLineComma(line[24+25:])
		} else if strings.HasPrefix(line[24:], `Type de formation : `) {
			types = splitHTMLNewLineComma(line[24+20:])
		} else if strings.HasPrefix(line[24:], `Niveau de sortie : `) {
			parsedDiv, err := goquery.NewDocumentFromReader(strings.NewReader(line[24+19:]))
			if err != nil {
				continue
			}
			outputLevel = strings.Trim(parsedDiv.Text(), trimStr)
		}
	}
	return
}

/*
	RomeCode   []Job            `json:"RomeCode"`
*/
func parseJobs(s *goquery.Selection) (result []string) {
	lines := strings.Split(s.Text(), "\n")
	for _, line := range lines {
		splittedJob := strings.Split(line, ` : `)
		if len(splittedJob) < 2 {
			continue
		}
		result = append(result, strings.Trim(splittedJob[0], trimStrWithNumber))
	}
	return
}

/*
PedagogicalTerms      string           `json:"pedagogicalterms"`
Duration              []string           `json:"duration"`
WorkStudyTerms        []string           `json:"workstudyterms"`
PublicServiceContract bool             `json:"publicservicecontract"`
Financers              []string         `json:"financers"`

*/
func parseDurationAndTerms(s *goquery.Selection) (pedagogicalTerms string, duration []string, workStudyTerms []string, publicServiceContract bool, financers []string) {
	lines := strings.Split(s.Text(), "\n")
	lastKey := ``
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, `Modalites pédagogiques : `) {
			pedagogicalTerms = line[25:]
			lastKey = `Modalites pédagogiques : `
		} else if strings.HasPrefix(line, `Durée : `) {
			splittedLine := strings.Split(line[8:], `,`)
			for _, val := range splittedLine {
				duration = append(duration, strings.Trim(val, trimStrWithNumber))
			}
			lastKey = `Durée : `
		} else if strings.HasPrefix(line, `Modalités de l'alternance : `) {
			splittedLine := strings.Split(line[28:], `-`)
			for _, val := range splittedLine {
				workStudyTerms = append(workStudyTerms, strings.Trim(val, trimStrWithNumber))
			}
			lastKey = `Modalités de l'alternance : `
		} else if strings.HasPrefix(line, `Conventionnement : `) {
			publicServiceContract = func(str string) bool {
				if str == `Oui` {
					return true
				}
				return false
			}(strings.Trim(line[19:], trimStr))
		} else if strings.HasPrefix(line, `Financeur(s) : `) {
			// Dirty stuff, this </br> case is annoying
			outerHtml, _ := goquery.OuterHtml(s)
			splittedHtml := strings.Split(outerHtml, "\n")
			for _, line := range splittedHtml {
				if strings.HasPrefix(line, `<span class="titreInfo">Financeur(s) : </span>`) {
					splittedFinancers := strings.Split(line[46:], `<br/>`)
					for _, financer := range splittedFinancers {
						if len(financer) == 0 {
							continue
						}
						parsedDiv, err := goquery.NewDocumentFromReader(strings.NewReader(financer))
						if err != nil {
							continue
						}
						financers = append(financers, strings.Trim(parsedDiv.Text(), trimStr))
					}
				}
			}
		} else {
			if lastKey == `Durée : ` {
				duration = append(duration, strings.Trim(line, trimStrWithNumber))
			}
			lastKey = ``
		}
	}
	return
}

/*
PublicAccess          []string         `json:"publicaccess"`
AdmissionTerm         []string         `json:"admissionterm"`
EntryLevel            string         `json:"entrylevel"`
Prerequisite          []string           `json:"prerequisite"`
*/
func parseConditions(s *goquery.Selection) (publicAccess []string, admissionTerm []string, entryLevel string, prerequisite []string) {
	lines := strings.Split(s.Text(), "\n")
	lastKey := ``
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, `Public(s) : `) {
			splittedComma := strings.Split(line[12:], `,`)
			for _, val := range splittedComma {
				publicAccess = append(publicAccess, strings.Trim(val, trimStr))
			}
			lastKey = `Public(s) : `
		} else if strings.HasPrefix(line, `Modalités de recrutement et d'admission : `) {
			splittedSemiColon := strings.Split(line[42:], `;`)
			for _, val := range splittedSemiColon {
				admissionTerm = append(admissionTerm, strings.Trim(val, trimStr))
			}
			lastKey = `Modalités de recrutement et d'admission : `
		} else if strings.HasPrefix(line, `Niveau d'entrée : `) {
			entryLevel = strings.Trim(line[18:], trimStr)
			lastKey = `Niveau d'entrée : `
		} else if strings.HasPrefix(line, `Conditions spécifiques et prérequis : `) {
			outerHtml, _ := goquery.OuterHtml(s)
			splittedHtml := strings.Split(outerHtml, "\n")
			for _, line := range splittedHtml {
				if strings.HasPrefix(line, `<span class="titreInfo">Conditions spécifiques et prérequis : </span>`) {
					splittedBr := strings.Split(line[71:], `<br/>`)
					for _, valBr := range splittedBr {
						if len(valBr) == 0 {
							continue
						}
						splittedDot := strings.Split(valBr, `.`)
						for _, valDot := range splittedDot {
							parsedDiv, err := goquery.NewDocumentFromReader(strings.NewReader(valDot))
							if err != nil || len(parsedDiv.Text()) == 0 {
								continue
							}
							prerequisite = append(prerequisite, strings.Trim(parsedDiv.Text(), trimStr))
						}
					}
				}
			}
			lastKey = `Conditions spécifiques et prérequis : `
		} else {
			if lastKey == `Public(s) : ` {
				publicAccess = append(publicAccess, strings.Trim(line, trimStr))
			}
			lastKey = ``
		}
	}
	return
}

func parseFormationLocation(s *goquery.Selection) (result models.Contact) {
	adressLines := strings.Split(s.Children().First().Text(), "\n")
	// Approximative ! No rules for this
	if len(adressLines) >= 3 {
		result.Adress.Street = adressLines[0]
		result.Adress.Locality = adressLines[1]
		result.Adress.Zip = adressLines[2]
	}
	lines := strings.Split(s.Text(), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, `Site web : `) {
			result.Website = strings.Trim(line[11:], trimStrWithNumber)
		} else if strings.HasPrefix(line, `Téléphone fixe : `) {
			result.Tel = strings.Trim(line[17:], trimStrWithNumber)
		} else if strings.HasPrefix(line, `m_a_i_l_trompe_robot`) {
			result.Mail = parseMailFromJS(line)
		}
	}
	return
}

func parseFormationContact(s *goquery.Selection) (result models.Contact) {
	lines := strings.Split(s.Text(), "\n")
	for _, line := range lines {
		// FTM, let's ignore this weird double
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, `Contacter l'organisme formateur`) {
			return
		}
		if strings.HasPrefix(line, `Contact sur la formation : `) {
			line = strings.Trim(line[27:], trimStr)
		}
		if strings.HasPrefix(line, `Responsable : `) {
			result.Name = strings.Trim(line[14:], trimStr)
		} else if strings.HasPrefix(line, `Téléphone fixe : `) {
			result.Tel = strings.Trim(line[17:], trimStrWithNumber)
		} else if strings.HasPrefix(line, `Site web : `) {
			result.Website = strings.Trim(line[11:], trimStr)
		} else if strings.HasPrefix(line, `m_a_i_l_trompe_robot`) {
			result.Mail = parseMailFromJS(line)
		}
	}
	return
}

func parseEligibility(s *goquery.Selection) (result []models.CPFEligibility) {
	regHref := regexp.MustCompile(`<a href="([^"]*)`)
	outerHtml, err := goquery.OuterHtml(s)
	if err != nil {
		return
	}
	linesHtml := strings.Split(outerHtml[52:], `<br/>`)
	currentEligibility := models.CPFEligibility{}
	for _, line := range linesHtml {
		line = strings.TrimLeft(line, "\n")
		if strings.HasPrefix(line, `Code CPF `) {
			if len(currentEligibility.Code) > 0 {
				result = append(result, currentEligibility)
				currentEligibility = models.CPFEligibility{}
			}
			nextDash := strings.Index(line, `-`)
			if nextDash <= 9 {
				continue
			}
			currentEligibility.Code = strings.Trim(line[9:nextDash], trimStrWithNumber)
		} else if strings.HasPrefix(line, `Validité du `) {
			nextDash := strings.Index(line, `-`)
			if nextDash <= 12 {
				continue
			}
			splittedLine := strings.Split(line[12:], `-`)
			if len(splittedLine) < 2 {
				continue
			}
			currentEligibility.Region = strings.Trim(splittedLine[1], trimStr)
			splittedDates := strings.Split(splittedLine[0], `au`)
			if len(splittedDates) < 2 {
				continue
			}
			currentEligibility.StartPeriodValidity = strings.Trim(splittedDates[0], trimStrWithNumber)
			currentEligibility.EndPeriodValidity = strings.Trim(splittedDates[1], trimStrWithNumber)
		} else if strings.HasPrefix(line, `<div class="LiensAouvrirEnPopupReduite">`) {
			href := regHref.FindString(line)
			currentEligibility.Branchs = parseBranchs(throughCaptcha(string(href[9:])))
		} else if strings.HasPrefix(line, `Branches professionnelles : `) {
			currentEligibility.Branchs = append(currentEligibility.Branchs,
				models.ProfessionalBranch{
					NafCode: ``,
				})
		}
	}
	if len(currentEligibility.Code) > 0 {
		result = append(result, currentEligibility)
	}
	return
}

func parseBranchs(doc *goquery.Document) (result []models.ProfessionalBranch) {
	doc.Find("#content_onglet").Each(func(i int, s *goquery.Selection) {
		s.ChildrenFiltered(`ul`).Each(func(i int, s *goquery.Selection) {
			s.ChildrenFiltered(`li`).Each(func(i int, s *goquery.Selection) {
				splittedLine := strings.Split(s.Text(), `:`)
				if len(splittedLine) < 2 {
					return
				}
				result = append(result, models.ProfessionalBranch{
					NafCode: strings.Trim(splittedLine[0], trimStrWithNumber),
				})
			})
		})
	})
	return
}

func parseSessions(s *goquery.Selection) (sessions []models.Session) {
	outerHtml, err := goquery.OuterHtml(s)
	if err != nil {
		return
	}
	linesHtml := strings.Split(outerHtml, `<br/>`)
	currentSession := models.Session{}
	adressIndicator := 0
	for _, line := range linesHtml {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, `<div id="ACB_Sessions" class="blocAfficheCache">`) {
			splittedLn := strings.Split(line, "\n")
			for _, line := range splittedLn {
				if strings.HasPrefix(line, `Du `) {
					splittedDates := strings.Split(line[3:], `au`)
					if len(splittedDates) < 2 {
						continue
					}
					currentSession.StartPeriodValidity = strings.Trim(splittedDates[0], trimStrWithNumber)
					currentSession.EndPeriodValidity = strings.Trim(splittedDates[1], trimStrWithNumber)
					break
				}
			}
		} else if strings.HasPrefix(line, `<b>Etat du recrutement :</b>`) {
			currentSession.RecruitmentState = strings.Trim(line[28:], trimStr)
		} else if strings.HasPrefix(line, `<b>Modalité :</b>`) {
			currentSession.Terms = strings.Trim(HTMLToText(line[17:]), trimStr)
		} else if strings.HasPrefix(line, `<b>Adresse d&#39;inscription :</b>`) {
			adressIndicator++
		} else if adressIndicator == 1 {
			currentSession.Adress.Locality = strings.Trim(line, trimStrWithNumber)
			adressIndicator++
		} else if adressIndicator == 2 {
			currentSession.Adress.Street = strings.Trim(line, trimStrWithNumber)
			adressIndicator++
		} else if adressIndicator == 3 {
			currentSession.Adress.Zip = strings.Trim(line, trimStrWithNumber)
			adressIndicator = 0
		}
	}
	sessions = append(sessions, currentSession)
	return
}
