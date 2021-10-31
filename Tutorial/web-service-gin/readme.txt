1. Begin tracking the Gin module as a dependency.
$ go get .

2. From the command line in the directory containing main.go, run the code. 
Use a dot argument to mean “run code in the current directory.”
$ go run .

3. From a new command line window, use curl to make a request to your running web service.
$ curl http://localhost:8080/albums


The command should display the data you seeded the service with.
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
]