---
title: posture-api_ZeroTrustProduct
page_id: schema-posture-api-zerotrustproduct-93acb743
path: schemas
description: Information about a Zero Trust product integration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_ZeroTrustProduct

Information about a Zero Trust product integration.

```yaml
{"description": "Information about a Zero Trust product integration.", "type": "object", "properties": {"description": {"description": "Brief description of the Zero Trust Product.", "type": "string", "example": "example"}, "display_name": {"description": "The verbose name of the Zero Trust Product.", "type": "string", "example": "Cloud Access Security Broker"}, "enabled": {"description": "Flag to enable/disable access to the listed integration from the corresponding Cloudflare product.", "type": "boolean", "example": true, "default": false}, "id": {"description": "The internal identifier of the Zero Trust Product.", "type": "string", "example": "casb"}}, "example": {"description": "example", "display_name": "Cloud Access Security Broker", "enabled": true, "id": "casb"}}
```
