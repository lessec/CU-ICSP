<script>
    import { navigate } from "svelte-routing";
    
	export let user;
    export let getAlls;

    let id = user.id;
	
    // request to delete an user
    async function deleteUser(e) {
        let  url = 'http://127.0.0.1:5555/api/v1/user/delete/' + id;

        const options = {
            method: 'POST',
        };
        
        let res = await fetch(url, options).then(res => {
            console.log("new group added!");
            alert("Success!");
            getAlls();
        });
    }

    // go to edit user page
    function editUser() {
		navigate("user/edit/" + id, {replace: true });
	}
	
</script>

<div>
	<h3>{user.name}</h3>
    <ul>
        <li>email = {user.email}</li>
        <li>type = {user.type}</li>
        <li>group = {user.group}</li>
    </ul>
    <button on:click={deleteUser}>x</button>
    <button on:click={editUser}>edit</button>
</div>