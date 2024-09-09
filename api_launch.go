package main

import (
        "fmt"

        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
        ssa "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssa/v20180608"
)

func main() {
        // 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
        // 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
        // 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
        credential := common.NewCredential(
                "AKIDskYK1Z8Zs35ykOuIXdaP6e3LnO1W3Sks",
                "SecretKey",
        )
        // 实例化一个client选项，可选的，没有特殊需求可以跳过
        cpf := profile.NewClientProfile()
        cpf.HttpProfile.Endpoint = "ssa.tencentcloudapi.com"
        // 实例化要请求产品的client对象,clientProfile是可选的
        client, _ := ssa.NewClient(credential, "", cpf)

        // 实例化一个请求对象,每个接口都会对应一个request对象
        request := ssa.NewDescribeLeakDetectionListRequest()
        
        request.Filters = []*ssa.Filter {
                &ssa.Filter {
                        Name: common.StringPtr(""),
                        Values: common.StringPtrs([]string{ "" }),
                        ExactMatch: common.BoolPtr(false),
                },
        }
        request.Limit = common.Int64Ptr(10)
        request.Page = common.Int64Ptr(1)
        request.StartTime = common.StringPtr("2024-06-06")
        request.EndTime = common.StringPtr("2024-07-07")

        // 返回的resp是一个DescribeLeakDetectionListResponse的实例，与请求对象对应
        response, err := client.DescribeLeakDetectionList(request)
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
                fmt.Printf("An API error has returned: %s", err)
                return
        }
        if err != nil {
                panic(err)
        }
        // 输出json格式的字符串回包
        fmt.Printf("%s", response.ToJsonString())
} 
