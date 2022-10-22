# Example testing code

This is the exercise to download and handle with large data [tatoeba.org](https://tatoeba.org/en/downloads)

## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
go get
```

## Env
Copy file .env.example to .env

```go
# Port running Application
PORT=:3000
# Information of MysqlDB
SQL_URI=mysql://root:password@tcp(127.0.0.1:3306)/excercise
```

## Running
```go
# Script make 3 file sentences.tar.bz2, links.tar.bz2, sentences_with_audio.tar.bz2 to 1 csv file
go run makeFileCsv.go

# Script to import file csv to mysqlDB
go run seedCsvDB.go

# Script to run app
go run main.go
```

## HTTP API
```go
# Request
http://localhost:3000/api/translations/?page_number=2&page_size=10
# Response
[
  {
    "sentenceId": "1277",
    "text": "I have to go to sleep.",
    "vieSentenceId": "5662",
    "vieText": "Tôi phải đi ngủ.",
    "audioUrl": "http://www.manythings.org/tatoeba/eng/1277.mp3"
  },
  {
    "sentenceId": "1280",
    "text": "Today is June 18th and it is Muiriel's birthday!",
    "vieSentenceId": "5665",
    "vieText": "Hôm nay là ngày 18 tháng sáu, và cũng là ngày sinh nhật của Muiriel!",
    "audioUrl": ""
  }
]
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.