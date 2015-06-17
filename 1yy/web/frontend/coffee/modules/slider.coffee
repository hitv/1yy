(->
	Slider = (opt) ->
		that = this
		touchX = 0
		moveX = 0
		offsetX = 0
		itemNum = 0
		sliderWidth = 0
		current = 0
		points = []
		timer = false
		hm = that.html()

		hm = hm.replace /\n|\r/g, ''
		hm = hm.replace />\s*</g, '><'
		hm = hm.replace /\s*</g, '<'
		that.html hm

		opt = opt or {}
		transformDuration = opt.transformDuration or 500
		autoPlayInterval = opt.autoPlayInterval or 3000
		opt.autoPlay = if typeof opt.autoPlay isnt 'undefined' then opt.autoPlay else true
		wrapper = that.find opt.wrapper or '.wrap'
		items = wrapper.find opt.itemSelector or '.item'
		pointsWrapper = that.find opt.pointsWrapper or '.points'

		resize = ->
			sliderWidth = that.width()
			items.width sliderWidth
			showSlider current
			return

		init = ->
			items.each () ->
				pointsWrapper.append '<span></span>'
				return

			itemNum = items.length

			points = pointsWrapper.find 'span'

			that.on 'touchstart', touchStart
			that.on 'touchmove', touchMove
			that.on 'touchend', touchEnd

			setTimer() if opt.autoPlay
			resize()
			return

		setTimer = ->
			timer = setInterval(showNext, autoPlayInterval)
			return

		setTransformX = (pixel, duration) ->
			wrapper.css {
				'-webkit-transition-duration': duration + 'ms',
				'-webkit-transform': 'translate3d(' + pixel + 'px,0,0)'
			}
			return

		setPointActive = (i) ->
			points.removeClass 'active'
			$(points.get(current)).addClass 'active'
			return

		showSlider = (i) ->
			if i < 0
				i = itemNum + i
			else if i is itemNum
				i = 0
			current = i % itemNum
			offsetX = -sliderWidth * current
			setTransformX offsetX, transformDuration

			setPointActive current

			that.trigger 'show', current
			return

		showNext = ->
			showSlider current + 1
			return

		showPrev = ->
			showSlider current - 1
			return

		touchStart = (e) ->
			moveX = 0
			moveY = 0
			# 获取第一次touch坐标值
			if typeof e.touches isnt 'undefined'
				touchX = e.touches[0].pageX
			else
				touchX = e.pageX

			# 清除定时
			clearInterval timer if timer isnt false
			return

		touchMove = (e) ->
			if typeof e.touches isnt 'undefined'
				moveX = e.touches[0].pageX - touchX
			else
				moveX = e.pageX - touchX

			if Math.abs(moveX) > 5
				setTransformX moveX+offsetX, transformDuration

			return

		touchEnd = (e) ->
			if Math.abs(moveX) > 10
				if moveX > 0
					showPrev()
				else if moveX < 0
					showNext()
			else
				item = items.get current
				setTransformX offsetX, transformDuration
				that.trigger 'click', item

			setTimer() if opt.autoPlay

			return

		init()

		$.extend that, {
			prev: showPrev,
			next: showNext,
			show: showSlider,
			pause: ->
				opt.autoPlay = false
				clearInterval timer if timer isnt false
				return

			play: ->
				opt.autoPlay = true
				setTimer()
				return

			resize: resize
		}
		return that

	$.extend $.fn, Slider: Slider
	return
)()
