# microservice-template
**Complete Over-engineered Golang Microservice Template**

## Features
- Golang Server
    - [x] Static compiling
    - [ ] Clean Architecture
    - [ ] Golang standard layout
    - [ ] Complete Mock testing
    - [x] Makefile
    - [x] Styling checks with GolangCI-Lint
    - [x] Dockerfile
    - [x] Echo Web Framework
    - [ ] Plain sql or GORM
        - [ ] Migration manager as golang-migrate
        - [ ] Embeded migrations with go-bindata
    - [x] Panic-Recovery
    - [x] Logrus
    - [ ] Documentation (Grindsome Docc)
    - [ ] Automatic Locales Internationalization
    - [ ] Graceful shutdown
    - [ ] Casbin RBAC Auth Controller
    - [ ] Load tester as Vegeta
- Security
    - [ ] AutoTLS with Let's Encrypt
    - [ ] CORS management
    - [ ] Services protected with Authorization
    - [ ] AWS Secret Manager for environmental variables
    - [ ] Different database users for admin, app and metrics
    - [ ] Gosec automatic checks
- Services
    - [x] Docker-Compose that inits all services
    - [x] Postgres
    - [x] PgAdmin4 (Note: don't use this in prod)
    - [x] Metabase
    - [x] Prometheus
    - [x] Jaeger
    - [x] Grafana
    - [ ] NewRelic
    - [x] Sentry (SaaS)
        - [ ] Sentry (Self-Hosted)
    - [ ] Celery or other distributed task system
    - [ ] Redis cache
    - [ ] Swagger
    - [ ] Weblate/Traduora (Self-Hosted)
    - [ ] Fossa
    - [ ] Helm charts for deployment
    - [ ] Nginx/Traefik for load balancing
    - [ ] Codecov or similar
    - [ ] Terraform plan

## Structure
```
.
├── config
├── connector
│   └── fortune
├── database
│   ├── migrations
│   └── repository
│       └── cookie
├── dependencies
├── logic
│   └── entity
│       └── cookie
├── server
│   ├── handler
│   │   └── cookie
│   └── middleware
└── utils
```
