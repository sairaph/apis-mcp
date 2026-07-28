---
title: cache-rules_origin_cloud_region_v2_entry
page_id: schema-cache-rules-origin-cloud-region-v2-entry-e4a7b382
path: schemas
description: A single origin IP-to-cloud-region mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cache-rules_origin_cloud_region_v2_entry

A single origin IP-to-cloud-region mapping.

```yaml
{"description": "A single origin IP-to-cloud-region mapping.", "type": "object", "properties": {"modified_on": {"description": "Time this mapping was last modified.", "type": "string", "format": "date-time", "x-auditable": true}, "origin_ip": {"description": "The origin IP address (IPv4 or IPv6). Normalized to canonical form (RFC 5952 for IPv6).", "type": "string", "example": "192.0.2.1", "x-auditable": true}, "region": {"description": "Cloud vendor region identifier.", "type": "string", "example": "us-east-1", "x-auditable": true}, "vendor": {"description": "Cloud vendor hosting the origin.", "type": "string", "example": "aws", "enum": ["aws", "azure", "gcp", "oci"], "x-auditable": true}}, "required": ["origin_ip", "vendor", "region"]}
```
