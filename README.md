
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