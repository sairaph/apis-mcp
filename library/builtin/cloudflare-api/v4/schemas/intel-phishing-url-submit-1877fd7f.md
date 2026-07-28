---
title: intel_phishing-url-submit
page_id: schema-intel-phishing-url-submit-1877fd7f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel_phishing-url-submit

```yaml
{"type": "object", "properties": {"excluded_urls": {"description": "URLs that were excluded from scanning because their domain is in our no-scan list.", "type": "array", "items": {"properties": {"url": {"description": "URL that was excluded.", "type": "string", "example": "https://developers.cloudflare.com", "x-auditable": true}}, "type": "object"}}, "skipped_urls": {"description": "URLs that were skipped because the same URL is currently being scanned.", "type": "array", "items": {"properties": {"url": {"description": "URL that was skipped.", "type": "string", "example": "https://www.cloudflare.com/developer-week/", "x-auditable": true}, "url_id": {"description": "ID of the submission of that URL that is currently scanning.", "type": "integer", "example": 2, "x-auditable": true}}, "type": "object"}}, "submitted_urls": {"description": "URLs that were successfully submitted for scanning.", "type": "array", "items": {"properties": {"url": {"description": "URL that was submitted.", "type": "string", "example": "https://www.cloudflare.com", "x-auditable": true}, "url_id": {"description": "ID assigned to this URL submission. Used to retrieve scanning results.", "type": "integer", "example": 1, "x-auditable": true}}, "type": "object"}}}}
```
