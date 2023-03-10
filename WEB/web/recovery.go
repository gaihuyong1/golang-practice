package web

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3,pcs[:])
	var sb strings.Builder
	sb.WriteString(message+"\nTrance:")
	for _,pc:=range pcs[:n]{
		fn:=runtime.FuncForPC(pc)
		file,line:=fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n\t%s:%d",file,line))
	}

	return sb.String()
}

func Recovery()HandlerFunc{
	return func(c *Context){
		defer func(){
			if err:=recover();err!=nil{
				message:=fmt.Sprintf("%s",err)
				log.Printf("%s\n\n",trace(message))
				c.Fail(http.StatusInternalServerError,"Internal Server Error")
			}
		}()
	}
}