package instance

import (
	"fmt"

	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack/dms/v1/instances"
	"github.com/huaweicloud/huaweicloud-service-broker/pkg/database"
	"github.com/huaweicloud/huaweicloud-service-broker/pkg/models"
	"github.com/pivotal-cf/brokerapi"
)

// LastOperation implematation
func (b *DMSBroker) LastOperation(instanceID string, operationData database.OperationDetails) (brokerapi.LastOperation, error) {

	// Log opts
	b.Logger.Debug(fmt.Sprintf("lastoperation dms instance opts: instanceID: %s operationData: %v", instanceID, models.ToJson(operationData)))

	// Handle different cases
	if (operationData.OperationType == models.OperationProvisioning) ||
		(operationData.OperationType == models.OperationUpdating) {
		// OperationProvisioning || OperationUpdating
		instance, err, serviceErr := SyncStatusWithService(b, instanceID, operationData.ServiceID,
			operationData.PlanID, operationData.TargetID)
		if err != nil {
			return brokerapi.LastOperation{}, err
		}
		if serviceErr != nil {
			return brokerapi.LastOperation{
				State:       brokerapi.Failed,
				Description: fmt.Sprintf("get dms instance failed. Error: %s", serviceErr),
			}, nil
		}
		// ErrorCode
		if instance.ErrorCode != "" {
			return brokerapi.LastOperation{
				State:       brokerapi.Failed,
				Description: fmt.Sprintf("ErrorCode: %s", instance.ErrorCode),
			}, nil
		}
		// Status
		if instance.Status == "RUNNING" {
			return brokerapi.LastOperation{
				State:       brokerapi.Succeeded,
				Description: fmt.Sprintf("Status: %s", instance.Status),
			}, nil
		} else if (instance.Status == "CREATEFAILED") ||
			(instance.Status == "ERROR") {
			return brokerapi.LastOperation{
				State:       brokerapi.Failed,
				Description: fmt.Sprintf("ErrorCode: %s", instance.ErrorCode),
			}, nil
		} else {
			return brokerapi.LastOperation{
				State:       brokerapi.InProgress,
				Description: fmt.Sprintf("Status: %s", instance.Status),
			}, nil
		}
	} else if operationData.OperationType == models.OperationDeprovisioning {
		// OperationDeprovisioning
		// Init dms client
		dmsClient, err := b.CloudCredentials.DMSV1Client()
		if err != nil {
			return brokerapi.LastOperation{}, fmt.Errorf("create dms client failed. Error: %s", err)
		}
		// Invoke sdk get
		instance, serviceErr := instances.Get(dmsClient, operationData.TargetID).Extract()
		if serviceErr != nil {
			e, ok := serviceErr.(golangsdk.ErrDefault404)
			if ok {
				return brokerapi.LastOperation{
					State:       brokerapi.Succeeded,
					Description: fmt.Sprintf("Status: %s", e.Error()),
				}, nil
			} else {
				return brokerapi.LastOperation{
					State:       brokerapi.Failed,
					Description: fmt.Sprintf("get dms instance failed. Error: %s", serviceErr),
				}, nil
			}
		} else {
			return brokerapi.LastOperation{
				State:       brokerapi.InProgress,
				Description: fmt.Sprintf("Name: %s", instance.Name),
			}, nil
		}
	} else {
		b.Logger.Debug(fmt.Sprintf("unknown operation data type: %s", operationData.OperationType))
	}

	// Log result
	b.Logger.Debug(fmt.Sprintf("lastoperation dms instance success: instanceID: %s", instanceID))

	return brokerapi.LastOperation{}, nil
}
