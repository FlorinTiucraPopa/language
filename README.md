First we need to install the required package:
go get github.com/rylans/getlang

As we cannot assure any versioning on "go get" we will install the dep utility.

Setting vendor using:
curl https://raw.githubusercontent.com/golang/dep/master/install.sh |sh

Initialization of dep:
dep init

Refresh the packages:
dep ensure