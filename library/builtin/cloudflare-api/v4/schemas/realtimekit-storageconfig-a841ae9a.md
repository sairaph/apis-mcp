---
title: realtimekit_StorageConfig
page_id: schema-realtimekit-storageconfig-a841ae9a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_StorageConfig

```yaml
{"type": "object", "properties": {"access_key": {"description": "Access key of the storage medium. Access key is not required for the `gcs` storage media type.\n\nNote that this field is not readable by clients, only writeable.", "type": "string", "writeOnly": true}, "auth_method": {"description": "Authentication method used for \"sftp\" type storage medium\n", "type": "string", "enum": ["KEY", "PASSWORD"]}, "bucket": {"description": "Name of the storage medium's bucket.", "type": "string"}, "host": {"description": "SSH destination server host for SFTP type storage medium", "type": "string"}, "password": {"description": "SSH destination server password for SFTP type storage medium when auth_method is \"PASSWORD\". If auth_method is \"KEY\", this specifies the password for the ssh private key.", "type": "string"}, "path": {"description": "Path relative to the bucket root at which the recording will be placed.", "type": "string"}, "port": {"description": "SSH destination server port for SFTP type storage medium", "type": "number"}, "private_key": {"description": "Private key used to login to destination SSH server for SFTP type storage medium, when auth_method used is \"KEY\"", "type": "string"}, "region": {"description": "Region of the storage medium.", "type": "string", "example": "us-east-1"}, "secret": {"description": "Secret key of the storage medium. Similar to `access_key`, it is only writeable by clients, not readable.", "type": "string"}, "type": {"description": "Type of storage media.", "type": "string", "enum": ["aws", "azure", "digitalocean", "gcs", "sftp"]}, "username": {"description": "SSH destination server username for SFTP type storage medium", "type": "string"}}, "nullable": true, "required": ["type"], "title": "StorageConfig"}
```
