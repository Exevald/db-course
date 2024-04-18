local project = import 'brewkit/project.libsonnet';

local appIDs = [
    'orgchart',
    'integrationaltests'
];

local openAPI = [
    'api/server/orgchartpublic/orgchartpublic.yaml'
];

project.project(appIDs, openAPI)