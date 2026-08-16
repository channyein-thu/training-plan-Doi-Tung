# Manager - Training and Record Management System

This repository contains the manager and staff web app for training and record management. It is built with the Next.js App Router and focuses on department staff management, training plan registration, OJT record updates, and certificate handling.

## Tech Stack

- Next.js 16 App Router + React 19 + TypeScript
- Tailwind CSS 4, Radix UI primitives, class-variance-authority
- react-hook-form and zod for forms and validation
- lucide-react for icons

## Local Development

- Install dependencies: `npm install`
- Start dev server: `npm run dev`

The backend is expected at `http://localhost:8080` with API base `http://localhost:8080/api/v1` (see [lib/api/api.ts](lib/api/api.ts) and [app/api/api.ts](app/api/api.ts)).

## App Flow (High-Level)

- Root entry redirects to `/dashboard` in [app/page.tsx](app/page.tsx).
- App shell and fonts are defined in [app/layout.tsx](app/layout.tsx).
- Auth pages share [app/(auth)/layout.tsx](app/(auth)/layout.tsx).
- Authenticated routes share [app/(root)/layout.tsx](app/(root)/layout.tsx) which reads the current user with `getMe()` and redirects to `/login` if unauthenticated.

## Routes (App Router)

- `/login` - Manager login form in [app/(auth)/login/page.tsx](app/(auth)/login/page.tsx)
- `/register` - Manager registration form in [app/(auth)/register/page.tsx](app/(auth)/register/page.tsx)
- `/onboarding` - Complete profile flow in [app/onboarding/page.tsx](app/onboarding/page.tsx)
- `/dashboard` - Dashboard overview in [app/(root)/dashboard/page.tsx](app/(root)/dashboard/page.tsx)
- `/department-staff` - Staff list + search in [app/(root)/department-staff/page.tsx](app/(root)/department-staff/page.tsx)
- `/department-staff/[staffId]` - Staff detail shell in [app/(root)/department-staff/[staffId]/page.tsx](app/(root)/department-staff/[staffId]/page.tsx)
- `/register-staff` - Training plan list in [app/(root)/register-staff/page.tsx](app/(root)/register-staff/page.tsx)
- `/register-staff/[planId]` - Plan details + staff selection in [app/(root)/register-staff/[planId]/page.tsx](app/(root)/register-staff/[planId]/page.tsx)
- `/ojt-records` - OJT record table in [app/(root)/ojt-records/page.tsx](app/(root)/ojt-records/page.tsx)
- `/ojt-records/[ojtId]` - OJT record edit page in [app/(root)/ojt-records/[ojtId]/page.tsx](app/(root)/ojt-records/[ojtId]/page.tsx)
- `/my-training-records` - Staff training records in [app/(root)/my-training-records/page.tsx](app/(root)/my-training-records/page.tsx)
- `/my-training-records/[recordId]` - Upload certificate for record in [app/(root)/my-training-records/[recordId]/page.tsx](app/(root)/my-training-records/[recordId]/page.tsx)
- `/my-certificates` - Certificate gallery in [app/(root)/my-certificates/page.tsx](app/(root)/my-certificates/page.tsx)
- `/notifications` - Notification list in [app/(root)/notifications/page.tsx](app/(root)/notifications/page.tsx)

## API Helpers and Server Actions (Function Index)

- `authFetch()` - Authenticated fetch helper in [lib/api/authFetch.ts](lib/api/authFetch.ts)
- `getMe()` - Load current user session in [lib/api/getMe.ts](lib/api/getMe.ts)
- `getDepartmentList()` - Department options for registration in [lib/api/getDepartment.ts](lib/api/getDepartment.ts)
- `LoginAction()` - Login and set cookie token in [lib/actions/login.ts](lib/actions/login.ts)
- `RegisterAction()` - Manager registration in [lib/actions/register.ts](lib/actions/register.ts)
- `logoutAction()` - Clear auth cookie in [lib/actions/logout.ts](lib/actions/logout.ts)
- `completeProfileAction()` - Submit onboarding data in [lib/actions/completeProfile.ts](lib/actions/completeProfile.ts)
- `CreateStaffAction()` (default export) - Create department staff in [lib/actions/department-staff/createStaffAction.ts](lib/actions/department-staff/createStaffAction.ts)
- `getStaff()` - Fetch department staff list in [lib/actions/department-staff/getStaff.ts](lib/actions/department-staff/getStaff.ts)
- `getPlans()` and `getPlanById()` - Training plans in [lib/actions/register-staff/getPlanAction.ts](lib/actions/register-staff/getPlanAction.ts)
- `registerStaffAction()` - Register staff to plan in [lib/actions/register-staff/registerStaffAction.ts](lib/actions/register-staff/registerStaffAction.ts)
- `getStaff()` - Fetch staff list for registration in [lib/actions/register-staff/getStaffAction.ts](lib/actions/register-staff/getStaffAction.ts)
- `getOjtRecords()` - Manager OJT records in [lib/actions/ojt-records/getOjtAction.ts](lib/actions/ojt-records/getOjtAction.ts)
- `UpdateOjtRecordAction()` - Update OJT details in [lib/actions/ojt-records/updateOjtRecord.ts](lib/actions/ojt-records/updateOjtRecord.ts)
- `DeleteOjtAction()` - Delete OJT record in [lib/actions/ojt-records/deleteOjt.ts](lib/actions/ojt-records/deleteOjt.ts)
- `getRecords()` - Staff training records in [lib/actions/training-records/getRecords.ts](lib/actions/training-records/getRecords.ts)
- `getRecordById()` - Training record detail in [lib/actions/training-records/getRecordById.ts](lib/actions/training-records/getRecordById.ts)
- `uploadCertificate()` - Upload training certificate in [lib/actions/training-records/uploadCertificate.ts](lib/actions/training-records/uploadCertificate.ts)
- `getCertificates()` - Certificate fetch (API helper) in [lib/api/getCertificates.ts](lib/api/getCertificates.ts)
- `getCertificates()` - Certificate fetch (server action) in [lib/actions/my-certificates/getMyCertificate.ts](lib/actions/my-certificates/getMyCertificate.ts)
- `cn()` - Tailwind class helper in [lib/utils.ts](lib/utils.ts)

## Component Map (Feature Inventory)

### Dashboard

- [components/dashboard/pageHeader.tsx](components/dashboard/pageHeader.tsx)
- [components/dashboard/dashboardCard.tsx](components/dashboard/dashboardCard.tsx)
- [components/dashboard/logoCard.tsx](components/dashboard/logoCard.tsx)
- [components/dashboard/userInfo.tsx](components/dashboard/userInfo.tsx)
- [components/dashboard/searchBar.tsx](components/dashboard/searchBar.tsx)
- [components/dashboard/searchBarUI.tsx](components/dashboard/searchBarUI.tsx) (empty placeholder)

### Auth

- [components/login/logo.tsx](components/login/logo.tsx)
- [components/login/passwordInput.tsx](components/login/passwordInput.tsx)
- [components/register/register-form.tsx](components/register/register-form.tsx)

### Department Staff

- [components/department-staff/departmentCard.tsx](components/department-staff/departmentCard.tsx)
- [components/department-staff/addStaffForm.tsx](components/department-staff/addStaffForm.tsx)
- [components/department-staff/editStaffDialog.tsx](components/department-staff/editStaffDialog.tsx)
- [components/department-staff/deleteStaffDialog.tsx](components/department-staff/deleteStaffDialog.tsx)
- [components/department-staff/buttonDialog.tsx](components/department-staff/buttonDialog.tsx)

### Register Staff to Plans

- [components/register-staff/registerCard.tsx](components/register-staff/registerCard.tsx)
- [components/register-staff/planDetails.tsx](components/register-staff/planDetails.tsx)
- [components/register-staff/infoDetail.tsx](components/register-staff/infoDetail.tsx)
- [components/register-staff/staffList.tsx](components/register-staff/staffList.tsx)
- [components/register-staff/routerButton.tsx](components/register-staff/routerButton.tsx)
- [components/register-staff/registerStaffForm.tsx](components/register-staff/registerStaffForm.tsx)
- [components/register-staff/buttonDialog.tsx](components/register-staff/buttonDialog.tsx)

### OJT Records

- [components/ojt-records/ojtTable.tsx](components/ojt-records/ojtTable.tsx)
- [components/ojt-records/editOjtForm.tsx](components/ojt-records/editOjtForm.tsx)
- [components/ojt-records/deleteOJTdialog.tsx](components/ojt-records/deleteOJTdialog.tsx)

### Training Records and Certificates

- [components/my-training-records/trainingRecordTable.tsx](components/my-training-records/trainingRecordTable.tsx)
- [components/my-training-records/uploadCertificateForm.tsx](components/my-training-records/uploadCertificateForm.tsx)
- [components/my-certificates/certificateCard.tsx](components/my-certificates/certificateCard.tsx)

### Notifications

- [components/notification/notificationCard.tsx](components/notification/notificationCard.tsx)

### Shared UI Primitives

- [components/ui/app-sidebar.tsx](components/ui/app-sidebar.tsx)
- [components/ui/sidebar.tsx](components/ui/sidebar.tsx)
- [components/ui/button.tsx](components/ui/button.tsx)
- [components/ui/card.tsx](components/ui/card.tsx)
- [components/ui/checkbox.tsx](components/ui/checkbox.tsx)
- [components/ui/dialog.tsx](components/ui/dialog.tsx)
- [components/ui/dropdown-menu.tsx](components/ui/dropdown-menu.tsx)
- [components/ui/field.tsx](components/ui/field.tsx)
- [components/ui/form.tsx](components/ui/form.tsx)
- [components/ui/input.tsx](components/ui/input.tsx)
- [components/ui/label.tsx](components/ui/label.tsx)
- [components/ui/progress.tsx](components/ui/progress.tsx)
- [components/ui/select.tsx](components/ui/select.tsx)
- [components/ui/separator.tsx](components/ui/separator.tsx)
- [components/ui/sheet.tsx](components/ui/sheet.tsx)
- [components/ui/skeleton.tsx](components/ui/skeleton.tsx)
- [components/ui/table.tsx](components/ui/table.tsx)
- [components/ui/tooltip.tsx](components/ui/tooltip.tsx)
- [components/ui/avatar.tsx](components/ui/avatar.tsx)

## Data and Types

- Mock data for demos in [data/data.ts](data/data.ts)
- API and UI types in:
	- [types/certificate.ts](types/certificate.ts)
	- [types/course.ts](types/course.ts)
	- [types/department.ts](types/department.ts)
	- [types/records.ts](types/records.ts)
	- [types/staff.ts](types/staff.ts)

## Utilities

- Mobile breakpoint helper in [hooks/use-mobile.ts](hooks/use-mobile.ts)

## Scripts

See [package.json](package.json) for all scripts. Common ones:

- `npm run dev`
- `npm run build`
- `npm run start`
- `npm run lint`
