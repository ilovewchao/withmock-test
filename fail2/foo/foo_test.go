package foo

import (
    "testing"
    "code.google.com/p/gomock/gomock"
	
    "github.com/ilovewchao/withmock-test/fail2/baz" //mock
    "encoding/json" //mock
)

func TestTryMe(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    baz.MOCK().SetController(ctrl)
    json.MOCK().SetController(ctrl)
	
    fo := &Foo{}
    bz := &baz.Baz{}
	
    json.EXPECT().Unmarshal(gomock.Any(), bz).Return(nil)
    bz.EXPECT().DoSomething().Return(nil)
	
    fo.TryMe()
}
