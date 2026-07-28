---
title: email-auth_ApprovedSourceResponse
page_id: schema-email-auth-approvedsourceresponse-e18d249b
path: schemas
description: A single approved sending source
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email-auth_ApprovedSourceResponse

A single approved sending source

```yaml
{"description": "A single approved sending source", "type": "object", "properties": {"created": {"description": "Deprecated, use created_at", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00.12345Z", "deprecated": true, "x-stainless-deprecation-message": "Use `created_at` instead."}, "created_at": {"description": "Creation timestamp", "type": "string", "format": "date-time", "example": "2024-01-15T10:30:00.12345Z"}, "domain": {"description": "The source domain", "type": "string", "example": "sendgrid.net"}, "ips": {"description": "Resolved IP addresses from SPF", "type": "array", "items": {"type": "string"}, "example": ["192.168.1.1", "10.0.0.1"]}, "modified": {"description": "Deprecated, use modified_at", "type": "string", "format": "date-time", "example": "2024-01-15T11:45:00.12345Z", "deprecated": true, "x-stainless-deprecation-message": "Use `modified_at` instead."}, "modified_at": {"description": "Last modification timestamp", "type": "string", "format": "date-time", "example": "2024-01-15T11:45:00.12345Z"}, "name": {"description": "Source name (typically same as domain)", "type": "string", "example": "SendGrid"}, "slug": {"description": "URL-friendly identifier", "type": "string", "example": "sendgrid-net"}, "tag": {"description": "Source UUID", "type": "string", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}}}
```
