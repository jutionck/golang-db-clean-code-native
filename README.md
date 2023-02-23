## Golang Clean Code Native Query

- Case: Vehicle

### Setup Configuration

1. Copy and rename name `.env.example` become `.env`
2. Configure here

```env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=db_car_shop
DB_USER=postgres
DB_PASSWORD=
API_PORT=8888
API_HOST=localhost
```

### Setup Database

```sql
create table vehicle (
	id varchar(100) primary key,
	brand varchar(200) not null,
	model varchar(200) not null,
	year int,
	weight float8
);
```

### Run Project

```bash
go run .
```

### Api Documentation

- `POST` Vehicle
- `localhost:8888/vehicle`
- Payload

```json
{
  "brand": "Toyota",
  "model": "New Avanza GT",
  "year": 2023,
  "weight": 230
}
```

- `PUT` Vehicle
- `localhost:8888/vehicle`
- Payload

```json
{
  "id": "52adb410-12d6-4d43-90ac-7b4b46e5550a",
  "brand": "Toyota",
  "model": "New Avanza G",
  "year": 2023,
  "weight": 230
}
```

- `GET` Vehicle -> `localhost:8888/vehicle`

- `GET` Vehicle -> `localhost:8888/vehicle/<id>`

- `DELETE` Vehicle -> `localhost:8888/vehicle/<id>`
