---
title: iam_oauth_client_update_request
page_id: schema-iam-oauth-client-update-request-c451df89
path: schemas
description: Partial update request for an OAuth client. Only include fields you want to update.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_update_request

Partial update request for an OAuth client. Only include fields you want to update.

```yaml
{"description": "Partial update request for an OAuth client. Only include fields you want to update.", "allOf": [{"$ref": "#/components/schemas/iam_oauth_client_common"}, {"properties": {"visibility": {"description": "Promote the OAuth client from private to public visibility. Only `public` is accepted; demotion to `private` is not supported. Promotion requires a non-empty client name, logo URI, verified client URI host, and at least one non-identity scope.", "type": "string", "example": "public", "enum": ["public"]}}, "type": "object"}], "title": "Update OAuth Client Request"}
```
