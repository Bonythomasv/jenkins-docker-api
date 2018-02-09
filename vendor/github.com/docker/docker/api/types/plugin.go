package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Plugin A plugin for the Engine API
// swagger:ignore
// swagger:model Plugin
type Plugin struct {

	// config
	// Required: true
	Config PluginConfig `json:"Config"`

	// True if the plugin is running. False if the plugin is not running, only installed.
	// Required: true
	Enabled bool `json:"Enabled"`

	// Id
	ID string `json:"Id,omitempty"`

	// name
	// Required: true
	Name string `json:"Name"`

	// plugin remote reference used to push/pull the plugin
	PluginReference string `json:"PluginReference,omitempty"`

	// settings
	// Required: true
	Settings PluginSettings `json:"Settings"`
}

// PluginConfig The config of a plugin.
// swagger:ignore
// swagger:model PluginConfig
type PluginConfig struct {

	// args
	// Required: true
	Args PluginConfigArgs `json:"Args"`

	// description
	// Required: true
	Description string `json:"Description"`

	// Docker Version used to create the plugin
	DockerVersion string `json:"DockerVersion,omitempty"`

	// documentation
	// Required: true
	Documentation string `json:"Documentation"`

	// entrypoint
	// Required: true
	Entrypoint []string `json:"Entrypoint"`

	// env
	// Required: true
	Env []PluginEnv `json:"Env"`

	// interface
	// Required: true
	Interface PluginConfigInterface `json:"Interface"`

	// ipc host
	// Required: true
	IpcHost bool `json:"IpcHost"`

	// linux
	// Required: true
	Linux PluginConfigLinux `json:"Linux"`

	// mounts
	// Required: true
	Mounts []PluginMount `json:"Mounts"`

	// network
	// Required: true
	Network PluginConfigNetwork `json:"Network"`

	// pid host
	// Required: true
	PidHost bool `json:"PidHost"`

	// propagated mount
	// Required: true
	PropagatedMount string `json:"PropagatedMount"`

	// user
	User PluginConfigUser `json:"User,omitempty"`

	// work dir
	// Required: true
	WorkDir string `json:"WorkDir"`

	// rootfs
	Rootfs *PluginConfigRootfs `json:"rootfs,omitempty"`
}

// PluginConfigArgs plugin config args
// swagger:ignore
// swagger:model PluginConfigArgs
type PluginConfigArgs struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// name
	// Required: true
	Name string `json:"Name"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// value
	// Required: true
	Value []string `json:"Value"`
}

// PluginConfigInterface The interface between Docker and the plugin
// swagger:ignore
// swagger:model PluginConfigInterface
type PluginConfigInterface struct {

	// socket
	// Required: true
	Socket string `json:"Socket"`

	// types
	// Required: true
	Types []PluginInterfaceType `json:"Types"`
}

// PluginConfigLinux plugin config linux
// swagger:ignore
// swagger:model PluginConfigLinux
type PluginConfigLinux struct {

	// allow all devices
	// Required: true
	AllowAllDevices bool `json:"AllowAllDevices"`

	// capabilities
	// Required: true
	Capabilities []string `json:"Capabilities"`

	// devices
	// Required: true
	Devices []PluginDevice `json:"Devices"`
}

// PluginConfigNetwork plugin config network
// swagger:ignore
// swagger:model PluginConfigNetwork
type PluginConfigNetwork struct {

	// type
	// Required: true
	Type string `json:"Type"`
}

// PluginConfigRootfs plugin config rootfs
// swagger:ignore
// swagger:model PluginConfigRootfs
type PluginConfigRootfs struct {

	// diff ids
	DiffIds []string `json:"diff_ids"`

	// type
	Type string `json:"type,omitempty"`
}

// PluginConfigUser plugin config user
// swagger:ignore
// swagger:model PluginConfigUser
type PluginConfigUser struct {

	// g ID
	GID uint32 `json:"GID,omitempty"`

	// UID
	UID uint32 `json:"UID,omitempty"`
}

// PluginSettings Settings that can be modified by users.
// swagger:ignore
// swagger:model PluginSettings
type PluginSettings struct {

	// args
	// Required: true
	Args []string `json:"Args"`

	// devices
	// Required: true
	Devices []PluginDevice `json:"Devices"`

	// env
	// Required: true
	Env []string `json:"Env"`

	// mounts
	// Required: true
	Mounts []PluginMount `json:"Mounts"`
}
