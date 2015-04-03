<div id="detail">
  <div class="box">
    <div class="box_t">
      <span class="txt">您的位置：</span>
      <a href="/" title="返回首页">首页</a> &gt;
      <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">《${video.Title}》详情</a> &gt;
      ${currentUrl.Title}&nbsp;在线播放
    </div>
    <div class="box_c">
      <div id="player">
	  $if currentUrl.Type == "bdhd":
        <script language="javascript">
          var BdPlayer = new Array();
          BdPlayer['time'] = 10;
          BdPlayer['buffer'] = '';
          BdPlayer['pause'] = '';
          BdPlayer['end'] = '';
          BdPlayer['width'] = 640;
          BdPlayer['height'] = 480;
          BdPlayer['showclient'] = 1;
          BdPlayer['url'] = '$currentUrl.Url';
          BdPlayer['nextcacheurl'] = '';
          BdPlayer['lastwebpage'] = '';
          BdPlayer['nextwebpage'] = '';
        </script>
        <script language="javascript" src="http://php.player.baidu.com/bdplayer/player.js" charset="utf-8"></script>
	  $elif currentUrl.Type == "qvod":
	    <object name="QvodPlayer" width="100%" height="100%" classid="clsid:F3D0D36F-23F8-4682-A195-74C92B03D4AF">
	      <param name="URL" value="$currentUrl.Url">
	      <param name="Autoplay" value="1">
	      <embed width="100%" height="100%" url="$currentUrl.Url" type="application/qvod-plugin" autoplay="1"/>
	    </object>
	  $end
      </div>
    </div>
  </div>
  <div class="box tip">
    【电影观看小贴士】&nbsp;&nbsp;[DVD:标准清晰版]&nbsp;&nbsp;[BD:1024高清版]&nbsp;&nbsp;[HD:1280高清版]&nbsp;&nbsp;[TS:抢先版，不清晰] - BD和HD版本不太适合网速过慢的用户观看
	  <div class="bdlikebutton"></div>
  </div>
  <div class="box">
    <div class="ch">
	$for _,typeUrl in typeUrlList:
      <div class="box_t $typeUrl.Code">
        <span class="txt">$typeUrl.Name</span>
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l">
            <div class="box_l playlinks">
              <ul>
                $for _,url in typeUrl.Slice:
                <li><a href="/play/${video.Id}-${url.Type}-${url.Sn}.html" title="《${video.Title}》${url.Title}在线播放">${url.Title}</a></li>
                $end
              </ul>
            </div>
          </div>
        </div>
      </div>
	$end
    </div>
  </div>
  <div class="box">
    <div class="box_t">
      <span class="txt">其他推荐的影片</span>
    </div>
    <div class="box_c thumb">
      <ul>
	    $for _,video in recVideos:
        <li>
          <a class="img" href="/detail/${video.Id}.html" target="_blank" title="《${video.Title}》详情">
            <img src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报" />
          </a>
            <p class="t_p">
              <a target="_blank" href="/detail/${video.Id}.html" title="《${video.Title}》详情">$video.Title</a>
            </p>
        </li>
		$end
      </ul>
    </div>
  </div>
  <div class="clear"></div>
  <script id="bdlike_shell"></script>
  <script>
    var bdShare_config = {
	  "type":"small",
	  "color":"red",
	  "uid":"6509569",
	  "likeText":"喜欢",
	  "likedText":"喜欢过"
  };
  document.getElementById("bdlike_shell").src="http://bdimg.share.baidu.com/static/js/like_shell.js?t=" + Math.ceil(new Date()/3600000);
  </script>
  <img src="http://122.10.87.242/stat/play/${video.Id}/${currentUrl.Type}/${currentUrl.Sn}.gif"/>
<div>
