# Building Web Applications with Go - Intermediate Level

## Introduction
### Introduction
Building web applications with Go - Intermediate Level
- What we'll cover
  - Multiple applications from a single code base
  - Building a front end in Go
  - Building a back end API in Go
  - Building a microservice
  - Credit Card transactions using [Stripe](https://www.stripe.com)
    - Processing
    - Refunding
    - Creating subscriptions
  - Authentication on front and back ends
  - Session auth with username/password
  - Stateful tokens for API authentications
  - Password resets
  - User management
  - Microservices
### A bit about me
### Mistakes, we all make them
### How to ask for help
1. Search online
2. Compare your code to mine
3. Look in Q&A
4. Ask me, and provide your code

## Setting up Our environment
### Installing Go
- Download Go [here](https://go.dev/dl/)
- WinGet
  - Install: ```winget install --id GoLang.Go```
  - Upgrade: ```winget upgrade --id GoLang.Go```
- Verify  
  ```go version```
### Installing an IDE
- Visual Studio Code
  - Install: ```winget install --id Microsoft.VisualStudioCode```
  - Upgrade: ```winget upgrade --id Microsoft.VisualStudioCode```
  - Add extensions:
    - [Go]
      - Also, press `shift+ctl` and search for `Go: Install/Update Tools`
        - Click on it, select all associated checkboxes and click OK to install them 
    - [goTemplate-syntax]
- GoLand
  - Install using toolbox: ```winget install --id JetBrains.Toolbox```
  - Install directly: ```winget install --id JetBrains.GoLand```
### Get a free Stripe account
- Go to [stripe.com](https://stripe.com)
- Click button [Start now] to create you account
### Installing make
- Install: ```winget install -e --id GnuWin32.Make```
### Installing MariaDB
- [Installing and Using MariaDB via Docker](https://mariadb.com/kb/en/installing-and-using-mariadb-via-docker/)
  ````
  cd docker
  docker-compose up mariadb -d
  docker-compose down mariadb
  `````
### Getting a database client
- HeidiSQL
  - [Website](https://www.heidisql.com/)
  - Install: ```winget install -e --id HeidiSQL.HeidiSQL```

## Building a Virtual credit card terminal
### What we're going to build
- A Virtual Terminal, only used locally for testing with live credit card numbers
### Setting up a (trivial) web application
- Chi - A lightweight, idiomatic and composable router for building Go HTTP services. [link](https://github.com/go-chi/chi)
- Create a new folder as the project folder for our stripe app and go into it
  ```shell
  md go-stripe
  cd go-stripe
  ```
- Then create a go module for it:  
  ```go mod init github.com/johnwr-response/golang-build-web-applications-intermediate-level/go-stripe```
- 
- Then create a go module for it:  
  ```shell
  md cmd/web
  md cmd/api
  md internal
  ni cmd/web/main.go -type file -Value "package main`n`nfunc main() {`n`n}"
  ni cmd/web/routes.go -type file -Value "package main`n`nfunc routes() {`n`n}"
  go get -u github.com/go-chi/chi/v5
  ```
- Start the app:
  ```go run ./cmd/web/.```
### Setting up routes and building  a render function
- Sidenote: Added hostInterface config and command line parameter to make server listen to specific interface.  
  As an added benefit during development, windows defender will not try to block this if set to localhost.  
  If left blank it will listen to all interfaces.






## Selling a Product online
## Setting up and charging a recurring payment using Stripe Plans
## Authentication
## Protecting routes on the Front End and improving authentication
## Mail and Password Resets
## Building Admin pages to manage purchases
## Refunds
## Cancelling Subscriptions
## Paginating Data
## Managing Users
## Microservices
## Validation
## Where to go next
