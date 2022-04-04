package config

/// User

// Login url params: user password
const Login = "/alfresco/service/api/login.json?u=%s&pw=%s"

// UserInfo url params: ticket
const UserInfo = "/alfresco/service/api/login/ticket/%s"

//Workflows

//Get
const GetAllDef = "/alfresco/service/api/workflow-definitions?alf_ticket=%s"
const GetWorkfInst = "/alfresco/service/api/workflow-instances?definitionId=%s&state=active&alf_ticket=%s"
const GetWorkfInstTask = "/alfresco/service/api/workflow-instances/%s?includeTasks=true&alf_ticket=%s"

///Node

//NodeSearchPath params: path ticket
const NodeSearchPath = "/nodeservice/node/path?path=%s&alf_ticket=%s"

//NodeSearchId params: path ticket
const NodeSearchId = "/nodeservice/tenant/%s/node/%s?alf_ticket=%s"

//NodeSecondaryChildren params: path ticket
const NodeSecondaryChildren = "/nodeservice/tenant/totalcheck/node/%s/children/secondary?alf_ticket=%s"

//NodeDownload params: path ticket
const NodeDownload = "/nodeservice/download/tenant/totalcheck/node/%s?alf_ticket=%s"

//NodeTextEditor url params: nodeId   ticket
const NodeTextEditor = `/api/nodeservice/preview/tenant/totalcheck/node/%s?alf_ticket=%s`
