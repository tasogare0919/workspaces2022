const client = require('@sendgrid/client');
client.setApiKey("");
let date = new Date();
const startTime = Math.floor( date.getTime());
const endTime = Math.floor( date.getTime() - 300000 )
const req = {
    method: "POST",
    url: "/v3/suppression/bounces",
    start_time: startTime,
    end_time: endTime
}
let bounceDetailText = ""

try {
    const res = client.request(req)
    if (res.length === 0) {
        console.log('No Bounce mail.')
    } else {
        res.map( bounce,i => {
            bounceDetailText = "```\n対象email: " + bounce[i].email + "\nバウンスの原因: " + bounce[i].reason +"\n```"
        })
    }
    console.log(bounceDetailText)
    console.log('Check Bounce mail sucessfully')
} catch (error) {
    console.error('Failed to chek Bounce mail.')
}