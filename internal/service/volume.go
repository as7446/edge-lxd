package service

import "log"

type VolumeService struct{}

func NewVolumeService() *VolumeService {
	return &VolumeService{}
}

func (s *VolumeService) Create(data map[string]interface{}) error {
	log.Println("创建存储卷:", data)
	return nil
}

func (s *VolumeService) Clone(data map[string]interface{}) error {
	log.Println("克隆存储卷:", data)
	return nil
}

func (s *VolumeService) Delete(data map[string]interface{}) error {
	log.Println("删除存储卷:", data)
	return nil
}

func (s *VolumeService) Attach(data map[string]interface{}) error {
	log.Println("绑定存储卷:", data)
	return nil
}

func (s *VolumeService) Detach(data map[string]interface{}) error {
	log.Println("解绑存储卷:", data)
	return nil
}

func (s *VolumeService) ChangeSize(data map[string]interface{}) error {
	log.Println("调整存储卷大小:", data)
	return nil
}
