export function convertAPIData(data: any) {
    switch (typeof data)
    {
        case "object": {
            if (data instanceof Array) {
                console.log("Array")
                break
            }
            console.log("Object")
            break
        }
        default: {
            console.log(data)
        }
    }
}