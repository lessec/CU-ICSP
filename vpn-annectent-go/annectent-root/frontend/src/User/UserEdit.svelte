<script>
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import FormField from '@smui/form-field';
    import Radio from '@smui/radio';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import { navigate } from "svelte-routing";
    import { onMount } from "svelte";
    
    export let id;

    let name = '';
    let email = '';
    let type = '';
    
    // get an user from server
    async function getUser(id) {
        let url = `http://127.0.0.1:5555/api/v1/user/get/id/${id}`;
        
        let response = await fetch(url);
        console.log(response);
        let data = await response.json();
        console.log(data);
        return data;
    }

    // lifecycle function
    onMount(async ()=> {
        getUser(id).then(data => {
            console.log(data);
            name = data.name;
            email = data.email;
            type = data.type;
        });
    });
    
    // request to update an user
    async function updateUser() {
        let url = `http://127.0.0.1:5555/api/v1/user/update/${id}`;

        const params = {
            name,
            email,
            type,
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
            navigate('/user/detail/' + id, {replace: true});
        });
    }

    function goToBack() {
        navigate('/user/detail/' + id, {replace: true});
    }

    async function deleteUser() {
        let  url = 'http://127.0.0.1:5555/api/v1/user/delete/' + id;

        const options = {
            method: 'POST',
        };
        
        let res = await fetch(url, options).then(res => {
            console.log("new group added!");
            alert("Success!");
            navigate('/user', {replace: true});
        });
    }

</script>

<CommonAppBar isDelete={true} action={deleteUser}/>

<main>
    <div class="inputs">
        <div class="lable-input">
            <div class="mdc-typography--headline4">Name</div>
            <Textfield bind:value={name} label="Name" style="width: 70%" />
        </div>
        <div class="lable-input">
            <div class="mdc-typography--headline4">User type</div>
            {#each ['Admin', 'Manager', 'Staff'] as option}
            <FormField style="margin-right: 1em;">
                <Radio bind:group={type} value={option} />
                <span slot="label">{`${option[0].toUpperCase()}${option.slice(1)}`}</span>
            </FormField>
            {/each}
        </div>
    </div>
    <div class="buttons">
        <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
            <Label>Cancel</Label>
        </Button>
        <Button on:click={updateUser} variant="unelevated" style="min-width: 170px">
            <Label>Save</Label>
        </Button>
    </div>
</main>

<style>
    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        margin: 10%;
        margin-top: 5%;
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
        margin-bottom: 30px;
    }
</style>
