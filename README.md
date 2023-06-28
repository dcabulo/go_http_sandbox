# RSS Agregator

This project is a server to store RSS feeds in a postgres DB, allow to follow feeds and store feeds from RSS its like a scraper for rss.

## How to run it

The first step is to clone this project and create _.env_ file in the root of the project.

- You need to have a postgres DB up with a rssagg DB created, for this you can run a docker postgres image or just download the postgres, create the postgres with a user and password you can remember for the next part
- Install go/gooose, after installing use the this command _goose -dir "./sql/schema/" postgres postgres://user:password@host:port/rssagg up_
- the user, password and host match the postgres credentials.
- then you can go build the project in the root.
- The last part is just run the executable './rssagg' and start hitting the endpoints.

That's it enjoy the project.
