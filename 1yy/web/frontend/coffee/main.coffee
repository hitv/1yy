(->
	$.InitNav = ->
		overlay = $ '#overlay'
		nav = $('nav.list-nav').NavPlugin('.cat>.item', '.selector')
		overlay.on 'click', (e) ->
			nav.reset()
			return

		nav.on 'active', ->
			overlay.trigger 'show'
			return

		nav.on 'reset', ->
			overlay.trigger 'hide'
			return

		return

	return
)()