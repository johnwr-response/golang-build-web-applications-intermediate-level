# golang-build-web-applications-intermediate-level
Building Web Applications with Go - Intermediate Level

## What you'll learn
- How to build a front end website using Go
- How to build a back end API using Go
- How to build multiple applications from a single code base
- How to build microservices in Go
- User authentication in Go
- API authentication using stateful tokens
- How to allow users to reset a password in a safe, secure manner
- How to integrate Stripe credit card processing with a Go back end
- Make one time or recurring payments with Stripe
- Best practices for making secure credit card transactions

## Requirements
- A basic understanding of the Go programming language
- A basic understanding of HTML, CSS, and JavaScript
- A basic understanding of SQL syntax

## Course content
- Introduction
- Setting up Our environment
- Building a Virtual credit card terminal
- Selling a Product online
- Setting up and charging a recurring payment using Stripe Plans
- Authentication
- Protecting routes on the Front End and improving authentication
- Mail and Password Resets
- Building Admin pages to manage purchases
- Refunds
- Cancelling Subscriptions
- Paginating Data
- Managing Users
- Microservices
- Validation
- Where to go next

## Description
This course is the followup to Building Modern Web Applications in Go. In this course, we go further than we did the first time around. We will build a sample E-Commerce application that consists of multiple, separate applications: a front end (which services content to the end user as web pages); a back end API (which is called by the front end as necessary), and a microservice that performs only one task, but performs it extremely well (dynamically building PDF invoices and sending them to customers as an email attachment).

The application will sell individual items, as well as allow users to purchase a monthly subscription. All credit card transactions will be processed through Stripe, which is arguably one of the most popular payment processing systems available today, and for good reason: developers love it. Stripe offers a rich API (application programming interface), and it is available in more than 35 countries around the world, and works with more than 135 currencies. Literally millions of organizations and businesses use Stripe’s software and APIs to accept payments, send payouts, and manage their businesses online with the Stripe dashboard. However, in many cases, developers want to be able to build a more customized solution, and not require end users to log in to both a web application and the Stripe dashboard. That is precisely the kind of thing that we will be covering in this course.

We will start with a simple Virtual Terminal, which can be used to process so-called "card not present" transactions. This will be a fully functional web application, built from the ground up on Go (sometimes referred to as Golang). The front end will be rendered using Go's rich html/template package, and authenticated users will be able to process credit card payments from a secure form, integrated with the Stripe API. In this section of the course, we will cover the following:
  - How to build a secure, production ready web application in Go 
  - How to capture the necessary information for a secure online credit card transaction 
  - How to call the Stripe API from a Go back end to create a paymentIntent (Stripe's object for authorizing and making a transaction)

Once we have that out of the way, we'll build a second web application in the next section of the course, consisting of a simple website that allows users to purchase a product, or purchase a monthly subscription. Again, this will be a web application built from the ground up in Go. In this section of the course, we'll cover the following:

  - How to allow users to purchase a single product 
  - How to allow users to purchase a recurring monthly subscription (a Stripe Plan)
  - How to handle cancellations and refunds 
  - How to save all transaction information to a database (for refunds, reporting, etc.) 
  - How to refund a transaction 
  - How to cancel a subscription 
  - How to secure access to the front end (via session authentication)
  - How to secure access to the back end API (using stateful tokens)
  - How to manage users (add/edit/delete)
  - How to allow users to reset their passwords safely and securely 
  - How to log a user out and cancel their account instantly, over websockets

Once this is complete, we'll start work on the microservice. A microservice is a particular approach to software development that has the basic premise of building very small applications that do one thing, but do it very well. A microservice does not care in the slightest about what application calls it; it is completely separate, and completely agnostic. We'll build a microservice that does the following:

  - Accepts a JSON payload describing an individual purchase 
  - Produces a PDF invoice with information from the JSON payload 
  - Creates an email to the customer, and attaches the PDF to it 
  - Sends the email

All of these components (front end, back end, and microservice) will be built using a single code base that produces  multiple binaries, using Gnu Make.

## Who this course is for:
- Developers who want to integrate Stripe into their applications
- Developers who want to learn how to build a back end API in Go
- Developers who want to learn best practices for building modern applications in Go (and JavaScript)

## Stripe app
- Folder is [go-stripe](go-stripe)
- Built in Go version 1.22.4
- Uses the chi router [link](https://github.com/go-chi/chi)
  - Also, the chi cors middleware [link](https://github.com/go-chi/cors)
- Uses the Go library for the Stripe API [link](https://github.com/stripe/stripe-go)
- Uses the Go MySQL Driver [link](https://github.com/go-sql-driver/mysql)
  - Indirectly uses: [filippo.io/edwards25519](filippo.io/edwards25519)
- Uses the bcrypt adaptive hashing algorithm [link](golang.org/x/crypto/bcrypt)
- Uses the SCS - HTTP Session Management for Go [link](https://github.com/alexedwards/scs)
  - Also, the session store for mysql [link](https://github.com/alexedwards/scs/tree/master/mysqlstore)


