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

listIgnoreTitle: true
weight: 1
title: "ome provider"
linkTitle: "Provider"
page_title: "ome Provider"
subcategory: ""
description: |-
  The Terraform Provider for OpenManage Enterprise (OME) is a plugin for Terraform that allows the resource management of PowerEdge servers using OME
---

# ome Provider

The Terraform Provider for OpenManage Enterprise (OME) is a plugin for Terraform that allows the resource management of PowerEdge servers using OME

## Example Usage

```terraform
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

terraform {
  required_providers {
    ome = {
      version = "1.0.0"
      source  = "registry.terraform.io/dell/ome"
    }
  }
}

provider "ome" {
  username = "username"
  password = "password"
  host     = "yourhost.host.com"
  skipssl  = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `host` (String) OpenManage Enterprise IP address or hostname.
- `password` (String, Sensitive) OpenManage Enterprise password.
- `username` (String) OpenManage Enterprise username.

### Optional

- `port` (Number) OpenManage Enterprise HTTPS port. Default value is `443`.
- `skipssl` (Boolean) Skips SSL certificate validation on OpenManage Enterprise. Default value is `false`.
- `timeout` (Number) HTTPS timeout in seconds for OpenManage Enterprise client. Default value is `30`.