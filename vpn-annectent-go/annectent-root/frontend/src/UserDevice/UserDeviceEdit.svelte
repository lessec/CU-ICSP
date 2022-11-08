<script>
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import {navigate} from "svelte-routing";
    import {onMount} from "svelte";

    export let id;
    export let did;

    let name = '';
    let ip = '';

    // get a device from server
    async function getData(did) {
        let url = `http://127.0.0.1:5555/api/v1/device/get/id/${did}`;

        let response = await fetch(url);
        console.log(response);
        return await response.json();
    }

    // lifecycle function
    onMount(async () => {
        getData(did).then(data => {
            name = data.name;
            ip = data.ip;
        });
    });

    // request to update a device by id
    async function updateUserDevice() {
        let url = `http://127.0.0.1:5555/api/v1/device/update/${did}/n/${name}/ip/${ip}`;
        
        const options = {
            mode: 'no-cors',
            method: 'POST',
        }

        await fetch(url, options).then(response => {
            console.log(response);
            alert("Success!");
            navigate('/user/detail/' + id, { replace: true });
        });
    }

    function goToBack() {
        navigate('/user/detail/' + id, { replace: true });
    } 

    async function deleteUserDevice() {
        let url = `http://127.0.0.1:5555/api/v1/device/delete/${did}`;

        const options = {
            method: 'POST',
        }

        await fetch(url, options).then(response => {
            console.log(response);
            alert("Device deleted");
            navigate('/user/detail/' + id, { replace: true });
        });
    }
</script>

<CommonAppBar isDelete={true} action={deleteUserDevice}/>

<main>
    <div class="inputs">
        <div class="lable-input">
            <div class="mdc-typography--headline4">Device Name</div>
            <Textfield bind:value={name} label="Enter the device name" style="width: 70%" />
        </div>
        <div class="lable-input">
            <div class="mdc-typography--headline4">Address</div>
            <Textfield bind:value={ip} label="Enter internal IP address" style="width: 70%" />
        </div>
    </div>
    <div class="buttons">
        <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
            <Label>Cancel</Label>
        </Button>
        <Button on:click={updateUserDevice} variant="unelevated" style="min-width: 170px">
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
