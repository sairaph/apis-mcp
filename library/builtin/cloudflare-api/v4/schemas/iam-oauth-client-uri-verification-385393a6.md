---
title: iam_oauth_client_uri_verification
page_id: schema-iam-oauth-client-uri-verification-385393a6
path: schemas
description: Client URI domain control verification state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# iam_oauth_client_uri_verification

Client URI domain control verification state.

```yaml
{"description": "Client URI domain control verification state.", "type": "object", "properties": {"status": {"description": "Current verification status for the client URI host.", "type": "string", "example": "in_progress", "enum": ["pending", "in_progress", "verified", "failed"]}, "text": {"description": "Exact TXT record value that must be added to DNS to prove ownership of the client URI host.", "type": "string", "example": "cloudflare_oauth_client_publisher=example"}}, "title": "OAuth Client URI Verification"}
```
