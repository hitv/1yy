<?xml version="1.0" encoding="UTF-8"?><?xml-stylesheet type="text/xsl" href="http://www.1yingyuan.com/static/sitemap.xsl"?>
<urlset xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc><![CDATA[http://www.1yingyuan.com/]]></loc>
        <lastmod><![CDATA[${timeFmt.Now()}]]></lastmod>
        <changefreq>daily</changefreq>
        <priority>1.0</priority>
    </url>
$for _,v in sitemapVideos:
    <url>
        <loc><![CDATA[http://www.1yingyuan.com/detail/${v.Id}.html]]></loc>
        <lastmod><![CDATA[${timeFmt.Format(v.ModifiedTime)}]]></lastmod>
        <changefreq>weekly</changefreq>
        <priority>0.6</priority>
    </url>
$end
</urlset>
