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
### Displaying one page
- For Visual Studio Code
  - Try adding extension: [Go Template Support]
- In GoLand (and possibly also Visual Studio Code)
  - Always replace .tmpl with .gohtml in template filenames. This will enable syntax highlighting and code completion. 
### Addition for Live reload
Several alternatives are available. Also, just using the `go run ./cmd/web/.` works fine
- Air - Live reload for Go apps
  [GitHub](https://github.com/air-verse/air)
  - Install and use:
    ```shell
    go install github.com/air-verse/air@latest
    ni .air.toml -type file
    ```
- Nodemon - Live reload for Go apps
  - Install Nodemon
    ```npm i -g nodemon```
  - Run Go file
    ```nodemon --exec go run ./main.go --signal SIGTERM```
### A better extension for Go templates and VS Code
- Rename all templates and references to them from `*.tmpl` to `*.gohtml`
- For Visual Studio Code
  - Remove extension: [Go Template Support]
  - Add extension: [goTemplate-syntax]
### Creating the form
### Connecting our form to stripe.js
### Client side validation
### Getting the paymentIntent - setting up the back end package
- Go Stripe - Go library for the Stripe API.
  [GitHub](https://github.com/stripe/stripe-go)
  - Install and use:
    ```shell
    go get -u github.com/stripe/stripe-go/v79
    md internal/card
    ni internal/card/card.go -type file -Value "package card`n`n"
    ```
- Sidenote: Old ruby library for the internationalization of the error codes from stripe, there are probably better solutions for go as well out there
  [GitHub](https://github.com/ekosz/stripe-i18n)
### Getting the paymentIntent - starting work on the back end api
- Chi CORS net/http middleware - a chi net/http compatible middleware for performing preflight CORS checks on the server side. [link](https://github.com/go-chi/cors)
    ```shell
    ni cmd/api/api.go -type file -Value "package main`n`nfunc main() {`n`n}"
    ni cmd/api/routes-api.go -type file -Value "package main`n`n"
    go get github.com/go-chi/cors
    ```
### Getting the paymentIntent - starting up a route and handler, and using make
- Create the api handler
    ```shell
    ni cmd/api/handlers-api.go -type file -Value "package main`n`n"
    ```
### Getting the paymentIntent - finishing up our handler
- Refactor our internal card package to avoid name conflicts 
### Updating the front end JavaScript to call our paymentIntent handler
### Getting the payment intent, and completing the transaction
- List of Stripe Test Cards: [Stripe Testing](https://docs.stripe.com/testing)
  - Look for suitable test cards, i.e. in the `Declined payments` section
### Generating a receipt
- Create the succeeded page template
    ```shell
    ni cmd/web/templates/succeeded.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Cleaning up the API url and Stripe Publishable Key on our form
- NOTE: temporarily putting test key in struct
### Quiz 1: Test your knowledge
- What is the purpose of Stripe's PaymentIntent?
  - A PaymentIntent transitions through multiple statuses throughout its lifetime as it interfaces with Stripe.js to 
    perform authentication flows and ultimately creates at most one successful charge.
- What is the Stripe Publishable key for?
  - It is used on public facing web pages to identify what Stripe account is associated with the 
    transactions that takes place.
- When dealing with amounts (such as prices, or totals, or something that is going to be processed as a currency 
  transaction), what is the best data type to store that amount in if you are using Go?
  - Use an int type

## Selling a Product online
### What are we going to build?
- A simple form that allows someone to buy a widget spinner
  (Built mostly in the front end) 
### Create the database
- Connect to the database server to create database and database user 
    ```mariadb
    CREATE DATABASE IF NOT EXISTS widgets;
    GRANT all ON widgets.* TO 'widgets'@'%' IDENTIFIED BY 'secret';
    ```
### Connecting to the database
- Go-MySQL-Driver - Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package.
  [GitHub](https://github.com/go-sql-driver/mysql)
- Create the driver package
    ```shell
    md internal/driver
    ni internal/driver/driver.go -type file -Value "package driver`n`n"
    go get github.com/go-sql-driver/mysql
    ```
### Creating a product page
- Adding files and folders
    ```shell
    md static
    ni cmd/web/templates/buy-once.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Creating the product form
- Sidenote: Added restart, restart-front and restart-back to Makefile to simplify restarting using make.
  Air is not serving the static content.
### Moving JavaScript to a reusable file
- Adding files and folders
    ```shell
    ni cmd/web/templates/stripe-js.partial.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Modifying the handler to take a struct
- Adding files and folders
    ```shell
    md internal/models
    ni internal/models/models.go -type file -Value "package models`n`n"
    ```
### Update the Widget page to use data passed to the template
### Creating a formatCurrency template function
### Testing the transaction functionality
### Creating a database table for items for sale
- Temporarily creating and populating a widgets table
    ```mariadb
    CREATE TABLE widgets(id int,name varchar(64));
    INSERT INTO widgets(id, name) VALUES(1,'Widget');
    ```
### Running database migrations
- Database Migrations - An intelligent means of managing the structure of our database
- Soda is part of the Buffalo framework (also called Pop)
- To install:  
  `go install github.com/gobuffalo/pop/v6/soda@latest`
- Remove temporarily created widgets table
    ```mariadb
    DROP TABLE widgets;
    ```
- Create a folder to hold migrations
    ```shell
    md migrations
    ```
- Add content to up and down files for the migration
- Run migration:
  `soda migrate`
- Revert migration:
  `soda migrate down`
### Creating database models
### Working on database functions
### Inserting a new transaction
### Inserting a new order
### An aside: fixing a problem with calculating the amount
### Getting more information about a transaction
### Customers
- Generate a migration for our customer table
  `soda generate fizz CreateCustomerTable`
  `soda generate fizz AddColsToTransactionsTable`
  `soda generate fizz AddCustomerIDToOrdersTable`
### Getting started saving customer and transaction information
- Adding stub home page
    ```shell
    ni cmd/web/templates/home.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
- SCS - HTTP Session Management for Go. [link](https://github.com/alexedwards/scs)
  - Setup:
    ```shell
    go get github.com/alexedwards/scs/v2
    ```
- Add middleware to handle loading and saving of sessions
    ```shell
    ni cmd/web/middleware.go -type file -Value "package main`n`n"
    ```
### Create the save customer database method
### Saving the customer, transaction and order from the handler
### Running a test transaction
### Fixing a database error, and saving more details
- Generate a migration for adding payment intent and method fields to our transaction table
  `soda generate fizz AddColsToTransactionTable`
### Redirecting after post
### Simplifying our PaymentSucceeded handler
### Revisiting our Virtual Terminal
- Adding a receipt page for virtual terminal
    ```shell
    ni cmd/web/templates/virtual-terminal-receipt.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Fixing a mistake in the formatCurrency template function

## Setting up and charging a recurring payment using Stripe Plans
### What are we going to build in this section?
- A simple form that allows someone to buy a monthly subscription for widget spinners
  (Built mostly in the back end this time)
### Creating a plan on the Stripe Dashboard
- To enable a subscription we need to set up a Stripe Plan
  - Log in to `Stripe Dashboard`
  - Go to `Product catalog` to `Add product`
    - Name: `Bronze Widget Plan`
    - Prize: `NOK 20.00`
    - Choose `Recurring`
    - Billing Period: `Monthly`
  - `Add product`
  - Make a note of it's ID
### Creating stubs for the front end page and handler
- Adding a template page for the bronze-plan
    ```shell
    ni cmd/web/templates/bronze-plan.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Setting up the form
- Adding a widget in the database
    ```mariadb
    INSERT INTO widgets(name, description, image, inventory_level, price, created_at, updated_at) VALUES('Bronze Plan','Get three widgets for the price of two every month','',100000,2000,now(), now());
    ```
- Generate a migration for adding plan_id and is_recurring fields to our widgets table
  `soda generate fizz AddPlanIDRecurringColsToWidgetsTable`
  `soda migrate`
- Update the plan_id column in widgets table with the correct plan_id 
### Working on the JavaScript for plans
- Create a Stripe Customer [link](https://docs.stripe.com/api/customers)
### Continuing with the JavaScript for subscribing to a plan
### Create a handler for the POST request after a user is subscribed
### Create methods to create a Stripe customer and subscribe to a plan
### Updating our handler to complete a subscription
- Changing plan_id stored in the database from Product ID for plan to Price ID
- Side note: later make sure stripe key and secret are properly set from config
### Saving transaction & customer information to the database 
### Saving transaction & customer information II 
### Displaying a receipt page for the Bronze Plan
- Adding a template page for the receipt page of bronze-plan
    ```shell
    ni cmd/web/templates/receipt-plan.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
## Authentication
### Introduction
How to ensure our users are valid
- How authentication works
  - Front end - `session Auth`
  - Back end - `tokens`
- Authentication types
  - HTTP Basic
  - Tokens
  - Stateful tokens
  - Stateless tokens (JWT)
  - API keys
  - OAuth 2.0
### Creating a login page
- Adding a template page for the login page
    ```shell
    ni cmd/web/templates/login.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Writing the stub JavaScript to authenticate against the back end
### Create a route and handler for authentication
- Add a helpers file to hold various helper functions
    ```shell
    ni cmd/api/helpers.go -type file -Value "package main`n`n"
    ```
### Create a writeJSON helper function
### Starting the authentication process
### Creating an invalidCredentials helper function
### Creating an passwordMatches helper function
- Package bcrypt implements bcrypt adaptive hashing algorithm
  - Install and use:
    ```shell
    go get golang.org/x/crypto/bcrypt
    ```
### Making sure that everything works
- sample user: admin@example.com:password
### Create a function to generate a token
- Add a file to hold token helper functions
    ```shell
    ni internal/models/tokens.go -type file -Value "package models`n`n"
    ```
### Generating and sending back a token
### Saving the token to the database
- Generate a migration for our tokens table
  `soda generate fizz CreateTokensTable`
### Saving the token to local storage
### Changing the login link based on authentication status
### Checking authentication on the back end
### A bit of housekeeping
### Creating stub functions to validate a token
### Extracting the token from the authorization header
### Validating the token on the back end
### Testing out our token validation
### Challenge: Checking for expiry
### Solution to challenge
- Generate a migration for adding expiry to our tokens table
  `soda generate fizz AddExpiryToTokensTable`
### Implementing middleware to protect specific routes
- Add middleware to handle protection of api routes
    ```shell
    ni cmd/api/middleware.go -type file -Value "package main`n`n"
    ```
### Trying out a protected route
### Converting the Virtual Terminal post to use the back end
### Changing the Virtual Terminal page to use fetch
### Verifying saved transaction

## Protecting routes on the Front End and improving authentication
### Writing middleware on the front end to check authentication
### Protecting routes on the front end
### Logging out from the front end
### Saving sessions in the database
- SCS Feature: Session Stores
  - By default, SCS uses an in-memory store for session data. This is convenient (no setup!) and very fast, but all 
    session data will be lost when your application is stopped or restarted.  In most production applications you will 
    want to use a persistent session store instead. SCS currently supports most major different DBMSs like MySQL, 
    PostgresSql, Consul, Etcd, Redis and more. And in addition you can make your own custom store.
    - [MySQL/MariaDB](https://github.com/alexedwards/scs/tree/master/mysqlstore)
    - [PostgresSQL](https://github.com/alexedwards/scs/tree/master/postgresstore)
    - ...
  - To install:  
    `go get github.com/alexedwards/scs/mysqlstore`
- Generate a migration for creating sessions table0
  `soda generate sql CreateSessionsTable`

## Mail and Password Resets
### Password resets
- Adding a template page for the forgot password page
    ```shell
    ni cmd/web/templates/forgot-password.page.gohtml -type file -Value "{{template `u{0022}base`u{0022} . }}`n`n{{define `u{0022}title`u{0022}}}`n`n{{end}}`n`n{{define `u{0022}content`u{0022}}}`n`n{{end}}"
    ```
### Sending mail Part I
- Go Simple Mail - Golang package for send email. Support keep alive connection, TLS and SSL. Easy for bulk SMTP.
  [GitHub](https://github.com/xhit/go-simple-mail)
- Import into project
  ```go get github.com/xhit/go-simple-mail/v2```
- Add mailer helper functions
    ```shell
    ni cmd/api/mailer.go -type file -Value "package main`n`n"
    md cmd/api/templates
    ni cmd/api/templates/password-reset.html.gohtml -type file -Value "{{define `u{0022}body`u{0022}}}`n`n{{end}}"
    ni cmd/api/templates/password-reset.plain.gohtml -type file -Value "{{define `u{0022}body`u{0022}}}`n`n{{end}}"
    ```
### MailTrap.io
- MailTrap Email Delivery Platform is the toolset to test, send, and control your emails in one place.
  [Link](https://mailtrap.io/)
- Sample Credentials
  - Host: `sandbox.smtp.mailtrap.io`
  - Port: `25, 465, 587 or 2525`
  - Username: `25853d08526311`
  - Password: `399982fbb4cbe9`
### Sending mail Part II
### Creating our mail templates and sending a test email
### Implementing signed links for our email message
- Go-alone - A simple to use, high-performance, Go (golang) MAC (Message authentication code) signer.
  [Link](https://github.com/bwmarrin/go-alone)
- Import into project
  ```go get github.com/bwmarrin/go-alone```
- Add urlSigner internal package
    ```shell
    md internal/urlSigner
    ni internal/urlSigner/signer.go -type file -Value "package urlSigner`n`n"
    ```
### Using our urlSigner package
### Creating the reset password route and handler





## Building Admin pages to manage purchases
## Refunds
## Cancelling Subscriptions
## Paginating Data
## Managing Users
## Microservices
## Validation
## Where to go next
