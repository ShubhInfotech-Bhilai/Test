//go mod init github.com/ShubhInfotech-Bhilai/redis-connect
package main

import (
	"context"
	"fmt"
     //"encoding/json"
	"github.com/go-redis/redis/v8"
)
type Server struct{
ServerName string `json:"servername"`
Port int    `json:"port"`
 }
var ctx = context.Background()
func main(){
	client:=redis.NewClient(&redis.Options{
		     Addr: "127.0.0.1:6379",
			 Password: "",
			 DB:0,
            })
	/*client.Append(ctx,"mykey","10")
	client.Append(ctx,"mykey","11")
	client.Append(ctx,"mykey","12")
	client.Append(ctx,"mykey","13")*/
/*
json1,err:=json.Marshal(Server{ServerName:"mylocalhost"  ,Port:8091 })
if  err!=nil{
	fmt.Println(err)
}
json2,err:=json.Marshal(Server{ServerName:"mylocalhost"  ,Port:8091 })
if  err!=nil{
	fmt.Println(err)
}   
err=client.Set(ctx,"id001",json1,0).Err()
	if  err!=nil{
		fmt.Println(err)
	}
	err=client.Set(ctx,"id002",json2,0).Err()
	if  err!=nil{
		fmt.Println(err)
	}
*/
val,err:=client.Get(ctx,"id001").Result()
if  err!=nil{
	fmt.Println(err)
}
fmt.Println(val)

val,err=client.Get(ctx,"id002").Result()
if  err!=nil{
	fmt.Println(err)
}
fmt.Println(val)

}
