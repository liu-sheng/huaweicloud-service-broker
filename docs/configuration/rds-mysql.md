# RDS MySQL Service

## Services

| Service Name                   | Description
|:-------------------------------|:-----------
| rds-mysql                      | RDS MySQL Service

## Plans

| Plan Name                      | Description
|:-------------------------------|:-----------
| 5.6.39                         | RDS MySQL 5.6.39
| 5.7.21                         | RDS MySQL 5.7.21

## Plan Metadata Parameters Configuration

| Parameter Name         | Type       | Description
|:-----------------------|:-----------|:-----------
| datastore_type         | string     | Specifies the DB engine. Currently, MySQL, PostgreSQL, SQLServer and HWSQL are supported. The value is MySQL.
| datastore_version      | string     | Specifies the DB instance version. 5.6.39 and 5.7.21 are supported.
| speccode               | string     | Indicates the resource specifications code. Use rds.mysql.m1.xlarge as an example. rds indicates RDS, mysql indicates the DB engine, and m1.xlarge indicates the performance specification (large-memory). The parameter containing rr indicates the read replica specifications. The parameter not containing rr indicates the single or primary/standby DB instance specifications. If you enable HA, the suffix ```.ha``` need be added to the DB instance name. For example, the DB instance name is ```rds.db.s1.xlarge.ha```.
| volume_type            | string     | Specifies the volume type. Valid value: It must be COMMON (SATA) or ULTRAHIGH (SSD) and is case-sensitive.
| volume_size            | int        | Specifies the volume size. Its value must be a multiple of 10 and the value range is 100 GB to 2000 GB.
| availability_zone      | string     | Specifies the ID of the AZ. Valid value: The value cannot be empty. For details about how to obtain this parameter value, see [Regions and Endpoints](https://developer.huaweicloud.com/endpoint).
| vpc_id                 | string     | Specifies the VPC ID.
| subnet_id              | string     | Specifies the UUID for nics information.
| security_group_id      | string     | Specifies the security group ID which the RDS DB instance belongs to.
| database_username      | string     | Specifies the username of the database. Currently the only supported value of username is root.

These plans are only differently configured by the Parameter Name [```datastore_version```].
