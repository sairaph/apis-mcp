---
title: iam_oauth_client
page_id: schema-iam-oauth-client-1fd1407a
path: schemas
description: An OAuth 2.0 client registration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client

An OAuth 2.0 client registration.

```yaml
{"description": "An OAuth 2.0 client registration.", "allOf": [{"$ref": "#/components/schemas/iam_oauth_client_common"}, {"properties": {"client_id": {"allOf": [{"$ref": "#/components/schemas/iam_oauth_client_identifier"}], "readOnly": true}, "client_uri_verification": {"$ref": "#/components/schemas/iam_oauth_client_uri_verification"}, "created_at": {"description": "Timestamp when the OAuth client was created.", "type": "string", "format": "date-time", "example": "2025-01-01T00:00:00Z", "readOnly": true, "x-auditable": true}, "has_rotated_secret": {"description": "Indicates whether the client has a rotated secret that has not yet been deleted.", "type": "boolean", "example": false, "readOnly": true}, "promoted_at": {"description": "Timestamp when the OAuth client was promoted to public visibility.", "type": "string", "format": "date-time", "example": "2026-05-13T12:00:00Z", "readOnly": true, "x-auditable": true}, "updated_at": {"description": "Timestamp when the OAuth client was last updated.", "type": "string", "format": "date-time", "example": "2025-01-01T00:00:00Z", "readOnly": true, "x-auditable": true}, "visibility": {"description": "Visibility of the OAuth client.", "type": "string", "example": "private", "enum": ["public", "private"], "readOnly": true, "x-auditable": true, "x-stainless-terraform-configurability": "computed_optional"}}, "required": ["client_id", "visibility"], "type": "object"}], "title": "OAuth Client"}
```
