// Package go_hi is an autogenerated binder stub for package hi.
//
// File is generated by gobind. Do not edit.
package go_hi

import (
	"code.google.com/p/go.mobile/bind/seq"
	"code.google.com/p/go.mobile/example/hello/hi"
)

func proxy_Hello(out, in *seq.Buffer) {
	param_name := in.ReadUTF16()
	hi.Hello(param_name)
}

func init() {
	seq.Register("hi", 1, proxy_Hello)
}