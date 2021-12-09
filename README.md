# BannerId library

This library contains various helpers to identify a software from it's TCP/HTTP banner.

See the `example` directory.

```
$ go run ./example/main.go
{
 "name": "Apache",
 "version": "2.4.18",
 "modules": null,
 "os": "Ubuntu"
}
{
 "name": "Apache",
 "version": "2.4.41",
 "modules": [
  {
   "name": "OpenSSL",
   "version": "1.0.2k-fips"
  },
  {
   "name": "PHP",
   "version": "5.6.40"
  },
  {
   "name": "mod_perl",
   "version": "2.0.7"
  },
  {
   "name": "Perl",
   "version": "v5.16.3"
  }
 ],
 "os": "Amazon"
}
```