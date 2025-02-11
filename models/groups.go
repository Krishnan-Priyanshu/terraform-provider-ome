/*
Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://mozilla.org/MPL/2.0/
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Groups - list of group response from OME
type Groups struct {
	Value []Group `json:"value"`
}

// Group - embedded group response from the groups
type Group struct {
	ID                    int64  `json:"Id,omitempty"`
	Name                  string `json:"Name"`
	Description           string `json:"Description"`
	MembershipTypeID      int64  `json:"MembershipTypeId"`
	ParentID              int64  `json:"ParentId"`
	GlobalStatus          int64  `json:"GlobalStatus"`
	IDOwner               int64  `json:"IdOwner"`
	CreationTime          string `json:"CreationTime"`
	UpdatedTime           string `json:"UpdatedTime"`
	CreatedBy             string `json:"CreatedBy"`
	UpdatedBy             string `json:"UpdatedBy"`
	Visible               bool   `json:"Visible"`
	DefinitionID          int64  `json:"DefinitionId"`
	DefinitionDescription string `json:"DefinitionDescription"`
	TypeID                int64  `json:"TypeId"`
	HasAttributes         bool   `json:"HasAttributes"`
	IsAccessAllowed       bool   `json:"IsAccessAllowed"`
	// SubGroups are only present when expanded
	SubGroups []Group `json:"SubGroups,omitempty"`
}

// GroupMemberPayload - Payload struct for adding or removing devices from static groups
type GroupMemberPayload struct {
	GroupID   int64   `json:"GroupId"`
	DeviceIds []int64 `json:"MemberDeviceIds,omitempty"`
}

// NewGroupMemberPayload initializes a GroupMemberPayload struct
func NewGroupMemberPayload(gid int64) GroupMemberPayload {
	return GroupMemberPayload{
		GroupID:   gid,
		DeviceIds: make([]int64, 0),
	}
}

// RegisterDevice - helper function to add devices to the payload
func (plan *GroupMemberPayload) RegisterDevice(id int64) {
	plan.DeviceIds = append(plan.DeviceIds, id)
}

// GroupDevicesData - schema for data source groupdevices
type GroupDevicesData struct {
	ID                types.String        `tfsdk:"id"`
	DeviceIDs         types.List          `tfsdk:"device_ids"`
	DeviceServicetags types.List          `tfsdk:"device_servicetags"`
	DeviceGroupNames  types.Set           `tfsdk:"device_group_names"`
	DeviceGroups      map[string]OmeGroup `tfsdk:"device_groups"`
}

// SetDevices - sets a list of devices in the device_ids and device_servicetags fields
func (g *GroupDevicesData) SetDevices(devices []Device) {
	devIDs := []attr.Value{}
	devSvcTags := []attr.Value{}

	for _, device := range devices {
		devIDs = append(devIDs, types.Int64Value(device.ID))
		devSvcTags = append(devSvcTags, types.StringValue(device.DeviceServiceTag))
	}

	devIDsTfsdk, _ := types.ListValue(
		types.Int64Type,
		devIDs,
	)

	g.DeviceIDs = devIDsTfsdk

	devSTsTfsdk, _ := types.ListValue(
		types.StringType,
		devSvcTags,
	)

	g.DeviceServicetags = devSTsTfsdk
}

// SetGroup - registers a group and its child devices in the device_groups map
func (g *GroupDevicesData) SetGroup(group Group, devices []Device) {
	if g.DeviceGroups == nil {
		g.DeviceGroups = make(map[string]OmeGroup)
	}

	omeDevices := make([]OmeDeviceIDData, 0)
	for _, dev := range devices {
		omeDevices = append(omeDevices, OmeDeviceIDData{
			ID:         types.Int64Value(dev.ID),
			ServiceTag: types.StringValue(dev.DeviceServiceTag),
		})
	}

	subGroupTfsdk := make([]OmeSubGroup, 0)
	for _, subGroup := range group.SubGroups {
		subGroupTfsdk = append(subGroupTfsdk, OmeSubGroup{
			ID:   types.Int64Value(subGroup.ID),
			Name: types.StringValue(subGroup.Name),
		})
	}

	g.DeviceGroups[group.Name] = OmeGroup{
		ID:                    types.Int64Value(group.ID),
		Name:                  types.StringValue(group.Name),
		Description:           types.StringValue(group.Description),
		MembershipTypeID:      types.Int64Value(group.MembershipTypeID),
		ParentID:              types.Int64Value(group.ParentID),
		GlobalStatus:          types.Int64Value(group.GlobalStatus),
		IDOwner:               types.Int64Value(group.IDOwner),
		CreationTime:          types.StringValue(group.CreationTime),
		UpdatedTime:           types.StringValue(group.UpdatedTime),
		CreatedBy:             types.StringValue(group.UpdatedBy),
		UpdatedBy:             types.StringValue(group.CreatedBy),
		Visible:               types.BoolValue(group.Visible),
		DefinitionID:          types.Int64Value(group.ParentID),
		DefinitionDescription: types.StringValue(group.Name),
		TypeID:                types.Int64Value(group.TypeID),
		HasAttributes:         types.BoolValue(group.HasAttributes),
		IsAccessAllowed:       types.BoolValue(group.IsAccessAllowed),
		Devices:               omeDevices,
		SubGroups:             subGroupTfsdk,
	}
}

// OmeGroup - schema for group info in data source groupdevices
type OmeGroup struct {
	ID                    types.Int64       `tfsdk:"id"`
	Name                  types.String      `tfsdk:"name"`
	Description           types.String      `tfsdk:"description"`
	MembershipTypeID      types.Int64       `tfsdk:"membership_type_id"`
	ParentID              types.Int64       `tfsdk:"parent_id"`
	GlobalStatus          types.Int64       `tfsdk:"global_status"`
	IDOwner               types.Int64       `tfsdk:"id_owner"`
	CreationTime          types.String      `tfsdk:"creation_time"`
	UpdatedTime           types.String      `tfsdk:"updated_time"`
	CreatedBy             types.String      `tfsdk:"created_by"`
	UpdatedBy             types.String      `tfsdk:"updated_by"`
	Visible               types.Bool        `tfsdk:"visible"`
	DefinitionID          types.Int64       `tfsdk:"definition_id"`
	DefinitionDescription types.String      `tfsdk:"definition_description"`
	TypeID                types.Int64       `tfsdk:"type_id"`
	HasAttributes         types.Bool        `tfsdk:"has_attributes"`
	IsAccessAllowed       types.Bool        `tfsdk:"is_access_allowed"`
	Devices               []OmeDeviceIDData `tfsdk:"devices"`
	SubGroups             []OmeSubGroup     `tfsdk:"sub_groups"`
}

// OmeDeviceIDData - schema for device info info in data source groupdevices
type OmeDeviceIDData struct {
	ID         types.Int64  `tfsdk:"id"`
	ServiceTag types.String `tfsdk:"servicetag"`
}

// OmeSubGroup - schema for sub groups in data source groupdevices
type OmeSubGroup struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// StaticGroup - schema for resource static group
type StaticGroup struct {
	ID               types.Int64  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	MembershipTypeID types.Int64  `tfsdk:"membership_type_id"`
	ParentID         types.Int64  `tfsdk:"parent_id"`
	DeviceIds        types.Set    `tfsdk:"device_ids"`
}

// NewStaticGroup - marshalls api response group and devices structs in a static group state
func NewStaticGroup(plan Group, devs Devices) (StaticGroup, diag.Diagnostics) {
	devidVals := make([]attr.Value, 0)
	for _, device := range devs.Value {
		devidVals = append(devidVals, types.Int64Value(device.ID))
	}
	deviceIds, dgs := types.SetValue(types.Int64Type, devidVals)
	return StaticGroup{
		ID:               types.Int64Value(plan.ID),
		Name:             types.StringValue(plan.Name),
		Description:      types.StringValue(plan.Description),
		MembershipTypeID: types.Int64Value(plan.MembershipTypeID),
		ParentID:         types.Int64Value(plan.ParentID),
		DeviceIds:        deviceIds,
	}, dgs
}

// GetPayload - returns the diff between plan and state sttaic groups in terms of its modify api request
func (plan *StaticGroup) GetPayload(state StaticGroup) (Group, bool) {
	ret := Group{
		ID:               state.ID.ValueInt64(),
		Name:             plan.Name.ValueString(),
		Description:      plan.Description.ValueString(),
		MembershipTypeID: 12,
		ParentID:         plan.ParentID.ValueInt64(),
	}
	return ret, ret.Name == state.Name.ValueString() && ret.Description == state.Description.ValueString()
}

// GetMemberPayload - returns the diff between plan and state sttaic groups in terms of
// its device add and remove api requests
func (plan *StaticGroup) GetMemberPayload(ctx context.Context, state StaticGroup) (GroupMemberPayload,
	GroupMemberPayload, diag.Diagnostics) {
	var d diag.Diagnostics
	toAdd, toRmv := NewGroupMemberPayload(state.ID.ValueInt64()), NewGroupMemberPayload(state.ID.ValueInt64())
	if plan.DeviceIds.Equal(state.DeviceIds) {
		return toAdd, toRmv, d
	}
	planDevIds, dgs1 := plan.GetDeviceIDMap(ctx)
	d.Append(dgs1...)
	stateDevIds, dgs2 := state.GetDeviceIDMap(ctx)
	d.Append(dgs2...)

	// Loop over all devices in state
	for sid := range stateDevIds {
		if _, ok := planDevIds[sid]; !ok {
			// Register all devices to remove
			toRmv.RegisterDevice(sid)
		} else {
			// Flag all devices that are already in state
			planDevIds[sid] = false
		}
	}

	// Register all devices to add (ie, devices not flagged)
	for pid, toAddFlag := range planDevIds {
		if toAddFlag {
			toAdd.RegisterDevice(pid)
		}
	}

	return toAdd, toRmv, d
}

// GetDeviceIDMap - helper function that converts a static group's list of devices into a map of its ids
// this helps in quick comparison of device lists of two static groups
func (plan *StaticGroup) GetDeviceIDMap(ctx context.Context) (map[int64]bool, diag.Diagnostics) {
	var d diag.Diagnostics
	ret, devIds := make(map[int64]bool), make([]int64, 0)
	d.Append(plan.DeviceIds.ElementsAs(ctx, &devIds, false)...)
	for _, id := range devIds {
		ret[id] = true
	}
	return ret, d
}
