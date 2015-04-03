<script type="text/javascript">
  ga('send', 'event', 'play', '$video.Title', '$currentUrl.Type', $currentUrl.Sn);
</script>
<div class="detail_box">
  <div class="t1">
    <a class="back" href="/detail/${video.Id}.html">返回</a>
	$video.Title - $currentUrl.Title
  </div>
  <section class="play $currentUrl.Type">
    $if currentUrl.Type == "bdhd":
	<a class="btn" href="javascript:void(0)"></a>
        <embed id="player" type="application/x-bdvideoplugin" width="1" height="1" uc_plugin_title="$currentUrl.Title" uc_plugin_url="$currentUrl.Url"></embed>
	$elif currentUrl.Type == "qvod":
	<object id="player" name="QvodPlayer" width="100%" height="100%" classid="clsid:F3D0D36F-23F8-4682-A195-74C92B03D4AF">
	  <param name="URL" value="${currentUrl.Url}"/>
	  <param name="Autoplay" value="1"/>
          <param name="Showcontrol" value="1"/>
	  <embed width="100%" height="100%" url="${currentUrl.Url}" type="application/qvod-plugin" autoplay="1" showcontrol="1"/>
	</object>
	$end
  </section>
  <section class="tip">
    <div class="bdlikebutton"></div>
  </section>
  <section class="tip">
    <a href="http://filestorage.qiniudn.com/ninegame_v1.3.2.apk" onclick="down('play/ninegame')" class="hehelink android">[打美女屁屁]好痛！官人轻点哦！</a>
    $#<div id="cm_01">
      <script type="text/javascript">
        $$(function(){function a(a){try{-1!==a.indexOf("移动")&&(ga("send","event","show_cm_ad","wuyeyouhuo"),$$("#cm_01").html('<a onclick="down(\'play/wuyeyouhuo\')" href="http://filestorage.qiniudn.com/night.apk" class="hehelink android">[午夜诱惑]是男人都无法抗拒</a>'))}catch(b){}}var c,b=window.localStorage.getItem("ucparam");if(b)try{c=JSON.parse(b),a(c.cp)}catch(d){}else $$.getJSON("http://hao.uc.cn/getucparam.php?uc_param_str=vednfrpfcpssne&callback=?",function(b){window.localStorage.setItem("ucparam",JSON.stringify(b)),a(b.cp)})});
      </script>
    </div>
    <a href="http://www.orz-j.com" onclick="ga('send', 'event', 'ad', 'orz-j')" class="hehelink">[美图]巨乳美腿清纯总有一款合你口味！</a>#$
  </section>
  <section class="links">
  $for _,typeUrl in typeUrlList:
    <dl class="$typeUrl.Code">
      <dt>
        <i class="icon"></i> $typeUrl.Name
      </dt>
      <dd class="col5 playlinks">
        $for _,url in typeUrl.Slice:
        <a $if url.IsCurrent(currentUrl.Type,currentUrl.Sn):class="current" ${end}href="/play/${video.Id}-${url.Type}-${url.Sn}.html">$url.Title</a>
        $end
        <div class="clear"></div<>
      </dd>
    </dl>
  $end
  </section>
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
</div>
