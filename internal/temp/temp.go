package temp

const Tmpl string = `
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{ .RssTitle }}</title>
    <link>{{ .RssLink }}</link>
    <atom:link href="{{ .RssLink }}" rel="self" type="application/rss+xml" />
    <description>{{ .RssDesc }}</description>
    <pubDate>{{ .RssNow }}</pubDate>
		<generator>Ysoding</generator>
    <lastBuildDate>{{ .RssNow }}</lastBuildDate>

{{ range .Items }}
      <item>
        <title>{{ .Title }}</title>{{ if .Desc }}
        <description>{{ .Desc }}</description>{{ end }}{{ if .PubDate }}
        <pubDate>{{ .PubDate }}</pubDate>{{ end }}
        <guid>{{ .Link }}</guid>
        <link>{{ .Link }}</link>
      </item>
{{ end }}      

  </channel>
</rss>
`
