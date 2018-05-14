# Lift

## Deploy to Heroku
Download and install [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli#download-and-install)
```shell
git clone https://github.com/teddykoker/lift.git && cd lift
heroku login
heroku create
heroku buildpacks:set heroku/go
heroku buildpacks:add --index 2 heroku/nodejs
heroku addons:create heroku-postgresql:hobby-dev
git push heroku master
heroku open
```

## Local Installation
Requires:
* [Postgres](https://www.postgresql.org/), Ubuntu: `sudo apt install postgresql`, Mac: `brew install postgresql`
* [node + npm](https://nodejs.org/en/), Ubuntu: `sudo sudo apt install nodejs npm`, Mac: `brew install node`
* [Go](https://golang.org/dl/), Ubuntu: `sudo apt install golang-go`, Mac: `brew install golang`
* [dep](https://golang.github.io/dep/), Ubuntu: `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`, Mac: `brew install dep`

```shell
mkdir -p ~/go/src && cd ~/go/src
git clone https://github.com/teddykoker/lift.git && cd lift
dep ensure # ensures all dependencies are installed
go build # builds backend
cd client
npm install # install client dependencies
npm start # runs developent server
npm run build # (optional) build production client
```


## Common Issues
*Cannot connect to postgres using ssl*
```shell
cd /usr/local/var/postgres # or wherever postgres config is on your machine
openssl req -new -x509 -days 365 -nodes -text -out server.crt \
  -keyout server.key -subj "/CN=localhost"
chmod og-rwx server.key
vim postgresql.conf
# Uncomment ssl = false, change to ssl = true
createdb
```