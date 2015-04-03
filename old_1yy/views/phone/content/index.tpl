  <section class="channel hot">
    <h1>
      <i class="icon"></i>
      <span class="txt">热播影视</span>
	</h1>
    <div class="list_box">
      <ul class="list-inner row2x">
	  $for _,video in hotVideos:
        <li>
          <div class="poster">
            <a href="/detail/${video.Id}.html" title="《${video.Title}》详情">
              <img class="img" data-src="$utils.PosterUrl(video.Poster)"/>
            </a>
          </div>
          <a class="title" href="/detail/${video.Id}.html" title="《${video.Title}》详情">$video.Title</a>
        </li>
      $end
	  </ul>
	</div>
  </section>
$for _, channel in channelVideos:
  <section class="channel $channel.Code">
    <h1>
    <a href="/genre/${channel.Code}/1.html">
      <i class="icon"></i>
      <span class="txt">$channel.Name</span>
    </a>
  </h1>
  <div class="list_box">
    <ul class="list-inner row2x">
      $for _,video in channel.Videos:
        <li>
          <div class="poster">
            <a href="/detail/${video.Id}.html" title="《${video.Title}》详情">
              <img class="img" data-src="$utils.PosterUrl(video.Poster)"/>
            </a>
          </div>
          <a class="title" href="/detail/${video.Id}.html" title="《${video.Title}》详情">$video.Title</a>
        </li>
      $end
      </ul>
      <a href="/genre/$channel.Code/1.html" class="more">查看更多$channel.Name</a>
    </div>
  </section>
$end
