//
//  AppMain.h
//  GoTest
//
//  Created by Hajime Hoshi on 6/13/13.
//  Copyright (c) 2013 Hajime Hoshi. All rights reserved.
//

#ifndef GoTest_AppMain_h
#define GoTest_AppMain_h

int GoTest_AppMain(int argc, const char** argv);

typedef void(*GoTest_SendMessageToGoFunc)(const char* message);

void GoTest_SetSendMessageToGoFunc(GoTest_SendMessageToGoFunc func);
void GoTest_SendMessageToGo(const char* message);

void GoTest_SendMessageToUI(const char* message);

#endif
