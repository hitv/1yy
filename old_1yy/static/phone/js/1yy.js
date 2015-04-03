//base64
var Base64 = (function () {
    // private property
    _keyStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";

    // private method for UTF-8 encoding
    function _utf8_encode(string) {
        string = string.replace(/\r\n/g,"\n");
        var utftext = "";
        for (var n = 0; n < string.length; n++) {
            var c = string.charCodeAt(n);
            if (c < 128) {
                utftext += String.fromCharCode(c);
            } else if((c > 127) && (c < 2048)) {
                utftext += String.fromCharCode((c >> 6) | 192);
                utftext += String.fromCharCode((c & 63) | 128);
            } else {
                utftext += String.fromCharCode((c >> 12) | 224);
                utftext += String.fromCharCode(((c >> 6) & 63) | 128);
                utftext += String.fromCharCode((c & 63) | 128);
            }

        }
        return utftext;
    }

    // private method for UTF-8 decoding
    function _utf8_decode(utftext) {
        var string = "";
        var i = 0;
        var c = c1 = c2 = 0;
        while ( i < utftext.length ) {
            c = utftext.charCodeAt(i);
            if (c < 128) {
                string += String.fromCharCode(c);
                i++;
            } else if((c > 191) && (c < 224)) {
                c2 = utftext.charCodeAt(i+1);
                string += String.fromCharCode(((c & 31) << 6) | (c2 & 63));
                i += 2;
            } else {
                c2 = utftext.charCodeAt(i+1);
                c3 = utftext.charCodeAt(i+2);
                string += String.fromCharCode(((c & 15) << 12) | ((c2 & 63) << 6) | (c3 & 63));
                i += 3;
            }
        }
        return string;
    }
    // public method for encoding
	return {
    		encode: function (input) {
	        var output = "";
	        var chr1, chr2, chr3, enc1, enc2, enc3, enc4;
	        var i = 0;
	        input = _utf8_encode(input);
	        while (i < input.length) {
	            chr1 = input.charCodeAt(i++);
	            chr2 = input.charCodeAt(i++);
	            chr3 = input.charCodeAt(i++);
	            enc1 = chr1 >> 2;
	            enc2 = ((chr1 & 3) << 4) | (chr2 >> 4);
	            enc3 = ((chr2 & 15) << 2) | (chr3 >> 6);
	            enc4 = chr3 & 63;
	            if (isNaN(chr2)) {
	                enc3 = enc4 = 64;
	            } else if (isNaN(chr3)) {
	                enc4 = 64;
	            }
	            output = output +
	            _keyStr.charAt(enc1) + _keyStr.charAt(enc2) +
	            _keyStr.charAt(enc3) + _keyStr.charAt(enc4);
	        }
	        return output;
	    },

	    // public method for decoding
	    decode: function (input) {
	        var output = "";
	        var chr1, chr2, chr3;
	        var enc1, enc2, enc3, enc4;
	        var i = 0;
	        input = input.replace(/[^A-Za-z0-9\+\/\=]/g, "");
	        while (i < input.length) {
	            enc1 = _keyStr.indexOf(input.charAt(i++));
	            enc2 = _keyStr.indexOf(input.charAt(i++));
	            enc3 = _keyStr.indexOf(input.charAt(i++));
	            enc4 = _keyStr.indexOf(input.charAt(i++));
	            chr1 = (enc1 << 2) | (enc2 >> 4);
	            chr2 = ((enc2 & 15) << 4) | (enc3 >> 2);
	            chr3 = ((enc3 & 3) << 6) | enc4;
	            output = output + String.fromCharCode(chr1);
	            if (enc3 != 64) {
	                output = output + String.fromCharCode(chr2);
	            }
	            if (enc4 != 64) {
	                output = output + String.fromCharCode(chr3);
	            }
	        }
	        output = _utf8_decode(output);
	        return output;
	    }
	};
})();
function setCookie(name, value, days){
    var exp = new Date();
	days = days || 1;
    exp.setTime(exp.getTime() + days * 86400000);
    document.cookie = name + "="+ escape(value) +";expires="+ exp.toGMTString();
}
function getCookie(name){
    var arr = document.cookie.match(new RegExp("(^| )"+name+"=([^;]*)(;|$$)"));
    if(arr != null){
        return unescape(arr[2]);
    }
    return null;
}
//1yy
!function(win, doc){
	function loadUcParam(fn){
		var ucparam = getCookie('ucparam');
        if(!ucparam){
           $.getJSON('http://hao.uc.cn/getucparam.php?uc_param_str=vednfrpfcpssne&callback=?',function(json){
				try{
		            setCookie('ucparam', JSON.stringify(json), 365);
				}
				catch(e){fn({})}
		        fn(json);
	        });
        } else {
			try{
               fn(JSON.parse(ucparam));
			} catch(e){fn({})}
        }
	}
	//替换.poster .img为iframe
	function loadPosterImg(imgs){
		imgs.each(function(i, img){
			var img = $(img),
			    inner = [],
				iframe = [],
				timer = null,
				src = img.data('src'),
				poster = img.parents('.poster'),
				win = window;
			inner.push('<html><head><style>body,img{width:100%;');
			inner.push('height:100%;margin:0;padding:0}</style></head>');
			inner.push('<body><img src="');
			inner.push('http://imgsize.ph.126.net/?enlarge=true&imgurl=');
			inner.push(src);
			inner.push('_');
			inner.push(win.poster.width);
			inner.push('x');
			inner.push(win.poster.height);
			inner.push('x1.jpg');
			inner.push('"/></body></html>');
			iframe.push('<iframe src="data:text/html;base64,');
			iframe.push(Base64.encode(inner.join('')));
			iframe.push('" class="img"></iframe>');
			iframe = $(iframe.join(''));
			img.hide();
			poster.append(iframe);
			iframe.on('load', function(){
				if(timer !== null){
			        clearTimeout(timer);
				}
			});
			timer = setTimeout(function(){
				img.attr('src', src);
				iframe.hide();
				poster.append(img.remove())
				img.show();
			},3000);
		});
	}
    function tplRender(tpl, data){
        return tpl.replace(/\{(\w*)\}/g, function(a, b) {
            return data[b];
        });
    }
    function tplListRender(tpl, data){
        var parts = [];
        for(var i = 0;i < data.length;++i){
            var html = tplRender(tpl, data[i]);
            parts.push(html);
        }
        return parts.join('');
    }
	$(function(){
		loadUcParam(function(ucparam){
			if(ucparam && ucparam.fr === 'android'){
				function addFav(){
					win.location.href="ext:add_favorite";
				}
			    if(win.autoAf && getCookie('auto_af') === null){
		            addFav();
					setCookie('auto_af', 'true', 1);
		        }
				$('#addFav').show().click(addFav);
			}
		});
		//图片懒加载
	    /*$("img.lazy").lazyload();*/

		//初始化海报图盗链加载
		loadPosterImg($('.poster img.img'));

		//初始化搜索框清空按钮
	    !function(){
	        var cross = $('#search-cross'),
	            input = $('#search_input');
	        cross.on('click', function(){
	            input.val('');
	            cross.hide();
	        });
	        function crossSwitch(){
	            setTimeout(function(){
	                if(input.val().length > 0){
	                    cross.show();
	                }
	                else{
	                    cross.hide();
	                }
	            }, 0);
	        }
			input.on('keydown', crossSwitch);
			input.on('click', crossSwitch);
		}();
		//初始化加载更多按钮
		!function(){
			var loadMore = $('#loadMore')
	        if(loadMore.length > 0){
	           var containerSelector = loadMore.data('container'),
	                container = $(containerSelector),
	                loadPlugins = loadMore.find('span'),
	                tpl = $('#item-tpl').html(),
	                pageSize = parseInt(loadMore.data('size')),
	                lock = false;
	            loadMore.click(function(e){
	                if(!lock){
	                    lock = true;
	                    var page = loadMore.data('page'),
	                        url = [pageSize, page].join('/') + '.json';
	                    if(page > 0){
	                        $(loadPlugins[0]).css('display', 'inline-block');
	                        $(loadPlugins[1]).html('正在加载...');

	                        $.getJSON(url, function(json){
	                            var html = '', msg = '';
	                            if(json.Status == 200){
								    json = json.Data || {};
									videos = json.Videos || []
	                                 html = $(tplListRender(tpl, videos));
									loadPosterImg(html.find('.poster img.img'));
	                                container.append(html);
	                                if(json.Total > json.Page * pageSize){
	                                    page = json.Page + 1;
	                                    msg = '加载更多';
										 loadMore.attr('href', page +'.html')
	                                }
	                                else{
	                                    page = 0;
	                                    msg = '已到最后一页';
	                                }
	                                loadMore.data('page', page);
	                            }
	                            else{
	                                msg = json.Message
	                            }
	                            $(loadPlugins[0]).hide();
	                            $(loadPlugins[1]).html(msg);
	                            lock = false;
	                        });
	                    }
	                }
	                e.preventDefault();
	            });
	        }
		}();
		//自动点击播放器控件
		!function(){
			var player = $('#player');
			if(player.length == 0) return;
			player = player[0];
			function play(e){
				var click = doc.createEvent("MouseEvent"),
				    touchend = doc.createEvent("TouchEvent");
		    		player.focus();
	        		click.initEvent("click", true, true);
	        		touchend.initEvent("touchend", true, true);
		    		player.dispatchEvent(click);
		    		player.dispatchEvent(touchend);
			}
			$('section.play .btn').click(play);
			play();
		}();
	});
}(window, document);
