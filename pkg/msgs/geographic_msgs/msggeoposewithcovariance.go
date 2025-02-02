//autogenerated:yes
//nolint:revive,lll
package geographic_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type GeoPoseWithCovariance struct {
	msg.Package `ros:"geographic_msgs"`
	Pose        GeoPose
	Covariance  [36]float64
}
