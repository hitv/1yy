<div id="index">
  <div class="box">
    <div class="ch ch1" id="hot">
      <div class="box_t">
        <span class="txt">热播大片</span>
        <!-- <div class="more">
          <span class="page_txt"><em>1</em> / <em>2</em></span>
          <div class="page_btn">
            <a href="javascript:void(0)" onclick="prevPage()">&nbsp;</a>
            <a href="javascript:void(0)" onclick="nextPage()">&nbsp;</a>
          </div>
        </div> -->
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l thumb">
            <ul>
              $for _,video in hotVideos:
              <li>
			    <div class="poster">
                  <a href="/detail/${video.Id}.html" target="_blank" title="点击查看《${video.Title}》详情">
                    <img class="img" data-src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报" />
                  </a>
				</div>
               <p class="t_p">
                    <a target="_blank" href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">$video.Title</a>
                </p>
              </li>
              $end
            </ul>
          </div>
        </div>
      </div>
    </div>
    <div class="side" id="tip">
      <div class="box_t">
        <span class="txt">温馨提示</span>
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l">
            <div class="download">
              <p>『<a href="/" title="1影院">1影院</a>』采用基于P2P技术的快播播放器，高清视频任你在线观看！</p>
              <a href="http://dl.p2sp.baidu.com/BaiduPlayerAB.php" target="_blank" title="下载百度影音播放器">
                <img src="$utils.StaticUrl("pc/images/down_bdplayer.gif")" />
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  $for chIndex, channel in channelVideos:
  <div class="box">
    <div class="ch $if chIndex < 2:ch1${else:}ch2${end}">
      <div class="box_t">
        <span class="txt">$channel.Name</span>
        <span class="type">
        </span>
        <span class="more">
          <a href="/genre/$channel.Code/1.html" title="查看更多$channel.Name">更多&gt;&gt;</a>
        </span>
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l">
            <div class="box_l thumb">
              <ul>
                $for n,video in channel.Videos:
                  $if chIndex < 2:
                    $if n < 6:
                      <li>
					    <div class="poster">
                         <a href="/detail/${video.Id}.html" target="_blank" title="点击查看《${video.Title}》详情">
                           <img class="img" data-src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报" />
                         </a>
						</div>
                        <p class="t_p">
                          <a target="_blank" href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">$video.Title</a>
                        </p>
                      </li>
					$end
				  $else:
				    $if n < 4:
                      <li>
					    <div class="poster">
                         <a href="/detail/${video.Id}.html" target="_blank" title="点击查看《${video.Title}》详情">
                           <img class="img" data-src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报" />
                         </a>
						</div>
                        <p class="t_p">
                          <a target="_blank" href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">$video.Title</a>
                        </p>
                      </li>
					$end
				  $end
				$end
              </ul>
            </div>
            <div class="list">
              <ul>
                $for n,video in channel.Videos:
				   $if n >= 6:
                    <li>
                      $if chIndex < 2:<em>[$video.Genres]</em>${end}<a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">$video.Title</a>
                    </li>
				   $end
                $end
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
	$if rankVideos[channel.Name]:
	<div class="side">
      <div class="box_t">
        <span class="txt">${channel.Name}排行榜</span>
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l">
            <ul>
			  $for _,video in rankVideos[channel.Name]:
              <li>
                <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">$video.Title</a>
				<span class="play_num">$video.TotalPlayCount</span>
              </li>
			  $end
            </ul>
          </div>
        </div>
      </div>
    </div>
	$end
  </div>
  $end
  <div class="box">
    <div class="ch ch3" id="friendlink">
      <div class="box_t">
        <span class="txt">友情链接</span>
      </div>
      <div class="box_bg_c box_l">
        <div class="box_bg_l box_l">
          <div class="box_bg_r box_l">
            <div class="box_l links">
              <p class="link_text">
                <a href="http://www.baidu.com/" target="_blank">百度搜索</a>
                <a href="http://www.google.com/" target="_blank">Google搜索</a>
                $#<a href="http://www.xn--iorw51ad9b0v3f.com" target="_blank">友情链接</a>#$
                $#<a href="http://www.98xiaoshuo.com" target="_blank">98小说</a>#$
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="clear"></div>
</div>
