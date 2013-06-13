//
//  AppMain.m
//  GoTest
//
//  Created by Hajime Hoshi on 6/13/13.
//  Copyright (c) 2013 Hajime Hoshi. All rights reserved.
//

#import "AppMain.h"

#import <Cocoa/Cocoa.h>

GoTest_SendMessageToGoFunc GoTest_SendMessageToGo_;

int GoTest_AppMain(int argc, const char** argv)
{
    return NSApplicationMain(argc, argv);
}

void GoTest_SendMessageToGo(char* message)
{
    GoTest_SendMessageToGo_(message);
}