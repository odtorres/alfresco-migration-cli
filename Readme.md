 # alfmigcli CLI

 ## Modules
 1. user
 ``` 
 	login ( user pass) # alfmigcli user -login <user> <password>
```
 2. node
 ```
 	get # alfmigcli node get <unique_id>
```
 3. workflow
 
 ## example 
 1. alfmigcli <module> <action>


## examples
1. alfmigcli cluster -add QA https://alfrescoip:port
2. alfmigcli cluster -current 1
3. alfmigcli user -login system changeme

## Build
go build -o build/alfmigcli

## Run 
go run main.go cluster -describe

go run main.go workflow -getalldef  
go run main.go workflow -getinst 'activiti$procesoPrenda2:91:15264316'
go run main.go workflow -getinstask 'activiti$16264080'
go run main.go workflow -getallandsave 'activiti$procesoPrenda2:91:15264316'