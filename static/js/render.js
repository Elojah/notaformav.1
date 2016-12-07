/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   render.js                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/27 16:16:07 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 20:48:19 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

function renderProgrammeItems(obj, id) {
	var result = ""
	obj.forEach(function (elem) {
		result += `
		<div class="ui item action input">
			<input value="` + elem + `" type="text" id="edit-formation-` + id + `-programme">
			<button class="ui teal labeled icon left attached button" onclick="EditInfoFormation('programme', '` + id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
			<button class="ui red labeled icon right attached button" onclick="EditInfoFormation('programme', '` + id + `')">
			  <i class="delete icon"></i>
			  Supprimer
			</button>
		</div>`
	})
	return result
}

function renderTagFields(obj) {
	var result = ""
	for (var prop in obj) {
		result += `<div class="ui image label">` + prop + `</div>`
	}
	return result
}

function renderTagFieldsValues(obj) {
	var result = ""
	obj.forEach(function (elem) {
		if (elem.length > 50) {return;}
		result += `<div class="ui image label">` + elem + `</div>`
	})
	return result
}

function	renderOrganismPagination(i, current) {
	var isActive = ""
	if (current == i) {
		isActive = " active"
	}
	return (`
		<a class="item ` + isActive + `" href="#" onclick="refresh('list-organisms', {page:` + i + `, limit:25, sortOrder:null})">` + i + `</a>
		`)
}

function	renderFormationPagination(i, current) {
	var isActive = ""
	if (current == i) {
		isActive = " active"
	}
	return (`
		<a class="item ` + isActive + `" href="#" onclick="refresh('list-formations', {page:` + i + `, limit:25, sortOrder:null})">` + i + `</a>
		`)
}

function	renderSearchFormationPagination(i, current) {
	var isActive = ""
	if (current == i) {
		isActive = " active"
	}
	return (`
		<a class="item ` + isActive + `" href="#" onclick="refresh('list-formations', {page:` + i + `, limit:25, sortOrder:null, q:$.urlParam('q')})">` + i + `</a>
		`)
}

function	renderCommentFormationPagination(i, current) {
	var isActive = ""
	if (current == i) {
		isActive = " active"
	}
	return (`
		<a class="item ` + isActive + `" href="#" onclick="refresh('list-comments', {page:` + i + `, limit:25, sortOrder:null})">` + i + `</a>
		`)
}

function	renderCommentOrganismPagination(i, current) {
	var isActive = ""
	if (current == i) {
		isActive = " active"
	}
	return (`
		<a class="item ` + isActive + `" href="#" onclick="refresh('list-comments', {page:` + i + `, limit:25, sortOrder:null})">` + i + `</a>
		`)
}

function	renderImageRole(role) {
	switch (role) {
		case "student":
			return (`<img class="ui mini image left floated" src="/static/img/student.png">`)
		case "graduated":
			return (`<img class="ui mini image left floated" src="/static/img/graduate.png">`)
		case "teacher":
			return (`<img class="ui mini image left floated" src="/static/img/teacher.png">`)
		case "employee":
			return (`<img class="ui mini image left floated" src="/static/img/employee.png">`)
		case "external":
			return (`<img class="ui mini image left floated" src="/static/img/confspeaker.png">`)
		case "none":
			return (`<img class="ui mini image left floated" src="/static/img/usergroup.png">`)
		default :
			return (`<img class="ui mini image left floated" src="/static/img/usergroup.png">`)
	}
}

function	renderTitleRole(role) {
	switch (role) {
		case "student":
			return (`Étudiant`)
		case "graduated":
			return (`Diplômé`)
		case "teacher":
			return (`Enseignant`)
		case "employee":
			return (`Employé`)
		case "external":
			return (`Intervenant externe`)
		case "none":
			return (`Autre`)
		default :
			return (``)
	}
}

/*"author"
"content"
"organismid"
"hygienerating"
"sizerating"
"adminrating"
"extraactivityrating"
"easyaccessrating"
"openinghoursrating"
"environmentalrating"
"externalrating"
*/
function	renderCommentOrganism(com) {
	return (`
		<div class="ui segments">
			<div class="comment ui horizontal segments">
				<a class="avatar avatarcom">`+ renderImageRole(com.attributes.role) +`</a>
				<div class="ui content">
					<a class="author ui basic segment">` + renderTitleRole(com.attributes.role) + `</a>
					<div class="metadata ui basic segment">
						<span class="date"></span>
						<p>Note globale
						<div class="ui rating disable" data-rating="` + com.attributes.globalrating + `" data-max-rating="5"></div>
						</p>
					</div>
					<div class="metadata ui basic segment">
						<a href="#" onclick="Upvote('` + com.attributes.id + `');">
							<i class="thumbs up icon teal"></i><p class="colorturquoise" id="upvote-` + com.attributes.id + `">` + com.attributes.upvote + `</p>
						</a>
						<a href="#" onclick="Downvote('` + com.attributes.id + `');">
							<i class="thumbs down icon grey"></i><p class="colorgrey"  id="downvote-` + com.attributes.id + `">` + com.attributes.downvote + `</p>
						</a>
					</div>
				</div>
			</div>
			<div class="text ui segment">
				` + com.attributes.content + `
			</div>
		<button class="ui fluid button show-details-comment" id="show-` + com.attributes.id + `">
			Voir les notes en detail
			<i class="dropdown icon fluid right"></i>
		</button>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Hygiène</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.hygienerating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Effectifs</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.sizerating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Administration</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.adminrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Facilité d'accès</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.accessibilityrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Cadre environnemental</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.environmentalrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Équipements</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.stuffrating + `" data-max-rating="5"></div>
		</div>
		</div>
		</div>
		`)
}

/*"author"
"content"
*/
function	renderCommentFormation(com) {
	return (`

		<div class="ui segments">
			<div class="comment ui horizontal segments">
				<a class="avatar avatarcom">`+ renderImageRole(com.attributes.role) +`</a>
				<div class="ui content">
					<a class="author ui basic segment">` + renderTitleRole(com.attributes.role) + `</a>
					<div class="metadata ui basic segment">
						<span class="date"></span>
						<p>Note globale
						<div class="ui rating disable" data-rating="` + com.attributes.globalrating + `" data-max-rating="5"></div>
						</p>
					</div>
					<div class="metadata ui basic segment">
						<a href="#" onclick="Upvote('` + com.attributes.id + `');">
							<i class="thumbs up icon teal"></i><p class="colorturquoise" id="upvote-` + com.attributes.id + `">` + com.attributes.upvote + `</p>
						</a>
						<a href="#" onclick="Downvote('` + com.attributes.id + `');">
							<i class="thumbs down icon grey"></i><p class="colorgrey" id="downvote-` + com.attributes.id + `">` + com.attributes.downvote + `</p>
						</a>
					</div>
				</div>
			</div>
			<div class="text ui segment">
				` + com.attributes.content + `
			</div>
		<button class="ui fluid button show-details-comment" id="show-` + com.attributes.id + `">
			Voir les notes en detail
			<i class="dropdown icon fluid right"></i>
		</button>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Pédagogie</h4>
		</div>
		<div class="column ui">
		<div class="ui rating disable" data-rating="` + com.attributes.qualityrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Professeurs</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.teachersrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Rapport qualité/prix</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.affordabilityrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Effectifs</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.headcountrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Suivi de la formation</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.monitoringrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Équipements</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.equipmentrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Liens externes</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.externalrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Professionalisation</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.professionalisationrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Salaire de sortie</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.salaryrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Reconnaissance</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.recognitionrating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Ambiance</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.ambiancerating + `" data-max-rating="5"></div>
		</div>
		</div>
		<div class="ui stackable two column grid segment rating-comment transition hidden showable-` + com.attributes.id + `">
		<div class="column">
		<h4>Activités extra scolaires</h4>
		</div>
		<div class="column">
		<div class="ui rating disable" data-rating="` + com.attributes.extraactivityrating + `" data-max-rating="5"></div>
		</div>
		</div>
		</div>
		`)
}

/*"author"
"content"
*/
function	renderFormation(item) {
	return (`
		<div class="ui segments">
			<div class="ui container segment big grid">
				<a href="formation/` + item.attributes.id + `" class="ui top label attached container big eight column floated left colorturquoise">` + item.attributes.name + `
				<div class="ui rating disable three column" data-rating="` + item.attributes.globalrating + `" data-max-rating="5"></div>
				<div class="ui three column floated right">` + item.attributes.contact.adress.locality + `</div>
				</a>
			</div>
			<div class="ui container segment fluid">
				<div>` + renderTagFieldsValues(item.attributes.programme) + `</div>
			</div>
		</div>
		`)
}


/*
"name"
"sigle"
"type"
"sector"
"url"
*/
function	renderOrganism(item) {
	return (`
		<div class="ui segments">
			<div class="ui container segment big grid">
				<a href="organism/` + item.attributes.id + `" class="ui top label attached container big eigth floated column colorturquoise">` + item.attributes.name + `
				<div class="ui rating disable three floated column" data-rating="` + item.attributes.globalrating + `" data-max-rating="5"></div>
				<div class="ui three floated column">` + item.attributes.contact.adress.locality + `</div>
				</a>
			</div>
			<div class="ui container segment fluid">
				<div>` + renderTagFields(item.attributes.fields) + `</div>
			</div>
		</div>
		`)
}

function	renderFormationEdition(item) {
	return (`
		<div class="rows ui form ten wide column">
		 <div class="field">
		  <label>Nom de l'organisme</label>
		  <div class="ui action input">
		   <input value="` + item.data[0].attributes.name + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-name">
		   <button class="ui teal right labeled icon button" onclick="EditInfoFormation('name', '` + item.data[0].attributes.id + `')">
			<i class="edit icon"></i>
			Modifier
		  </button>
		</div>
		</div>
		<div class="field">
		 <label>Objectifs</label>
		 <div class="ui action input">
		   <input value="` + item.data[0].attributes.objectives + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-objectives">
		   <button class="ui teal right labeled icon button" onclick="EditInfoFormation('objectives', '` + item.data[0].attributes.id + `')">
			<i class="edit icon"></i>
			Modifier
		  </button>
		</div>
		</div>
		<div class="field">
		  <label>Programme</label>
		  <div class="ui dropdown fluid selection">
			  <input type="hidden" name="item-programme">
			  <i class="dropdown icon"></i>
			  <div class="default text">Programme</div>
			  <div class="menu">
			  ` + renderProgrammeItems(item.data[0].attributes.programme, item.data[0].attributes.id) + `
			  </div>
		  </div>
		</div>
		<div class="field">
		  <label>Adresse</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.adress.street + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-street">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('street', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Code postal</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.adress.zip + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-zip">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('zip', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Ville</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.adress.locality + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-locality">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('locality', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Telephone</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.tel + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-tel">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('tel', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Telephone</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.mail + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-mail">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('mail', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Telephone</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.name + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-contact-name">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('contact-name', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Fax</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.fax + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-fax">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('fax', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		<div class="field">
		  <label>Website</label>
		  <div class="ui action input">
			<input value="` + item.data[0].attributes.contact.website + `" type="text" id="edit-formation-` + item.data[0].attributes.id + `-website">
			<button class="ui teal right labeled icon button" onclick="EditInfoFormation('website', '` + item.data[0].attributes.id + `')">
			  <i class="edit icon"></i>
			  Modifier
			</button>
		  </div>
		</div>
		</div>
		</div>
		</div>
`)
}
