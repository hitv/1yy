(->
	Overlay = ->
		@on 'show', ->
			overlay.removeClass 'hide'
			return
			
		@on 'hide', ->
			overlay.addClass 'hide'
			return
)()