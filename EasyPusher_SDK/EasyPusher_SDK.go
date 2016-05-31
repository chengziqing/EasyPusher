package main

import "github.com/chengziqing/EasyPusher"
import "fmt"
import (
	"unsafe"
)

func main() {
	i := EasyPusher.GO_EasyPusher_Activate("6A34714D6C3469576B5A7541662F5257715174355875314659584E355548567A6147567958314E455379356C6547572B56752B7141506A655A57467A65513D3D")
	var j = (*int32)(unsafe.Pointer(i.Swigcptr()))
	fmt.Println(*j)
	easy_Pusher_Handle := EasyPusher.GO_EasyPusher_Create()
	fmt.Println(easy_Pusher_Handle.Swigcptr())
}
