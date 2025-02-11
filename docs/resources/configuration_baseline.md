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

title: "ome_configuration_baseline resource"
linkTitle: "ome_configuration_baseline"
page_title: "ome_configuration_baseline Resource - terraform-provider-ome"
subcategory: ""
description: |-
  Resource for managing configuration baselines on OpenManage Enterprise.
---

# ome_configuration_baseline (Resource)

Resource for managing configuration baselines on OpenManage Enterprise.

~> **Note:** Exactly one of `ref_template_name` and `ref_template_id` and exactly one of `device_ids` and `device_servicetags` are required.

~> **Note:** When `schedule` is `true`, following parameters are considered: `notify_on_schedule`, `cron`, `email_addresses`, `output_format`.

~> **Note:** Updates are supported for all the parameters.

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

# Manage baseline using Device Servicetags
resource "ome_configuration_baseline" "baseline_name" {
  baseline_name      = "Baseline Name"
  device_servicetags = ["MXL1234", "MXL1235"]
}

# Create Baseline using device ids
resource "ome_configuration_baseline" "baseline1" {
  baseline_name   = "baseline1"
  ref_template_id = 745
  device_ids      = [10001, 10002]
  description     = "baseline description"
}


# Create Baseline using device servicetag with daily notification scheduled 
resource "ome_configuration_baseline" "baseline2" {
  baseline_name      = "baseline2"
  ref_template_id    = 745
  device_servicetags = ["MXL1234", "MXL1235"]
  description        = "baseline description"
  schedule           = true
  notify_on_schedule = true
  email_addresses    = ["test@testmail.com"]
  cron               = "0 30 11 * * ? *"
  output_format      = "csv"
}


# Create Baseline using device ids with daily notification on status changing to non-compliant 
resource "ome_configuration_baseline" "baseline3" {
  baseline_name   = "baseline3"
  ref_template_id = 745
  device_ids      = [10001, 10002]
  description     = "baseline description"
  schedule        = true
  email_addresses = ["test@testmail.com"]
  output_format   = "pdf"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `baseline_name` (String) Name of the Baseline.

### Optional

- `cron` (String) Cron expression for notification schedule. Can be set only when both `schedule` and `notify_on_schedule` are set to `true`.
- `description` (String) Description of the baseline.
- `device_ids` (Set of Number) List of the device id on which the baseline compliance needs to be run. Conflicts with `device_servicetags`.
- `device_servicetags` (Set of String) List of the device servicetag on which the baseline compliance needs to be run. Conflicts with `device_ids`.
- `email_addresses` (Set of String) Email addresses for notification. Can be set only when `schedule` is `true`.
- `job_retry_count` (Number) Number of times the job has to be polled to get the final status of the resource. Default value is `30`.
- `notify_on_schedule` (Boolean) Schedule notification via cron or any time the baseline becomes non-compliant. Default value is `false`.
- `output_format` (String) Output format type, the input is case senitive. Valid values are `html`, `csv`, `pdf`and `xls`. Default value is `html`.
- `ref_template_id` (Number) Reference template ID. Conflicts with `ref_template_name`.
- `ref_template_name` (String) Reference template name. Conflicts with `ref_template_id`.
- `schedule` (Boolean) Schedule notification via email. Default value is `false`.
- `sleep_interval` (Number) Sleep time interval for job polling in seconds. Default value is `20`.

### Read-Only

- `id` (Number) ID of the configuration baseline resource.
- `task_id` (Number) Task id associated with baseline.

## Import

Import is supported using the following syntax:

```shell
# /*
# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://mozilla.org/MPL/2.0/
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# */

terraform import ome_configuration_baseline.create_baseline "<existing_baseline_name>"
```