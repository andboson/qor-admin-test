# QOR example application

This is an example application to show and explain features of [QOR](http://getqor.com).

## Quick Started

```shell
# Get example app
$ go get -u github.com/andboson/qor-admin-test

# Setup database
$ mysql -uroot -p
mysql> CREATE DATABASE qor_example;

# Setup configs
$ cp configs/database.example.yml configs/database.yml
$ cp configs/smtp.example.yml configs/smtp.yml

# Run Application
$ go run main.go
```


## Admin Management Interface

[Qor Example admin configuration](https://go-cat/blob/master/config/admin/admin.go)

Online Demo Website: [demo.getqor.com/admin](http://demo.getqor.com/admin)

## RESTful API

[Qor Example API configuration](https://go-cat/blob/master/config/api/api.go)

Online Example APIs:

* Users: [http://demo.getqor.com/api/users.json](http://demo.getqor.com/api/users.json)

## License

Released under the MIT License.

[@QORSDK](https://twitter.com/qorsdk)
