#include <stdio.h>
#include "EasyPusherAPI.h"

Easy_I32 GO_EasyPusher_Activate(char *license){
	return EasyPusher_Activate(license);
}
Easy_Pusher_Handle GO_EasyPusher_Create(){
	return EasyPusher_Create();
}
Easy_U32  GO_EasyPusher_Release(Easy_Pusher_Handle handle){
	return EasyPusher_Release(handle);
}
Easy_U32 Easy_APICALL GO_EasyPusher_SetEventCallback(Easy_Pusher_Handle handle,  EasyPusher_Callback callback, int id, void *userptr){
	return EasyPusher_SetEventCallback(handle,callback,id,userptr);
}
Easy_U32 GO_EasyPusher_StartStream(Easy_Pusher_Handle handle, char* serverAddr, Easy_U16 port, char* streamName, char *username, char *password, EASY_MEDIA_INFO_T*  pstruStreamInfo, Easy_U32 bufferKSize, Easy_Bool createlogfile ){
	return EasyPusher_StartStream(handle,serverAddr,port,streamName,username,password,pstruStreamInfo,bufferKSize,createlogfile);
}
Easy_U32 GO_EasyPusher_StopStream(Easy_Pusher_Handle handle){
	return EasyPusher_StopStream(handle);
}
Easy_U32 GO_EasyPusher_PushFrame(Easy_Pusher_Handle handle, EASY_AV_Frame* frame ){
	return EasyPusher_PushFrame(handle,frame);
}