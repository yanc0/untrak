package kubernetes

var DefaultNonNamespacedResources = []string{
	"componentstatuse",
	"namespace",
	"node",
	"persistentvolume",
	"mutatingwebhookconfiguration",
	"validatingwebhookconfiguration",
	"customresourcedefinition",
	"apiservice",
	"certificatesigningrequest",
	"runtimeclass",
	"podsecuritypolicy",
	"clusterrolebinding",
	"clusterrole",
	"priorityclass",
	"csidriver",
	"csinode",
	"storageclass",
	"volumeattachment",
}