# Kejar Mimpi API

Provided API to Kejar Mimpi App

### Dependency
* Go Language (Golang)
* Setting GOPATH
* Beego Framework Golang
* Postgresql

Migration with Goose 

`https://bitbucket.org/liamstask/goose/`

### Running in local
* Clone this repo
* Replace the database connection code with your database connection in the `main.go` file
`orm.RegisterDataBase("default", "postgres", "user=postgres password=yourpassword host=127.0.0.1 dbname=yourdb")`

run
```
bee run -downdoc=true -gendoc=true
```
open your browser 
```
http://127.0.0.1:8080/swagger/
```
testing local ==> documentation video
```
https://www.youtube.com/watch?v=TCcldd7AnzY
```


## Built With

* [Beego](https://beego.me/docs/intro/) - The web frameworks used
* [Glide](https://glide.sh/) - Dependency Management


## Authors

* **Fany** - *Back end team B* -
* **Indra** - *Back end team B* -


## Acknowledgments

* Binar Academy
* Mentors backend: mas Prima, mas Gean, mas Andi
