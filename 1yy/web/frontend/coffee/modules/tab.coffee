(->
	Tab = (opt)->
		navSel = opt.navSelector || '.tab-nav'
		contentSel = opt.contentSelector || '.tab-content'
		navEl = @.find navSel
		contentEl = @.find contentSel

		select = (n) ->
			fn = ()->
				children = @.children()
				children.removeClass 'active'
				$(children.get(n)).addClass 'active'
				return

			fn.call navEl
			fn.call contentEl
			return

		navEl.children().each (i, el)->
			$(el).on 'click', (e) ->
				select i
				e.preventDefault()
				return
			return

		$.extend @, select: select

		return

	$.extend $.fn, Tab: Tab
)()