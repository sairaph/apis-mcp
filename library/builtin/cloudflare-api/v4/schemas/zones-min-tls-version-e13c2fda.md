---
title: zones_min_tls_version
page_id: schema-zones-min-tls-version-e13c2fda
path: schemas
description: Only accepts HTTPS requests that use at least the TLS protocol version specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_min_tls_version

Only accepts HTTPS requests that use at least the TLS protocol version specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.

```yaml
{"description": "Only accepts HTTPS requests that use at least the TLS protocol version specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.", "default": "1.0", "allOf": [{"$ref": "#/components/schemas/zones_base"}, {"properties": {"id": {"description": "ID of the zone setting.", "example": "min_tls_version", "enum": ["min_tls_version"]}, "value": {"$ref": "#/components/schemas/zones_min_tls_version_value"}}}], "title": "Zone Minimum TLS Version value"}
```
