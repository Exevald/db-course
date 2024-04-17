export function convertAPIData(data: any) {
    switch (data.type)
    {
        case "Object": {
            console.log("Object")
            break
        }
        default: {
            console.log(data.type)
        }
    }
}