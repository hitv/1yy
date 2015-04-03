<!DOCTYPE>
<html>
  <head>
    <!-- $utils.Now() -->
    <meta charset="utf-8" />
    <meta name="viewport" content="initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0, user-scalable=no, width=device-width" />
    <meta name="apple-mobile-web-app-status-bar-style" content="black" />
    <meta name="apple-mobile-web-app-capable" content="yes"/>
    <meta name="format-detection" content="telephone=no"/>
    <meta name="layoutmode" content="standard" />
    <link id="favicon" href="$utils.StaticUrl("favicon.ico")" rel="icon" type="image/x-icon">
    $# development #$
    <link type="text/css" rel="stylesheet" href="/phone/css/1yy.css"></link>
    <script type="text/javascript" src="/phone/js/zepto.js"></script>
    <script type="text/javascript" src="/phone/js/1yy.js"></script>
    $# //development #$
    $# production # $
    <link type="text/css" rel="stylesheet" href="$utils.StaticUrl("phone/css/1yy.03020035.min.css")"></link>
    <script src="$utils.StaticUrl("phone/js/1yy.03020035.min.js")" type="text/javascript"></script>
    $ # //production #$
    $headTpl.Render(head)
    <script type="text/javascript">
    (function(i,s,o,g,r,a,m){
    i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
    })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
    ga('create', 'UA-42032903-2', '1yingyuan.com');
    ga('send', 'pageview');

    var _hmt = _hmt || [];
    (function() {
      var hm = document.createElement("script");
      hm.src = "//hm.baidu.com/hm.js?4a90194e549cd55e2feef47be7dc56e8";
      var s = document.getElementsByTagName("script")[0];
      s.parentNode.insertBefore(hm, s);
    })();
    function down(name){
      ga('send', 'event', 'download', 'apk', name);
    }
    </script>
  </head>
  <body>
      <header>
        <div class="logo">
          <a href="/">
            <img src="$utils.StaticUrl("phone/images/logo.png")" width="89" height="40"/>
          </a>
        </div>
        <form action="/search" class="search-form" method="get">
          <div class="search">
            <div class="search-warp">
              <div class="search-input">
                <input type="text" name="q" autocomplete="off" autocorrect="off" placeholder="$search.defaultKeyword" value="$search.keyword" id="search_input"/>
                <input type="hidden" name="p" value="1"/>
              </div>
            </div>
            <div class="search-button">
          <div id="search-cross"></div>
            <input type="submit" value=""></div>
          </div>
        </form>
      </header>
      <nav class="nav">
        $for _, item in nav.items:
          <div class="nav_item $if item.IsActive(nav.activeGenre):current${end}">
            <a class="link" href="$if item.Code=="index":/${else:}/genre/$item.Code/1.html${end}">$item.Name</a>
          </div>
        $end
      </nav>
      $if nav.activeItem:
        $if len(nav.activeItem.Children)>0:
            <nav class="sub">
                $for _, child in nav.activeItem.Children:
                <a class="link $if child.IsActive(nav.activeGenre):current${end}" href="/genre/$child.Code/1.html">$child.Name</a>
                $end
            </nav>
        $end
      $end
      <div id="content">
        $if browser=="OTHER":
        <div class="message">
          看本站爽片需要下载并安装UC浏览器<br/>
          <a class="down_btn" href="http://filestorage.qiniudn.com/UCBrowser_V9.5.2.394_android_pf145_bi800_(Build1402131914).apk" onclick="down('ucbrowser_android')">下载Android版</a>&nbsp;&nbsp;<a class="down_btn" href="http://filestorage.qiniudn.com/apk/UCBrowser_V9.3.2.356_IP4X_pf41_bi800_(Build13103114).ipa" onclick="down('ucbrowser_iphone')">下载iPhone版</a>
           </div>
           $end
        $contentTpl.Render(content)
      </div>
      <div class="share">
        <span class="txt">分享到:</span>
        <!-- Baidu Button BEGIN -->
        <div id="bdshare" class="bdshare_t bds_tools_32 get-codes-bdshare">
        <a class="bds_tsina"></a>
        <a class="bds_qzone"></a>
        <a class="bds_tqq"></a>
        <a class="bds_renren"></a>
        <a class="bds_t163"></a>
        <span class="bds_more"></span>
        <a class="shareCount"></a>
        </div>
        <script type="text/javascript" id="bdshare_js" data="type=tools&amp;uid=6509569" ></script>
        <script type="text/javascript" id="bdshell_js"></script>
        <script type="text/javascript">
        document.getElementById("bdshell_js").src = "http://bdimg.share.baidu.com/static/js/shell_v2.js?cdnversion=" + Math.ceil(new Date()/3600000)
        </script>
        <!-- Baidu Button END -->
        <div class="clear"></div>
        <!-- Baidu Button END -->
      </div>
      $if showFriendLink:
      <div class="friend_link">
        <h2 class="t1">友情链接</h2>
        <ul>
          <li><a href="http://tv.uc.cn/" target="_blank">UC视频</a></li>
          <li><a href="http://m.baidu.com/video?src=video" target="_blank">百度视频</a></li>
          <li><a href="http://m.iyiqi.com/" target="_blank">爱奇艺</a></li>
          <li><a href="http://m.tv.sohu.com/" target="_blank">搜狐</a></li>
          <li><a href="http://m1905.cn/" target="_blank">电影网</a></li>
          <li><a href="http://www.tudou.com/" target="_blank">土豆</a></li>
          <li><a href="http://3g.pptv.com/" target="_blank">PPTV</a></li>
          <li><a href="http://m.56.com/" target="_blank">56</a></li>
          <li><a href="http://m.ku6.com/" target="_blank">酷6</a></li>
          <li><a href="http://www.youku.com/" target="_blank">优酷</a></li>
        </ul>
      </div>
      $end
      <footer>
        <p>本网站影音源于网络，不提供内容存储，</p>
        <p>仅供测试和学习交流，请大家支持正版。</p>
        <p>Copyright © 2013 冀ICP备12002867号-9</p>
        <p><a href="/" title="1影院">『1影院』</a> 版权所有</p>
      </footer>
  </body>
</html>
