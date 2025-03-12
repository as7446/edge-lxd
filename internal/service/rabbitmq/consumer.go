package rabbitmq

import handlers "github.com/as7446/lxd/internal/handler"

func RegisterQueueHandlers(rabbit *RabbitMQ) {
	rabbit.RegisterConsumer("api.instance.create", handlers.HandleInstanceCreate)
	rabbit.RegisterConsumer("api.instance.delete", handlers.HandleInstanceDelete)
	rabbit.RegisterConsumer("api.instance.start", handlers.HandleInstanceStart)
	rabbit.RegisterConsumer("api.instance.shutdown", handlers.HandleInstanceShutdown)
	rabbit.RegisterConsumer("api.instance.attach.gpu", handlers.HandleInstanceAttachGPU)
	rabbit.RegisterConsumer("api.instance.detach.gpu", handlers.HandleInstanceDetachGPU)
	rabbit.RegisterConsumer("api.instance.modify.config", handlers.HandleInstanceModifyConfig)
	rabbit.RegisterConsumer("api.instance.enable.network.proxy", handlers.HandleInstanceEnableNetworkProxy)
	rabbit.RegisterConsumer("api.instance.disable.network.proxy", handlers.HandleInstanceDisableNetworkProxy)
	rabbit.RegisterConsumer("api.instance.setup.applications", handlers.HandleInstanceSetupApplications)
	rabbit.RegisterConsumer("api.instance.setup.applications.v2", handlers.HandleInstanceSetupApplicationsV2)
	rabbit.RegisterConsumer("api.instance.modify.password", handlers.HandleInstanceModifyPassword)

	rabbit.RegisterConsumer("api.local.volume.create", handlers.HandleLocalVolumeCreate)
	rabbit.RegisterConsumer("api.local.volume.clone", handlers.HandleLocalVolumeClone)
	rabbit.RegisterConsumer("api.local.volume.delete", handlers.HandleLocalVolumeDelete)
	rabbit.RegisterConsumer("api.local.volume.attach", handlers.HandleLocalVolumeAttach)
	rabbit.RegisterConsumer("api.local.volume.detach", handlers.HandleLocalVolumeDetach)
	rabbit.RegisterConsumer("api.local.volume.change.size", handlers.HandleLocalVolumeChangeSize)

	rabbit.RegisterConsumer("api.image.export", handlers.HandleImageExport)
	rabbit.RegisterConsumer("api.image.delete", handlers.HandleImageDelete)
	rabbit.RegisterConsumer("api.image.copy", handlers.HandleImageCopy)

	rabbit.RegisterConsumer("api.ops.ping.edge", handlers.HandleOpsPingEdge)
	rabbit.RegisterConsumer("api.ops.ping.edge.instance", handlers.HandleOpsPingEdgeInstance)
	rabbit.RegisterConsumer("api.ops.get.edge.info", handlers.HandleOpsGetEdgeInfo)
	rabbit.RegisterConsumer("api.ops.run.edge.instance.command", handlers.HandleOpsRunEdgeInstanceCommand)
}
