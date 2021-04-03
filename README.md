# prometheus
Backend project of code:phoenix.

## Introduction

This part will be released after project's achievement.

## Deploy requirement

- [Go](https://golang.org/) 1.16
- [Node.js](https://nodejs.org/) 14.15.4
- [MySQL](https://downloads.mysql.com/archives/installer/) 5.7

## Installation

### Clone frontend and backend repo

```shell
git clone https://github.com/TzeHimSung/prometheus.git
git clone https://github.com/TzeHimSung/phoenix-eye.git
```

### Set database configuration
create database `prometheus`
```mysql
CREATE DATABASE `prometheus` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
```

create `config.json` at prometheus root path
```json
{
  "username": "username",
  "password": "password",
  "network": "tcp",
  "server": "localhost",
  "port": 3306,
  "database": "prometheus"
}
```

### Build frontend project

```shell
npm run build
```

then move `dist` folder to prometheus root path

### Build backend project

```shell
go generate
go build
```

### Launch

#### Windows
```shell
# add -initdb tag when launching for first time
prometheus.exe -initdb -runserver
```

#### Linux

```shell
./prometheus -initdb -runserver
```

Open your browser and visit `http://localhost:8000/`.

## Model Development

See more detail in [tutorial](https://github.com/TzeHimSung/prometheus/blob/main/doc/devModelEN.md)
