//
//  AppDelegate.m
//  GoTest
//
//  Created by Hajime Hoshi on 6/11/13.
//  Copyright (c) 2013 Hajime Hoshi. All rights reserved.
//

#import "AppDelegate.h"
#import "AppMain.h"

@implementation AppDelegate

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification
{
    // Insert code here to initialize your application
}

- (IBAction)pushButton:(id)sender {
    GoTest_SendMessageToGo("Hello, World!");
}

@end
