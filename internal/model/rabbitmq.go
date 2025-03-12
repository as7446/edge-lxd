package model

// 监听的队列
var queues = []string{
	"api.instance.create",
	"api.instance.delete",
	"api.instance.start",
	"api.instance.shutdown",
	"api.instance.attach.gpu",
	"api.instance.detach.gpu",
	"api.instance.modify.config",
	"api.instance.enable.network.proxy",
	"api.instance.disable.network.proxy",
	"api.instance.setup.applications",
	"api.instance.setup.applications.v2",
	"api.instance.modify.password",
	"api.ops.ping.edge",
	"api.ops.ping.edge.instance",
	"api.ops.get.edge.info",
	"api.ops.run.edge.instance.command",
	"api.local.volume.create",
	"api.local.volume.clone",
	"api.image.export",
	"api.image.delete",
	"api.image.copy",
	"api.local.volume.delete",
	"api.local.volume.attach",
	"api.local.volume.detach",
	"api.local.volume.change.size",
}
