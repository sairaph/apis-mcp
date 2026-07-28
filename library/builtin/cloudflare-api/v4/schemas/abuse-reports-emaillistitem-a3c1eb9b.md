---
title: abuse-reports_EmailListItem
page_id: schema-abuse-reports-emaillistitem-a3c1eb9b
path: schemas
description: An email sent to the customer for an abuse report.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# abuse-reports_EmailListItem

An email sent to the customer for an abuse report.

```yaml
{"description": "An email sent to the customer for an abuse report.", "type": "object", "properties": {"body": {"description": "Body content of the email.", "type": "string"}, "id": {"description": "Unique identifier of the email.", "type": "string"}, "recipient": {"description": "Email address of the recipient.", "type": "string"}, "sent_at": {"description": "When the email was sent. Time in RFC 3339 format (https://www.rfc-editor.org/rfc/rfc3339.html)", "type": "string", "example": "2009-11-10T23:00:00Z"}, "subject": {"description": "Subject line of the email.", "type": "string"}}, "required": ["id", "subject", "body", "recipient", "sent_at"]}
```
