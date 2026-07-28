---
title: spectrum-config_edge_ips
page_id: schema-spectrum-config-edge-ips-df3669e4
path: schemas
description: The anycast edge IP configuration for the hostname of this application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-config_edge_ips

The anycast edge IP configuration for the hostname of this application.

```yaml
{"description": "The anycast edge IP configuration for the hostname of this application.", "default": {"connectivity": "all", "type": "dynamic"}, "oneOf": [{"properties": {"connectivity": {"description": "The IP versions supported for inbound connections on Spectrum anycast IPs.", "type": "string", "example": "all", "enum": ["all", "ipv4", "ipv6"]}, "type": {"description": "The type of edge IP configuration specified. Dynamically allocated edge IPs use Spectrum anycast IPs in accordance with the connectivity you specify. Only valid with CNAME DNS names.", "type": "string", "example": "dynamic", "enum": ["dynamic"]}}, "type": "object"}, {"properties": {"ips": {"description": "The array of customer owned IPs we broadcast via anycast for this hostname and application.", "type": "array", "items": {"description": "Edge anycast IPs.", "example": "192.0.2.1", "type": "string"}, "example": ["192.0.2.1"]}, "type": {"description": "The type of edge IP configuration specified. Statically allocated edge IPs use customer IPs in accordance with the ips array you specify. Only valid with ADDRESS DNS names.", "type": "string", "example": "static", "enum": ["static"]}}, "type": "object"}]}
```
