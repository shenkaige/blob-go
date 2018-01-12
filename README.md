# BLOB ![build status](https://img.shields.io/travis/blob-go/blob-go.svg?style=for-the-badge)![iris](https://img.shields.io/badge/iris-powered-2196f3.svg?style=for-the-badge)![fast](https://img.shields.io/badge/fast-loading-2196f3.svg?style=for-the-badge)![responsive](https://img.shields.io/badge/responsive-ready-2196f3.svg?style=for-the-badge)
A fast, lightweight and extensible blogging platform, written in Go.
# TODO List
- [x] Basic Interface
- [x] Admin Interface
- [x] MVC structure
- [x] IRIS 10
- [x] Responsive
- [x] Easy Theme Structure
- [x] ORM Database
- [x] Cache support
- [ ] Install script
- [ ] Comment
- [ ] Plugin System
- [x] More, contribution welcomed!
# Features
__BLOB is a backend blogging platform with minimum CPU and RAM requirements.__

__BLOB is designed to be fast.__ It can handle 10'000 connections concurrently for an average load time of 40 ms on an Intel i7-6700HQ clocked at 1.80 GHz.

BLOB supports a variety of databases, including Sqlite, Mysql/MariaDB, MSSql, etc.

BLOB is written in Go, which makes it __work everywhere__, from server (GNU/Linux or Windows) to Raspberry Pi.
Everywhere Go runs, BLOB works.

BLOB is written using __IRIS MVC__, enabling user to tweak around.

BLOB uses __template__ design to make it easier to write a theme for it.

BLOB uses minimum Javascript for minor functions. It can currently work without Javascript.

BLOB works out-of-box.

BLOB will support __extension ecosystem__ in the near future. 
Extensions make blogging platform friendly for everyone.
# Compile
Make sure you have Go >=1.9. Go 1.9 can be downloaded from https://golang.org
> Building BLOB
>
> `go get -u github.com/blob-go/blob-go`

And that's done! This will also install IRIS V10(for HTTP, MVC and template) and BlackFriday(for Markdown) ahead.
# Bug Report
Please open an issue when you encounter an error or discover a bug. 
Pull request is welcomed!