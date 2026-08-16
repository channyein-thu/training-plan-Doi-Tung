# Training Plan API

A Go Fiber REST API for managing training plans, staff registrations, training records, and certificate approvals with role-based access (Admin, Manager, Staff). It uses MySQL (GORM), JWT authentication, optional Google OAuth2 login, Google Calendar integration for training events, and local file storage for certificate uploads.

## Architecture
- **Entry point:** `main.go` wires config, DB, middleware, routes, and storage.
- **Layers:** Controller → Service → Repository → Model.
- **DTOs:** `data/request` and `data/response` define API contracts and response shapes.
- **Integrations:** Google OAuth2 and Google Calendar via service account.

## Folder Overview
- **config/**: env config and DB connection/migrations.
- **container/**: dependency injection setup for controllers/services/repos.
- **controller/**: HTTP handlers (auth, user, training plan, record, certificate, department).
- **service/**: business logic, validation, orchestration.
- **repository/**: DB access (GORM) and query logic.
- **model/**: domain entities and enums.
- **router/**: route registration and role-based groups.
- **middleware/**: JWT auth, role checks, profile completion, error handler.
- **helper/**: utilities (password hashing, JWT, calendar, storage, error helpers).
- **mapper/**: model <-> DTO transformations.
- **data/**: request/response DTOs.
- **seed/**: development seed data (admin user).

## Key Workflows
- **Auth**: Role-based login plus generic `/api/v1/auth/login`, and Google OAuth2 via root-level endpoints. JWT is used for protected routes. Staff access requires profile completion.
- **Training Plans**: Admin creates/updates/deletes plans; optionally syncs with Google Calendar.
- **Records**: Managers register staff for training, update attendance and scores; staff view their own records; admin can search/export.
- **Certificates**: Staff upload certificates; admin approves or rejects. Managers also have staff-like certificate endpoints.

## Running Locally
### Prerequisites
- Go 1.25+
- MySQL 8+
- Google service account JSON (for calendar features)

### Run with Go
1. Create or update `app.env` with required variables.
2. Start MySQL and create the database.
3. Run:
   ```bash
   go run ./main.go
   ```

### Run with Docker
```bash
docker compose up --build
```

## Environment Variables
These keys are read from `app.env` and/or environment variables. Use safe values in production.
- **Database**: `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DB`
- **Uploads**: `UPLOAD_PATH` (local folder for certificate files)
- **JWT**: `JWT_SECRET` (read directly from env in token helper)
- **CORS**: `ALLOWED_ORIGINS` (currently not wired; CORS is hardcoded to `http://localhost:3000` in `main.go`)
- **Google OAuth2**: `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`, `GOOGLE_REDIRECT_URL`
- **Google Calendar**: `GOOGLE_SERVICE_ACCOUNT_FILE`, `GOOGLE_CALENDAR_ID`, `TIMEZONE`
- **Other app.env entries**: `GO_ENV`, `COOKIE_DOMAIN`, `COOKIE_SAMESITE`, `COOKIE_SECURE` are present but not used by the current code.

## API Routes
### Base prefix
- `/api/v1` for auth, admin, manager, and staff groups.

### Public/Auth (under `/api/v1`)
- `POST /auth/admin/login`
- `POST /auth/manager/login`
- `POST /auth/manager/register`
- `POST /auth/staff/login`
- `POST /auth/staff/register`
- `POST /auth/login`
- `GET /auth/me` (JWT)
- `GET /departments-list`

### Root-level public routes
- `GET /auth/google/login`
- `POST /auth/google/exchange`
- `POST /user/complete-profile` (JWT)
- `GET /uploads/*` (static files)

### Admin (JWT + AdminOnly)
- Departments: `POST /admin/departments`, `PUT /admin/departments/:id`, `DELETE /admin/departments/:id`, `GET /admin/departments`, `GET /admin/departments/:id`
- Department List: `GET /admin/departments-list`
- Users: `POST /admin/users`, `PUT /admin/users/:id`, `DELETE /admin/users/:id`, `GET /admin/users`, `GET /admin/users/:id`
- Training Plans: `POST /admin/training-plans`, `PUT /admin/training-plans/:trainingPlanId`, `DELETE /admin/training-plans/:trainingPlanId`, `GET /admin/training-plans`, `GET /admin/training-plans/:trainingPlanId`
- Records: `POST /admin/records/search`, `POST /admin/records/export`
- Certificates: `GET /admin/certificates`, `PUT /admin/certificates/:id/approve`, `PUT /admin/certificates/:id/reject`

### Manager (JWT + ManagerOnly)
- Users: `POST /manager/users`, `GET /manager/users`
- Training Plans: `GET /manager/training-plans`, `GET /manager/training-plans/:trainingPlanId`
- Records: `GET /manager/records`, `GET /manager/records/:id`, `PUT /manager/records/:id`, `DELETE /manager/records/:id`
- Registration: `POST /manager/training-plans/:trainingPlanId/registrations`
- Staff-like endpoints: `GET /manager/staffrecords`, `GET /manager/staffrecords/:id`, `GET /manager/certificates`, `POST /manager/certificates`, `DELETE /manager/certificates/:id`

### Staff (JWT + RequireProfileComplete)
- Records: `GET /staff/records`, `GET /staff/records/:id`
- Certificates: `GET /staff/certificates`, `POST /staff/certificates`, `DELETE /staff/certificates/:id`

### Healthcheck (per group)
- `GET /auth/healthchecker`, `GET /admin/healthchecker`, `GET /manager/healthchecker`, `GET /staff/healthchecker`

## Data Models (Summary)
- **User**: role/status, department link, OAuth fields, profile-complete flag, created-by metadata (`CreatedBy`, `CreatedByID`).
- **Department**: name and fixed `Division` enum, with staff list.
- **TrainingPlan**: name/category/type/date, plus content, speaker, duration, location, cost, budget code, and calendar event ID.
- **Record**: staff enrollment with status, evaluation, and pre/post test scores.
- **Certificate**: training-linked certificate with `TrainingID`, status (Pending/Approved/Rejected), and image storage.

## Interfaces (Contracts)
- `service/interfaces.go` defines service contracts used by controllers.
- `repository/interfaces.go` defines repository contracts used by services.

## Function Index (By File)
### main.go
- `main` - App bootstrap: config, DB, middleware, routes, storage, server startup.

### config/database.go
- `ConnectionDB` - Connect to MySQL and auto-migrate models.

### config/load_env.go
- `LoadConfig` - Load `app.env` into `Config`.

### container/app.go
- `NewAppDependencies` - Build repos, services, controllers, and return dependencies.

### seed/admin_seeder.go
- `SeedAdmin` - Create HR department and default admin (dev-only).

### router/router.go
- `RegisterRoutes` - Register public and role-based routes.

### router/auth_routes.go
- `AuthRoutes` - Register auth endpoints.

### router/admin_routes.go
- `AdminRoutes` - Register admin endpoints.

### router/manager_routes.go
- `ManagerRoutes` - Register manager endpoints.

### router/staff_routes.go
- `StaffRoutes` - Register staff endpoints.

### middleware/error_handler.go
- `ErrorHandler` - Standard error response mapping.

### middleware/jwt.go
- `JWTProtected` - Validate bearer token and set user locals.
- `AdminOnly` - Restrict to admin role.
- `ManagerOnly` - Restrict to manager role.
- `StaffOnly` - Restrict to staff role.
- `AdminOrManager` - Restrict to admin or manager.
- `GetJWTSecret` - Read JWT secret from env.

### middleware/completeProfile.go
- `RequireProfileComplete` - Block incomplete OAuth profiles.

### helper/app_error.go
- `(*AppError).Error` - Error string output.
- `ValidationError` - Build validation error response.
- `FormatValidationError` - Convert validator errors to field map.
- `BadRequest`, `NotFound`, `Internal`, `Unauthorized`, `Forbidden`, `InternalServerError` - Error helpers.

### helper/password.go
- `GeneratePassword` - Hash password with bcrypt.
- `ComparePassword` - Verify bcrypt hash.

### helper/token.go
- `GenerateAccessToken` - Create JWT access token.
- `VerifyAccessToken` - Validate and parse access token.
- `ExtractUserID` - Read user ID from claims.
- `ExtractUserRole` - Read role from claims.
- `GenerateToken` - Alias for access token creation.
- `VerifyToken` - Alias for access token verification.

### helper/google_calendar.go
- `NewGoogleCalendarService` - Initialize Calendar client using service account.
- `LoadLocation` - Resolve timezone.
- `CreateTrainingPlanCalendarEvent` - Create calendar event.
- `UpdateTrainingPlanCalendarEvent` - Update calendar event.
- `DeleteTrainingPlanCalendarEvent` - Delete calendar event.

### helper/upload.go
- `NewLocalStorage` - Create local storage adapter.
- `(*LocalStorage).Upload` - Save file to disk.
- `(*LocalStorage).Delete` - Remove file from disk.

### helper/error.go
- `ErrorPanic` - Panic on error.

### model/user.go
- `Role.IsValid` - Validate role enum.

### mapper/training_plan_mapper.go
- `ToTrainingPlanModel` - Request -> model.
- `ToTrainingPlanResponse` - Model -> response.
- `UpdateTrainingPlanFromRequest` - Apply updates to model.
- `ToTrainingPlanResponseList` - Batch conversion.

### mapper/record_mapper.go
- `ToRecordResponse` - Build full record response.

### data/response/user_response.go
- `ToUserTableResponse` - Model -> table response.
- `ToUserTableResponses` - Batch conversion.
- `ToUserResponse` - Model -> detail response.
- `ToUserListResponse` - Model -> list response.
- `ToUserListResponses` - Batch conversion.

### controller/auth_controller.go
- `NewAuthController` - Create auth controller.
- `AdminLogin` - Admin login handler.
- `ManagerLogin` - Manager login handler.
- `ManagerRegister` - Manager registration handler.
- `StaffLogin` - Staff login handler.
- `StaffRegister` - Staff registration handler.
- `Login` - Generic login handler (any role).
- `GetMe` - Current user profile handler.
- `handleLogin` - Shared login logic.
- `handleRegister` - Shared registration logic.
- `validateRegisterRequest` - Validate registration payload.
- `checkDuplicateUser` - Ensure unique email/employee ID.
- `createUserFromRequest` - Build user model from request.
- `buildUserResponse` - Build login response payload.

### controller/auth_oauth_controller.go
- `NewAuthOAuthController` - Create OAuth controller.
- `GoogleLogin` - Start OAuth login flow.
- `GoogleExchange` - Exchange OAuth code for JWT.
- `generateStateToken` - CSRF state token helper.

### controller/user_controller.go
- `NewUserController` - Create user controller.
- `AdminCreate` - Admin user creation.
- `AdminUpdate` - Admin user update.
- `AdminDelete` - Admin user delete.
- `AdminFindAll` - Admin list users.
- `AdminFindById` - Admin get user by id.
- `ManagerCreate` - Manager creates staff user.
- `ManagerFindDepartmentUsers` - Manager lists department staff.
- `getManagerDepartmentID` - Helper for manager department.
- `CompleteProfile` - OAuth profile completion.

### controller/department_controller.go
- `NewDepartmentController` - Create department controller.
- `Create` - Create department.
- `Update` - Update department.
- `Delete` - Delete department.
- `FindById` - Get department by id.
- `FindPaginated` - Paginated department list.
- `GetDepartmentsList` - Department dropdown list.

### controller/training_plan_controller.go
- `NewTrainingPlanController` - Create training plan controller.
- `Create` - Create training plan.
- `Update` - Update training plan.
- `Delete` - Delete training plan.
- `FindById` - Get training plan by id.
- `FindPaginated` - Paginated training plans.

### controller/record_controller.go
- `NewRecordController` - Create record controller.
- `RegisterStaff` - Register staff to training plan.
- `FindById` - Get record by id.
- `Update` - Update record details.
- `Delete` - Delete record.
- `FindRecordByCurrentDepartment` - Manager record list by department.
- `FindByCurrentUser` - Staff record list.
- `Search` - Admin search records.
- `Export` - Export records to Excel.

### controller/certificate_controller.go
- `NewCertificateController` - Create certificate controller.
- `FindByCurrentUser` - Staff list of certificates.
- `Upload` - Staff upload certificate.
- `Delete` - Staff delete certificate.
- `FindAllPending` - Admin list pending certificates.
- `Approve` - Admin approve certificate.
- `Reject` - Admin reject certificate.

### service/auth_oauth_service.go
- `NewAuthOAuthServiceImpl` - Create OAuth service.
- `GetGoogleLoginURL` - Build OAuth login URL.
- `HandleGoogleCallback` - Exchange code, upsert user, return JWT.
- `fetchGoogleUserInfo` - Fetch Google profile data.

### service/user_service.go
- `NewUserServiceImpl` - Create user service.
- `AdminCreate` - Admin create user.
- `AdminUpdate` - Admin update user.
- `AdminDelete` - Admin delete user.
- `AdminFindAll` - Admin list users.
- `AdminFindById` - Admin get user.
- `AdminFindAllForTable` - Admin list for table view.
- `ManagerCreate` - Manager create staff.
- `ManagerFindByDepartment` - Manager list department staff.
- `CompleteProfile` - OAuth user profile completion.

### service/department_service.go
- `NewDepartmentServiceImpl` - Create department service.
- `Create` - Create department.
- `Delete` - Delete department.
- `FindPaginated` - Paginated department list.
- `FindById` - Get department by id.
- `FindDepartmentList` - Department dropdown list.
- `Update` - Update department.

### service/training_plan_service.go
- `NewTrainingPlanServiceImpl` - Create training plan service.
- `Create` - Create training plan (and calendar event).
- `Delete` - Delete training plan.
- `FindById` - Get training plan.
- `FindPaginated` - Paginated training plans.
- `Update` - Update training plan.

### service/record_service.go
- `NewRecordServiceImpl` - Create record service.
- `Search` - Admin record search.
- `FindByUser` - Staff record list.
- `RegisterStaff` - Enroll staff.
- `FindById` - Get record by id.
- `Update` - Update record status/scores.
- `Delete` - Delete record.
- `FindByManager` - Manager record list by department.
- `Export` - Export records to Excel.

### service/certificate_service.go
- `NewCertificateServiceImpl` - Create certificate service.
- `Approve` - Approve certificate.
- `Reject` - Reject certificate.
- `FindAllPending` - Admin pending certificates.
- `FindByCurrentUser` - Staff list of certificates.
- `Upload` - Upload certificate image.
- `Delete` - Delete certificate.

### repository/department_repository.go
- `NewDepartmentRepositoryImpl` - Create repository.
- `Save` - Create department.
- `FindById` - Get department.
- `FindByIdWithStaffCount` - Get department with staff count.
- `FindDepartmentList` - List departments.
- `Update` - Update department.
- `Delete` - Delete department.
- `FindAllPaginated` - Paginated department list with staff count.

### repository/training_plan_repository.go
- `NewTrainingPlanRepositoryImpl` - Create repository.
- `FindPaginated` - Paginated training plans.
- `Delete` - Delete training plan.
- `FindAll` - List all training plans.
- `FindById` - Get training plan.
- `Save` - Create training plan.
- `Update` - Update training plan.

### repository/user_repository.go
- `NewUserRepositoryImpl` - Create repository.
- `Save` - Create user.
- `Update` - Update user.
- `UpdateProfile` - Partial profile update.
- `UpdateOAuthFields` - Update OAuth data.
- `Delete` - Delete user.
- `FindById` - Get user.
- `FindByIdWithDepartment` - Get user with department.
- `FindByEmail` - Find user by email.
- `FindByEmployeeID` - Find by employee ID.
- `FindAllPaginated` - Paginated users.
- `FindByDepartmentPaginated` - Paginated users by department.
- `ExistsByEmail` - Email uniqueness check.
- `ExistsByEmployeeID` - Employee ID uniqueness check.
- `FindAllWithFilters` - Search/sort/filter users.

### repository/record-repository.go
- `NewRecordRepositoryImpl` - Create repository.
- `Search` - Complex record filtering.
- `FindByUserId` - Records for staff.
- `FindByManagerDepartment` - Records for manager department.
- `Delete` - Delete record.
- `Exists` - Check duplicate enrollment.
- `FindById` - Get record.
- `Save` - Create record.
- `Update` - Update record.

### repository/certificate_repository.go
- `NewCertificateRepositoryImpl` - Create repository.
- `Save` - Create certificate.
- `FindById` - Get certificate.
- `FindByUserId` - List user certificates.
- `Delete` - Delete certificate.
- `UpdateStatus` - Update approval status.
- `FindAllPending` - Paginated pending certificates.
