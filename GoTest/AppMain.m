//
//  AppMain.m
//  GoTest
//
//  Created by Hajime Hoshi on 6/13/13.
//  Copyright (c) 2013 Hajime Hoshi. All rights reserved.
//

#import "AppMain.h"

#import <Cocoa/Cocoa.h>

static GoTest_SendMessageToGoFunc GoTest_SendMessageToGo_;

int GoTest_AppMain(int argc, const char** argv)
{
    return NSApplicationMain(argc, argv);
}

void GoTest_SetSendMessageToGoFunc(GoTest_SendMessageToGoFunc func)
{
    GoTest_SendMessageToGo_ = func;
}

void GoTest_SendMessageToGo(char* message)
{
    GoTest_SendMessageToGo_(message);
}
