# simplebank PostgreSQL client


Designing, developing, and deploying a complete backend system from scratch using Golang, PostgreSQL.


Steps Taken
1. Design DB schema and generate SQL code with dbdiagram.io
2. Generate scripts for running database migrations in Golang
3. Automatically generate Golang code to perform CRUD operations on the database with sqlc.
4. Write Goalgn unit tests for database CRUD with randomized data
5. Implement clean database transactions that combine operations from several tables in Golang
6. Implement safeguards to avoid deadlock in DB transactions
   
