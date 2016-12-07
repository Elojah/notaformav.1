/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   list.js                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hdezier <hdezier@student.42.fr>            +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2016/10/13 11:32:25 by hdezier           #+#    #+#             */
/*   Updated: 2016/12/07 20:51:53 by hdezier          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

var loadingDiv = `
	<div class="ui segment">
		<div class="ui active centered inline loader"></div>
	</div>
		`

$(window).on('end-listing-list-comments', function(){
		$('.ui.rating.disable').rating('disable');
		$('.show-details-comment').click(ShowHideCommentDetails)
})
$(window).on('end-listing-list-organisms', function(){
		$('.ui.rating.disable').rating('disable');
})
$(window).on('end-listing-list-formations', function(){
		$('.ui.rating.disable').rating('disable');
})

function OnEndListing(id, state) {
	var evt = $.Event('end-listing-' + id);
	evt.state = state;
	$(window).trigger(evt);
}

function	refreshPagination(meta, id) {
	var maxPages = 5
	var current = 1 + Math.trunc(meta.offset / meta.limit)
	if (meta.total_pages < maxPages) {
		maxPages = meta.total_pages
	}
	const minOffset = 3;
	var startPagination;
	if (current < minOffset) {
		startPagination = 1
	} else if (current > meta.total_pages - minOffset) {
		startPagination = meta.total_pages - maxPages + 1
	} else {
		startPagination = current - minOffset + 1
	}

	var DOMElementList = $("#" + id)
	var DOMElement = $("#" + id + "-pagination")
	DOMElement.html('')
	for (var i = startPagination; i < startPagination + maxPages; ++i) {
		DOMElement.append($.data(DOMElementList[0], "renderPagination")(i, current))
	}
}

function	renderElems(data, id) {
	console.log(data, id)
	var DOMElement = $("#" + id)
	DOMElement.html('')
	if (data == null
		|| data.data == null
		|| data.meta == null) {
		return
	}
	data.data.forEach(function (elem) {
		DOMElement.append($.data(DOMElement[0], "renderElem")(elem))
	})
	refreshPagination(data.meta, id)
	OnEndListing(id, null)
}

function	refresh(id, params) {
	var DOMElement = $("#" + id)
	DOMElement.html(loadingDiv)
	$.get(
		$.data(DOMElement[0], "url"),
		params,
		function(data) {
			renderElems(data, id)
		}
	);
}

function	refreshOne(id, params, url, renderFn) {
	var DOMElement = $("#" + id)
	DOMElement.html(loadingDiv)
	$.get(
		url,
		params,
		function(data) {
			console.log(data)
			DOMElement.html('')
			if (data == null
				|| data.data == null
				|| data.meta == null) {
				return
			}
			DOMElement.html(renderFn(data))

			$('.ui.dropdown').dropdown();
			showFormationEdition()
			$('.menu .item').tab();
		}
	);
}

function	init_list(id, url, renderElem, renderPagination) {
	var DOMElement = $("#" + id)
	$.data(DOMElement[0], "url", url)
	$.data(DOMElement[0], "renderElem", renderElem)
	$.data(DOMElement[0], "renderPagination", renderPagination)
}

function	init_dropdown_nelem(id) {
	$('.ui.dropdown')
	  .dropdown({
	    action: function(text, value) {
			refresh(id, {page: 1, limit:value, sortOrder:null})
	    }
	  })
	;
}

function	test_match() {
	$.get(
		"www.notaforma.fr/search/match",
		{"romecode": ["H1203","H1210","H2504","H2908","H3203"]},
		function(data) {
			console.log("GOT IT")
			console.log(data)
		}
	);
}
