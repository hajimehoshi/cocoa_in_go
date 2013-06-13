//
//  AppMain.m
//  GoTest
//
//  Created by Hajime Hoshi on 6/13/13.
//  Copyright (c) 2013 Hajime Hoshi. All rights reserved.
//

#import "AppMain.h"

#import <Cocoa/Cocoa.h>

static GoTest_SendMessageToGoFunc sendMessageToGo_;

int GoTest_AppMain(int argc, const char** argv)
{
    return NSApplicationMain(argc, argv);
}

void GoTest_SetSendMessageToGoFunc(GoTest_SendMessageToGoFunc func)
{
    sendMessageToGo_ = func;
}

void GoTest_SendMessageToGo(const char* message)
{
    sendMessageToGo_(message);
}

void GoTest_SendMessageToUI(const char* message)
{
    // TODO: Use performeSelectorOnMainThread
    NSLog(@"%s", message);
}
