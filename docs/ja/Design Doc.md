# Design Doc: hypersonic

## Objective

<!--  -->

## Goal, Non goal

### Goal

<!--  -->

### Non goal

<!--  -->

## High Level Structure

- cmd
  - server
- internal
  - infrastructure
    - api
      - http
    - datasource
      - filesystem
      - inmemory
  - interface-adapter
    - handler
      - graphql
  - usecase
    - upload
    - stream
    - search
  - domain

## Open Issues

- [bogem/id3v2](github.com/bogem/id3v2) package does not support versions older than `id3v2.2`.

## References

- mp3
  - [ID3タグ](https://ja.wikipedia.org/wiki/ID3%E3%82%BF%E3%82%B0)
- m4a
  - [QuickTime File Format | Apple Developer Documentation](https://developer.apple.com/documentation/quicktime-file-format)
  - [QuickTime Tags (ItemList)](https://exiftool.org/TagNames/QuickTime.html#ItemList)
