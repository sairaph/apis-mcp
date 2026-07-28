---
title: waf-product-api-bundle_custom-detection
page_id: schema-waf-product-api-bundle-custom-detection-489c7a32
path: schemas
description: Defines a custom set of username/password expressions to match Leaked Credential Checks on.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-product-api-bundle_custom-detection

Defines a custom set of username/password expressions to match Leaked Credential Checks on.

```yaml
{"description": "Defines a custom set of username/password expressions to match Leaked Credential Checks on.", "type": "object", "properties": {"id": {"$ref": "#/components/schemas/waf-product-api-bundle_detection-id"}, "password": {"description": "Defines ehe ruleset expression to use in matching the password in a request.", "type": "string", "example": "lookup_json_string(http.request.body.raw, \"secret\")", "x-auditable": true}, "username": {"description": "Defines the ruleset expression to use in matching the username in a request.", "type": "string", "example": "lookup_json_string(http.request.body.raw, \"user\")", "x-auditable": true}}}
```
