# log4go
Very useful log package,this log4go support compress base on google orig log4go,the compress log package name like *.zip. 

# Useage:

go get github.com/xianlezheng/log4go

# Demo
```go
package main
import (
	log "github.com/xianlezheng/log4go/log4go"
)

const (
	XML_FILE_PATH    = "./log4go.xml"
)

func Init(){
   log.LoadConfiguration(XML_FILE_PATH)
}

func main(){
   log.Info("%s RowsAffected is :%d", insertSql, count)
}
```

# Have fun
:)
