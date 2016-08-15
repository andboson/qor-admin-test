# QOR example application

This is an example application to show and explain features of [QOR](http://getqor.com).

## Quick Started

```shell
# Get example app
$ go get -u go-cat

# Setup database
$ mysql -uroot -p
mysql> CREATE DATABASE qor_example;

# Run Application
$ cd $GOPATH/src/go-cat
$ go run main.go
```

#### Generate sample data

```go
$ go run db/seeds/main.go
```

## Admin Management Interface

[Qor Example admin configuration](https://go-cat/blob/master/config/admin/admin.go)

Online Demo Website: [demo.getqor.com/admin](http://demo.getqor.com/admin)

## RESTful API

[Qor Example API configuration](https://go-cat/blob/master/config/api/api.go)

Online Example APIs:

* Users: [http://demo.getqor.com/api/users.json](http://demo.getqor.com/api/users.json)
* User 1: [http://demo.getqor.com/api/users/1.json](http://demo.getqor.com/api/users/1.json)
* Orders: [http://demo.getqor.com/api/orders.json](http://demo.getqor.com/api/orders.json)
* Products: [http://demo.getqor.com/api/products.json](http://demo.getqor.com/api/products.json)
* Product 1's ColorVariations [http://demo.getqor.com/api/products/1/color_variations.json](http://demo.getqor.com/api/products/1/color_variations.json)
* Product 1's ColorVariation 1 [http://demo.getqor.com/api/products/1/color_variations/1.json](http://demo.getqor.com/api/products/1/color_variations/1.json)

## License

Released under the MIT License.

[@QORSDK](https://twitter.com/qorsdk)
