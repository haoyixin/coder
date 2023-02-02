package authzquery_test

import (
	"testing"
	"time"

	"github.com/coder/coder/coderd/database"
	"github.com/coder/coder/coderd/database/dbgen"
	"github.com/coder/coder/coderd/rbac"
)

func (suite *MethodTestSuite) TestAPIKey() {
	suite.Run("DeleteAPIKeyByID", func() {
		suite.RunMethodTest(func(t *testing.T, db database.Store) MethodCase {
			key, _ := dbgen.APIKey(t, db, database.APIKey{})
			return methodCase(inputs(key.ID), asserts(key, rbac.ActionDelete))
		})
	})
	suite.Run("GetAPIKeyByID", func() {
		suite.RunMethodTest(func(t *testing.T, db database.Store) MethodCase {
			key, _ := dbgen.APIKey(t, db, database.APIKey{})
			return methodCase(inputs(key.ID), asserts(key, rbac.ActionRead))
		})
	})
	suite.Run("GetAPIKeysByLoginType", func() {
		suite.RunMethodTest(func(t *testing.T, db database.Store) MethodCase {
			a, _ := dbgen.APIKey(t, db, database.APIKey{LoginType: database.LoginTypePassword})
			b, _ := dbgen.APIKey(t, db, database.APIKey{LoginType: database.LoginTypePassword})
			_, _ = dbgen.APIKey(t, db, database.APIKey{LoginType: database.LoginTypeGithub})
			return methodCase(inputs(database.LoginTypePassword), asserts(a, rbac.ActionRead, b, rbac.ActionRead))
		})
	})
	suite.Run("GetAPIKeysLastUsedAfter", func() {
		suite.RunMethodTest(func(t *testing.T, db database.Store) MethodCase {
			a, _ := dbgen.APIKey(t, db, database.APIKey{LastUsed: time.Now().Add(time.Hour)})
			b, _ := dbgen.APIKey(t, db, database.APIKey{LastUsed: time.Now().Add(time.Hour)})
			_, _ = dbgen.APIKey(t, db, database.APIKey{LastUsed: time.Now().Add(-time.Hour)})
			return methodCase(inputs(time.Now()), asserts(a, rbac.ActionRead, b, rbac.ActionRead))
		})
	})
	suite.Run("InsertAPIKey", func() {
		suite.RunMethodTest(func(t *testing.T, db database.Store) MethodCase {
			u := dbgen.User(t, db, database.User{})
			return methodCase(inputs(database.InsertAPIKeyParams{
				UserID:    u.ID,
				LoginType: database.LoginTypePassword,
				Scope:     database.APIKeyScopeAll,
			}), asserts(rbac.ResourceAPIKey.WithOwner(u.ID.String()), rbac.ActionCreate))
		})
	})
	suite.Run("UpdateAPIKeyByID", func() {
		suite.RunMethodTest(func(t *testing.T, db database.Store) MethodCase {
			a, _ := dbgen.APIKey(t, db, database.APIKey{})
			return methodCase(inputs(database.UpdateAPIKeyByIDParams{
				ID:       a.ID,
				LastUsed: time.Now(),
			}), asserts(a, rbac.ActionUpdate))
		})
	})
}
