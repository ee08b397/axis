/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"github.com/blackrock/axis/pkg/apis/sensor/v1alpha1"
)

// getNodeByName returns a copy of the node from this sensor for the nodename
// for signals this nodename should be the name of the signal
func (soc *sOperationCtx) getNodeByName(nodename string) *v1alpha1.NodeStatus {
	nodeID := soc.s.NodeID(nodename)
	node, ok := soc.s.Status.Nodes[nodeID]
	if !ok {
		return nil
	}
	return node.DeepCopy()
}
