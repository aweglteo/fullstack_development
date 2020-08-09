# 📖 Full Stack Develoment Tutorial Application
This application uses various kinds of modern software such as React+Redux, TypeScript, Docker, API Server(Golang), BFF(express), gRPC(HTTP/2) and so on.   
  
Although this application is for technical education purpose, the core feature is very simple. You can stock your favorite Japanese restaurants via posting it's link of Tabelog(https://tabelog.com/) pages.
After you posting it, your link is passed to scraping server(implemented by Golang) via Internal gRPC and your shop info is stocked RDBMS(PostgreSQL). You can view your stocked restaurants via web Frontend(implemented by React+Redux+TypeScript)  
  
Whole System is composed as MicroService Architecture, so it's good materials to learn how to create API Based microservices, how to use Docker and k8s and how to deploy cloud infra(like GCP, AWS).

## System Configuration
![infrastructure](https://github.com/aweglteo/fullstack_development/blob/master/doc/images/infra.png?raw=true)

# Getting Start
```
docker-compose build
```
  
```
dockr-compose up -d
```
then, backend services are running.  
  
```
cd bff/
npm install
npm run webpack
npm run dev
```

That's all.  
Let's enjoy your full stack travel !!
  
  
# How to deploy in cloud services(GCP, AWS)
T.B.C  
