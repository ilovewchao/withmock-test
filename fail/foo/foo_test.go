package foo

import (
    "testing"
    "code.google.com/p/gomock/gomock"
	
    "github.com/ilovewchao/withmock-test/fail/baz" //mock
    mockfmt "fmt" //mock
)

func TestTryMe(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    baz.MOCK().SetController(ctrl)
    mockfmt.MOCK().SetController(ctrl)
	
    fo := &Foo{}
    bz := &baz.Baz{}
    bz.EXPECT().DoSomething().Return("doing sth")
	
    mockfmt.EXPECT().Println(gomock.Any())
	
    fo.TryMe()
}
