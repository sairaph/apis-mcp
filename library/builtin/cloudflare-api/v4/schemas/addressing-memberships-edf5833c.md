---
title: addressing_memberships
page_id: schema-addressing-memberships-edf5833c
path: schemas
description: Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# addressing_memberships

Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.

```yaml
{"description": "Zones and Accounts which will be assigned IPs on this Address Map. A zone membership will take priority over an account membership.", "type": "array", "items": {"$ref": "#/components/schemas/addressing_address-maps-membership"}}
```
