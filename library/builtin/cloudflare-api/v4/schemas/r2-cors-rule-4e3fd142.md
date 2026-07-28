---
title: r2_cors-rule
page_id: schema-r2-cors-rule-4e3fd142
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_cors-rule

```yaml
{"type": "object", "properties": {"allowed": {"description": "Object specifying allowed origins, methods and headers for this CORS rule.", "type": "object", "properties": {"headers": {"description": "Specifies the value for the Access-Control-Allow-Headers header R2 sets when requesting objects in this bucket from a browser. Cross-origin requests that include custom headers (e.g. x-user-id) should specify these headers as AllowedHeaders.", "type": "array", "items": {"example": "x-requested-by", "type": "string", "x-auditable": true}}, "methods": {"description": "Specifies the value for the Access-Control-Allow-Methods header R2 sets when requesting objects in a bucket from a browser.", "type": "array", "items": {"enum": ["GET", "PUT", "POST", "DELETE", "HEAD"], "type": "string", "x-auditable": true}}, "origins": {"description": "Specifies the value for the Access-Control-Allow-Origin header R2 sets when requesting objects in a bucket from a browser.", "type": "array", "items": {"example": "http://localhost:3000", "type": "string", "x-auditable": true}}}, "required": ["methods", "origins"]}, "exposeHeaders": {"description": "Specifies the headers that can be exposed back, and accessed by, the JavaScript making the cross-origin request. If you need to access headers beyond the safelisted response headers, such as Content-Encoding or cf-cache-status, you must specify it here.", "type": "array", "items": {"example": "Content-Encoding", "type": "string", "x-auditable": true}}, "id": {"description": "Identifier for this rule.", "type": "string", "example": "Allow Local Development", "x-auditable": true}, "maxAgeSeconds": {"description": "Specifies the amount of time (in seconds) browsers are allowed to cache CORS preflight responses. Browsers may limit this to 2 hours or less, even if the maximum value (86400) is specified.", "type": "number", "example": 3600, "x-auditable": true}}, "required": ["allowed"]}
```
