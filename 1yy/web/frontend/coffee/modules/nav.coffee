(->
	NavPlugin = (itemsSel, selectorSel) ->
		that = this
		navItems = that.find itemsSel
		navSelectors = that.find selectorSel

		iScroll = (el) ->
			if el._iscroll_
				el._iscroll_.refresh
			else
				el._iscroll_ = new $.IScroll el, {
					mouseWheel: yes, 
					click: yes, 
					scrollbars: 'custom'
				}
				el._iscroll_.scrollToElement el, 0
			return

		active = (index) ->
			item = navItems.eq index
			selector = navSelectors.eq index

			item.addClass 'active' if item

			if selector
				selector.removeClass('hide');
				iScroll selector.get 0;
			return

		unactive = (index) ->
			item = navItems.eq index
			selector = navSelectors.eq index

			item?.removeClass 'active'
			selector?.addClass 'hide'
			return

		that.on 'click', '.item', (e) ->
			index = navItems.indexOf e.currentTarget
			if index != -1
				that.reset()
				that.active index if that.actived
			return

		$.extend that, {
			actived: (index) ->
				item = navItems.eq index
				item.hasClass 'active'
				return

			active: (index) ->
				active index
				that.trigger 'active', index
				return

			unactive: (index) ->
				unactive index
				that.trigger 'unactive', index
				return

			reset: ->
				navItems.each (i) ->
					unactive i
					return

				that.trigger 'reset'
				return
			}
		return that
	$.extend $.fn, NavPlugin: NavPlugin
	return
)()