<script id="item-tpl" type="text/template">
  <li>
    <div class="poster">
      <a href="/detail/{Id}.html" title="《{Title}》详情">
        <img class="img" data-src="$utils.PosterUrl("{Poster}")"/>
      </a>
    </div>
    <div class="detail">
      <h2 class="title">
        <a href="/detail/{Id}.html" title="《{Title}》详情">{Title}</a>
      </h2>
    </div>
  </li>
</script>
<div class="list_box">
  <ul class="list-inner" id="list">
  $for _, video in videos:
      <li class="item">
        <div class="poster">
          <a href="/detail/${video.Id}.html" title="《${video.Title}》详情">
            <img class="img" data-src="$utils.PosterUrl(video.Poster)"/>
          </a>
        </div>
        <div class="detail">
          <h2 class="title">
            <a href="/detail/${video.Id}.html" title="《${video.Title}》详情">$video.Title</a>
          </h2>
        </div>
      </li>
  $end
  </ul>
  <a href="${pager.nextPage}.html" class="more" id="loadMore" data-size="${pager.pageSize}" data-page="${if pager.hasNextPage:}2${else:}0${end}" data-container="#list">
    <span class="loading">&nbsp;</span>
    <span class="text">${if pager.hasNextPage:}加载更多${else:}没有更多结果${end}</span>
  </a>
</div>
