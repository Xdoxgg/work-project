
# CHECK API'S

---
### USER-SERVICE
curl "http://localhost:8081/api/user?email=&password="

curl -X POST "http://localhost:8081/api/user?email=newuser@example.com&password=newpassword"

---

### MOVIE-SERVICE

curl "http://localhost:8082/api/movies"

curl -X POST "http://localhost:8082/api/movies?title=Inception&description=A%20thief%20who%20steals%20corporate%20secrets%20through%20the%20use%20of%20dream-sharing%20technology%20is%20given%20the%20inverse%20task%20of%20planting%20an%20idea%20into%20the%20mind%20of%20a%20CEO.&year=2010"

---


### RECOMENDATIONS-SERVICE

curl -X GET "http://localhost:8083/api/recomendations" \
-H "Content-Type: application/json" \
-d '{
    "tags": [{"id": 1, "title": "Action"}, {"id": 2, "title": "Adventure"}],
    "genres": [{"id": 1, "title": "Sci-Fi"}, {"id": 2, "title": "Drama"}]
}'