# Contributing to hypersonic

## Code of Conduct

Please adhere to the [Go Community Code of Conduct](https://go.dev/conduct) when interacting with others in the project.

## How to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b my-feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin my-feature-branch`).
6. Create a new Pull Request.

## Run

```sh
go run cmd/server/main.go -d ~/Downloads/music
```

- graphql playground
  - [http://localhost:8080/hypersonic.v1graphql.MusicLibrary/playground](http://localhost:8080/hypersonic.v1graphql.MusicLibrary/playground)

## Testing

### Test code

```sh
go test ./... -parallel 10
```

## Issue Reporting

If you encounter a bug or have a feature request, please open an issue in the GitHub repository.
