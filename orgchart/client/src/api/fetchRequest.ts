const fetchPostRequest = async (
    requestUrl: string,
    dataObject?: object,
    ignoredStatuses?: number[]
) => {
    return fetch(requestUrl, {
        method: 'POST',
        body: JSON.stringify(dataObject)
    }).then(response => {
        if (!response.ok) {
            if (ignoredStatuses && ignoredStatuses.includes(response.status)) {
                return response
            }
            throw new Error()
        }
        return response
    })
}

const fetchGetRequest = async (
    requestUrl: string
) => {
    return fetch(requestUrl, {
        method: 'GET',
    }).then(response => {
        if (!response.ok) {
            throw new Error(`Failed with status ${response.status}: ${response.statusText}`);
        }
        return response.json()
    })
}

const fetchUpdateRequest = async (
    requestUrl: string,
    dataObject?: object,
    ignoredStatuses?: number[]
) => {
    return fetch(requestUrl, {
        method: 'PATCH',
        body: JSON.stringify(dataObject)
    }).then(response => {
        if (!response.ok) {
            if (ignoredStatuses && ignoredStatuses.includes(response.status)) {
                return response
            }
            throw new Error()
        }
        return response
    })
}


const fetchDeleteRequest = async (
    requestUrl: string,
    ignoredStatuses?: number[]
) => {
    console.log(requestUrl)
    return fetch(requestUrl, {
        method: 'DELETE',
    }).then(response => {
        if (!response.ok) {
            if (ignoredStatuses && ignoredStatuses.includes(response.status)) {
                return response
            }
            throw new Error()
        }
        return response
    })
}


export {fetchPostRequest, fetchGetRequest, fetchDeleteRequest, fetchUpdateRequest}