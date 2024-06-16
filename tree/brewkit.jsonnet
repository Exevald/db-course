local project = import 'brewkit/project.libsonnet';

local appIDs = [
    'tree',
];

local openAPI = [
    'api/server/treepublic/treepublic.yaml'
];

project.project(appIDs, openAPI)