---
title: cc_ModifyUserDeploymentConfiguration
page_id: schema-cc-modifyuserdeploymentconfiguration-f00a0a5e
path: schemas
description: Properties required to modify a cloudchamber deployment specified by the user.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_ModifyUserDeploymentConfiguration

Properties required to modify a cloudchamber deployment specified by the user.

```yaml
{"description": "Properties required to modify a cloudchamber deployment specified by the user.", "type": "object", "properties": {"authorized_keys": {"type": "array", "items": {"$ref": "#/components/schemas/cc_UserSSHPublicKey"}}, "command": {"$ref": "#/components/schemas/cc_Command"}, "disk": {"$ref": "#/components/schemas/cc_Disk"}, "dns": {"$ref": "#/components/schemas/cc_DNSConfiguration"}, "entrypoint": {"$ref": "#/components/schemas/cc_Entrypoint"}, "environment_variables": {"description": "Container environment variables", "type": "array", "items": {"$ref": "#/components/schemas/cc_EnvironmentVariable"}}, "experimental_flags": {"description": "Opt-in experimental flags for this application. Only a subset of\nexperimental flags can be set by users; unsupported values are rejected.\n", "type": "array", "items": {"type": "string"}, "maxItems": 10}, "image": {"$ref": "#/components/schemas/cc_Image"}, "instance_type": {"$ref": "#/components/schemas/cc_InstanceType"}, "labels": {"description": "Deployment labels", "type": "array", "items": {"$ref": "#/components/schemas/cc_Label"}}, "lifecycle": {"$ref": "#/components/schemas/cc_DeploymentLifecycle"}, "memory": {"$ref": "#/components/schemas/cc_MemorySizeWithUnit"}, "memory_mib": {"description": "Specify the memory to be used for the deployment, in MiB. The default will be the one configured for the account.", "type": "integer"}, "metadata_service": {"$ref": "#/components/schemas/cc_MetadataService"}, "observability": {"$ref": "#/components/schemas/cc_DeploymentObservability"}, "secrets": {"description": "A list of objects with secret names and the their access types from the account", "type": "array", "items": {"$ref": "#/components/schemas/cc_DeploymentSecretMap"}}, "ssh_public_key_ids": {"description": "A list of SSH public key IDs from the account", "type": "array", "items": {"$ref": "#/components/schemas/cc_SSHPublicKeyID"}}, "trusted_user_ca_keys": {"type": "array", "items": {"$ref": "#/components/schemas/cc_UserSSHPublicKey"}}, "vcpu": {"description": "Specify the vcpu to be used for the deployment. Vcpu must be at least 1. The input value will be rounded to\nthe nearest 0.0001. The default will be the one configured for the account.\n", "type": "number", "format": "float"}, "wrangler_ssh": {"$ref": "#/components/schemas/cc_WranglerSSHConfig"}}}
```
