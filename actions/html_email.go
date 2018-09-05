package action

import "fmt"

func GetHTMLEmail(name, filename, linkUrl string) string {

	return fmt.Sprintf(`
<html>
<font face="raleway">
<font size=12>
<p>
Hi %s,

Thanks for adding your event to the Advocacy website.
</p>
<p>
As promised, here is everything you need to be the best Heptio Advocate you can be!

You can check out your page <a href="%s">here</a>.
This is a trackable link to your event markdown page, %s.

Please use this link at the event and edit the markdown as much as you'd like.
It is common for folks to use their markdown page to share things like the following:

<ul>
<li>A link to your personal blog</li>
<li>A link to your slides</li>
<li>A link to your GitHub, Twitter, or Linkedin</li>
<li>A link to resources you mention in your presentation (GitHub repos, blogs, books, etc)</li>
</ul>
</p>
<p>
Thanks for helping us keep Heptio looking smart. Feel free to ping Kris Nova <a href="mailto:knova@heptio.com">knova@heptio.com</a> if you have any questions.
</p>
</font>
</html>
`, name, linkUrl, filename)

}
