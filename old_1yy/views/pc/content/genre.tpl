<div id="type">
  <div class="ch ch1">
    <div class="box_t">
      <span class="txt">您的位置：</span>
      <a href="/" title="首页">首页</a> &gt;
	  <a href="/genre/$activeNavItem.Code/1.html" title="点击查看 ${activeNavItem.Name} 列表">$activeNavItem.Name</a>
      $if genre.Code!=activeNavItem.Code: &gt; <a href="/genre/${genre.Code}/1.html" title="点击查看 ${genre.Name} 列表">${genre.Name}</a>${end} 第${Pager.CurrentPage}页
    </div>
    <div class="box_bg_c box_l">
      <div class="box_bg_l box_l">
        <div class="box_bg_r box_l">
        $for _, video in videos:
          <!--item{-->
          <div class="item">
		    <div class="poster">
              <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">
                <img class="img" data-src="$utils.PosterUrl(video.Poster)" alt="《${video.Title}》海报"/>
              </a>
		    </div>
			<div class="info">
              <ul>
                <li>
                  <h2>
                    <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">${video.Title}</a>
                  </h2>
                </li>
                 $if len(video.Actors):<li>主演：<span class="txt">$video.Actors</span></li>$end
	             $if len(video.Genres):<li>类型：<span class="txt">$video.Genres</span></li>$end
	             $if len(video.Area):<li>地区：<span class="txt">$video.Area</span></li>$end
	             $if video.Year:<li>年份：<span class="txt">$video.Year</span></li>$end
              </ul>
			  <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">
                <img src="$utils.StaticUrl("pc/images/play_btn.gif")" alt="点击查看《${video.Title}》详情"/>
              </a>
			</div>
          </div>
          <!--}item-->
        $end
        <!--page{-->
        <div class="page">
          共${pager.recordCount}部&nbsp;
          当前:${pager.currentPage}/${pager.pageCount}页&nbsp;
	      $if pager.hasPrevPage:
            <a href="1.html" title="跳至${genre.Name} 首页">首页</a>
            <a href="${pager.prevPage}.html" title="跳至${genre.Name} 第${pager.prevPage}页">上一页</a>
          $else:
            <em>首页</em>&nbsp;<em>上一页</em>&nbsp;
          $end
          $for _, i in pager.pageNumbers:
              $if i==pager.currentPage:
                <em class="current">${i}</em>
              $else:
                <a href="${i}.html" title="跳至${genre.Name} 第${i}页">$i</a>
              $end
		    $end
            $if pager.hasNextPage:
              &nbsp;<a href="${pager.nextPage}.html" title="跳至${genre.Name} 第${pager.nextPage}页">下一页</a>
              <a href="${pager.pageCount}.html" title="跳至${genre.Name} 尾页">尾页</a>
            $else:
              <em>下一页</em>
              <em>尾页</em>
            $end
          </div>
          <!--}page-->
        </div>
      </div>
    </div>
  </div>
  <div class="side">
    <div class="box_t">
      <span class="txt">热播$if genre.Code!=activeNavItem.Code:${genre.Name}${end}${activeNavItem.Name}</span>
    </div>
    <div class="box_bg_c box_l">
      <div class="box_bg_l box_l">
        <div class="box_bg_r box_l">
          <ul>
          $for _,video in hotVideos:
            <li>
              <a href="/detail/${video.Id}.html" title="点击查看《${video.Title}》详情">${video.Title}</a>
			  <span class="play_num">${video.TotalPlayCount}</span>
            </li>
          $end
          </ul>
        </div>
      </div>
    </div>
  </div>
  <div class="clear"></div>
</div>
