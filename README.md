# Go-Framework

> GWF `(Golang Web Framework)` require `iris`、`xorm`

1. `xdb` Used to database manager.
1. `xlog` Used to logger manager.

### Application directory

```text
.
└── sketch
   ├── app
   │   ├── controllers
   │   │   ├── index_controller.go
   │   │   └── example_controller.go
   │   ├── logics
   │   │   ├── index_index_logic.go
   │   │   ├── example_add_logic.go
   │   │   └── example_list_logic.go
   │   ├── middlewares
   │   │   └── example_middleware.go
   │   ├── models
   │   │   └── example.go
   │   └── services
   │       └── example_service.go
   └── config
       ├── app.yaml
       ├── logger.yaml
       └── service.yaml
```