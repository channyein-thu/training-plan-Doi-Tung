# Training Plan — Mae Fah Luang Foundation (Doi Tung)

A training-management platform for an organization with three roles — **HR Admin**, **Department Manager**, and **Staff** — covering training plans, staff enrollment, on-the-job training (OJT) records, certificate approval, and dashboards.

## layout

```
training-plan-Doi-Tung/
├── backend/        # Go (Fiber) REST API + MySQL (GORM)
├── admin-web/      # Next.js admin dashboard (HR Admin)
├── manager-web/    # Next.js app for Department Managers and Staff
├── docker-compose.yml         # production stack (backend, admin-web, manager-web, db, caddy)
├── docker-compose.local.yml   # local stack (no TLS / no Caddy)
├── Caddyfile                  # reverse proxy + automatic HTTPS
├── DEPLOY.md                  # deployment guide
└── .env.example               # compose environment template
```

## Tech stack

- **Backend:** Go 1.25, Fiber, GORM, MySQL 8, JWT auth, Google OAuth2, Google Calendar, Web Push (VAPID), local file storage for certificate uploads.
- **Frontends:** Next.js 16 (App Router) + React 19 + TypeScript, Tailwind CSS 4, Radix UI / shadcn, React Hook Form + Zod, Recharts (dashboards), lucide-react icons.

## Roles

| Role | App | Highlights |
|------|-----|-----------|
| HR Admin | admin-web | Departments, users (all roles), training plans, register users to plans, records search/export, certificate approval |
| Department Manager | manager-web | Department staff, register staff to plans, OJT records (scores/evaluation), own certificates |
| Staff | manager-web | Own training records, certificates, dashboard (training performance) |

## Key features

- **Auth:** JWT-based; login with **email or phone**; phone is required and unique per user.
- **Users:** create/edit with department, role, status, work start date; department dropdowns show `Name (Division)`.
- **Training plans & records:** enroll staff (with "select all"), track status (Register / Attended / Absent), pre/post-test scores and evaluation, filter & export records.
- **Certificates:** staff upload, admin approve/reject; served via an authenticated endpoint (not public).
- **Dashboards:** per-role stats; staff see a training-status donut (Register/Attended/Absent) and certificate/training counts.
- **Notifications:** in-app + Web Push.
- **Integrations:** Google OAuth2 login, Google Calendar sync for training events.

---

## Local development

### Prerequisites
- Go 1.25+, Node.js 22+, MySQL 8+ (or Docker Desktop)

### Option A — run each part directly

**Backend**
```bash
cd backend
cp app.env.example app.env    # then fill in values (see Environment below)
go run ./main.go              # serves on :8080
```

**Admin web**
```bash
cd admin-web
npm install
npm run dev                   # http://localhost:3000
```

**Manager/Staff web**
```bash
cd manager-web
npm install
npm run dev                   # http://localhost:3001 (or 3000 per script)
```

### Option B — full stack in Docker
```bash
docker compose -f docker-compose.local.yml up --build
```
- Admin → http://localhost:3000
- Manager/Staff → http://localhost:3001
- API → http://localhost:8080

## Environment

Backend config is read from `backend/app.env` (see `app.env.example`). Key variables:

- **Database:** `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_DB`, `MYSQL_USER`, `MYSQL_PASSWORD`
- **Auth:** `JWT_SECRET` (required, ≥32 chars — the server refuses to start without it)
- **Server:** `PORT`, `ALLOWED_ORIGINS` (comma-separated frontend origins)
- **Environment:** `GO_ENV` (`development` enables the admin seeder), `ADMIN_SEED_PASSWORD`
- **Uploads:** `UPLOAD_PATH`
- **Google OAuth2:** `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`, `GOOGLE_REDIRECT_URL`
- **Google Calendar:** `GOOGLE_SERVICE_ACCOUNT_FILE`, `GOOGLE_CALENDAR_ID`, `TIMEZONE`
- **Web Push:** `VAPID_PUBLIC_KEY`, `VAPID_PRIVATE_KEY`, `VAPID_SUBJECT`

Frontends read the backend origin from `NEXT_PUBLIC_BACKEND_ORIGIN` (defaults to `http://localhost:8080` in dev).

> **Never commit** `backend/app.env` or `backend/service-account.json` — they hold secrets and are gitignored.


## API overview

Base prefix `/api/v1`, with role-scoped groups:

- `/auth/*` — login (admin/manager/staff + generic), register, `me`
- `/admin/*` — departments, users, training plans, plan registrations, records search/export, certificates
- `/manager/*` — department users, training plans, plan registrations, records, certificates, notifications
- `/staff/*` — own records, certificates, dashboard stats, notifications


