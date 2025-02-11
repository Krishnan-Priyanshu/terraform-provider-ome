---
# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
# 
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://mozilla.org/MPL/2.0/
# 
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

title: "ome_discovery resource"
linkTitle: "ome_discovery"
page_title: "ome_discovery Resource - terraform-provider-ome"
subcategory: ""
description: |-
  Resource for managing discovery on OpenManage Enterprise.
---

# ome_discovery (Resource)

Resource for managing discovery on OpenManage Enterprise.




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `discovery_config_targets` (Attributes Set) - Provide the list of discovery targets.
      			- Each discovery target is a set of "network_address_detail", "device_types", and one or more protocol credentials. (see [below for nested schema](#nestedatt--discovery_config_targets))
- `name` (String) Name of the discovery configuration job

### Optional

- `cron` (String) Provide a cron expression based on Quartz cron format
- `email_recipient` (String) - Enter the email address to which notifications are to be sent about the discovery job status.
				- Configure the SMTP settings to allow sending notifications to an email address.
- `enable_community_strings` (Boolean) - Enable the use of SNMP community strings to receive SNMP traps using Application Settings in OpenManage Enterprise. 
				- This option is available only for the discovered iDRAC servers and MX7000 chassis.
- `schedule` (String) Provides the option to schedule the discovery job. If `RunLater` is selected, then attribute `cron` must be specified.
- `trap_destination` (Boolean) - Enable OpenManage Enterprise to receive the incoming SNMP traps from the discovered devices. 
				- This is effective only for servers discovered by using their iDRAC interface.

### Read-Only

- `id` (String) ID of the discovery configuration group
- `job_id` (Number) Discovery Job ID.

<a id="nestedatt--discovery_config_targets"></a>
### Nested Schema for `discovery_config_targets`

Required:

- `device_type` (List of String) - Provide the type of devices to be discovered.
				- The accepted types are SERVER, CHASSIS, NETWORK SWITCH, and STORAGE.
				- A combination or all of the above can be provided.
				- "Supported protocols for each device type are:"
				- SERVER - "redfish", "snmp", and "ssh".
				- CHASSIS - "redfish".
				- NETWORK SWITCH - "snmp".
				- STORAGE - "snmp".
- `network_address_detail` (List of String) - "Provide the list of IP addresses, host names, or the range of IP addresses of the devices to be discoveredor included."
         		- "Sample Valid IP Range Formats"
         		- "   192.35.0.0"
         		- "   192.36.0.0-10.36.0.255"
         		- "   192.37.0.0/24"
         		- "   2345:f2b1:f083:135::5500/118"
         		- "   2345:f2b1:f083:135::a500-2607:f2b1:f083:135::a600"
         		- "   hostname.domain.tld"
         		- "   hostname"
         		- "   2345:f2b1:f083:139::22a"
         		- "Sample Invalid IP Range Formats"
         		- "   192.35.0.*"
         		- "   192.36.0.0-255"
         		- "   192.35.0.0/255.255.255.0"
         		- NOTE: The range size for the number of IP addresses is limited to 16,385 (0x4001).
         		- NOTE: Both IPv6 and IPv6 CIDR formats are supported.

Optional:

- `redfish` (Attributes) REDFISH protocol (see [below for nested schema](#nestedatt--discovery_config_targets--redfish))
- `snmp` (Attributes) Simple Network Management Protocol (SNMP) (see [below for nested schema](#nestedatt--discovery_config_targets--snmp))
- `ssh` (Attributes) Secure Shell (SSH) (see [below for nested schema](#nestedatt--discovery_config_targets--ssh))
- `wsman` (Attributes) WSMAN protocol (see [below for nested schema](#nestedatt--discovery_config_targets--wsman))

<a id="nestedatt--discovery_config_targets--redfish"></a>
### Nested Schema for `discovery_config_targets.redfish`

Required:

- `password` (String) Provide a password for the protocol.
- `username` (String) Provide a username for the protocol.

Optional:

- `ca_check` (Boolean) Enable the Certificate Authority (CA) check.
- `cn_check` (Boolean) Enable the Common Name (CN) check.
- `port` (Number) Enter the port number that the job must use to discover the devices.
- `retries` (Number) Enter the number of repeated attempts required to discover a device
- `timeout` (Number) Enter the time in seconds after which a job must stop running.


<a id="nestedatt--discovery_config_targets--snmp"></a>
### Nested Schema for `discovery_config_targets.snmp`

Required:

- `community` (String) Community string for the SNMP protocol.

Optional:

- `port` (Number) Enter the port number that the job must use to discover the devices.
- `retries` (Number) Enter the number of repeated attempts required to discover a device.
- `timeout` (Number) Enter the time in seconds after which a job must stop running.


<a id="nestedatt--discovery_config_targets--ssh"></a>
### Nested Schema for `discovery_config_targets.ssh`

Required:

- `password` (String) Provide a password for the protocol.
- `username` (String) Provide a username for the protocol.

Optional:

- `check_known_hosts` (Boolean) Verify the known host key.
- `is_sudo_user` (Boolean) Use the SUDO option
- `port` (Number) Enter the port number that the job must use to discover the devices.
- `retries` (Number) Enter the number of repeated attempts required to discover a device.
- `timeout` (Number) Enter the time in seconds after which a job must stop running.


<a id="nestedatt--discovery_config_targets--wsman"></a>
### Nested Schema for `discovery_config_targets.wsman`

Required:

- `password` (String) Provide a password for the protocol.
- `username` (String) Provide a username for the protocol.

Optional:

- `ca_check` (Boolean) Enable the Certificate Authority (CA) check.
- `cn_check` (Boolean) Enable the Common Name (CN) check.
- `port` (Number) Enter the port number that the job must use to discover the devices.
- `retries` (Number) Enter the number of repeated attempts required to discover a device
- `timeout` (Number) Enter the time in seconds after which a job must stop running.

