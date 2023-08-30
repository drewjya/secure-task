rebuild:
	docker rmi secure-task-app --force && docker compose  up -d

generate:
	sqlc generate
	sed -i -e 's/sql.NullString/*string/g' ./repository/*.go
	sed -i -e 's/sql.NullInt32/*int32/g' ./repository/*.go
	sed -i -e 's/sql.NullInt64/*int64/g' ./repository/*.go
	sed -i -e 's/sql.NullBool/*bool/g' ./repository/*.go
	sed -i -e 's/sql.NullFloat64/*float64/g' ./repository/*.go
	sed -i -e 's/sql.NullTime/*time.Time/g' ./repository/*.go
	
	goimports -w ./repository/