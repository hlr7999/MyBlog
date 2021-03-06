module MyBlog

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.2.8 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/valyala/fasttemplate v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190402142545-baf5eb976a8c
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190402142545-baf5eb976a8c
)
