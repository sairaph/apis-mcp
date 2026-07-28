---
title: resource-tagging_etag
page_id: schema-resource-tagging-etag-1945e0cf
path: schemas
description: |-
    ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>" where
    the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the tags map
    canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients should treat
    ETags as opaque strings and pass them back via the If-Match header on write operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_etag

ETag identifier for optimistic concurrency control. Formatted as "v1:<hash>" where
the hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the tags map
canonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients should treat
ETags as opaque strings and pass them back via the If-Match header on write operations.

```yaml
{"description": "ETag identifier for optimistic concurrency control. Formatted as \"v1:<hash>\" where\nthe hash is the base64url-encoded SHA-256 (truncated to 128 bits) of the tags map\ncanonicalized using RFC 8785 (JSON Canonicalization Scheme). Clients should treat\nETags as opaque strings and pass them back via the If-Match header on write operations.\n", "type": "string", "example": "v1:RBNvo1WzZ4oRRq0W9-hkng", "readOnly": true}
```
