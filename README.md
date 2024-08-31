# Lister of Repos
This is a simple tool that lists the top repositories on GitHub for a given programming language, by stars. It is written in [Golang](https://go.dev/learn/).

## Usage

To fetch all dependencies, please clone the repository first, and then run
```
cd listerofrepos
go mod tidy
```
To fetch the top 18 repositories for a language, for example C++, please run
```
cd listerofrepos
go run . --lang cpp --n 18
```
At this time, the number of repositories must be between 1 and 1000. The output looks like this:
![output](https://github.com/user-attachments/assets/9bbc4b53-0f54-4325-8f43-a9f23d46fa1c)

To run unit tests, please run
```
cd listerofrepos
go test -v
```

## License

MIT
