//autogenerated:yes
//nolint:revive,lll
package mavros_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type CommandTOLReq struct {
	msg.Package `ros:"mavros_msgs"`
	MinPitch    float32
	Yaw         float32
	Latitude    float32
	Longitude   float32
	Altitude    float32
}

type CommandTOLRes struct {
	msg.Package `ros:"mavros_msgs"`
	Success     bool
	Result      uint8
}

type CommandTOL struct {
	msg.Package `ros:"mavros_msgs"`
	CommandTOLReq
	CommandTOLRes
}
