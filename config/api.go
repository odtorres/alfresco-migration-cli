package config

/// User

// Login url params: user password
const Login = "/alfresco/service/api/login.json?u=%s&pw=%s"

// UserInfo url params: ticket
const UserInfo = "/alfresco/service/api/login/ticket/%s"

//Workflows

//Get
const GetAllDef = "/alfresco/service/api/workflow-definitions?alf_ticket=%s"
const GetWorkfInst = "/alfresco/service/api/workflow-instances?definitionName=%s&state=active&alf_ticket=%s"
const GetWorkfInstTask = "/alfresco/service/api/workflow-instances/%s?includeTasks=true&state=active&alf_ticket=%s"
const PostCreateWorkflow = "/alfresco/api/-default-/public/workflow/versions/1/processes?alf_ticket=%s"
const PutUpdateTask = "/alfresco/api/-default-/public/workflow/versions/1/tasks/%s?select=state,variables&alf_ticket=%s"

///Node

//NodeSearchPath params: path ticket
const NodeSearchUUID = "/alfresco/service/api/version?nodeRef=%s&alf_ticket=%s"
