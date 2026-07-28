---
title: cache-rules_origin_cloud_region_request
page_id: schema-cache-rules-origin-cloud-region-request-b46b5489
path: schemas
description: Request body for creating or updating an origin cloud region mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_request

Request body for creating or updating an origin cloud region mapping.

```yaml
{"description": "Request body for creating or updating an origin cloud region mapping.", "type": "object", "properties": {"ip": {"description": "Origin IP address (IPv4 or IPv6). Normalized to canonical form before storage (RFC 5952 for IPv6).", "type": "string", "example": "192.0.2.1", "x-auditable": true}, "region": {"description": "Cloud vendor region identifier. Must be a valid region for the specified vendor as returned by the supported_regions endpoint.", "type": "string", "example": "us-east-1", "x-auditable": true}, "vendor": {"description": "Cloud vendor hosting the origin. Must be one of the supported vendors.", "type": "string", "example": "aws", "enum": ["aws", "azure", "gcp", "oci"], "x-auditable": true}}, "required": ["ip", "vendor", "region"]}
```
