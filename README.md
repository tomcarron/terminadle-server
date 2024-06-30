# terminadle-server

## Database stuff 

Verifying database:
docker exec -it postgres_db bash
psql -U my_user -d my_database
\dt
SELECT * FROM my_table;


removing volumes
docker-compose down -v


starting docker container:
docker-compose up --build
