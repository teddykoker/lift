# Lift

## Install Postgres (Mac)
```bash
brew install postgresql
cd /usr/local/var/postgres
openssl req -new -x509 -days 365 -nodes -text -out server.crt \
  -keyout server.key -subj "/CN=localhost"
chmod og-rwx server.key
vim postgresql.conf
# Uncomment ssl = false, change to ssl = true
createdb
```