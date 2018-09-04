/*
Copyright 2018 Yunify, Inc.

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

package neonsan

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	SnapTestSnapshotName     = "test"
	SnapTestFakeSnapshotName = "fake"
	SnapTestPoolName         = "csi"
	SnapTestVolumeName       = "foo"
	SnapTestFakeVolumeName   = "fake"
)

func TestCreateSnapshot(t *testing.T) {
	tests := []struct {
		name     string
		snapInfo *snapshotInfo
		err      error
	}{
		{
			name: "Succeed to create snapshot",
			snapInfo: &snapshotInfo{
				snapName:         SnapTestSnapshotName,
				pool:             SnapTestPoolName,
				sourceVolumeName: SnapTestVolumeName,
			},
			err: nil,
		},
		{
			name: "Recreate snapshot",
			snapInfo: &snapshotInfo{
				snapName:         SnapTestSnapshotName,
				pool:             SnapTestPoolName,
				sourceVolumeName: SnapTestFakeVolumeName,
			},
			err: fmt.Errorf("Raise error"),
		},
		{
			name: "Failed to create snapshot",
			snapInfo: &snapshotInfo{
				snapName:         SnapTestSnapshotName,
				pool:             SnapTestPoolName,
				sourceVolumeName: SnapTestFakeVolumeName,
			},
			err: fmt.Errorf("Raise error"),
		},
	}
	for _, v := range tests {
		snapInfo, err := CreateSnapshot(v.snapInfo.snapName, v.snapInfo.sourceVolumeName, v.snapInfo.pool)
		if (v.err != nil && err == nil) || (v.err == nil && err != nil) {
			t.Errorf("name %s: error expect %v, but actually %v", v.name, v.err, err)
		} else if !reflect.DeepEqual(v.snapInfo, snapInfo) {
			t.Errorf("name %s: error expect %v, but actually %v", v.name, v.snapInfo, snapInfo)
		}
	}
}
