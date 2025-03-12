package service

import "log"

type OpsService struct{}

func NewOpsService() *OpsService {
	return &OpsService{}
}

func (s *OpsService) PingEdge(data map[string]interface{}) error {
	log.Println("Ping Edge:", data)
	return nil
}

func (s *OpsService) PingEdgeInstance(data map[string]interface{}) error {
	log.Println("Ping Edge Instance:", data)
	return nil
}

func (s *OpsService) GetEdgeInfo(data map[string]interface{}) error {
	log.Println("获取 Edge 信息:", data)
	return nil
}

func (s *OpsService) RunEdgeInstanceCommand(data map[string]interface{}) error {
	log.Println("运行 Edge 命令:", data)
	return nil
}
