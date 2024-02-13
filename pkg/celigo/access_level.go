package celigo

const (
	ManageAccessLevel  = "manage"
	MonitorAccessLevel = "monitor"
	AdminAccessLevel   = "administrator"
)

var (
	AccessLevels             = []string{ManageAccessLevel, MonitorAccessLevel, AdminAccessLevel}
	IntegrationsAccessLevels = []string{ManageAccessLevel, MonitorAccessLevel}
)
