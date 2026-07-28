---
title: access_app_resp_embedded_policies
page_id: schema-access-app-resp-embedded-policies-e6388edf
path: schemas
description: The policies that Access applies to the application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_app_resp_embedded_policies

The policies that Access applies to the application.

```yaml
{"description": "The policies that Access applies to the application.", "type": "object", "properties": {"policies": {"type": "array", "items": {"$ref": "#/components/schemas/access_app_policy_response"}}}}
```
