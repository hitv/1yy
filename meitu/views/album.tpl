<!DOCTYPE HTML>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta http-equiv="Cache-Control" content="no-cache">
    <meta name="viewport" content="width=device-width,user-scalable=0,initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0">
    <meta name="MobileOptimized" content="320">
    <meta name="description" content="后入式orz-j">
    <title>后入式orz-j</title>
    <link rel="stylesheet" type="text/css" href="/css/mt.min.css?2112228" media="all">
    <script type="text/javascript" src="/js/mt.min.js?2112228"></script>
    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
      ga('create', 'UA-47880107-1', 'orz-j.com');
      ga('send', 'pageview');
    </script>
    <script type="text/javascript">
    $$(function(){
      var viewer = MtAlbumViewer();
      viewer.load($albumId);
    });
    </script>
  </head>
    <body>
      <section class="album">
        <header>
            <a class="back" href="javascript:history.back()">&nbsp;</a>
            <h2></h2>
            <div class="count"></div>
        </header>
        <script type="text/template" id="slider_item_tpl">
          <div class="item">
            <img class="pic"/>
          </div>
        </script>
        <script type="text/template" id="slider_recommend_tpl">
          <div class="recommend">
            <% for(var i = 0; i < this.Recommend.length; ++i){ %>
            <% (function(rec){ rec.Link = rec.Link ? rec.Link : (rec.Id + '.html'); %>
            <a class="cover" href="<%rec.Link%>">
              <img src="<%rec.Cover%>"/>
              <% if(!!rec.Title){ %>
              <div class="title"><h3><%rec.Title%></h3></div>
              <% } %>
            </a>
            <% })(this.Recommend[i])} %>
            <div class="clear"></div>
          </div>
	  <div id="cm_01"></div>
        </script>
        <script type="text/javascript">
          $$(function(){function a(a){try{-1!==a.cp.indexOf("移动")&&(a.fr=='android')&&(ga("send","event","show_cm_ad","wuyeyouhuo"),$$("#cm_01").html('<a onclick="ga(\'send\',\'event\',\'download',\'wuyeyouhuo\')" href="http://filestorage.qiniudn.com/night.apk" class="hehelink android">[午夜诱惑]是男人都无法抗拒</a>'))}catch(b){}}var c,b=window.localStorage.getItem("ucparam");if(b)try{c=JSON.parse(b),a(c)}catch(d){}else $$.getJSON("http://hao.uc.cn/getucparam.php?uc_param_str=vednfrpfcpssne&callback=?",function(b){window.localStorage.setItem("ucparam",JSON.stringify(b)),a(b)})});
        </script>
        <ul class="panel" id="slider">
            <li class="paper loading"></li>
            <li class="paper loading"></li>
            <li class="paper loading"></li>
        </ul>
        <a href="javascript:void(0)" class="play" id="play_btn">&nbsp;</a>
      </section>
    </body>
</html>
