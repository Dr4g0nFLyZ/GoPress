# GoPress
A modern blogging solution

This is currently in the early stages. This is so far a zero-dependancy solution, so everything is written in pure Go. Application will run normally, and you can read/update/add/delete posts via a REST API. Work primarily needs to be done securing the admin page using some combination of password hashing and session cookies, as well as TLS support to encrypt uploaded data. Also potential bottlenecks need to be considered. For example I am using a simple JSON file for database persistence, but obviously a more comprehensive solution will be needed as the amount and complexity of content grows.

Test it out with ```go run .```
Access from your web browser at http://localhost:8080

I recommend you set your theme to dark or use a dark reader plugin for your browser when using this app.

The URL will bring you to the root page, which lists all the existing posts one after the other. The link element shows the title of the post, and underneath the link the content is displayed for only the first post. Clicking on the title brings you to a page where you can view the post by itself. Click Login at the bottom to access the admin section, where you can manage your posts.

More todo ..

Add some AJAX to reduce full page loads to a minimum
Improve the templating structure
Enhance the default theme

Source of demo text - loremipsum.io
