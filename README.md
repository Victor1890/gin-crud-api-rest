# API REST Golang / GIN

## Why have I made the decision to choose this architecture?

- Improves the organization and structure of the code, which makes it easier to understand and maintain.

- Helps ensure system scalability and adaptability in the face of changing business or user requirements

- Reduce costs and increase efficiency in software development and maintenance.

---

# How to clone

Before cloning the repository, you need to have these technologies.

### Technologies

* [Postman](https://www.postman.com/downloads/)
* [Golang](https://go.dev/)
* [Gin](https://gin-gonic.com/)
* [Gorm](https://gorm.io/)
* [Cosmtrek/air](https://github.com/cosmtrek/air)
* [Postgrest](https://www.postgresql.org/download/)
* [VS Code](https://code.visualstudio.com/)

# Installation
### To install and run this project, you must follow these steps.
    

```bash
git clone https://github.com/Victor1890/gin-crud-api-rest
```

```bash
cd gin-crud-api-rest
```

```bash
go mod download
```

You need the `.env` file to be able to establish the connection from the database to the project.

```md
GIN_MODE=debug

DB_CONNECTION=postgresql://postgres:password@localhost:5432/database-name
```

And you also need [air](https://github.com/cosmtrek/air) to be able to run the project in `debug` mode.

```sh
go install github.com/cosmtrek/air@latest
```

Run the project
```sh
air
```

# Routes

### Get all products
```sh
curl --location --request GET 'http://localhost:3000/api/'
```

### Get one product
```sh
curl --location --request GET 'http://localhost:3000/api/1'
```

### Update one product
```sh
curl --location --request PATCH 'http://localhost:3000/api/1' \
--form 'name="Soap"' \
--form 'brand="Gorgeous"' \
--form 'price="209.09"' \
--form 'image=@"/C:/Users/Victor1890/Pictures/product-test.jpg"'
```

### Create one product
```sh
curl --location --request POST 'http://localhost:3000/api/' \
--form 'name="Bacon"' \
--form 'brand="Fantastic"' \
--form 'price="795.23"' \
--form 'image=@"/C:/Users/victo/Pictures/4k-wallpaper-black-and-white-bone-1270184.jpg"'
```

# Routes Code

```golang
router.GET("/", controllers.GetAllProductHandler)
router.GET("/:id", controllers.GetOneProductHandler)
router.POST("/", controllers.CreateProductHandler)
router.PATCH("/:id", controllers.UpdateOneProductHandler)
router.DELETE("/:id", controllers.DeleteOneProductHandler)
```

# Postman Collection

We have a postman collection in order to do tests with the requests already created [collection](collection/Products.postman_collection.json)