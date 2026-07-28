---
title: teams-devices_split_tunnel_address
page_id: schema-teams-devices-split-tunnel-address-c839c5e2
path: schemas
description: The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_split_tunnel_address

The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.

```yaml
{"description": "The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.", "type": "string", "example": "192.0.2.0/24", "x-stainless-terraform-configurability": "computed_optional"}
```
