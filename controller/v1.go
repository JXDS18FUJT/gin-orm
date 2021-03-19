package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"

	bda "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bda/v20200324"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func CreateV1Todo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})

}

func V1bda(c *gin.Context) {

	credential := common.NewCredential(
		"AKIDxgHa1CC4LM7hT4jGWieJpzcy3sfjK6S5",
		"WmYxhAVbEbniOXR57sCEySP7ZppgHcy3",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "bda.tencentcloudapi.com"
	client, _ := bda.NewClient(credential, "ap-guangzhou", cpf)

	request := bda.NewSegmentCustomizedPortraitPicRequest()

	request.Image = common.StringPtr("")
	request.SegmentationOptions = &bda.SegmentationOptions{
		Background:   common.BoolPtr(false),
		Hair:         common.BoolPtr(true),
		LeftEyebrow:  common.BoolPtr(true),
		RightEyebrow: common.BoolPtr(true),
		LeftEye:      common.BoolPtr(true),
		RightEye:     common.BoolPtr(true),
		Nose:         common.BoolPtr(true),
		UpperLip:     common.BoolPtr(true),
		LowerLip:     common.BoolPtr(true),
		Tooth:        common.BoolPtr(true),
		Mouth:        common.BoolPtr(true),
		LeftEar:      common.BoolPtr(true),
		RightEar:     common.BoolPtr(true),
		Face:         common.BoolPtr(true),
		Head:         common.BoolPtr(true),
		Body:         common.BoolPtr(true),
	}

	response, err := client.SegmentCustomizedPortraitPic(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
	// fmt.Printf("%s", response.ToJsonString())

}
