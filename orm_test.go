package unittest

import (
	"code.google.com/p/gomock/gomock"
	"testing"

	"github.com/astaxie/beego/orm" //mock

	"github.com/stretchr/testify/assert"
	"rnd-github.huawei.com/paas/paas-controller/models/apps"

	//_ "github.com/cloudfoundry/cli/cf/models"
	//_ "github.com/go-sql-driver/mysql"
	//_ "rnd-github.huawei.com/paas/paas-controller/routers"
	//_ "rnd-github.huawei.com/paas/paas-controller/util/cfagent"
)

//func init() {
//	orm.RegisterDriver("mysql", orm.DR_MySQL)

//	// set default database
//	{
//		sqldriver := "mysql"
//		url := "root:root@tcp(9.91.18.211:3306)/paas?charset=utf8"
//		orm.RegisterDataBase("default", sqldriver, url, 30)
//	}
//	//automatically sync database schema
//	orm.RunSyncdb("default", false, true)
//}

func TestGetAppRecipesByAppGuid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	orm.MOCK().SetController(ctrl)
	appmodels.MOCK().SetController(ctrl)
	appmodels.MOCK().MockAll(false)
	appmodels.MOCK().EnableMock("GetCustomAppByGuid")

	customApp := appmodels.CustomApp{}
	customApp.Id = 123456
	appmodels.EXPECT().GetCustomAppByGuid("appGuid").Return(customApp, nil)

	ormer := orm.MOCK().NewOrmer()
	query := orm.MOCK().NewQuerySeter()

	orm.EXPECT().NewOrm().Return(ormer)
	ormer.EXPECT().QueryTable(gomock.Any()).Return(query)
	query.EXPECT().Filter("app_id", 123456).Return(query)
	query.EXPECT().OrderBy("-create_at").Return(query)
	query.EXPECT().All(gomock.Any()).Return(int64(2), nil)

	_, err := appmodels.GetAppRecipesByAppGuid("appGuid")
	assert.Nil(t, err)
}

//func TestGetCustomAppByAppId(t *testing.T) {
//	assert := assert.New(t)

//	_, err := GetCustomAppByAppId("12345")
//	assert.Equal(err, orm.ErrNoRows) //, "The result should not be nil."
//}

//func TestGetAppRecipesByAppGuid(t *testing.T) {
//	assert := assert.New(t)

//	_, err := GetAppRecipesByAppGuid("12345")
//	assert.Equal(err, orm.ErrNoRows) //, "The result should not be nil."
//}
