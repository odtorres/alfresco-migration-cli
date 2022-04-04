 # alfmigcli CLI

 ## Modules
 1. user
 ``` 
 	login ( user pass) # albertocs user login <user> <password>
```
 2. node
 ```
 	get # albertocs node get <unique_id>
 	path # albertocs node path <path>
 	download # albertocs node download unique_id
 	modificar
 	upload # Albertocs node upload <path to file> —type <doc type>  —metadata ‘“{“hello”: “world”}”’
 	asociar # Albertocs node assoc child_id —to parent_id
 	update: modificar metadato #Albertocs node update unique_id —metadata ‘{}’
```
 3. iam
 4. workflow
 5. model (indices)
 6. config

 ## example 
 1. alfmigcli <module> <action>


## examples
1. alfmigcli cluster -add QA https://alfrescoip:port
2. alfmigcli cluster -current 1
3. alfmigcli user -login system changeme
4. alfmigcli node -json -sid=5cb15b2c-78da-11e8-a807-0a580a200520 totalchek
5. alfmigcli node -spath /tenants/totalcheck
6. alfmigcli <module> -h

## Build
go build -o build/alfmigcli

## Run 
go run main.go cluster -describe

go run main.go workflow -getalldef  
go run main.go workflow -getinst 'activiti$procesoPrenda2:91:15264316'
go run main.go workflow -getinstask 'activiti$16264080'
go run main.go workflow -getallandsave 'activiti$procesoPrenda2:91:15264316'