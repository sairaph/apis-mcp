---
title: teams-devices_client_version
page_id: schema-teams-devices-client-version-03cdfea2
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_client_version

```yaml
{"type": "object", "properties": {"package_size": {"description": "Size of the package in bytes.", "type": "integer", "format": "int64", "example": 123125760}, "package_url": {"description": "URL to download the package.", "type": "string", "example": "https://downloads.cloudflareclient.com/v1/download/windows/version/2024.11.309.0"}, "release_date": {"description": "The release date timestamp.", "type": "string", "example": "2024-11-18T21:57:58.478Z"}, "release_notes": {"description": "Release notes for this version.", "type": "string", "example": "This release contains minor fixes and improvements."}, "version": {"description": "The client version string.", "type": "string", "example": "2024.11.309.0"}}, "required": ["version", "release_date"]}
```
