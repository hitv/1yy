<script type="text/javascript">
  ga('send', 'event', 'detail', '$video.Channel', '$video.Title');
</script>
<div class="detail_box">
  <div class="t1">
    <a class="back" href="/genre/${channel.Code}/1.html">返回</a>
	$video.Title
  </div>
  <section class="intro">
    <div class="poster">
	  <a href="/detail/${video.Id}.html" title="《${video.Title}》详情">
	    <img class="img" data-src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报"/>
	  </a>
	</div>
	<div class="info">
	  $if len(video.Actors):<p>主演：<span class="txt">$video.Actors</span></p>$end
	  $if len(video.Directors):<p>导演：<span class="txt">$video.Directors</span></p>$end
	  $if len(video.Genres):<p>类型：<span class="txt">$video.Genres</span></p>$end
	  $if len(video.Area):<p>地区：<span class="txt">$video.Area</span></p>$end
	  $if video.Year:<p>年份：<span class="txt">$video.Year</span></p>$end
	  $if video.Duration:<p>时长：<span class="txt">$video.Duration</span></p>$end
	  <p><div class="bdlikebutton"></div></p>
	</div>
  </section>
  <section class="tip">
    <a href="http://filestorage.qiniudn.com/ninegame_v1.3.2.apk" onclick="down('detail/ninegame')" class="hehelink android">[打美女屁屁]好痛！官人轻点哦！</a>
    $#<div id="cm_01">
      <script type="text/javascript">
        $$(function(){function a(a){try{-1!==a.indexOf("移动")&&(ga("send","event","show_cm_ad","wuyeyouhuo"),$$("#cm_01").html('<a onclick="down(\'detail/wuyeyouhuo\')" href="http://filestorage.qiniudn.com/night.apk" class="hehelink android">[午夜诱惑]是男人都无力抗拒</a>'))}catch(b){}}var c,b=window.localStorage.getItem("ucparam");if(b)try{c=JSON.parse(b),a(c.cp)}catch(d){}else $$.getJSON("http://hao.uc.cn/getucparam.php?uc_param_str=vednfrpfcpssne&callback=?",function(b){window.localStorage.setItem("ucparam",JSON.stringify(b)),a(b.cp)})});
      </script>
    </div>
    <a href="http://www.orz-j.com" onclick="ga('send', 'event', 'ad', 'orz-j')" class="hehelink">巨乳美腿清纯总有一款合你口味！</a>#$
  </section>
  <section class="links">
	$for _,typeUrl in typeUrlList:
	<dl class="$typeUrl.Code">
	  <dt>
	    <i class="icon"></i> $typeUrl.Name
	  </dt>
	  <dd class="col5 playlinks">
	    $for _,url in typeUrl.Slice:
	    <a href="/play/${video.Id}-${url.Type}-${url.Sn}.html">$url.Title</a>
	    $end
	    <div class="clear"></div<>
	  </dd>
	</dl>
	$end
  </section>
  <section class="summary">
    $:video.Summary
  </section>
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
  <section class="links">
    <dl class="recommend">
	  <dt>其他推荐视频</dt>
	  <dd class="col3">
	    $for _,video in recVideos:
        <a target="_blank" href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">$video.Title</a>
	    $end
		<div class="clear"></div<>
	  </dd>
    </dl>
  </section>
</div>
