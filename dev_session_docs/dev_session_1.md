# FamLedger Go + Supabase Project Summary

## **1. Go Project & Private Repo**
- Backend and frontend exist in the same repository (`famledger`).  
- Private GitHub repo imports require correct module paths and SSH/HTTPS credentials:  

```go
import "github.com/TheIronRock95/famledger/internal/db"
import "github.com/TheIronRock95/famledger/internal/api"
import "github.com/TheIronRock95/famledger/internal/models"
```

- Go modules require that the module path in `go.mod` matches import paths.  

**Lessons Learned:**
- Private repos need SSH keys or HTTPS credentials.  
- Go modules fetch dependencies based on **module paths**, not local directories.  

**Key Commands:**
```bash
go mod tidy        # Ensure all dependencies are resolved
go get <module>    # Fetch a specific Go module
go build ./cmd/server
go run cmd/server/main.go
```

---

## **2. Supabase Local vs Online**
- Local development DB works fine; online DB initially empty.  
- UUID for `user_id` caused errors: `22P02: invalid input syntax for type uuid: ""`.  
- Solution: change `user_id` column type to `text`:

```sql
ALTER TABLE public.income
ALTER COLUMN user_id TYPE text USING user_id::text;
```

**Lessons Learned:**
- Supabase CLI does **not automatically sync data** between local and remote.  
- UUIDs must be valid when inserting; otherwise, PostgreSQL rejects the insert.  
- Use **service_role key** for server-side writes to bypass RLS.  

---

## **3. Go + Supabase Connection**
- `.env` initially used local URL and anon key → not suitable for online write operations.  
- Correct setup uses **online SUPABASE_URL + service_role key**:  

```dotenv
SUPABASE_URL=https://<your-project>.supabase.co
SUPABASE_KEY=<service_role_key>
```

- Go must load `.env` before initializing the database:

```go
import "github.com/joho/godotenv"

err := godotenv.Load()
if err != nil {
    log.Fatal("Error loading .env file")
}
db.InitDatabase()
```

- Supabase-go v0.5.0: `Execute()` expects a pointer for the response.  

**Lessons Learned:**
- Service_role key is mandatory for server-side write access.  
- Environment variables must be loaded **before DB initialization**.  

---

## **4. Migrations & Schema Management**
- `supabase db push` pushes **schema only**.  
- `supabase db pull` pulls **remote schema** only, not the data.  
- Data migration requires **manual scripts** or **pg_dump/seed scripts**.  

**Lessons Learned:**
- Schema and data are separate in Supabase.  
- Seed scripts are recommended for moving test or production data between local and online.  

**Key Commands:**
```bash
# Push local migrations to the remote database
supabase db push

# Pull remote database schema to local
supabase db pull

# Repair migration history if inconsistent
supabase migration repair --status reverted <migration_number>

# Run SQL in Supabase local database
psql -h 127.0.0.1 -p 54322 -U postgres -d postgres -f migrations/001_init_schema.sql
```

---

## **5. API Testing via Curl**
**Local:**
```bash
curl -X POST http://localhost:8080/income \
-H "Content-Type: application/json" \
-d '{"user_id":"123","source":"Salary","amount":2500}'
```

**Online (after service_role key configured):**
```bash
curl -X POST http://localhost:8080/income \
-H "Content-Type: application/json" \
-d '{"user_id":"123","source":"Salary","amount":2500}'
```

- GET all entries:

```bash
curl http://localhost:8080/income
```

**Lessons Learned:**
- Curl testing is useful for quickly verifying backend functionality and DB writes.  
- Always ensure the correct Supabase key (service_role for server) is used when writing to the remote DB.  

---

## **6. Key Points & Takeaways**
1. Private Go repo → correct module path + SSH/HTTPS credentials.  
2. Local Supabase ≠ online Supabase; schema push does not sync data.  
3. UUID vs text → adjust schema to avoid insert errors.  
4. Use **service_role key** for server-side online DB writes.  
5. Migrations push → schema only; seed scripts needed for data.  
6. Environment variables must be correctly loaded for database initialization.  
7. Use curl for testing API endpoints against both local and online DB.  

---

### ✅ **Recommended CLI Reference**
```bash
# Go module management
go mod tidy
go get <module>
go build ./cmd/server
go run cmd/server/main.go

# Supabase
supabase db push                     # Push local migrations to remote
supabase db pull                     # Pull remote schema to local
supabase migration repair --status reverted <migration_number>

# Postgres local access
psql -h 127.0.0.1 -p 54322 -U postgres -d postgres -f migrations/001_init_schema.sql
```

