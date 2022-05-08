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

### windows
GOOS=windows GOARCH=amd64 go build -o build/alfmigcli main.go
GOARCH=386
### Mac
GOOS=darwin GOARCH=amd64 go build -o build/alfmigcli main.go
### Linux
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux app.go

## Run 
go run main.go cluster -describe

## migrate workflows
1. go run main.go workflow -getalldef  
2. go run main.go workflow -getallandsave 'activiti$procesoPrenda2:91:15264316'
3. go run main.go node -verifywfnodes
4. go run main.go workflow -create '[{ "processDefinitionKey": "activitiAdhoc", "variables": {"bpm_assignee": "fred"} }]'
+---+-----------+-------------------+
| # |    ID     | PROCESSDEFINITION | 
+---+-----------+-------------------+
| 1 | 160840729 | activitiAdhoc:1:4 |
+---+-----------+-------------------+
5. 

go run main.go workflow -getinst 'activiti$procesoPrenda2:91:15264316'
go run main.go workflow -getinstask 'activiti$16264080'

1 *cobra