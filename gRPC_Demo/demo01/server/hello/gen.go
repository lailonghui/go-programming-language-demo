/*
@Time : 2021/4/20 18:44
@Author : lai
@Description :
@File : gen
*/
package hello

//go:generate protoc --go_opt=paths=source_relative --go_out=plugins=grpc:.  *.grpc
