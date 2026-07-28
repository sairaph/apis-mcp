---
title: cache-rules_origin_cloud_region_v2_request
page_id: schema-cache-rules-origin-cloud-region-v2-request-54e81371
path: schemas
description: Request body for creating or replacing an origin cloud region mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_v2_request

Request body for creating or replacing an origin cloud region mapping.

```yaml
{"description": "Request body for creating or replacing an origin cloud region mapping.", "type": "object", "properties": {"origin_ip": {"description": "Origin IP address (IPv4 or IPv6). For the single PUT endpoint (`PUT /origin/cloud_regions/{origin_ip}`), this field must match the path parameter or the request will be rejected with a 400 error. For the batch PUT endpoint, this field identifies which mapping to upsert.", "type": "string", "example": "192.0.2.1", "x-auditable": true}, "region": {"description": "Cloud vendor region identifier. Must be a valid region for the specified vendor as returned by the supported_regions endpoint.", "type": "string", "example": "us-east-1", "x-auditable": true}, "vendor": {"description": "Cloud vendor hosting the origin. Must be one of the supported vendors.", "type": "string", "example": "aws", "enum": ["aws", "azure", "gcp", "oci"], "x-auditable": true}}, "required": ["origin_ip", "vendor", "region"]}
```
