package service

// 类比于java

// 定义类
type MultiSignService struct{}

// 创建实例
func NewMultiSignService() *MultiSignService {
	return &MultiSignService{}
}

// 实例方法
//func (c *MultiSignService) SetMultiSign(multiSign *request.SetMultiSign) (int, error) {
//	err := models.NewMultiSign().Set(multiSign)
//	if err != nil {
//		return statecode.CommonErrServerErr, err
//	}
//	return statecode.CommonSuccess, nil
//}
//
//// 值通过指针赋值到入参, 返回code和err
//func (c *MultiSignService) GetMultiSign(multiSign *response.MultiSign, chainId int) (int, error) {
//	multiSignModel := models.NewMultiSign()
//	// 获得的值已保存到multiSignModel
//	err := multiSignModel.Get(chainId)
//	if err != nil {
//		return statecode.CommonErrServerErr, err
//	}
//	var multiSignAccount []string
//
//	_ = json.Unmarshal([]byte(multiSignModel.MultiSignAccount), &multiSignAccount)
//
//	multiSign.SpName = multiSignModel.SpName
//	multiSign.SpToken = multiSignModel.SpToken
//	multiSign.SpHash = multiSignModel.SpHash
//	multiSign.SpAddress = multiSignModel.SpAddress
//
//	multiSign.JpName = multiSignModel.JpName
//	multiSign.JpToken = multiSignModel.JpToken
//	multiSign.JpHash = multiSignModel.JpHash
//	multiSign.JpAddress = multiSignModel.JpAddress
//	multiSign.MultiSignAccount = multiSignAccount
//	return statecode.CommonSuccess, nil
//}
