---
title: cc_InstanceType
page_id: schema-cc-instancetype-88865f89
path: schemas
description: |-
    The instance type will be used to configure vCPU, memory, and disk.

    - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
    - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
    - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
    - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
    - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
    - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_InstanceType

The instance type will be used to configure vCPU, memory, and disk.

- "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
- "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
- "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
- "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
- "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
- "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk

```yaml
{"description": "The instance type will be used to configure vCPU, memory, and disk.\n\n- \"lite\": 1/16 vCPU, 256 MiB memory, 2 GB disk\n- \"basic\": 1/4 vCPU, 1 GiB memory, 4 GB disk\n- \"standard-1\": 1/2 vCPU, 4 GiB memory, 8 GB disk\n- \"standard-2\": 1 vCPU, 6 GiB memory, 12 GB disk\n- \"standard-3\": 2 vCPU, 8 GiB memory, 16 GB disk\n- \"standard-4\": 4 vCPU, 12 GiB memory, 20 GB disk\n", "type": "string", "example": "lite", "default": "lite", "anyOf": [{"enum": ["lite", "basic", "standard-1", "standard-2", "standard-3", "standard-4"], "type": "string"}]}
```
