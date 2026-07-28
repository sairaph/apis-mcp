---
title: digital-experience-monitoring_ip_info
page_id: schema-digital-experience-monitoring-ip-info-b2535cf0
path: schemas
description: IP address information for the ISP hop. Fields marked as PII-gated (`name`, `address`, `netmask`, and all `location` sub-fields) will be returned as the literal string `"REDACTED"` for callers that do not have the PII permission. `asn`, `aso`, and `version` are always returned regardless of PII access.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_ip_info

IP address information for the ISP hop. Fields marked as PII-gated (`name`, `address`, `netmask`, and all `location` sub-fields) will be returned as the literal string `"REDACTED"` for callers that do not have the PII permission. `asn`, `aso`, and `version` are always returned regardless of PII access.

```yaml
{"description": "IP address information for the ISP hop. Fields marked as PII-gated (`name`, `address`, `netmask`, and all `location` sub-fields) will be returned as the literal string `\"REDACTED\"` for callers that do not have the PII permission. `asn`, `aso`, and `version` are always returned regardless of PII access.\n", "type": "object", "properties": {"address": {"description": "IP address. Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "203.0.113.1"}, "asn": {"description": "Autonomous System Number.", "type": "integer", "example": 13335}, "aso": {"description": "Autonomous System Organization name.", "type": "string", "example": "CLOUDFLARENET"}, "location": {"$ref": "#/components/schemas/digital-experience-monitoring_ip_location"}, "name": {"description": "Named IP address (reverse DNS hostname when available). Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "isp-gateway.example.com"}, "netmask": {"description": "Network mask. Returned as `\"REDACTED\"` without PII permission.", "type": "string", "example": "255.255.255.0"}, "version": {"description": "IP version (`1` for IPv4, `2` for IPv6, `0` if unknown).", "type": "integer", "example": 1}}}
```
