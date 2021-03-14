# prometheus
Backend project of code:phoenix.

## Deploy requirement

- Go 1.16
- Node.js 14.15.4
- Mysql 5.7

## Installation

### Clone frontend and backend repo

```shell
git clone https://github.com/TzeHimSung/prometheus.git
git clone https://github.com/TzeHimSung/phoenix-eye.git
```

### Set database configuration
create `config.json` at prometheus root path
```json
{
  "username": "username",
  "password": "password",
  "network": "tcp",
  "server": "localhost",
  "port": 3306,
  "database": "databaseName"
}
```

### Build frontend project

```shell
npm run build
```

then move `dist` folder to prometheus root path

### Launch backend project

```shell
go generate
go build
```

launch the `prometheus` binary file and visit `http://localhost:8000/` 