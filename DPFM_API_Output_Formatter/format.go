package dpfm_api_output_formatter

import (
	"data-platform-api-inspection-lot-confirmation-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.InspectionLot,
			&pm.Operations,
			&pm.OperationsItem,
			&pm.OperationID,
			&pm.ConfirmationCountingID,
			&pm.Inspection,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.InspectionType,
			&pm.InspectionTypeValueUnit,
			&pm.InspectionTypePlannedValue,
			&pm.InspectionTypeCertificateType,
			&pm.InspectionTypeCertificateValueInText,
			&pm.InspectionTypeCertificateValueInQuantity,
			&pm.InspectionLotInspectionText,
			&pm.InspectionTypeConfirmedValueInText,
			&pm.InspectionTypeConfirmedValueInQuantity,
			&pm.InspectionResult,
			&pm.OperationPlannedQuantityInBaseUnit,
			&pm.OperationPlannedQuantityInProductionUnit,
			&pm.OperationPlannedQuantityInOperationUnit,
			&pm.ProductBaseUnit,
			&pm.ProductProductionUnit,
			&pm.ProductOperationUnit,
			&pm.OperationPlannedScrapInPercent,
			&pm.ConfirmationEntryDate,
			&pm.ConfirmationEntryTime,
			&pm.ConfirmationText,
			&pm.IsFinalConfirmation,
			&pm.WorkCenter,
			&pm.EmployeeIDWhoConfirmed,
			&pm.ConfirmedExecutionStartDate,
			&pm.ConfirmedExecutionStartTime,
			&pm.ConfirmedSetupStartDate,
			&pm.ConfirmedSetupStartTime,
			&pm.ConfirmedProcessingStartDate,
			&pm.ConfirmedProcessingStartTime,
			&pm.ConfirmedExecutionEndDate,
			&pm.ConfirmedExecutionEndTime,
			&pm.ConfirmedSetupEndDate,
			&pm.ConfirmedSetupEndTime,
			&pm.ConfirmedProcessingEndDate,
			&pm.ConfirmedProcessingEndTime,
			&pm.ConfirmedWaitDuration,
			&pm.WaitDurationUnit,
			&pm.ConfirmedQueueDuration,
			&pm.QueueDurationUnit,
			&pm.ConfirmedMoveDuration,
			&pm.MoveDurationUnit,
			&pm.ConfirmedYieldQuantity,
			&pm.ConfirmedScrapQuantity,
			&pm.OperationVarianceReason,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsCancelled,
			&pm.IsCancelled,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			InspectionLot:                            data.InspectionLot,
			Operations:                               data.Operations,
			OperationsItem:                           data.OperationsItem,
			OperationID:                              data.OperationID,
			ConfirmationCountingID:                   data.ConfirmationCountingID,
			Inspection:                               data.Inspection,
			ProductionOrder:                          data.ProductionOrder,
			ProductionOrderItem:                      data.ProductionOrderItem,
			InspectionType:                           data.InspectionType,
			InspectionTypeValueUnit:                  data.InspectionTypeValueUnit,
			InspectionTypePlannedValue:               data.InspectionTypePlannedValue,
			InspectionTypeCertificateType:            data.InspectionTypeCertificateType,
			InspectionTypeCertificateValueInText:     data.InspectionTypeCertificateValueInText,
			InspectionTypeCertificateValueInQuantity: data.InspectionTypeCertificateValueInQuantity,
			InspectionLotInspectionText:              data.InspectionLotInspectionText,
			InspectionTypeConfirmedValueInText:       data.InspectionTypeConfirmedValueInText,
			InspectionTypeConfirmedValueInQuantity:   data.InspectionTypeConfirmedValueInQuantity,
			InspectionResult:                         data.InspectionResult,
			OperationPlannedQuantityInBaseUnit:       data.OperationPlannedQuantityInBaseUnit,
			OperationPlannedQuantityInProductionUnit: data.OperationPlannedQuantityInProductionUnit,
			OperationPlannedQuantityInOperationUnit:  data.OperationPlannedQuantityInOperationUnit,
			ProductBaseUnit:                          data.ProductBaseUnit,
			ProductProductionUnit:                    data.ProductProductionUnit,
			ProductOperationUnit:                     data.ProductOperationUnit,
			OperationPlannedScrapInPercent:           data.OperationPlannedScrapInPercent,
			ConfirmationEntryDate:                    data.ConfirmationEntryDate,
			ConfirmationEntryTime:                    data.ConfirmationEntryTime,
			ConfirmationText:                         data.ConfirmationText,
			IsFinalConfirmation:                      data.IsFinalConfirmation,
			WorkCenter:                               data.WorkCenter,
			EmployeeIDWhoConfirmed:                   data.EmployeeIDWhoConfirmed,
			ConfirmedExecutionStartDate:              data.ConfirmedExecutionStartDate,
			ConfirmedExecutionStartTime:              data.ConfirmedExecutionStartTime,
			ConfirmedSetupStartDate:                  data.ConfirmedSetupStartDate,
			ConfirmedSetupStartTime:                  data.ConfirmedSetupStartTime,
			ConfirmedProcessingStartDate:             data.ConfirmedProcessingStartDate,
			ConfirmedProcessingStartTime:             data.ConfirmedProcessingStartTime,
			ConfirmedExecutionEndDate:                data.ConfirmedExecutionEndDate,
			ConfirmedExecutionEndTime:                data.ConfirmedExecutionEndTime,
			ConfirmedSetupEndDate:                    data.ConfirmedSetupEndDate,
			ConfirmedSetupEndTime:                    data.ConfirmedSetupEndTime,
			ConfirmedProcessingEndDate:               data.ConfirmedProcessingEndDate,
			ConfirmedProcessingEndTime:               data.ConfirmedProcessingEndTime,
			ConfirmedWaitDuration:                    data.ConfirmedWaitDuration,
			WaitDurationUnit:                         data.WaitDurationUnit,
			ConfirmedQueueDuration:                   data.ConfirmedQueueDuration,
			QueueDurationUnit:                        data.QueueDurationUnit,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
