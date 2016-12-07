/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   url.js                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/09/24 15:07:05 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 19:20:21 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

function	getUrlID()
{
	var	path = window.location.pathname
	return (path.substr(path.lastIndexOf('/') + 1))
}

function	getUrlMain()
{
	var	path = window.location.pathname
	firstSlash = path.indexOf('/')
	secondSlash = path.substr(firstSlash + 1).indexOf('/')
	return (path.substr(firstSlash + 1, secondSlash - firstSlash))
}

function submitCommentFormation() {
	var formParams = {
		"globalrating": $('#globalrating').rating('get rating'),
		"qualityrating": $('#qualityrating').rating('get rating'),
		"teachersrating": $('#teachersrating').rating('get rating'),
		"affordabilityrating": $('#affordabilityrating').rating('get rating'),
		"headcountrating": $('#headcountrating').rating('get rating'),
		"monitoringrating": $('#monitoringrating').rating('get rating'),
		"equipmentrating": $('#equipmentrating').rating('get rating'),
		"externalrating": $('#externalrating').rating('get rating'),
		"professionalisationrating": $('#professionalisationrating').rating('get rating'),
		"salaryrating": $('#salaryrating').rating('get rating'),
		"recognitionrating": $('#recognitionrating').rating('get rating'),
		"ambiancerating": $('#ambiancerating').rating('get rating'),
		"extraactivityrating": $('#extraactivityrating').rating('get rating'),
		"role": $("input[name=role]").val(),
		"content": jQuery('textarea#comment-content').val()
	}
	$.post(
		noHash(window.location.href) + "/comments",
		formParams,
		function (data) {
			console.log(data, data.message)
			if (data.message === "Success") {
				$('#newcommentcontainer').html(`
					<div class="ui success message">
					<i class="close icon"></i>
					<div class="header">
					Votre commentaire a bien ete pris en compte
					</div>
					</div>
					`)
				refresh("list-comments", {page: 1, limit:25, sortOrder:null})
			} else {
				$('#newcommentcontainer').prepend(`
					<div class="ui negative message">
					<i class="close icon"></i>
					<div class="header">
					Une erreur s'est produite
					</div>
					<p>` + data.message + `
					</p></div>
					`)
			}
			$('.ui.rating.disable').rating('disable');
			$('.show-details-comment').click(ShowHideCommentDetails)
		})
}
function submitCommentOrganism() {
	var orgParams = {
		"globalrating": $('#globalrating').rating('get rating'),
		"hygienerating": $('#hygienerating').rating('get rating'),
		"sizerating": $('#sizerating').rating('get rating'),
		"adminrating": $('#adminrating').rating('get rating'),
		"accessibilityrating": $('#accessibilityrating').rating('get rating'),
		"environmentalrating": $('#environmentalrating').rating('get rating'),
		"stuffrating": $('#stuffrating').rating('get rating'),
		"role": $("input[name=role]").val(),
		"content": jQuery('textarea#comment-content').val()
	}
	$.post(
		noHash(window.location.href) + "/comments",
		orgParams,
		function (data) {
			console.log(data, data.message)
			if (data.message === "Success") {
				$('#newcommentcontainer').html(`
					<div class="ui success message">
					<i class="close icon"></i>
					<div class="header">
					Votre commentaire a bien été pris en compte
					</div>
					</div>
					`)
				refresh("list-comments", {page: 1, limit:25, sortOrder:null})
			} else {
				$('#newcommentcontainer').prepend(`
					<div class="ui negative message">
					<i class="close icon"></i>
					<div class="header">
					Une erreur s'est produite
					</div>
					<p>` + data.message + `
					</p></div>
					`)
			}
			$('.ui.rating.disable').rating('disable');
			$('.show-details-comment').click(ShowHideCommentDetails)
		})
}

$.urlParam = function(name){
	var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(noHash(window.location.href));
	if (results==null){
		return null;
	}
	else{
		return results[1] || 0;
	}
}

function	ShowHideAndScroll(actionDivID, ToDivID) {
	$('#' + actionDivID).transition('slide down');
	$('html,body').animate({scrollTop: $('#' + ToDivID).offset().top},'slow');
}

function	ShowHideCommentDetails(event) {
	$('.' + this.id.replace("show", "showable")).transition('slide down');
	$('html,body').animate({scrollTop: $('#' + this.id).offset().top},'slow');
}

function	noHash(s) {
	if (s.slice(-1) == '#') {
		return s.slice(0, -1)
	} else {
		return s
	}
}

function	Upvote(id) {
	CommentVote(id, 1)
}

function	Downvote(id) {
	CommentVote(id, 0)
}

function	CommentVote(id, vote) {
	$.post(
		noHash(window.location.href) + "/comment/" + id,
		{"vote":vote},
		function (data) {
			console.log(data)
			if (data.message === "Success") {
				$('#response-info').html(`
					<div class="ui success message">
					<i class="close icon"></i>
					<div class="header">
					Votre vote a bien ete pris en compte
					</div>
					</div>
					`)
				if (vote == 0) {
					$('#downvote-' + id).html(+$('#upvote-' + id).html() + 1)
				} else if (vote == 1) {
					$('#upvote-' + id).html(+$('#upvote-' + id).html() + 1)
				}
			} else {
				$('#response-info').html(`
					<div class="ui negative message">
					<i class="close icon"></i>
					<div class="header">
					Une erreur s'est produite
					</div>
					<p>` + data.message + `
					</p></div>
					`)
			}
		}
		)
}

function	EditInfo(field) {
	jsonData = {};
	jsonData[field] = $('#edit-field-' + field)[0].value
	$.post(
		noHash(window.location.href) + "/edit",
		jsonData,
		function (data) {
			console.log(data)
			if (data.message === "Success") {
				$('#response-info').html(`
					<div class="ui success message">
					<i class="close icon"></i>
					<div class="header">
					Edition ok
					</div>
					</div>
					`)
			} else {
				$('#response-info').html(`
					<div class="ui negative message">
					<i class="close icon"></i>
					<div class="header">
					Une erreur s'est produite
					</div>
					<p>` + data.message + `
					</p></div>
					`)
			}
		}
		)
}


function	EditInfoFormation(field, id) {
	jsonData = {};
	jsonData[field] = [id, $('#edit-formation-' + id + '-' + field)[0].value];
	$.post(
		noHash(window.location.href) + "/formation/edit",
		jsonData,
		function (data) {
			console.log(data)
			if (data.message === "Success") {
				$('#response-info-formation').html(`
					<div class="ui success message">
					<i class="close icon"></i>
					<div class="header">
					Edition ok
					</div>
					</div>
					`)
			} else {
				$('#response-info-formation').html(`
					<div class="ui negative message">
					<i class="close icon"></i>
					<div class="header">
					Une erreur s'est produite
					</div>
					<p>` + data.message + `
					</p></div>
					`)
			}
		}
		)
}
