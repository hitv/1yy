(function() {
  var NavPlugin;
  NavPlugin = function(itemsSel, selectorSel) {
    var active, iScroll, navItems, navSelectors, that, unactive;
    that = this;
    navItems = that.find(itemsSel);
    navSelectors = that.find(selectorSel);
    iScroll = function(el) {
      if (el._iscroll_) {
        el._iscroll_.refresh;
      } else {
        el._iscroll_ = new $.IScroll(el, {
          mouseWheel: true,
          click: true,
          scrollbars: 'custom'
        });
        el._iscroll_.scrollToElement(el, 0);
      }
    };
    active = function(index) {
      var item, selector;
      item = navItems.eq(index);
      selector = navSelectors.eq(index);
      if (item) {
        item.addClass('active');
      }
      if (selector) {
        selector.removeClass('hide');
        iScroll(selector.get(0));
      }
    };
    unactive = function(index) {
      var item, selector;
      item = navItems.eq(index);
      selector = navSelectors.eq(index);
      if (item != null) {
        item.removeClass('active');
      }
      if (selector != null) {
        selector.addClass('hide');
      }
    };
    that.on('click', '.item', function(e) {
      var index;
      index = navItems.indexOf(e.currentTarget);
      if (index !== -1) {
        that.reset();
        if (that.actived) {
          that.active(index);
        }
      }
    });
    $.extend(that, {
      actived: function(index) {
        var item;
        item = navItems.eq(index);
        item.hasClass('active');
      },
      active: function(index) {
        active(index);
        that.trigger('active', index);
      },
      unactive: function(index) {
        unactive(index);
        that.trigger('unactive', index);
      },
      reset: function() {
        navItems.each(function(i) {
          unactive(i);
        });
        that.trigger('reset');
      }
    });
    return that;
  };
  $.extend($.fn, {
    NavPlugin: NavPlugin
  });
})();

(function() {
  var Overlay;
  return Overlay = function() {
    this.on('show', function() {
      overlay.removeClass('hide');
    });
    return this.on('hide', function() {
      overlay.addClass('hide');
    });
  };
})();

(function() {
  var Slider;
  Slider = function(opt) {
    var autoPlayInterval, current, hm, itemNum, itemWidth, items, moveX, offsetX, points, setPointActive, setTimer, setTransformX, showNext, showPrev, showSlider, that, timer, touchEnd, touchMove, touchStart, touchX, transformDuration, wrapper;
    that = this;
    hm = that.html();
    hm = hm.replace(/\n|\r/g, '');
    hm = hm.replace(/>\s*</g, '><');
    hm = hm.replace(/\s*</g, '<');
    that.html(hm);
    opt = opt || {};
    opt.autoPlay = typeof opt.autoPlay !== 'undefined' ? opt.autoPlay : true;
    wrapper = that.find(opt.wrapSelector || '.wrap');
    items = wrapper.find(opt.itemSelector || '.item');
    points = that.find(opt.pointsSelector || '.points>span');
    itemWidth = that.width();
    itemNum = items.length;
    timer = 0;
    items.width(itemWidth);
    touchX = 0;
    moveX = 0;
    offsetX = 0;
    current = 0;
    transformDuration = opt.transformDuration || 500;
    autoPlayInterval = opt.autoPlayInterval || 3000;
    setTimer = function() {
      return timer = setInterval(showNext, autoPlayInterval);
    };
    setTransformX = function(pixel, duration) {
      wrapper.css({
        '-webkit-transition-duration': duration + 'ms',
        '-webkit-transform': 'translate3d(' + pixel + 'px,0,0)'
      });
    };
    setPointActive = function(i) {
      var point;
      points.removeClass('active');
      point = $(points.get(current));
      point.addClass('active');
    };
    showSlider = function(i) {
      if (i < 0) {
        i = itemNum + i;
      } else if (i === itemNum) {
        i = 0;
      }
      current = i % itemNum;
      offsetX = -itemWidth * current;
      setTransformX(offsetX, transformDuration);
      setPointActive(current);
      that.trigger('show', current);
    };
    showNext = function() {
      showSlider(current + 1);
    };
    showPrev = function() {
      showSlider(current - 1);
    };
    touchStart = function(e) {
      var moveY;
      moveX = 0;
      moveY = 0;
      if (typeof e.touches !== 'undefined') {
        touchX = e.touches[0].pageX;
      } else {
        touchX = e.pageX;
      }
      clearInterval(timer);
    };
    touchMove = function(e) {
      if (typeof e.touches !== 'undefined') {
        moveX = e.touches[0].pageX - touchX;
      } else {
        moveX = e.pageX - touchX;
      }
      if (Math.abs(moveX) > 5) {
        setTransformX(moveX + offsetX, transformDuration);
      }
    };
    touchEnd = function(e) {
      var item;
      if (Math.abs(moveX) > 10) {
        if (moveX > 0) {
          showPrev();
        } else if (moveX < 0) {
          showNext();
        }
      } else {
        item = items.get(current);
        setTransformX(offsetX, transformDuration);
        that.trigger('click', item);
      }
      if (opt.autoPlay) {
        setTimer();
      }
    };
    that.on('touchstart', touchStart);
    that.on('touchmove', touchMove);
    that.on('touchend', touchEnd);
    showSlider(0);
    if (opt.autoPlay) {
      setTimer();
    }
    $.extend(that, {
      prev: showPrev,
      next: showNext,
      show: showSlider,
      pause: function() {
        opt.autoPlay = false;
        return clearInterval(timer);
      },
      play: function() {
        opt.autoPlay = true;
        return setTimer();
      }
    });
    return that;
  };
  $.extend($.fn, {
    Slider: Slider
  });
})();

(function() {
  var Tab;
  Tab = function(opt) {
    var contentEl, contentSel, navEl, navSel, select;
    navSel = opt.navSelector || '.tab-nav';
    contentSel = opt.contentSelector || '.tab-content';
    navEl = this.find(navSel);
    contentEl = this.find(contentSel);
    select = function(n) {
      var fn;
      fn = function() {
        var children;
        children = this.children();
        children.removeClass('active');
        $(children.get(n)).addClass('active');
      };
      fn.call(navEl);
      fn.call(contentEl);
    };
    navEl.children().each(function(i, el) {
      $(el).on('click', function(e) {
        select(i);
        e.preventDefault();
      });
    });
    $.extend(this, {
      select: select
    });
  };
  return $.extend($.fn, {
    Tab: Tab
  });
})();

(function() {
  $.InitNav = function() {
    var nav, overlay;
    overlay = $('#overlay');
    nav = $('nav.list-nav').NavPlugin('.cat>.item', '.selector');
    overlay.on('click', function(e) {
      nav.reset();
    });
    nav.on('active', function() {
      overlay.trigger('show');
    });
    nav.on('reset', function() {
      overlay.trigger('hide');
    });
  };
})();
