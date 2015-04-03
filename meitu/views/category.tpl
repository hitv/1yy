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
      var viewer = MtCatViewer();
      viewer.load($categoryId);
    });
    </script>
  </head>
    <body>
      <section class="category">
        <header>
          <div class="type">
            <h2>美女图片</h2>
          </div>
          <nav>
          $for _,category in categories:
            <a class="link ${if category.Id==categoryId:}current${end}" href="/category/${category.Id}.html">$category.Title</a>
          $end
          </nav>
        </header>
        <div class="panel" id="waterflow">
          <script type="text/template" id="cat_item_tpl">
            <a class="item cover" href="/album/<%this.Id%>.html">
              <img data-src="<%this.Cover%>" data-width="<%this.Width%>" data-height="<%this.Height%>"/>
              <% if(!!this.Title){ %>
              <h3><%this.Title%></h3>
              <% } %>
              <div class="count"><%this.ImageCount%></div>
            </a>
          </script>
        </div>
        <div id="tip"></div>
      </section>
    </body>
</html>
