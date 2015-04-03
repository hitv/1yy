<div id="detail">
  <div class="box">
    <div class="box_t">
      <span class="txt">您的位置：</span>
      <a href="/" title="首页">首页</a> &gt;
      <a href="/genre/$channel.Code/1.html" title="$channel.Name 列表">$channel.Name</a> &gt;
      <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">《${video.Title}》&nbsp;详情</a>
    </div>
    <div class="box_c">
	  <div class="poster">
        <a href="/detail/${video.Id}.html" target="_blank" title="点击查看《${video.Title}》详情">
          <img class="img" data-src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报"/>
		</a>
      </div>
      <div class="info">
        <ul>
          <li><h1>$video.Title</h1></li>
          $if len(video.Actors):<li>主演：<span class="txt">$:utils.ActorLink(video.Actors)</span></li>$end
	      $if len(video.Directors):<li>导演：<span class="txt">$video.Directors</span></li>$end
	      $if len(video.Genres):<li>类型：<span class="txt">$video.Genres</span></li>$end
	      $if len(video.Area):<li>地区：<span class="txt">$video.Area</span></li>$end
	      $if video.Year:<li>年份：<span class="txt">$video.Year</span></li>$end
          <li><div class="bdlikebutton"></div></li>
        </ul>
      </div>
    </div>
  </div>
  <div class="box tip">
    【电影观看小贴士】&nbsp;&nbsp;[DVD:标准清晰版]&nbsp;&nbsp;[BD:1024高清版]&nbsp;&nbsp;[HD:1280高清版]&nbsp;&nbsp;[TS:抢先版，不清晰] - BD和HD版本不太适合网速过慢的用户观看
  </div>
  <div class="box">
	$for _,typeUrl in typeUrlList:
    <div class="ch">
      <div class="box_t">
        <span class="txt">${typeUrl.Name}点播列表</span>
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l">
            <div class="box_l playlinks">
              <ul>
                $for _,url in typeUrl.Slice:
                <li>
				  <a href="/play/${video.Id}-${url.Type}-${url.Sn}.html" title="《${video.Title}》${url.Title} ${typeUrl.Name}在线播放">$url.Title</a>
				</li>
                $end
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
	$end
  </div>
  <div class="box">
    <div class="box_t">
      <span class="txt">剧情介绍</span>
    </div>
    <div class="box_c">
      <div class="desc">
        $video.Summary
      </div>
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
	  "likeText":"喜欢",
	  "likedText":"喜欢过",
	  "uid":"6509569"
    };
    document.getElementById("bdlike_shell").src="http://bdimg.share.baidu.com/static/js/like_shell.js?t=" + Math.ceil(new Date()/3600000);
  </script>
  <script src="http://wm.lrswl.com/page/?s=98579"></script>
<div>
