# Distributed Cache Service for Redis

## Services

| Service Name                   | Description
|:-------------------------------|:-----------
| dcs-redis                      | Distributed Cache Service for Redis

## Plans

| Plan Name                      | Description
|:-------------------------------|:-----------
| SingleNode                     | Redis Single Node
| MasterStandby                  | Redis Master Standby
| Cluster                        | Redis Cluster

## Provision

Provision a new instance for Distributed Cache Service for Redis.

### Provision Parameters

| Parameter Name               | Type       | Required  | Description
|:-----------------------------|:-----------|:----------|:-----------
| capacity                     | int        | N         | Cache capacity. Unit: GB. For a Redis instance in single node or master standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. For a Redis instance in cluster mode, the cache capacity can be 64, 128, 256, 512, or 1024 GB. The default value is in the config file.
| vpc_id                       | string     | N         | Tenant's VPC ID. The default value is in the config file.
| subnet_id                    | string     | N         | Subnet ID. The default value is in the config file.
| security_group_id            | string     | N         | Tenant's security group ID. The default value is in the config file.
| availability_zones           | []string   | N         | IDs of the AZs where cache nodes reside. For details about how to obtain this parameter value, see [Regions and Endpoints](https://developer.huaweicloud.com/endpoint).  ```Note: In the current version of Orange Cloud, only one AZ ID can be set in the request.``` The default value is in the config file.
| password                     | string     | Y         | Password of a Redis instance. The password of a Redis instance must meet the following complexity requirements: A string of 6–32 characters. Contains at least two of the following character types: Uppercase letters; Lowercase letters; Digits; Special characters, such as `~!@#$%^&*()-_=+\|[{}]:'",<.>/?.
| name                         | string     | Y         | Redis instance name. An instance name is a string of 4–64 characters that contain letters, digits, underscores (_), and hyphens (-). An instance name must start with letters.
| description                  | string     | N         | Brief description of the Redis instance. A brief description supports up to 1024 characters.
| backup_strategy_savedays     | int        | N         | Retention time. Unit: day. Range: 1–7.
| backup_strategy_backup_type  | string     | N         | Backup type. Options: auto: automatic backup. manual: manual backup.
| backup_strategy_backup_at    | []int      | N         | Days in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday.
| backup_strategy_begin_at     | string     | N         | Time at which backup starts. "00:00-01:00" indicates that backup starts at 00:00:00.
| backup_strategy_period_type  | string     | N         | Interval at which backup is performed. Currently, only weekly backup is supported.
| maintain_begin               | string     | N         | Time at which the maintenance time window starts.
| maintain_end                 | string     | N         | Time at which the maintenance time window ends.

## Bind

Create a new credentials on the provisioned instance.
Bind returns the following connection details and credentials.

### Bind Credentials

| Parameter Name         | Type       | Description
|:-----------------------|:-----------|:-----------
| host                   | string     | The fully-qualified address of Redis instance.
| port                   | int        | The port number to connect to Redis instance.
| password               | string     | Password of a Redis instance.
| name                   | string     | Redis instance name.
| type                   | string     | The service type. The value is dcs-redis.

## Unbind

Remove the bind information from the provisioned instance.

## Update

Update a previously provisioned instance.

### Update Parameters

| Parameter Name               | Type       | Required  | Description
|:-----------------------------|:-----------|:----------|:-----------
| name                         | string     | N         | Redis instance name. An instance name is a string of 4–64 characters that contain letters, digits, underscores (_), and hyphens (-). An instance name must start with letters.
| description                  | string     | N         | Brief description of the Redis instance. A brief description supports up to 1024 characters.
| backup_strategy_savedays     | int        | N         | Retention time. Unit: day. Range: 1–7.
| backup_strategy_backup_type  | string     | N         | Backup type. Options: auto: automatic backup. manual: manual backup.
| backup_strategy_backup_at    | []int      | N         | Days in a week on which backup starts. Range: 1–7. Where: 1 indicates Monday; 7 indicates Sunday.
| backup_strategy_begin_at     | string     | N         | Time at which backup starts. "00:00-01:00" indicates that backup starts at 00:00:00.
| backup_strategy_period_type  | string     | N         | Interval at which backup is performed. Currently, only weekly backup is supported.
| maintain_begin               | string     | N         | Time at which the maintenance time window starts.
| maintain_end                 | string     | N         | Time at which the maintenance time window ends.
| security_group_id            | string     | N         | Subnet ID.
| new_capacity                 | int        | N         | New cache capacity. Unit: GB. For a Redis instance in single node or master standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. **For a Redis instance in cluster mode, it does not support extend the cache capacity.**
| old_password                 | string     | N         | The previous password of Redis instance.
| new_password                 | string     | N         | The new password of Redis instance.


## Deprovision

Delete the provisioned instance.

## Example on Cloud Foundry

### Provision

The following command will create a service.

```
cf create-service dcs-redis SingleNode myredis -c '{
    "password": "Password1234!",
    "name": "RedisSingleNode",
    "description": "Redis Single Node Test"
}'
```

You can check the status of the service instance using the `cf service` command.

```
cf service myredis
```

### Bind

Once the service has been successfully provisioned, you can bind to it by using
`cf bind-service` or by including it in a Cloud Foundry manifest.

```
cf bind-service myapp myredis
```

Use `cf restage` command to ensure your env variable changes take effect.

```
cf restage myapp
```

Once bound, you can view the environment variables for a given application using the `cf env` command.

```
cf env myapp
```

### Unbind

To unbind a service from an application, use the `cf unbind-service` command.

```
cf unbind-service myapp myredis
```

### Update

To update a service, use the `cf update-service` command.

```
cf update-service myredis -c '{
    "name": "RedisSingleNode1",
    "description": "Redis Single Node Test1",
    "new_capacity": 8
}'
```

You can also check the status of the service instance using the `cf service` command.

```
cf service myredis
```

### Deprovision

To deprovision the service, use the `cf delete-service` command.

```
cf delete-service myredis
```
