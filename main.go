package main

import "crud_fiber.com/m/apps"
import _ "crud_fiber.com/m/docs"

// @title CRUD Fiber API
// @version 1.0
// @description Dokumentasi API CRUD menggunakan go fiber
// @securityDefinitions.oauth2.password OAuth2Password
// @tokenUrl http://localhost:8080/api/v1/auth/token
// @scope read Grants read access
// @scope write Grants write access
func main() {
	apps.StartApps()
}
