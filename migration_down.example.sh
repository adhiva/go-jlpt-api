#bin/bash!
migrate -database "mysql://your_username:your_password@tcp(your_url_mysqldb:your_port_mysqldb)/your_db" -path db/migrations down