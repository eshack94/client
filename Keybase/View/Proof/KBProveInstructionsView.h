//
//  KBProveInstructionsView.h
//  Keybase
//
//  Created by Gabriel on 2/9/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import <Foundation/Foundation.h>

#import "KBUIDefines.h"
#import "KBRPC.h"
#import "KBProveType.h"

@interface KBProveInstructionsView : YONSView
@property KBNavigationView *navigation;
@property KBLabel *instructionsLabel;
@property KBLabel *proofLabel;
@property NSScrollView *scrollView;
@property KBButton *button;
@property KBButton *clipboardCopyButton;

@property NSString *proofText;
- (void)setInstructions:(KBRText *)instructions proofText:(NSString *)proofText targetBlock:(KBButtonTargetBlock)targetBlock;
@end

