---
title: aaa_unsubscribe_email_post
page_id: schema-aaa-unsubscribe-email-post-76ad462e
path: schemas
description: Response body for the POST unsubscribe endpoint. name and token are not returned.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_unsubscribe_email_post

Response body for the POST unsubscribe endpoint. name and token are not returned.

```yaml
{"description": "Response body for the POST unsubscribe endpoint. name and token are not returned.", "type": "object", "properties": {"account_id": {"$ref": "#/components/schemas/aaa_account-id"}, "email": {"type": "string", "format": "email"}, "id": {"$ref": "#/components/schemas/aaa_policy-id"}}}
```
