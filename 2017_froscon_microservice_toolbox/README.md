
Microservices
====================================
## What Microservices mean to me

* Independent Components and Self Contained Systems
* Independent teams develop and ship services
* Vertical Architecture, following the bounding contextes of your domain
* Loose Coupling
* Develop for replacement (not for reuse)

My Default Architecture
====================================

![Default Architecture](images/general-architecture.png#width: 100%; margin-top: -50px; margin-left: 15px;)

Challenges
====================================
## Infrastructure
* Distributed Deployment
* Service Discovery
* Edge-Services and HTTP Routing
* Centralized Logging
* Health checks and monitoring

## Application
* UI Composition
* Decentralized login management
* Data replication
* A huge number of small services

Docker Deployment
====================================

![Docker](images/docker.png#margin-left: 50px;)

## Every service is dockerized!

* All services are handled the same way
* Minimal requirements for the host setup

Docker Swarm
====================================

Multiple orchestration solutions are out there ...

![Docker](images/docker-swarm.jpg#float: left; margin-right: 50px;)

## I prefer __Docker Swarm__, because:

Simple to setup and use

Sufficient for most usecases

Easy local development and deployment

DNS and Docker Networking
====================================

## Don't make things complicated

* No complicated service registry!
* No IPs in config-files!
* No special port magic!

## Swarm gives you
* Overlay networks
* DNS
* Load balancing
* Health checks

Swarm Demo
====================================

![Demo](images/demo.png)

HTTP Termination and Routing
====================================

HTTP Router (aka edge service)

![Routing](images/routing.png#width: 80%; display: block; margin-left: auto; margin-right: auto; margin-top: 80px;)


Routing with Caddyserver
====================================

Caddy is a modern and easy to use webserver, written in golang.

![Caddy](images/caddy.png#width: 80%; display: block; margin-left: auto; margin-right: auto; margin-top: 30px; margin-bottom: 30px;)

* Build in Let's Encrypt support
* Reverse proxy and load balancing
* JWT login and access control
* Git checkout
* Hugo static page generator and Markdown rendering
* Extendable with golang plugins
* HTTP2 server push

Caddy Demo
====================================

![Demo](images/demo.png)    

Routing with fabio
====================================

[https://github.com/fabiolb/fabio](fabio)

Zero-conf load balancing HTTP(S) and TCP router

![Fabio](images/fabio.png#width: 70%;)

* Fast and stable
* Services discovery via consul


Routing with Traefic
====================================

* Several Configuration backends
* Let’s Encrypt support
* Circuit breaker and retry mechanism

![Traefic Architecture](images/traefik-architecture.png#width: 80%; display: block; margin-left: auto; margin-right: auto; margin-top: 40px;)

Traefic Demo
====================================

![Demo](images/demo.png)

Logging with microservices
====================================

## Logs are data: Handle them as data!

* Services write logs to `stdout`/`stderr`
* Collection and log shipping by `dockerd`
* Storage and retrieval by
  * ELK: Elasticsearch-Logstash-Kibana
  * or EFK: Elasticsearch-Fluentd-Kibana
* Structured logs (json records)
* Unified logging convention, were possible
  * Access logs
  * Application logs
  * Call logs
  * Lifecycle events

Logging example
====================================

## application log

    {
     "@timestamp": "2016-04-14T08:47:56.123456789Z",
      "level": "ERROR",
      "type": "application",
      "message": "Could not delete user",
      "user_correlation_id": "550e8400e29b",
      "correlation_id": "446655440000",
      "foo": "bar",
      "baz": "fuu"
    }

Logging example 
====================================

## access log
    {
      "@timestamp": "2016-04-14T08:47:56.123456789Z",
      "level": "INFO",
      "type": "access",
      "message": "GET /template/customer.html",
      "user_correlation_id": "550e8400e29b",
      "correlation_id": "446655440000",
      "path": "/template/customer.html",
      "remote": "123.122.122.123",
      "url": "/ressource",
      "method": "DELETE",
      "status": "500",
      "size": 9000
    }
    
Logging example
====================================

## call log
 
    {
      "@timestamp": "2016-04-14T08:47:56.123456789Z",
      "level": "error",
      "type": "call",
      "message": "Could not delete user",
      "user_correlation_id": "550e8400e29b",
      "correlation_id": "446655440000",
      "url": "https://someservice/foo/bar",
      "method": "GET",
      "duration": 50,
    }

Logging example
====================================

## lifecycle log
    {
      "@timestamp" : "2017-08-17T21:47:46.678870238+02:00",
      "level": "info",
      "type" : "lifecycle",
      "event" : "start",
      "message" : "starting application: loginsrv",
      "LoginPath" : "/login",
      "SuccessURL" : "/",
      "LogoutURL" : "",
      "TextLogging" : false,
      "Host" : "localhost",
      "CookieName" : "jwt_token",
      ..
    }

Avoid the Monolyth!
====================================
![Murten Monolith](images/murten-monolith.jpg#width:1024px; margin-left: -20px;)

[Micro-Services.pdf](Micro-Services.pdf)

UI Composition
====================================

## Design goal:

* Services should provide their UI themselves
* UI composition by a generic reverse proxy


![UI Composition](images/ui-composition.png#display: block; margin-left: auto; margin-right: auto;)

UI Composition - solution
====================================

## Lib Compose
* [github.com/tarent/lib-compose](https://github.com/tarent/lib-compose)
* MIT License
* Golang library and caddyserver plugin

## Features
* Fragment based UI composition
* Merging of script and meta tags
* Caching of parsed fragments
* Deduplication of meta tags (e.g. `<title>`)
* Configuration of upstream and downstream headers
* Fallback content if services are not available

Lib Compose
====================================

## How it works
* Every service provides valid HTML pages
* The HTML can contain some special directives like
  * `<uic-fragment ..>` - define a fragment
  * `<uic-include ..>` - include a fragment
  * `<div uic-remove ..>` - Remove parts before rendering
* A request is split up into multiple (parallel) service calls
* All responses are merged into one HTML page

UI Composition - example
====================================

layout.html

    <html>
      <head>
        <title>Default title</title>
        <link rel="stylesheet" href="/static/head/layout.css">
      </head>
      <body>
        <h1 uic-remove>This is the basic layout</h1>
        <uic-include src="header" required="true"/>
        <div>
          <uic-include src="content#subheader" required="true"/>
        <div>
        <uic-include src="content" required="true"/>
        <uic-include src="footer" required="true"/>
      </body>
    </html>

UI Composition - example
====================================

service response for `content`

    <html>
      <head>
        <title>This is the content title</title>
      </head>
      <body>
        <div uic-remove>Navigation for local service testing ..</div>
        <uic-fragment name="subheader">
          The subheader
        </uic-fragment>
        <div>
          Some content bla ....
        <div>
      </body>
    </html>


UI Composition - alternative
====================================

Zalando has a component with similar functionality:

[Zalando Tailor](https://github.com/zalando/tailor)


Authentication with JWT
====================================

Classical sessions carry a lot of problems with them:

* Server side state
* Dependencies between services
* Sessions in a microservice architecture end up with a central session service, which is a bad design

__Solution:__ Use a client side crypto token

JSON Web Token (JWT) ist a standard ([rfc7519](https://tools.ietf.org/html/rfc7519)) for signed JSON tokens.

![JWT Logo](images/jwt-logo.svg#float: right; margin-right: 130px; margin-top: 20px; width: 160px;)

* JSON payload + signature
* Can be used as cookie or HTTP header
* Symetric and asymetric algorithms
* APIs for every language


Loginsrv
====================================

[Loginsrv](https://github.com/tarent/loginsrv) is a JWT login microservice.

__Can be used as:__

* Standalone microservice
* Socker container
* Golang library
* [Caddyserver](http://caddyserver.com/) plugin

__Backends:__

* Htpasswd file
* OSIAM Identity Management
* Simple (user/password pairs by configuration)
* Httpupstream
* Oauth2: Github, Google

Loginsrv
====================================

Customizable UI

![Loginsrv](images/loginsrv.png#display: block; margin-left: auto; margin-right: auto;)

Loginsrv
====================================

REST interface

    curl -i -H 'Content-Type: application/json' \
      --data '{"username": "bob", "password": "secret"}' http://127.0.0.1:6789/login

    HTTP/1.1 200 OK
    Content-Type: application/jwt
    Date: Mon, 14 Nov 2016 21:35:42 GMT
    Content-Length: 100

    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJib2IifQ.-51G5JQmpJleARHp8rIljBczPFanWT93d_N_7LQGUXU


Loginsrv Demo
====================================

![Demo](images/demo.png)

Data Replication
====================================

> For us service orientation means encapsulating the data with the business logic that operates on the data [..] there’s no data sharing among the services.
>
> -- Werner Vogels, CTO Amazon

But, how handle data Replication?

![Kafka](images/kafka-logo.png)

Kafka Usage
====================================
![Kafka](images/kafka-apis.png#height: 80%; display: block; margin-left: auto; margin-right: auto;)

Kafka Log
====================================
![Kafka](images/kafka-log-anatomy.png#height: 80%; display: block; margin-left: auto; margin-right: auto;)

Kafka Consumer
====================================
![Kafka](images/kafka-log-consumer.png#height: 60%; display: block; margin-left: auto; margin-right: auto;)

Kafka Consumer Groups
====================================
![Kafka](images/kafka-consumer-groups.png#height: 60%; display: block; margin-left: auto; margin-right: auto;)

Thank You!
====================================

Slides:

[https://smancke.github.com/talks/2017\_froscon\_microservice\_toolbox/](https://smancke.github.com/talks/2017\_froscon\_microservice\_toolbox/)

Examples & Markdown:

[https://github.com/smancke/talks/](https://github.com/smancke/talks/tree/gh-pages/talks/2017\_froscon\_microservice\_toolbox/)
