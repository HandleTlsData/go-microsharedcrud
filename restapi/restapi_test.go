package restapi

import (
	sharedserver "sharedcrud/api"
	gdbmanager "sharedcrud/gormdb"
	"sharedcrud/models"
	"sharedcrud/restapi/alphaAPI"
	"sharedcrud/restapi/betaAPI"
	"sharedcrud/restapi/operations/entity"
	"testing"
)

func TestHandler_AlphaAPI(t *testing.T) {
	gdbmanager.CurrentAppConfig = "alpha"
	gdbmanager.InitDB()
	go sharedserver.StartAlphaGRPC()
	const predefinedTestUserName = "testingRecord"
	const predefinedTestUserStatus = "testing status"
	const predefinedTestUserDesc = "testing desc"

	const predefinedTestUserUpdateName = "testingRecordEdited"

	//get test on deleted user
	{
		var expected string
		var result string

		//delete our test case user
		err := alphaAPI.EntityDelete(entity.EntityDeleteParams{EntityName: predefinedTestUserName})

		if err != nil {
			if err.Error() != gdbmanager.StrNoRecords {
				t.Errorf("alphaAPI.EntityDelete failed with %s", err.Error())
			}
		}

		//delete second test case user
		err = alphaAPI.EntityDelete(entity.EntityDeleteParams{EntityName: predefinedTestUserUpdateName})

		if err != nil {
			if err.Error() != gdbmanager.StrNoRecords {
				t.Errorf("alphaAPI.EntityDelete failed with %s", err.Error())
			}
		}

		entity, err := alphaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})
		result = err.Error()
		expected = gdbmanager.StrNoRecords
		if result != expected {
			t.Errorf("alphaAPI.EntityGet expected %s got %s", expected, result)
		}
		if entity != nil {

		}
	}

	//create + get test
	{
		var expected = models.Entity{Name: predefinedTestUserName, Status: predefinedTestUserStatus, Description: predefinedTestUserDesc}
		var result models.Entity

		//storing entity to get in future
		err := alphaAPI.EntityStore(entity.EntityStoreParams{Body: &models.Entity{Name: predefinedTestUserName, Status: predefinedTestUserStatus, Description: predefinedTestUserDesc}})

		if err != nil {
			t.Errorf("alphaAPI.EntityStore failed with %s", err.Error())
		}

		entityS, err := alphaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})

		if err != nil {
			t.Errorf("alphaAPI.EntityGet failed with %s", err.Error())
		}

		result = *entityS[0]

		if result.Name != expected.Name || result.Status != expected.Status || result.Description != expected.Description {
			t.Errorf("alphaAPI.EntityGet expected:\n%+v\n got %+v", expected, result)
		}
	}

	//update existing entity test
	{
		var expected = models.Entity{Name: predefinedTestUserUpdateName, Status: predefinedTestUserStatus, Description: predefinedTestUserDesc}
		var result models.Entity

		arrOldEntity, err := alphaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})

		if err != nil {
			t.Errorf("alphaAPI.EntityGet failed with %s", err.Error())
		}

		oldEntity := *arrOldEntity[0]
		targetUserID := oldEntity.ID

		//update entity by passing ID
		err = alphaAPI.EntityStore(entity.EntityStoreParams{Body: &models.Entity{ID: targetUserID, Name: predefinedTestUserUpdateName, Status: predefinedTestUserStatus, Description: predefinedTestUserDesc}})

		if err != nil {
			t.Errorf("alphaAPI.EntityStore failed with %s", err.Error())
		}

		entityS, err := alphaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserUpdateName})

		if err != nil {
			t.Errorf("alphaAPI.EntityGet failed with %s", err.Error())
		}

		result = *entityS[0]

		if result.Name != expected.Name {
			t.Errorf("alphaAPI.EntityGet expected:\n%+v\n got %+v", expected, result)
		}

	}
}

func TestHandler_BetaAPI(t *testing.T) {
	gdbmanager.CurrentAppConfig = "beta"
	gdbmanager.InitDB()
	go sharedserver.StartBetaGRPC()
	const predefinedTestUserName = "testingRecord"
	const predefinedTestUserStatus = "testing status"
	const predefinedTestUserDesc = "testing desc"

	const predefinedTestUserUpdateDesc = "testing desc Edited"

	//get test on deleted user
	{
		var expected string
		var result string

		//delete our test case user
		err := betaAPI.EntityDelete(entity.EntityDeleteParams{EntityName: predefinedTestUserName})

		if err != nil {
			if err.Error() != gdbmanager.StrNoRecords {
				t.Errorf("betaAPI.EntityDelete failed with %s", err.Error())
			}
		}

		entity, err := betaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})
		result = err.Error()
		expected = gdbmanager.StrNoRecords
		if result != expected {
			t.Errorf("betaAPI.EntityGet expected %s got %s", expected, result)
		}
		if entity != nil {

		}
	}

	//create + get test
	{
		var expected = models.Entity{Name: predefinedTestUserName, Status: predefinedTestUserStatus, Description: predefinedTestUserDesc}
		var result models.Entity

		//storing entity to get in future
		err := betaAPI.EntityStore(entity.EntityStoreParams{Body: &models.Entity{Name: predefinedTestUserName, Status: predefinedTestUserStatus, Description: predefinedTestUserDesc}})

		if err != nil {
			t.Errorf("betaAPI.EntityStore failed with %s", err.Error())
		}

		entityS, err := betaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})

		if err != nil {
			t.Errorf("betaAPI.EntityGet failed with %s", err.Error())
		}

		result = *entityS[0]

		if result.Name != expected.Name || result.Status != expected.Status || result.Description != expected.Description {
			t.Errorf("betaAPI.EntityGet expected:\n%+v\n got %+v", expected, result)
		}
	}

	//update existing entity test
	{
		var expected = models.Entity{Name: predefinedTestUserName, Status: predefinedTestUserStatus, Description: predefinedTestUserUpdateDesc}
		var result models.Entity

		arrOldEntity, err := betaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})

		if err != nil {
			t.Errorf("betaAPI.EntityGet failed with %s", err.Error())
		}

		oldEntity := *arrOldEntity[0]
		targetUserID := oldEntity.ID

		//update entity by passing ID
		err = betaAPI.EntityStore(entity.EntityStoreParams{Body: &models.Entity{ID: targetUserID, Name: predefinedTestUserName,
			Status: predefinedTestUserStatus, Description: predefinedTestUserUpdateDesc}})

		if err != nil {
			t.Errorf("betaAPI.EntityStore failed with %s", err.Error())
		}

		entityS, err := betaAPI.EntityGet(entity.EntityGetParams{EntityName: predefinedTestUserName})

		if err != nil {
			t.Errorf("betaAPI.EntityGet failed with %s", err.Error())
		}

		result = *entityS[0]

		if result.Name != expected.Name {
			t.Errorf("betaAPI.EntityGet expected:\n%+v\n got %+v", expected, result)
		}

	}
}
