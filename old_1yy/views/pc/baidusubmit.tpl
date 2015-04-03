<?xml version="1.0" encoding="UTF-8"?>
<urlset>
$for _,v in sitemapVideos:
    <url>
        <loc><![CDATA[http://www.1yingyuan.com/detail/${v.Id}.html]]></loc>
        <lastmod><![CDATA[${timeFmt.Format(v.ModifiedTime)}]]></lastmod>
        <data>
            <thread>
                <threadUrl><![CDATA[http://www.1yingyuan.com/detail/${v.Id}.html]]></threadUrl>
                <threadTitle><![CDATA[${v.Title}]]></threadTitle>
                <post>
                    <postContent><![CDATA[${v.Summary}]]></postContent>
                    <createdTime><![CDATA[${timeFmt.Format(v.CreateTime)}]]></createdTime>
                    <viewAuthority><![CDATA[Allow]]></viewAuthority>
                    <postSequenceNumber>1</postSequenceNumber>
                </post>
                <replyCount>${v.HourlyPlayCount}</replyCount>
                <viewCount>${v.TotalPlayCount}</viewCount>
                <lastReplyTime><![CDATA[${timeFmt.Format(v.ModifiedTime)}]]></lastReplyTime>
                <forumIn>
                    <forumName><![CDATA[1影院]]></forumName>
                </forumIn>
            </thread>
        </data>
    </url>
$end
</urlset>
