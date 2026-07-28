---
title: teams-devices_input
page_id: schema-teams-devices-input-c1c355b3
path: schemas
description: The value to be checked against.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_input

The value to be checked against.

```yaml
{"description": "The value to be checked against.", "type": "object", "example": {"operating_system": "linux", "path": "/bin/cat", "thumbprint": "0aabab210bdb998e9cf45da2c9ce352977ab531c681b74cf1e487be1bbe9fe6e"}, "oneOf": [{"$ref": "#/components/schemas/teams-devices_file_input_request"}, {"$ref": "#/components/schemas/teams-devices_unique_client_id_input_request"}, {"$ref": "#/components/schemas/teams-devices_domain_joined_input_request"}, {"$ref": "#/components/schemas/teams-devices_os_version_input_request"}, {"$ref": "#/components/schemas/teams-devices_firewall_input_request"}, {"$ref": "#/components/schemas/teams-devices_sentinelone_input_request"}, {"$ref": "#/components/schemas/teams-devices_carbonblack_input_request"}, {"$ref": "#/components/schemas/teams-devices_access_serial_number_list_input_request"}, {"$ref": "#/components/schemas/teams-devices_disk_encryption_input_request"}, {"$ref": "#/components/schemas/teams-devices_application_input_request"}, {"$ref": "#/components/schemas/teams-devices_client_certificate_input_request"}, {"$ref": "#/components/schemas/teams-devices_client_certificate_v2_input_request"}, {"$ref": "#/components/schemas/teams-devices_antivirus_input_request"}, {"$ref": "#/components/schemas/teams-devices_workspace_one_input_request"}, {"$ref": "#/components/schemas/teams-devices_crowdstrike_input_request"}, {"$ref": "#/components/schemas/teams-devices_intune_input_request"}, {"$ref": "#/components/schemas/teams-devices_kolide_input_request"}, {"$ref": "#/components/schemas/teams-devices_tanium_input_request"}, {"$ref": "#/components/schemas/teams-devices_sentinelone_s2s_input_request"}, {"$ref": "#/components/schemas/teams-devices_custom_s2s_input_request"}]}
```
