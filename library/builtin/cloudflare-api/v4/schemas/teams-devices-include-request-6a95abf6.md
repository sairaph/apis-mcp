---
title: teams-devices_include_request
page_id: schema-teams-devices-include-request-6a95abf6
path: schemas
description: List of routes included in the WARP client's tunnel. Both 'exclude' and 'include' cannot be set in the same request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_include_request

List of routes included in the WARP client's tunnel. Both 'exclude' and 'include' cannot be set in the same request.

```yaml
{"description": "List of routes included in the WARP client's tunnel. Both 'exclude' and 'include' cannot be set in the same request.", "type": "array", "items": {"$ref": "#/components/schemas/teams-devices_split_tunnel_include"}, "x-stainless-terraform-configurability": "computed_optional"}
```
