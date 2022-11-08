// get all groups from server
export let getAllGroups = async () => {
    let response = await fetch("http://127.0.0.1:5555/api/v1/group/get/all");
	let commits = await response.json();
    return commits;
}