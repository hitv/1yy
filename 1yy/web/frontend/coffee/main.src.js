$(function(){
	var NavPlugin = function(itemsSel, selectorSel){
		var that = this,
		    navItems = that.find(itemsSel),
			navSelectors = that.find(selectorSel);

		var iScroll = function(el) {
            if (el._iscroll_)
                el._iscroll_.refresh();
            else {
                el._iscroll_ = new $.IScroll(el, {
                	mouseWheel: true, 
                	click: true, 
                	scrollbars: 'custom'
                });
                el._iscroll_.scrollToElement(el, 0)
            }
        }, active = function(index){
			var item = navItems.eq(index),
				selector = navSelectors.eq(index);

			if(item){
				item.addClass('active');
			}

			if(selector){
				selector.removeClass('hide');
				iScroll(selector.get(0));
			}
		}, unactive = function(index){
			var item = navItems.eq(index),
				selector = navSelectors.eq(index);

			if(item){
				item.removeClass('active');
			}
			if(selector){
				selector.addClass('hide');
			}
		};

		that.on('click', '.item', function(e){
			var index = navItems.indexOf(e.currentTarget);
			if(index != -1){
				if(!that.actived(index)){
					that.reset();
					that.active(index);
				} else{
					that.reset();
				}
			}
		});

		$.extend(that, {
			actived: function(index){
				var item = navItems.eq(index);
				return item.hasClass('active');
			},
			active: function(index){
				active(index);
				that.trigger('active', index);
			},
			unactive: function(index){
				unactive(index);
				that.trigger('unactive', index);
			},
			reset: function(){
				navItems.each(function(i){
					unactive(i);
				});
				that.trigger('reset');
			}
		});

		return that;
	}

	$.extend($.fn, {
		NavPlugin : NavPlugin
	});
});

$(function(){
	var overlay = $('.overlay');

	overlay.on('show', function(){
		overlay.removeClass('hide');
	});

	overlay.on('hide', function(){
		overlay.addClass('hide');
	});

	var initNav = function(){
		var nav = $('nav.list-nav').NavPlugin('.cat>.item', '.selector');
		overlay.on('click', function(e){
			if(e.currentTarget != overlay){
				nav.reset();
			}
		});

		nav.on('active', function(){
			overlay.trigger('show');
		});

		nav.on('reset', function(){
			overlay.trigger('hide');
		});
	};

	$.InitNav = initNav;
});