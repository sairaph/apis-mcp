---
title: posture-api_Vendor
page_id: schema-posture-api-vendor-ad245513
path: schemas
description: Information about a vendor/service provider.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# posture-api_Vendor

Information about a vendor/service provider.

```yaml
{"description": "Information about a vendor/service provider.", "type": "object", "properties": {"description": {"description": "Detailed information about what kinds of issues are detected for this vendor.", "type": "string", "example": "Identify important security issues across your Google Workspace account ranging from shadow IT, misconfigurations, user access, and more.", "nullable": true}, "display_name": {"description": "The display name of the vendor.", "type": "string", "example": "Google Workspace"}, "id": {"description": "The id of the vendor.", "type": "string", "example": "R09PR0xFX1dPUktTUEFDRQ=="}, "logo": {"description": "Logo URL for the vendor.", "type": "string", "format": "uri", "example": "https://cdn.vectrix-infra.com/DetectionPack_Logos/GoogleWorkspace/g.png"}, "name": {"description": "The name of the vendor.", "type": "string", "example": "GOOGLE_WORKSPACE"}, "policies": {"description": "The policies related to the vendor.", "type": "array", "items": {"additionalProperties": true, "type": "object"}, "example": []}, "static_logo": {"description": "Static logo URL for the vendor.", "type": "string", "format": "uri", "example": "https://onprem.cloudflare.come/DetectionPack_Logos/GoogleWorkspace/g.png"}, "zt_enrollments": {"description": "The vendor's compatible Zero Trust products.", "type": "array", "items": {"type": "string"}, "example": ["casb"]}}, "required": ["description", "display_name", "id", "logo", "name", "static_logo", "zt_enrollments"]}
```
