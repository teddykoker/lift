# Lift

## Deploy to Heroku

Download and install [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli#download-and-install)
In order to use the monglolab addon you will likely have to verify credit card information, but you will not be charged if you use the free `mongolab:sandbox` option as shown below.

```shell
git clone https://github.com/teddykoker/lift.git && cd lift
heroku login
heroku create
heroku buildpacks:set heroku/go
heroku buildpacks:add --index 2 heroku/nodejs
heroku addons:create mongolab:sandbox
git push heroku master
heroku open
```

## Local Installation

Requires:

* [Mongodb](https://www.mongodb.com/), Ubuntu: [Installation Guide](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/#install-mongodb-community-edition), Mac: `brew install mongodb`
  * Create data directory with proper permissions:

```shell
sudo mkdir -p /data/db
sudo chown -R $USER /data/db
sudo chmod -R u+rw /data/db
```

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
