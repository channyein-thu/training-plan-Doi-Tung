# Deployment Guide — Training Platform

Single-VM deployment: 5 Docker services (backend, admin-web, manager-web, db,
caddy) orchestrated by `docker-compose.yml`, with Caddy terminating TLS.

## 1. Assemble the monorepo

Place the three projects under this root with these exact folder names (the
compose build contexts depend on them):

```
training-platform/
  docker-compose.yml
  Caddyfile
  .env                       # from .env.example
  backend/                   # from goProjects/training-server
    app.env                  # from backend/app.env.example
    service-account.json     # Google service account (provided, not committed)
  admin-web/                 # from node/components
  manager-web/               # from node/manager
```

Move or clone each repo into place, e.g.:

```bash
mv /path/to/training-server  ./backend
mv /path/to/components        ./admin-web
mv /path/to/manager           ./manager-web
```

## 2. Configure

```bash
cp .env.example .env                     # domains, TLS email, DB creds
cp backend/app.env.example backend/app.env
```

Then edit both files with real values. Generate strong secrets:

```bash
# JWT signing key (backend/app.env -> JWT_SECRET). Must be >= 32 chars.
openssl rand -base64 48

# VAPID keys for Web Push (backend/app.env -> VAPID_PUBLIC_KEY / VAPID_PRIVATE_KEY)
# e.g. via:  npx web-push generate-vapid-keys
```

Set in `backend/app.env`: `JWT_SECRET`, `GOOGLE_CLIENT_ID`, `GOOGLE_CLIENT_SECRET`,
`GOOGLE_CALENDAR_ID`, `VAPID_*`. Leave `GO_ENV=production` (the compose file also
forces this) so the dev admin seeder never runs.

Set in `.env`: the three domains, `TLS_EMAIL`, and the DB credentials.

## 3. External prerequisites

- **DNS**: point `A`/`AAAA` records for `API_DOMAIN`, `APP_DOMAIN`, `ADMIN_DOMAIN`
  at this server. Caddy needs ports **80** and **443** reachable to issue certs.
- **Google OAuth**: in the Google Cloud console, add the redirect URI
  `https://<APP_DOMAIN>/api/auth/google/callback` to the OAuth client.

## 4. Deploy

```bash
docker compose up --build -d
docker compose ps
docker compose logs -f caddy    # watch certificate issuance
```

Caddy obtains Let's Encrypt certificates automatically on first request and
redirects HTTP→HTTPS. The apps are reachable at `https://<APP_DOMAIN>`,
`https://<ADMIN_DOMAIN>`, and the API at `https://<API_DOMAIN>`.

## 5. Create the first admin

The production build does not seed an admin. Create one directly (e.g. insert a
row with a bcrypt-hashed password, or temporarily run the backend with
`GO_ENV=development` + `ADMIN_SEED_PASSWORD` set, then revert).

---

## Secret & git hygiene (do BEFORE pushing anywhere public)

These files are currently tracked in the backend repo and should not be:

```bash
# In the backend repo:
git rm --cached dump.rdb "certificates/user_4/1779954293.jpeg"
git commit -m "Stop tracking local DB dump and an uploaded certificate"
```

`.gitignore` already ignores `app.env`, `service-account.json`, `uploads/`, and
`dump.rdb` for future changes, but the two files above were committed earlier and
remain in history.

**Rotate anything that was ever committed or shared**, since git history and any
pushed copies still contain the old values:

- Google OAuth client secret (`GOOGLE_CLIENT_SECRET`)
- VAPID private key (`VAPID_PRIVATE_KEY`)
- Database passwords

If the repo was pushed to a remote, also purge the two files from history
(`git filter-repo` or the BFG Repo-Cleaner) and force-push, then have all clones
re-clone.

> Actually removing/rotating these is tracked separately as security Task 13.

## Notes

- Only the `caddy` service publishes host ports (80/443); backend, web apps, and
  db are reachable only on the internal Docker network.
- Volumes `db_data` and `uploads` persist across `docker compose up`/`down`.
  Use `docker compose down -v` only if you intend to wipe them.
- Consider a managed MySQL (RDS/Cloud SQL) instead of the `db` container for
  automated backups and patching; point `MYSQL_HOST` at it and drop the `db`
  service.
- For local development you don't need any of this: `npm run dev` in each
  frontend and `go run ./main.go` in the backend use localhost fallbacks.
```
