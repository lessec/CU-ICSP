<script>
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import { navigate } from "svelte-routing";
    import { onMount } from "svelte";
    
    export let id;

    let name = "";
    let describe = "";
    let grantedServer = "";
    

    // get a group from server
    async function getUserGroup(id) {
        let url = `http://127.0.0.1:5555/api/v1/group/get/id/${id}`;
       
		let response = await fetch(url);
        console.log(response);
        let data = await response.json();
        console.log(data);
        return data;
	}

    // lifecycle function
    onMount(async () => {
            getUserGroup(id).then(data => {
            console.log(data);
            name = data.name;
            describe = data.describe;
            grantedServer = data.grantedServer;
        });
    });

    // request update a group
    async function updateUserGroup() {
        let url = `http://127.0.0.1:5555/api/v1/group/update/${id}`;

        const params = {
            name: name,
            describe: describe,
            grantedServer: grantedServer
        };

        const options = {
            mode: 'no-cors',
            method: 'POST',
            body: JSON.stringify(params),
            headers: {
                'Content-Type': 'application/json'
            }
        };

        await fetch(url, options).then(response => {
            console.log(response);
            alert("Success!");
            navigate("/usergroup/detail/" + id, {replace: true });
        });
    }

    function goToBack() {
        navigate("/usergroup/detail/" + id, {replace: true });
    }

    // request delete a group
    async function deleteUserGroup() {
        let  url = 'http://127.0.0.1:5555/api/v1/group/delete/' + id;

        const options = {
            method: 'POST',
        };
        
        let res = await fetch(url, options).then(res => {
            console.log("new group added!");
            alert("Success!");
            navigate('/usergroup', {replace: true });
        });
    }
</script>

<CommonAppBar isDelete={true} action={deleteUserGroup}/>

<main>
    <div class="inputs">
        <div class="lable-input">
            <div class="mdc-typography--headline4">Name</div>
            <Textfield bind:value={name} label="Name" style="width: 70%" />
        </div>
        <div class="lable-input">
            <div class="mdc-typography--headline4">Describe</div>
            <Textfield bind:value={describe} label="Describe" style="width: 70%" />
        </div>
    </div>
    <div class="buttons">
        <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
            <Label>Cancel</Label>
        </Button>
        <Button on:click={updateUserGroup} variant="unelevated" style="min-width: 170px">
            <Label>Save</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin: 5%;
    }
    main > div.inputs {
        display: flex;
        flex-direction: column;
    }
    main > div.buttons {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        margin-top: 18%;
    }
    
    div.inputs > div.lable-input {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
    }
</style>
