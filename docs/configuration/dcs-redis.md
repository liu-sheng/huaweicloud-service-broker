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

## Plan Metadata Parameters Configuration

| Parameter Name         | Type       | Description
|:-----------------------|:-----------|:-----------
| engine                 | string     | Cache engine, which is Redis.
| engine_version         | string     | Cache engine version, which is 3.0.7.
| speccode               | string     | DCS specifications. Values: dcs.single_node, dcs.master_standby and dcs.cluster.
| charging_type          | string     | Billing mode. Values: Yearly, Monthly and Hourly.
| capacity               | int        | Cache capacity. Unit: GB. For a DCS Redis instance in single node or master standby mode, the cache capacity can be 2 GB, 4 GB, 8 GB, 16 GB, 32 GB, or 64 GB. For a DCS Redis instance in cluster mode, the cache capacity can be 64, 128, 256, 512, or 1024 GB.
| vpc_id                 | string     | Tenant's VPC ID.
| subnet_id              | string     | Subnet ID.
| security_group_id      | string     | Tenant's security group ID.
| availability_zones     | []string   | IDs of the AZs where cache nodes reside. For details about how to obtain this parameter value, see [Regions and Endpoints](https://developer.huaweicloud.com/endpoint). ```Note: In the current version of Orange Cloud, only one AZ ID can be set in the request.```


These plans are only differently configured by the Parameter Name [```speccode```].
