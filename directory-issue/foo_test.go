package foo_test

import (
    "testing"
    "code.google.com/p/gomock/gomock"
	
    "github.com/ilovewchao/withmock-test/directory-issue" //mock
	"net/http" //mock
)

func TestTryMe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	foo.MOCK().SetController(ctrl)
	
	fo := &foo.Foo{}
	bz := &foo.Baz{}
	bz.EXPECT().DoSomething().Return("doone")
	
	fo.TryMe()
}

func TestHttpGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	http.MOCK().SetController(ctrl)

	res := &http.Response{}

    http.EXPECT().Get("http://www.google.com.hk").Return(res, nil)

	err := apps.HttpGet()

	if err != nil {
		t.Errorf("SetServiceCredentials return an error: %s", err)
	}
}
