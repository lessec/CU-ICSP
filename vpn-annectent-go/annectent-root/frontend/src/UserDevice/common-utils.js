// get all devices from server
export let getAllDevices = async () => {
    let response = await fetch("http://127.0.0.1:5555/api/v1/device/get/all");
    let commits = await response.json();
    return commits;
}