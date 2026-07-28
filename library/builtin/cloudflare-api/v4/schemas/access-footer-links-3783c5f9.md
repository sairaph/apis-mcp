---
title: access_footer_links
page_id: schema-access-footer-links-3783c5f9
path: schemas
description: The links in the App Launcher footer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_footer_links

The links in the App Launcher footer.

```yaml
{"description": "The links in the App Launcher footer.", "type": "array", "items": {"properties": {"name": {"description": "The hypertext in the footer link.", "type": "string", "example": "Cloudflare's Privacy Policy"}, "url": {"description": "the hyperlink in the footer link.", "type": "string", "example": "https://www.cloudflare.com/privacypolicy/"}}, "required": ["name", "url"], "type": "object"}, "example": [{"name": "Cloudflare's Privacy Policy", "url": "https://www.cloudflare.com/privacypolicy/"}]}
```
