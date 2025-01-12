/*
Copyright 2022 QuanxiangCloud Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package node

import (
	"context"
	"fmt"
	"github.com/quanxiang-cloud/flow/internal/convert"
	"github.com/quanxiang-cloud/flow/pkg"
	"github.com/quanxiang-cloud/flow/pkg/client"
	"github.com/quanxiang-cloud/flow/pkg/config"
	"github.com/quanxiang-cloud/flow/pkg/misc/logger"
	"github.com/quanxiang-cloud/flow/pkg/utils"
	"github.com/quanxiang-cloud/flow/rpc/pb"
)

// DataUpdate struct
type DataUpdate struct {
	*Node
}

// NewDataUpdate new
func NewDataUpdate(conf *config.Configs, node *Node) *DataUpdate {
	return &DataUpdate{
		Node: node,
	}
}

// InitBegin event
func (n *DataUpdate) InitBegin(ctx context.Context, eventData *EventData) (*pb.NodeEventRespData, error) {
	//logger.Logger.Debug("init update form begin")
	//flow, err := n.FlowRepo.FindByProcessID(n.Db, eventData.ProcessID)
	//if err != nil {
	//	return nil, err
	//}
	//formShape, err := convert.GetShapeByChartType(flow.BpmnText, convert.FormData)
	//if err != nil {
	//	return nil, err
	//}
	//formDefKey := formShape.ID
	//
	//instance, err := n.InstanceRepo.GetEntityByProcessInstanceID(n.Db, eventData.ProcessInstanceID)
	//if err != nil {
	//	return nil, err
	//}
	//if instance.Status == "REFUSE" {
	//	return nil, nil
	//}
	//instanceSteps, err := n.InstanceStepRepo.FindInstanceStepsByStatus(n.Db, instance.ProcessInstanceID, []string{"REFUSE"})
	//if err != nil {
	//	return nil, err
	//}
	//if len(instanceSteps) > 0 {
	//	return nil, nil
	//}
	//tasksReq := client.GetTasksReq{
	//	InstanceID: []string{instance.ProcessInstanceID},
	//}
	//taskResp, _ := n.ProcessAPI.GetHistoryTasks(ctx, tasksReq)
	//marshal, _ := json.Marshal(taskResp)
	//logger.Logger.Debug("taskResp==", string(marshal))
	//tasks := taskResp.Data
	//for k := range tasks {
	//	if strings.Contains(tasks[k].Comments, "REFUSE") {
	//		return nil, nil
	//	}
	//}
	//
	//bd := eventData.Shape.Data.BusinessData
	//if bd == nil {
	//	return nil, nil
	//}
	//
	//variables, err := n.Instance.GetInstanceVariableValues(ctx, instance)
	//if err != nil {
	//	return nil, err
	//}
	//
	//targetTableID := utils.Strval(bd["targetTableId"])
	//formQueryRef := bd["formQueryRef"]
	//triggerAgain := bd["silent"].(bool)
	//
	//var updateIDs []string
	//if instance.FormID != targetTableID { // 非本表支持过滤条件，本表只能更新当前记录
	//	var filterRule map[string]interface{}
	//	if v := bd["filterRule"]; v != nil {
	//		filterRule = v.(map[string]interface{})
	//	}
	//	if filterRule == nil {
	//		return nil, nil
	//	}
	//
	//	var conditions []map[string]interface{}
	//	if v := filterRule["conditions"]; v != nil {
	//		arr := v.([]interface{})
	//		for _, e := range arr {
	//			conditions = append(conditions, e.(map[string]interface{}))
	//		}
	//	}
	//	if conditions == nil {
	//		return nil, nil
	//	}
	//
	//	// reqConditions := make([]map[string]interface{}, 0)
	//	boolMap := make(map[string]interface{})
	//	queryMap := map[string]interface{}{
	//		"bool": boolMap,
	//	}
	//	terms := make([]map[string]interface{}, 0)
	//	if utils.Strval(filterRule["tag"]) == "or" {
	//		boolMap["should"] = &terms
	//	} else {
	//		boolMap["must"] = &terms
	//	}
	//
	//	for _, v := range conditions {
	//
	//		valueOf := v["value"]
	//		value, err := n.Instance.Cal(ctx, "currentFormValue", valueOf, nil, instance, variables, formQueryRef, formDefKey)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		// 等于eq, 不等于neq，包含in，不包含nin
	//		if utils.Strval(v["operator"]) == "eq" {
	//			term := map[string]interface{}{
	//				"term": map[string]interface{}{
	//					utils.Strval(v["fieldName"]): value,
	//				},
	//			}
	//			terms = append(terms, term)
	//		} else if utils.Strval(v["operator"]) == "neq" {
	//			mustNot := make([]map[string]interface{}, 0)
	//			mustNot = append(mustNot, map[string]interface{}{
	//				"term": map[string]interface{}{
	//					utils.Strval(v["fieldName"]): value,
	//				},
	//			})
	//			term := map[string]interface{}{
	//				"bool": map[string]interface{}{
	//					"mustNot": mustNot,
	//				},
	//			}
	//			terms = append(terms, term)
	//		} else if utils.Strval(v["operator"]) == "in" {
	//			// todo 如果是数组格式的需要修改in判断
	//			term := map[string]interface{}{
	//				"term": map[string]interface{}{
	//					utils.Strval(v["fieldName"]): value,
	//				},
	//			}
	//			terms = append(terms, term)
	//		} else if utils.Strval(v["operator"]) == "nin" {
	//			// todo 如果是数组格式的需要修改in判断
	//			mustNot := make([]map[string]interface{}, 0)
	//			mustNot = append(mustNot, map[string]interface{}{
	//				"term": map[string]interface{}{
	//					utils.Strval(v["fieldName"]): value,
	//				},
	//			})
	//			term := map[string]interface{}{
	//				"bool": map[string]interface{}{
	//					"mustNot": mustNot,
	//				},
	//			}
	//			terms = append(terms, term)
	//		}
	//	}
	//
	//	updateIDs, err = n.FormAPI.GetIDs(ctx, instance.AppID, targetTableID, queryMap)
	//	if err != nil {
	//		return nil, err
	//	}
	//} else { // if target form is current, can not trigger flow again
	//	triggerAgain = false
	//
	//	updateIDs = []string{instance.FormInstanceID} // if target form is current , only update current data
	//}
	//
	//var updateRules []map[string]interface{}
	//if v := bd["updateRule"]; v != nil {
	//	arr := v.([]interface{})
	//	for _, e := range arr {
	//		updateRules = append(updateRules, e.(map[string]interface{}))
	//	}
	//}
	//
	//updateReq := make(map[string]interface{}, 0)
	//for _, updateRule := range updateRules {
	//	fieldName := utils.Strval(updateRule["fieldName"])
	//	valueFrom := utils.Strval(updateRule["valueFrom"])
	//	valueOf := updateRule["valueOf"]
	//	formulaFields := updateRule["formulaFields"]
	//
	//	val, err := n.Instance.Cal(ctx, valueFrom, valueOf, formulaFields, instance, variables, formQueryRef, formDefKey)
	//	if err != nil {
	//		return nil, err
	//	}
	//	updateReq[fieldName] = val
	//}
	//
	//selectField := utils.Strval(bd["selectField"])               // 普通组件为空，高级组件为字段名
	//selectFieldType := utils.Strval(bd["selectFieldType"])       // 高级组件类型
	//selectFieldTableID := utils.Strval(bd["selectFieldTableId"]) // 高级组件涉及的tableId
	//ref := make(map[string]client.RefData, 0)
	//if len(selectField) > 0 && selectField != "normal" {
	//	if instance.FormID == targetTableID { // 本表
	//		dataReq := client.FormDataConditionModel{
	//			AppID:   instance.AppID,
	//			TableID: instance.FormID,
	//			DataID:  instance.FormInstanceID,
	//		}
	//		if selectFieldType == "associated_records" || selectFieldType == "foreign_table" || selectFieldType == "sub_table" {
	//			dataReq.Ref = map[string]interface{}{
	//				selectField: map[string]interface{}{
	//					"appID":   instance.AppID,
	//					"tableID": selectFieldTableID,
	//					"type":    selectFieldType,
	//				},
	//			}
	//		}
	//
	//		dataResp, err := n.FormAPI.GetFormData(ctx, dataReq)
	//		if err != nil {
	//			return nil, err
	//		}
	//		if dataResp == nil {
	//			return nil, nil
	//		}
	//
	//		fmt.Println("updateNode formData:" + utils.Strval(dataResp))
	//		fmt.Println("updateNode selectField:" + utils.Strval(dataResp[selectField]))
	//		fmt.Println("updateNode updateReq:" + utils.Strval(updateReq))
	//		if selectFieldType == "associated_records" || selectFieldType == "foreign_table" { // 外表
	//			targetTableID = selectFieldTableID
	//			updateIDs = utils.ChangeInterfaceToIDArray(dataResp[selectField])
	//		} else if selectFieldType == "associated_data" { // 外表
	//			targetTableID = selectFieldTableID
	//			// updateIDs = utils.ChangeInterfaceToValueArray(dataResp[selectField])
	//			associatedData := utils.ChangeObjectToMap(dataResp[selectField])
	//			if associatedData != nil {
	//				updateIDs = append(updateIDs, utils.Strval(associatedData["value"]))
	//			}
	//		} else if selectFieldType == "sub_table" { // 本表
	//			selectFieldDatas := utils.ChangeObjectToMapList(dataResp[selectField])
	//			newArr := make([]client.UpdateEntity, 0)
	//			for _, selectData := range selectFieldDatas {
	//				newArr = append(newArr, client.UpdateEntity{
	//					Entity: updateReq,
	//					Query: map[string]interface{}{
	//						"term": map[string]interface{}{
	//							"_id": selectData["_id"],
	//						},
	//					},
	//				})
	//			}
	//
	//			ref[selectField] = client.RefData{
	//				AppID:   instance.AppID,
	//				TableID: selectFieldTableID,
	//				Type:    selectFieldType,
	//				Updated: newArr,
	//			}
	//
	//			updateReq = make(map[string]interface{}, 0)
	//		}
	//
	//		fmt.Println("updateNode updateIDs:" + utils.Strval(updateIDs))
	//		fmt.Println("updateNode updateReq:" + utils.Strval(updateReq))
	//	} else { // 非本表
	//		fmt.Println("updateNode filter updateIDs:" + utils.Strval(updateIDs))
	//		if len(updateIDs) > 0 {
	//			for _, updateID := range updateIDs {
	//				updateReq2 := make(map[string]interface{}, 0)
	//				updateIDs2 := make([]string, 0)
	//				ref2 := make(map[string]client.RefData, 0)
	//				targetTableID2 := targetTableID
	//				dataReq := client.FormDataConditionModel{
	//					AppID:   instance.AppID,
	//					TableID: targetTableID,
	//					DataID:  updateID,
	//				}
	//				if selectFieldType == "associated_records" || selectFieldType == "foreign_table" || selectFieldType == "sub_table" {
	//					dataReq.Ref = map[string]interface{}{
	//						selectField: map[string]interface{}{
	//							"appID":   instance.AppID,
	//							"tableID": selectFieldTableID,
	//							"type":    selectFieldType,
	//						},
	//					}
	//				}
	//				dataResp, err := n.FormAPI.GetFormData(ctx, dataReq)
	//				if err != nil {
	//					return nil, err
	//				}
	//				if dataResp == nil {
	//					return nil, nil
	//				}
	//
	//				fmt.Println("updateNode formData:" + utils.Strval(dataResp))
	//				fmt.Println("updateNode selectField:" + utils.Strval(dataResp[selectField]))
	//				fmt.Println("updateNode updateReq:" + utils.Strval(updateReq))
	//
	//				if selectFieldType == "associated_records" || selectFieldType == "foreign_table" { // 外表
	//					targetTableID2 = selectFieldTableID
	//					updateIDs2 = utils.ChangeInterfaceToIDArray(dataResp[selectField])
	//					updateReq2 = updateReq
	//				} else if selectFieldType == "associated_data" { // 外表
	//					targetTableID2 = selectFieldTableID
	//					// updateIDs2 = utils.ChangeInterfaceToValueArray(dataResp[selectField])
	//					associatedData := utils.ChangeObjectToMap(dataResp[selectField])
	//					if associatedData != nil {
	//						updateIDs2 = append(updateIDs2, utils.Strval(associatedData["value"]))
	//					}
	//					updateReq2 = updateReq
	//				} else if selectFieldType == "sub_table" { // 本表
	//					selectFieldDatas := utils.ChangeObjectToMapList(dataResp[selectField])
	//					newArr := make([]client.UpdateEntity, 0)
	//					for _, selectData := range selectFieldDatas {
	//						newArr = append(newArr, client.UpdateEntity{
	//							Entity: updateReq,
	//							Query: map[string]interface{}{
	//								"term": map[string]interface{}{
	//									"_id": selectData["_id"],
	//								},
	//							},
	//						})
	//					}
	//					updateIDs2 = []string{updateID}
	//
	//					ref2[selectField] = client.RefData{
	//						AppID:   instance.AppID,
	//						TableID: selectFieldTableID,
	//						Type:    selectFieldType,
	//						Updated: newArr,
	//					}
	//				}
	//
	//				fmt.Println("updateNode updateIDs:" + utils.Strval(updateIDs2))
	//				fmt.Println("updateNode updateReq:" + utils.Strval(updateReq2))
	//
	//				ctx = pkg.SetRequestID2(ctx, instance.RequestID)
	//				err = n.FormAPI.UpdateData(ctx, instance.AppID, targetTableID2, "", client.UpdateEntity{
	//					Entity: updateReq2,
	//					Query: map[string]interface{}{
	//						"terms": map[string]interface{}{
	//							"_id": updateIDs2,
	//						},
	//					},
	//					Ref: ref2,
	//				}, triggerAgain)
	//				if err != nil {
	//					return nil, err
	//				}
	//			}
	//		}
	//		return nil, nil
	//	}
	//}
	//ctx = pkg.SetRequestID2(ctx, instance.RequestID)
	//err = n.FormAPI.UpdateData(ctx, instance.AppID, targetTableID, "", client.UpdateEntity{
	//	Entity: updateReq,
	//	Query: map[string]interface{}{
	//		"terms": map[string]interface{}{
	//			"_id": updateIDs,
	//		},
	//	},
	//	Ref: ref,
	//}, triggerAgain)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}

// InitEnd event
func (n *DataUpdate) InitEnd(ctx context.Context, eventData *EventData) (*pb.NodeEventRespData, error) {
	logger.Logger.Debug("init update form end")
	flow, err := n.FlowRepo.FindByProcessID(n.Db, eventData.ProcessID)
	if err != nil {
		return nil, err
	}
	formShape, err := convert.GetShapeByChartType(flow.BpmnText, convert.FormData)
	if err != nil {
		return nil, err
	}
	formDefKey := formShape.ID

	instance, err := n.InstanceRepo.GetEntityByProcessInstanceID(n.Db, eventData.ProcessInstanceID)
	if err != nil {
		return nil, err
	}
	if instance.Status == "REFUSE" {
		return nil, nil
	}
	//if !n.CheckRefuse(ctx, n.Db, instance.ProcessInstanceID) {
	//	return nil, nil
	//}

	bd := eventData.Shape.Data.BusinessData
	if bd == nil {
		return nil, nil
	}

	variables, err := n.Instance.GetInstanceVariableValues(ctx, instance)
	if err != nil {
		return nil, err
	}

	targetTableID := utils.Strval(bd["targetTableId"])
	formQueryRef := bd["formQueryRef"]
	triggerAgain := bd["silent"].(bool)

	var updateIDs []string
	if instance.FormID != targetTableID { // 非本表支持过滤条件，本表只能更新当前记录
		var filterRule map[string]interface{}
		if v := bd["filterRule"]; v != nil {
			filterRule = v.(map[string]interface{})
		}
		if filterRule == nil {
			return nil, nil
		}

		var conditions []map[string]interface{}
		if v := filterRule["conditions"]; v != nil {
			arr := v.([]interface{})
			for _, e := range arr {
				conditions = append(conditions, e.(map[string]interface{}))
			}
		}
		if conditions == nil {
			return nil, nil
		}

		// reqConditions := make([]map[string]interface{}, 0)
		boolMap := make(map[string]interface{})
		queryMap := map[string]interface{}{
			"bool": boolMap,
		}
		terms := make([]map[string]interface{}, 0)
		if utils.Strval(filterRule["tag"]) == "or" {
			boolMap["should"] = &terms
		} else {
			boolMap["must"] = &terms
		}

		for _, v := range conditions {

			valueOf := v["value"]
			value, err := n.Instance.Cal(ctx, "currentFormValue", valueOf, nil, instance, variables, formQueryRef, formDefKey)
			if err != nil {
				return nil, err
			}

			// 等于eq, 不等于neq，包含in，不包含nin
			if utils.Strval(v["operator"]) == "eq" {
				term := map[string]interface{}{
					"term": map[string]interface{}{
						utils.Strval(v["fieldName"]): value,
					},
				}
				terms = append(terms, term)
			} else if utils.Strval(v["operator"]) == "neq" {
				mustNot := make([]map[string]interface{}, 0)
				mustNot = append(mustNot, map[string]interface{}{
					"term": map[string]interface{}{
						utils.Strval(v["fieldName"]): value,
					},
				})
				term := map[string]interface{}{
					"bool": map[string]interface{}{
						"mustNot": mustNot,
					},
				}
				terms = append(terms, term)
			} else if utils.Strval(v["operator"]) == "in" {
				// todo 如果是数组格式的需要修改in判断
				term := map[string]interface{}{
					"term": map[string]interface{}{
						utils.Strval(v["fieldName"]): value,
					},
				}
				terms = append(terms, term)
			} else if utils.Strval(v["operator"]) == "nin" {
				// todo 如果是数组格式的需要修改in判断
				mustNot := make([]map[string]interface{}, 0)
				mustNot = append(mustNot, map[string]interface{}{
					"term": map[string]interface{}{
						utils.Strval(v["fieldName"]): value,
					},
				})
				term := map[string]interface{}{
					"bool": map[string]interface{}{
						"mustNot": mustNot,
					},
				}
				terms = append(terms, term)
			}
		}

		updateIDs, err = n.FormAPI.GetIDs(ctx, instance.AppID, targetTableID, queryMap)
		if err != nil {
			return nil, err
		}
	} else { // if target form is current, can not trigger flow again
		triggerAgain = false

		updateIDs = []string{instance.FormInstanceID} // if target form is current , only update current data
	}

	var updateRules []map[string]interface{}
	if v := bd["updateRule"]; v != nil {
		arr := v.([]interface{})
		for _, e := range arr {
			updateRules = append(updateRules, e.(map[string]interface{}))
		}
	}

	updateReq := make(map[string]interface{}, 0)
	for _, updateRule := range updateRules {
		fieldName := utils.Strval(updateRule["fieldName"])
		valueFrom := utils.Strval(updateRule["valueFrom"])
		valueOf := updateRule["valueOf"]
		formulaFields := updateRule["formulaFields"]

		val, err := n.Instance.Cal(ctx, valueFrom, valueOf, formulaFields, instance, variables, formQueryRef, formDefKey)
		if err != nil {
			return nil, err
		}
		updateReq[fieldName] = val
	}

	selectField := utils.Strval(bd["selectField"])               // 普通组件为空，高级组件为字段名
	selectFieldType := utils.Strval(bd["selectFieldType"])       // 高级组件类型
	selectFieldTableID := utils.Strval(bd["selectFieldTableId"]) // 高级组件涉及的tableId
	ref := make(map[string]client.RefData, 0)
	if len(selectField) > 0 && selectField != "normal" {
		if instance.FormID == targetTableID { // 本表
			dataReq := client.FormDataConditionModel{
				AppID:   instance.AppID,
				TableID: instance.FormID,
				DataID:  instance.FormInstanceID,
			}
			if selectFieldType == "associated_records" || selectFieldType == "foreign_table" || selectFieldType == "sub_table" {
				dataReq.Ref = map[string]interface{}{
					selectField: map[string]interface{}{
						"appID":   instance.AppID,
						"tableID": selectFieldTableID,
						"type":    selectFieldType,
					},
				}
			}

			dataResp, err := n.FormAPI.GetFormData(ctx, dataReq)
			if err != nil {
				return nil, err
			}
			if dataResp == nil {
				return nil, nil
			}

			fmt.Println("updateNode formData:" + utils.Strval(dataResp))
			fmt.Println("updateNode selectField:" + utils.Strval(dataResp[selectField]))
			fmt.Println("updateNode updateReq:" + utils.Strval(updateReq))
			if selectFieldType == "associated_records" || selectFieldType == "foreign_table" { // 外表
				targetTableID = selectFieldTableID
				updateIDs = utils.ChangeInterfaceToIDArray(dataResp[selectField])
			} else if selectFieldType == "associated_data" { // 外表
				targetTableID = selectFieldTableID
				// updateIDs = utils.ChangeInterfaceToValueArray(dataResp[selectField])
				associatedData := utils.ChangeObjectToMap(dataResp[selectField])
				if associatedData != nil {
					updateIDs = append(updateIDs, utils.Strval(associatedData["value"]))
				}
			} else if selectFieldType == "sub_table" { // 本表
				selectFieldDatas := utils.ChangeObjectToMapList(dataResp[selectField])
				newArr := make([]client.UpdateEntity, 0)
				for _, selectData := range selectFieldDatas {
					newArr = append(newArr, client.UpdateEntity{
						Entity: updateReq,
						Query: map[string]interface{}{
							"term": map[string]interface{}{
								"_id": selectData["_id"],
							},
						},
					})
				}

				ref[selectField] = client.RefData{
					AppID:   instance.AppID,
					TableID: selectFieldTableID,
					Type:    selectFieldType,
					Updated: newArr,
				}

				updateReq = make(map[string]interface{}, 0)
			}

			fmt.Println("updateNode updateIDs:" + utils.Strval(updateIDs))
			fmt.Println("updateNode updateReq:" + utils.Strval(updateReq))
		} else { // 非本表
			fmt.Println("updateNode filter updateIDs:" + utils.Strval(updateIDs))
			if len(updateIDs) > 0 {
				for _, updateID := range updateIDs {
					updateReq2 := make(map[string]interface{}, 0)
					updateIDs2 := make([]string, 0)
					ref2 := make(map[string]client.RefData, 0)
					targetTableID2 := targetTableID
					dataReq := client.FormDataConditionModel{
						AppID:   instance.AppID,
						TableID: targetTableID,
						DataID:  updateID,
					}
					if selectFieldType == "associated_records" || selectFieldType == "foreign_table" || selectFieldType == "sub_table" {
						dataReq.Ref = map[string]interface{}{
							selectField: map[string]interface{}{
								"appID":   instance.AppID,
								"tableID": selectFieldTableID,
								"type":    selectFieldType,
							},
						}
					}
					dataResp, err := n.FormAPI.GetFormData(ctx, dataReq)
					if err != nil {
						return nil, err
					}
					if dataResp == nil {
						return nil, nil
					}

					fmt.Println("updateNode formData:" + utils.Strval(dataResp))
					fmt.Println("updateNode selectField:" + utils.Strval(dataResp[selectField]))
					fmt.Println("updateNode updateReq:" + utils.Strval(updateReq))

					if selectFieldType == "associated_records" || selectFieldType == "foreign_table" { // 外表
						targetTableID2 = selectFieldTableID
						updateIDs2 = utils.ChangeInterfaceToIDArray(dataResp[selectField])
						updateReq2 = updateReq
					} else if selectFieldType == "associated_data" { // 外表
						targetTableID2 = selectFieldTableID
						// updateIDs2 = utils.ChangeInterfaceToValueArray(dataResp[selectField])
						associatedData := utils.ChangeObjectToMap(dataResp[selectField])
						if associatedData != nil {
							updateIDs2 = append(updateIDs2, utils.Strval(associatedData["value"]))
						}
						updateReq2 = updateReq
					} else if selectFieldType == "sub_table" { // 本表
						selectFieldDatas := utils.ChangeObjectToMapList(dataResp[selectField])
						newArr := make([]client.UpdateEntity, 0)
						for _, selectData := range selectFieldDatas {
							newArr = append(newArr, client.UpdateEntity{
								Entity: updateReq,
								Query: map[string]interface{}{
									"term": map[string]interface{}{
										"_id": selectData["_id"],
									},
								},
							})
						}
						updateIDs2 = []string{updateID}

						ref2[selectField] = client.RefData{
							AppID:   instance.AppID,
							TableID: selectFieldTableID,
							Type:    selectFieldType,
							Updated: newArr,
						}
					}

					fmt.Println("updateNode updateIDs:" + utils.Strval(updateIDs2))
					fmt.Println("updateNode updateReq:" + utils.Strval(updateReq2))
					ctx = pkg.SetRequestID2(ctx, instance.RequestID)
					err = n.FormAPI.UpdateData(ctx, instance.AppID, targetTableID2, "", client.UpdateEntity{
						Entity: updateReq2,
						Query: map[string]interface{}{
							"terms": map[string]interface{}{
								"_id": updateIDs2,
							},
						},
						Ref: ref2,
					}, triggerAgain)
					if err != nil {
						return nil, err
					}
				}
			}
			return nil, nil
		}
	}
	ctx = pkg.SetRequestID2(ctx, instance.RequestID)
	err = n.FormAPI.UpdateData(ctx, instance.AppID, targetTableID, "", client.UpdateEntity{
		Entity: updateReq,
		Query: map[string]interface{}{
			"terms": map[string]interface{}{
				"_id": updateIDs,
			},
		},
		Ref: ref,
	}, triggerAgain)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
