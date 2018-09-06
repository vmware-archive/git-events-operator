package action

import "fmt"

func getHTMLEmail(name, filename, link string) string {
	return fmt.Sprintf(`
<html>
<style>
body { font-family: "Comic Sans MS", cursive, sans-serif; }
</style>
<h1>
Hi there %s,
</h1>

<p>
Thank you for creating a new event page (%s) for the Heptio advocacy site. Because you created a new page, we automatically generated a link for you.
</p>
</br>
<p>
Please use this link on social media, and during your presentation to share meta information about the event.
</p>

<br><br>
<i><a href="%s">%s</a>"</i>
<br><br>

<h2>Common links that folks add to their event page include:</h2>

<ul>
  <li>Your Twitter.</li>
  <li>Your LinkedIn.</li>
  <li>Your GitHub.</li>
  <li>Your Personal blog.</li>
  <li>Your slides from the event.</li>
  <li>A link to the event page about your talk.</li>
  <li>Links to any resources you mention: books, repositories, videos, etc</li>
</ul>
</br>
<p>We can now easily track clicks through your event! You will never be measured on the amount of clicks you get, we just use the data to show that advocacy works.
</p>

<p>
<center>
<i>
For more information please reach out to Heptio advocacy at <a href="mailto:team-advocacy@heptio.com">team-advocacy@heptio.com</a>
</i>
<img src="https://avatars2.githubusercontent.com/u/22035492?s=200&v=4">
</p>
</html>`, name, filename, link, link)

}
