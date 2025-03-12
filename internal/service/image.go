package service

import "log"

type ImageService struct{}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (s *ImageService) Export(data map[string]interface{}) error {
	log.Println("导出镜像:", data)
	return nil
}

func (s *ImageService) Delete(data map[string]interface{}) error {
	log.Println("删除镜像:", data)
	return nil
}

func (s *ImageService) Copy(data map[string]interface{}) error {
	log.Println("复制镜像:", data)
	return nil
}
