package service

import "log"

type InstanceService struct{}

func NewInstanceService() *InstanceService {
	return &InstanceService{}
}

func (s *InstanceService) Create(data map[string]interface{}) error {
	log.Println("创建实例:", data)
	return nil
}

func (s *InstanceService) Delete(instanceID string) error {
	log.Println("删除实例:", instanceID)
	return nil
}

func (s *InstanceService) Start(instanceID string) error {
	log.Println("启动实例:", instanceID)
	return nil
}

func (s *InstanceService) Shutdown(instanceID string) error {
	log.Println("关闭实例:", instanceID)
	return nil
}

func (s *InstanceService) AttachGPU(instanceID string) error {
	log.Println("绑定 GPU:", instanceID)
	return nil
}

func (s *InstanceService) DetachGPU(instanceID string) error {
	log.Println("解绑 GPU:", instanceID)
	return nil
}

func (s *InstanceService) ModifyConfig(instanceID string, config map[string]interface{}) error {
	log.Println("修改配置:", instanceID, config)
	return nil
}

func (s *InstanceService) EnableNetworkProxy(instanceID string) error {
	log.Println("启用网络代理:", instanceID)
	return nil
}

func (s *InstanceService) DisableNetworkProxy(instanceID string) error {
	log.Println("禁用网络代理:", instanceID)
	return nil
}

func (s *InstanceService) SetupApplications(instanceID string, apps []string) error {
	log.Println("安装应用:", instanceID, apps)
	return nil
}

func (s *InstanceService) SetupApplicationsV2(instanceID string, apps []string) error {
	log.Println("安装应用 V2:", instanceID, apps)
	return nil
}

func (s *InstanceService) ModifyPassword(instanceID string, password string) error {
	log.Println("修改密码:", instanceID)
	return nil
}
