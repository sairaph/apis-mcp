---
title: workers_assets
page_id: schema-workers-assets-f28cfb65
path: schemas
description: Configuration for assets within a Worker.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_assets

Configuration for assets within a Worker.

```yaml
{"description": "Configuration for assets within a Worker.", "type": "object", "properties": {"config": {"description": "Configuration for assets within a Worker.", "type": "object", "properties": {"_headers": {"description": "The contents of a _headers file (used to attach custom headers on asset responses).", "type": "string", "example": "/dashboard/*\nX-Frame-Options: DENY\n\n/static/*\nAccess-Control-Allow-Origin: *"}, "_redirects": {"description": "The contents of a _redirects file (used to apply redirects or proxy paths ahead of asset serving).", "type": "string", "example": "/foo /bar 301\n/news/* /blog/:splat"}, "html_handling": {"description": "Determines the redirects and rewrites of requests for HTML content.", "type": "string", "example": "auto-trailing-slash", "enum": ["auto-trailing-slash", "force-trailing-slash", "drop-trailing-slash", "none"]}, "not_found_handling": {"description": "Determines the response when a request does not match a static asset, and there is no Worker script.", "type": "string", "example": "404-page", "enum": ["none", "404-page", "single-page-application"]}, "run_worker_first": {"oneOf": [{"description": "Contains a list path rules to control routing to either the Worker or assets. Glob (*) and negative (!) rules are supported. Rules must start with either '/' or '!/'. At least one non-negative rule must be provided, and negative rules have higher precedence than non-negative rules.", "example": "[\"/api/*\", \"/oauth/callback\", \"!/api/assets/*\"]", "items": {"type": "string"}, "type": "array"}, {"description": "Enables routing to always invoke the Worker script ahead of all requests. When true, this is equivalent to `[\"/*\"]` in the string array version of this field.", "example": true, "type": "boolean"}]}, "serve_directly": {"description": "When true and the incoming request matches an asset, that will be served instead of invoking the Worker script. When false, requests will always invoke the Worker script.", "type": "boolean", "example": true, "deprecated": true}}}, "jwt": {"description": "Token provided upon successful upload of all files from a registered manifest.", "type": "string", "x-sensitive": true}}}
```
