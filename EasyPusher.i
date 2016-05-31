/* EasyPusher.i */
 %module EasyPusher
 %{
   #include "EasyTypes.h"
   #include "EasyPusherAPI.h"
   extern Easy_I32 GO_EasyPusher_Activate(char *license);
   extern Easy_Pusher_Handle GO_EasyPusher_Create();
   extern Easy_U32 GO_EasyPusher_StartStream(Easy_Pusher_Handle handle, char* serverAddr, Easy_U16 port, char* streamName, char *username, char *password, EASY_MEDIA_INFO_T*  pstruStreamInfo, Easy_U32 bufferKSize, Easy_Bool createlogfile );
   extern Easy_U32 GO_EasyPusher_StopStream(Easy_Pusher_Handle handle);
   extern Easy_U32 GO_EasyPusher_PushFrame(Easy_Pusher_Handle handle, EASY_AV_Frame* frame );
%}
extern Easy_I32 GO_EasyPusher_Activate(char *license);
extern Easy_Pusher_Handle GO_EasyPusher_Create();
extern Easy_U32 GO_EasyPusher_StartStream(Easy_Pusher_Handle handle, char* serverAddr, Easy_U16 port, char* streamName, char *username, char *password, EASY_MEDIA_INFO_T*  pstruStreamInfo, Easy_U32 bufferKSize, Easy_Bool createlogfile );
extern Easy_U32 GO_EasyPusher_StopStream(Easy_Pusher_Handle handle);
extern Easy_U32 GO_EasyPusher_PushFrame(Easy_Pusher_Handle handle, EASY_AV_Frame* frame );