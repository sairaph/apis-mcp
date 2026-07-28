---
title: teams-devices_os_version_input_request
page_id: schema-teams-devices-os-version-input-request-a7bae076
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_os_version_input_request

```yaml
{"type": "object", "properties": {"operating_system": {"description": "Operating System.", "type": "string", "example": "windows", "enum": ["windows"], "x-auditable": true}, "operator": {"description": "Operator.", "type": "string", "example": "13.3.0", "enum": ["<", "<=", ">", ">=", "=="], "x-auditable": true}, "os_distro_name": {"description": "Operating System Distribution Name (linux only).", "type": "string", "example": "ubuntu", "x-auditable": true}, "os_distro_revision": {"description": "Version of OS Distribution (linux only).", "type": "string", "example": "11.3.1", "x-auditable": true}, "os_version_extra": {"description": "Additional operating system version details. For Windows, the UBR (Update Build Revision). For Mac or iOS, the Product Version Extra. For Linux, the distribution name and version.", "type": "string", "example": "(a) or 6889 or Ubuntu 24.04", "x-auditable": true}, "version": {"description": "Version of OS.", "type": "string", "example": "13.3.0", "x-auditable": true}}, "required": ["operating_system", "version", "operator"], "title": "OS Version"}
```
