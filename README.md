## Full Stack Develoment Tutorial Application
This sample application uses various kinds of modern software such as React+Redux, TypeScript, Docker, API Server(Golang), BFF(express) and so on.   
  
Although this application is for technical education purpose, the core feature is very simple. You can stock Japanese restaurant via posting your favorite it's link of Tabelog(https://tabelog.com/) pages.
after you posting it, your link is passed to scraping server(implemented by Golang) via Internal HTTP and your shop info isstocked RDBMS.  you can view your stocked restaurants via web Frontend(implemented by React+Redux+TypeScript)  
  
Whole System is composed as MicroService Architecture, so you can learn how to create API Based microservices and how to use Docker and something else. 

## System Configuration
![infrastructure](https://github.com/aweglteo/fullstack_development/blob/master/doc/images/infra.png?raw=true)

## Getting Start
1.  `docker-compose build`
2. `dockr-compose up -d`


That's all.  
Let's enjoy your full stack travel !!
