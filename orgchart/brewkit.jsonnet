local project = import 'brewkit/project.libsonnet';

local appIDs = [
    'orgchart',
];

local openAPI = [
    'api/server/orgchartpublic/orgchartpublic.yaml',
];

project.project(appIDs)