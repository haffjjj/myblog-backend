# About this project

This is my first backend project using golang :D

## Installation

Before you run the application, make sure the application config is set (config/config.json)

Build

```bash
make build
```

start

```bash
make start
```

run

```bash
make run
```

## Endpoints

get tags
```bash
[GET] /tags
```
get postsGroups, start & limit is optional
```bash
[GET] /postsGroups?start=0&limit=100
```
get postsGroups by tag, start & limit is optional
```bash
[GET] /postsGroups/tag/:tag?start=0&limit=100
```
get posts by id
```bash
[GET] /posts/:id
```
get jwt token for credential, Content-Type: application/x-www-form-urlencoded, key username, password

```bash
[POST] /auth
```
store/create post (need auth)
```bash
[POST] /posts
```
update post  (need auth)
```bash
[PUT]  /posts/:id
```
delete post  (need auth)
```bash
[DELETE] /posts/:id
```