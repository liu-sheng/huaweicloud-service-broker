package rds

import (
	"fmt"

	"github.com/huaweicloud/golangsdk/openstack/rds/v1/instances"
	"github.com/huaweicloud/huaweicloud-service-broker/pkg/database"
	"github.com/pivotal-cf/brokerapi"
)

// Deprovision implematation
func (b *RDSBroker) Deprovision(instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {

	// Check rds instance length in back database
	var length int
	err := database.BackDBConnection.
		Model(&database.InstanceDetails{}).
		Where("instance_id = ? and service_id = ? and plan_id = ?", instanceID, details.ServiceID, details.PlanID).
		Count(&length).Error
	if err != nil {
		return brokerapi.DeprovisionServiceSpec{}, fmt.Errorf("check rds instance length in back database failed. Error: %s", err)
	}
	// ErrInstanceDoesNotExist
	if length > 0 {
		return brokerapi.DeprovisionServiceSpec{}, brokerapi.ErrInstanceDoesNotExist
	}

	// get InstanceDetails in back database
	ids := database.InstanceDetails{}
	err = database.BackDBConnection.
		Where("instance_id = ? and service_id = ? and plan_id = ?", instanceID, details.ServiceID, details.PlanID).
		First(&ids).Error
	if err != nil {
		return brokerapi.DeprovisionServiceSpec{}, brokerapi.ErrInstanceDoesNotExist
	}

	// Log InstanceDetails
	b.Logger.Debug(fmt.Sprintf("rds instance in back database: %v", ids))

	// Init rds client
	rdsClient, err := b.CloudCredentials.RDSV1Client()
	if err != nil {
		return brokerapi.DeprovisionServiceSpec{}, fmt.Errorf("create rds client failed. Error: %s", err)
	}

	// Log opts
	b.Logger.Debug(fmt.Sprintf("deprovision rds instance opts: %s", instanceID))

	// Invoke sdk
	result := instances.Delete(rdsClient, ids.TargetID)
	if result.Err != nil {
		return brokerapi.DeprovisionServiceSpec{}, fmt.Errorf("deprovision rds instance failed. Error: %s", err)
	}

	// Log result
	b.Logger.Debug(fmt.Sprintf("deprovision rds instance success: %s", instanceID))

	// Return result
	return brokerapi.DeprovisionServiceSpec{IsAsync: false, OperationData: ""}, nil
}
