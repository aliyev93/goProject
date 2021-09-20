# EXAMPLE APPLICATION

The applicaton is terting purpose to deploy on kubernetes cluster.

## Database

For running the application we need Mysql Server, DB user and database.
After creating Db and a user create table with [this](../db/employee) sql script
Make appropirate changes on [config](../config.yaml) to connect application to DB

## Usage 


```bash
GET /health
```
Simple health check backend

```bash
POST /employee
```
with body of 
```bash
{
    "employee_name": "Name",
    "employee_age": "Age"
}
```
Create new database record

```bash
GET /employee/:id
```
Queries user wirh id

```bash
DELETE /employee/:id
```
Deletes user from database with id
