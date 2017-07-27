# gurlin
### go-url-link

A server that minifies and serves said redirect urls written in go.

## Api
```
/{path}

redirect or 404
```

```
/api

root of api
```

```
/api/available/{src}

returns
{ error: true|false, msg}
```

```
POST /api/register

accepts {
  to: redirect destination
  from: optional requested url, if omitted generates one
}

returns
{ error: true|false}
```

