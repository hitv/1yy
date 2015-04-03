<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <!-- $utils.Now() -->
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
    <link id="favicon" href="$utils.StaticUrl("favicon.ico")" rel="icon" type="image/x-icon" />
	$# development # $
    <link type="text/css" rel="stylesheet" href="/pc/css/1yy.min.css"></link>
    <script type="text/javascript" src="/pc/js/1yy.min.js"></script>
	$ # //development #$
	$# production #$
    <link type="text/css" rel="stylesheet" href="$utils.StaticUrl("pc/css/1yy.03020150.min.css")"></link>
    <script src="$utils.StaticUrl("pc/js/1yy.03020150.min.js")" type="text/javascript"></script>
	$# //production #$
	$headTpl.Render(head)
	<script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
      ga('create', 'UA-42032903-1', '1yingyuan.com');
      ga('send', 'pageview');
    </script>
    <script>
      var _hmt = _hmt || [];
      (function() {
        var hm = document.createElement("script");
        hm.src = "//hm.baidu.com/hm.js?4a90194e549cd55e2feef47be7dc56e8";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
      })();
    </script>
  </head>
  <body>
    <div class="header">
      <div class="info">
        <div class="w960">
          <div class="welcome">欢迎访问『<a href="/" title="1影院">1影院</a>』，本站及时收录百度影音和快播最新热门影片，并提供在线播放，好片要大家一起看！快来分享吧！</div>
          <!-- Baidu Button BEGIN -->
          <div id="bdshare" class="bdshare_t bds_tools get-codes-bdshare">
            <a class="bds_tsina"></a>
            <a class="bds_qzone"></a>
            <a class="bds_tqq"></a>
            <a class="bds_renren"></a>
            <a class="bds_t163"></a>
            <a class="bds_douban"></a>
            <a class="bds_tqf"></a>
            <span class="bds_more">更多</span>
            <a class="shareCount"></a>
          </div>
          <script type="text/javascript" id="bdshare_js" data="type=tools&amp;uid=6509569" ></script>
          <!-- Baidu Button END -->
        </div>
      </div>
      <div class="w960 top">
        <div id="logo">
          <a href="/" title="1影院">
            <img src="$utils.StaticUrl("pc/images/logo.png")" width="155" height="70" alt="1影院"/>
          </a>
        </div>
        <div id="search">
          <div class="form">
            <form id="searchForm" method="get" action="/search">
              <div class="s_l">
				<input type="text" class="keyword" name="q" id="keyword" autocomplete="off" autocorrect="off" placeholder="$search.defaultKeyword" title="请输入搜索关键字" value="$search.keyword"/>
                <input type="hidden" name="p" value="1"/>
              </div>
              <input type="submit" class="s_btn" id="dosearch" value=""/>
            </form>
          </div>
		  $if search.recKeywords:
            <div class="links">
              $for _, keyword in search.recKeywords:
              <a href="/search/q=$utils.QueryEscape(keyword)/1.html" title="点击查看《${keyword}》">$keyword</a>
              $end
            </div>
		  $end
        </div>
        <div class="s_rad">
          <ul>
            <li>今日更新影片：${count.todayUpdate}部</li>
            <li>总计收录影片：${count.total}部</li>
          </ul>
        </div>
        <div class="clear"></div>
      </div>
    </div>
    <div class="nav">
      $for navIndex, item in nav.items:
		$if len(item.Children) == 0:
		  <a class="nav_item $if item.IsActive(nav.activeGenre):on${end}" href="$if item.Code=="index":/${else:}/genre/$item.Code/1.html${end}" title="$item.Name">$item.Name</a>
		$end
		$for i, sub in item.Children:
		  $if i < 6:
            <a class="nav_item $if sub.IsActive(nav.activeGenre):on${end}" href="$if sub.Code=="index":/${else:}/genre/$sub.Code/1.html${end}" title="${sub.Name}">${sub.Name}$if item.Name=="电影":片${elif item.Name=="电视剧":}剧${end}</a>
		  $end
		$end
      $end
    </div>
    <div id="mainbg">
		$contentTpl.Render(content)
      <div id="declare">
        <p>© 本网站提供的最新电影资源均系收集于各大视频网站，本网站只提供web页面服务，并不提供影片资源存储，也不参与录制、上传。</p>
        <p>本站仅供测试和学习交流。请大家支持正版。</p>
        <p>Copyright&nbsp;©&nbsp;2012&nbsp;-&nbsp;2013&nbsp;『<a href="/">1影院</a>』&nbsp;&nbsp;冀ICP备12002867号-9</p>
      </div>
    </div>
	<!-- Baidu Button BEGIN -->
	<script type="text/javascript" id="bdshare_js" data="type=tools&amp;uid=6509569" ></script>
    <script type="text/javascript" id="bdshell_js"></script>
    <script type="text/javascript">
      document.getElementById("bdshell_js").src = "http://bdimg.share.baidu.com/static/js/shell_v2.js?cdnversion=" + Math.ceil(new Date()/3600000)
    </script>
    <!-- Baidu Button END -->
  </body>
</html>
