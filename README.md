# filecleaner

Clean a directory recursively of large assets:

(png|css|jpg|JPG|gif|GIF|doc|DOC|pdf|PDF|mov|mp4|mpg|wmv|3gp|swf|SWF|eot|svg|ttf|woff)



Installation
------------

```shell
go get github.com/techjacker/filecleaner
go install github.com/techjacker/filecleaner
```


Usage
-----

```shell
$ filecleaner

Cleaning dir: (/private/var/folders/2t/srny64c16n7836cdzv7zwv2m0000gp/T/filecleaner)

Success!

Before: 748 kb
Removed: 420 kb
Remaining: 328 kb
Deleted: 8 files
```


Tests
-----------------------------------------------------

```shell
go test
```



Find biggest files in dir + add to file removal regex
-----------------------------------------------------

```shell
find . -type f -size +1M -exec du {} \; | sort -n
```
