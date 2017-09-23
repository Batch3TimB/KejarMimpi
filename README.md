# Kejar Mimpi API

Provided API to Kejar Mimpi App

### Dependency
* [Go Language](https://golang.org/) - (Golang)
* [GOPATH](https://golang.org/doc/code.html#GOPATH) - Setting GOPATH
* [Beego](https://beego.me/docs/intro/) - The web frameworks used
* [Glide](https://glide.sh/) - Dependency Management

### Migration with Goose 

* [GOOSE](https://bitbucket.org/liamstask/goose/) - Setting Goose

### Running in local
* Clone this repo
* Replace the database connection code with your database connection in the `main.go` file:

`orm.RegisterDataBase("default", "postgres", "user=postgres password=yourpassword host=127.0.0.1 dbname=yourdb")`

* then run
```
bee run -downdoc=true -gendoc=true
```

![Peek recording itself](https://github.com/Batch3TimB/kejarmimpi/blob/master/run%20local.gif)

open your browser 
```
http://127.0.0.1:8080/swagger/
```
## testing local

* [video ](https://www.youtube.com/watch?v=yXFrLz9gV5A&feature=youtu.be) - Vide documentation in youtube


## Authors

* **Fany** - *Back end team B* -
* **Indra** - *Back end team B* -


## Acknowledgments

* Binar Academy
* Mentors backend: mas Prima, mas Gean, mas Andi