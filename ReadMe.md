
Tour of Go Notes

The aim of this repository is to document the [The Tour of Go](https://go.dev/tour/list)

For those who have not done any Go yet. 
You will need 

* Go 1.23.1
* [Delve installed](https://github.com/go-delve/delve)

The general thing is create example in a function, in a sensible namespace.
Then create a test to demonstrate it working. 
By convention the Golang test tool, will look for files of XXX_test.go
Then it will look for function signatures it recognises
Test_xxx(t *testing.\[B,T,M\])

  @Root of project
  go test ./...


---

For those who would like unit tests to run when they change

please install

* [watchexec](https://github.com/watchexec/watchexec)

  Then you can run this at the root

    watchexec       \
    -c clear      \
    -o do-nothing \
    -d 100ms      \
    --exts go     \
    'pkg=".${WATCHEXEC_COMMON_PATH/$PWD/}/..."; echo "running tests for $pkg"; go test "$pkg"'
