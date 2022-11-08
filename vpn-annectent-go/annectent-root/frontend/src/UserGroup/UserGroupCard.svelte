<script>
    /*
    This is UserGroupCard Component which contains UserGroup.
    */
    
    import { navigate } from "svelte-routing";

	export let group;
    export let getAlls;

    let id = group.id;
	
    // request delete a group
    async function deleteUserGroup(e) {
        let  url = 'http://127.0.0.1:5555/api/v1/group/delete/' + id;

        const options = {
            method: 'POST',
        };
        
        let res = await fetch(url, options).then(res => {
            console.log("new group added!");
            alert("Success!");
            getAlls();
        });
    }

    // go to edit group page
    function editUserGroup() {
		navigate("usergroup/edit/" + id, {replace: true });
	}
	
</script>

<div>
	<h3>{group.name}</h3>
    <ul>
        <li>describe = {group.describe}</li>
        <li>granted server = {group.grantedServer}</li>
    </ul>
    <button on:click={deleteUserGroup}>x</button>
    <button on:click={editUserGroup}>edit</button>
</div>