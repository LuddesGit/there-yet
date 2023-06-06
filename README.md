# There Yet

## Available Config
Service can be configured with ENV:

- OSRM_URL   
default http://router.project-osrm.org/route/v1/driving

- PORT  
default 8080

## To build and run as a docker image
```bash
docker build -t there-yet-image .
docker run -d -p 8080:8080 there-yet-image:latest
```

Service will now be available at localhost:8080/routes  
Example query: http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=14.428555,52.523219