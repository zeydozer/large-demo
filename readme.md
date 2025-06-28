````markdown
# large-demo

A hands-on demo for learning and simulating a high-traffic REST API with:

- Go (Golang) backend
- PostgreSQL master-replica streaming replication
- Redis caching
- Nginx load balancing and multiple API containers
- Docker Compose orchestration
- K6 load testing
- 1M+ demo user records

---

## Quick Start

**Requirements:**  
- Docker & Docker Compose  
- (Recommended) At least 8GB RAM

### 1. Clone

```sh
git clone https://github.com/zeydozer/large-demo.git
cd large-demo
````

### 2. Build & Run

```sh
docker-compose down -v
docker volume prune -f
docker-compose up --build
```

*First run may take several minutes (1M+ users generated).*

---

## Usage

* **API:**
  `http://localhost:8080/users/{id}`

* **Postgres (master):**
  Host: `localhost` — Port: `5432` — User: `postgres` — Password: `postgres`

* **Postgres (replica):**
  Host: `localhost` — Port: `5433` — User: `postgres` — Password: `postgres`

* **Redis:**
  `localhost:6379`

---

## Load Testing

1. Install [k6](https://k6.io/).
2. Edit `test.js` for your load scenario.
3. Run:

   ```sh
   k6 run test.js
   ```

---

## Architecture Diagram

```
[ k6 ] → [ NGINX ] → [ API1 | API2 | API3 ] → [ Redis ] → [ Postgres Master ] → [ Postgres Replica ]
```

---

## Notes

* High error rates are expected under heavy load in a single-machine Docker environment.
* For real-world use: run services on separate machines, or in the cloud.
* Tune connection pools, indexes, and cache settings for your own experiments.