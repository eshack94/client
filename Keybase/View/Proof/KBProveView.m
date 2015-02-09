//
//  KBTwitterView.m
//  Keybase
//
//  Created by Gabriel on 1/13/15.
//  Copyright (c) 2015 Gabriel Handford. All rights reserved.
//

#import "KBProveView.h"

//#import <Accounts/Accounts.h>
//#import <Social/Social.h>
#import <Slash/Slash.h>

#import "AppDelegate.h"

@implementation KBProveView

- (void)viewInit {
  [super viewInit];
  self.wantsLayer = YES;
  self.layer.backgroundColor = NSColor.whiteColor.CGColor;
  
  GHWeakSelf gself = self;

  _inputView = [[KBProveInputView alloc] init];
  _inputView.button.targetBlock = ^{
    [gself generateProof];
  };
  _inputView.cancelButton.targetBlock = ^{
    gself.completion(YES);
  };
  [self addSubview:_inputView];

  _instructionsView = [[KBProveInstructionsView alloc] init];
  _instructionsView.hidden = YES;
  [self addSubview:_instructionsView];

  [AppDelegate.client registerMethod:@"keybase.1.proveUi.promptUsername" requestHandler:^(NSString *method, NSArray *params, MPRequestCompletion completion) {
    //NSString *prompt = params[0][@"prompt"];
    completion(nil, gself.inputView.inputField.text);
  }];

  [AppDelegate.client registerMethod:@"keybase.1.proveUi.okToCheck" requestHandler:^(NSString *method, NSArray *params, MPRequestCompletion completion) {
    /*
    NSInteger attempt = [params[0][@"attempt"] integerValue];
    NSString *name = params[0][@"name"];
    NSString *prompt = NSStringWithFormat(@"Check %@%@?", name, attempt > 0 ? @" again" : @"");

    [KBAlert promptWithTitle:name description:prompt style:NSInformationalAlertStyle buttonTitles:@[@"OK", @"Cancel"] view:self completion:^(NSModalResponse response) {
      completion(nil, @(response == NSAlertFirstButtonReturn));
    }];
     */
    completion(nil, @(YES));
  }];

  [AppDelegate.client registerMethod:@"keybase.1.proveUi.promptOverwrite" requestHandler:^(NSString *method, NSArray *params, MPRequestCompletion completion) {

    NSString *account = params[0][@"account"];
    KBRPromptOverwriteType type = [params[0][@"typ"] integerValue];

    NSString *prompt;
    switch (type) {
      case KBRPromptOverwriteTypeSocial:
        prompt = NSStringWithFormat(@"You already have a proof for %@.", account);
        break;
      case KBRPromptOverwriteTypeSite:
        prompt = NSStringWithFormat(@"You already have claimed ownership of %@.", account);
        break;
    }

    [KBAlert promptWithTitle:@"Overwrite?" description:prompt style:NSWarningAlertStyle buttonTitles:@[NSStringWithFormat(@"Yes, Overwrite %@", account), @"Cancel"] view:self completion:^(NSModalResponse response) {
      completion(nil, @(response == NSAlertFirstButtonReturn));
    }];
  }];

  [AppDelegate.client registerMethod:@"keybase.1.proveUi.outputInstructions" requestHandler:^(NSString *method, NSArray *params, MPRequestCompletion completion) {
    // TODO: Verify sessionId?
    //sessionId = params[0][@"sessionId"];
    KBRText *instructions = [MTLJSONAdapter modelOfClass:KBRText.class fromJSONDictionary:params[0][@"instructions"] error:nil];
    NSString *proof = params[0][@"proof"];

    [self setInstructions:instructions proofText:proof targetBlock:^{
      completion(nil, @(YES));
    }];
  }];

  YOSelf yself = self;
  self.viewLayout = [YOLayout layoutWithLayoutBlock:^(id<YOLayout> layout, CGSize size) {
    CGFloat y = 40;

    [layout setFrame:CGRectMake(0, y, size.width, 0) view:yself.instructionsView];

    y += [layout sizeToFitVerticalInFrame:CGRectMake(0, y, size.width, 0) view:yself.inputView].size.height;

    return CGSizeMake(size.width, y);
  }];
}

+ (void)connectWithProveType:(KBProveType)proveType sender:(NSView *)sender completion:(KBProveCompletion)completion {
  KBProveView *proveView = [[KBProveView alloc] init];
  proveView.proveType = proveType;

  KBNavigationView *navigation = [[KBNavigationView alloc] initWithView:proveView];
  NSWindow *window = [KBWindow windowWithContentView:navigation size:CGSizeMake(420, 420) retain:NO];
  navigation.titleView = [KBTitleView titleViewWithTitle:NSStringWithFormat(@"Connect with %@", KBNameForProveType(proveType)) navigation:navigation];

  NSWindow *sourceWindow = sender.window ? sender.window : [NSApp mainWindow];
  [sourceWindow beginSheet:window completionHandler:^(NSModalResponse returnCode) {
    completion(returnCode == NSModalResponseCancel);
  }];

  proveView.completion = ^(BOOL canceled) {
    [sourceWindow endSheet:window returnCode:canceled ? NSModalResponseCancel : NSModalResponseContinue];
  };
}

- (void)viewDidAppear:(BOOL)animated {
  [self.window makeFirstResponder:_inputView];
}

- (void)setProveType:(KBProveType)proveType {
  _proveType = proveType;
  [_inputView setProveType:proveType];
}

- (void)setInstructions:(KBRText *)instructions proofText:(NSString *)proofText targetBlock:(KBButtonTargetBlock)targetBlock {
  [_instructionsView setInstructions:instructions proofText:proofText targetBlock:targetBlock];

  // TODO Animate change
  self.inputView.hidden = YES;
  self.instructionsView.hidden = NO;
}

- (void)generateProof {
  NSString *userName = [_inputView.inputField.text gh_strip];

  if ([NSString gh_isBlank:userName]) {
    // TODO Become first responder
    [AppDelegate setError:KBErrorAlert(@"You need to choose a username.") sender:_inputView];
    return;
  }

  NSString *service = KBServiceNameForProveType(self.proveType);
  NSAssert(service, @"No service");

  GHWeakSelf gself = self;
  [AppDelegate setInProgress:YES view:_inputView];
  KBRProveRequest *prove = [[KBRProveRequest alloc] initWithClient:AppDelegate.client];
  [self.navigation.titleView setProgressEnabled:YES];
  [prove proveWithService:service username:userName force:NO completion:^(NSError *error) {
    [AppDelegate setInProgress:NO view:gself.inputView];
    [self.navigation.titleView setProgressEnabled:NO];
    if (error) {
      [AppDelegate setError:error sender:gself.inputView];
      return;
    }

    [AppDelegate setInProgress:NO view:gself];

    self.completion(NO);
  }];
}

@end


