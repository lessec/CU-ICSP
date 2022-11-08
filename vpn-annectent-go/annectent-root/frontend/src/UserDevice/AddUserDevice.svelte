<script>
    import Button, { Label } from '@smui/button';
    import Textfield from '@smui/textfield';
    import CommonAppBar from '../Components/CommonAppBar.svelte';

    import {navigate} from "svelte-routing";
    import { onMount } from "svelte";

    export let id;

    let name = '';
    let ip = '';

    // request add a device
    async function addUserDevice() {
        let url = `http://localhost:5555/api/v1/device/add/${name}/ip/${ip}/u/${id}`;

        const options = {
            mode: "no-cors",
            method: 'POST',
        }

        await fetch(url, options).then(res => {
            console.log(res);
            alert("Success");
            navigate('/user/detail/' + id, {replace: true});
        });
    }

    function goToBack() {
        navigate('/user/detail/' + id, {replace: true});
    }

    async function getNewIp() {
		let response = await fetch("http://127.0.0.1:5555/api/v1/device/newip");
		let commits = await response.json();
		ip = commits.ip;
	}
    
    onMount(async () => {
        await getNewIp();
    });
</script>

<CommonAppBar />

<main>
    <div class="inputs">
        <div class="lable-input">
            <div class="mdc-typography--headline4">Device Name</div>
            <Textfield bind:value={name} label="Enter the device name" style="width: 70%" />
        </div>
        <div class="lable-input">
            <div class="mdc-typography--headline4">Address</div>
            <Textfield bind:value={ip} label="Enter internal IP address" style="width: 70%" disabled/>
        </div>
    </div>
    <div class="buttons">
        <Button on:click={goToBack} variant="outlined" style="min-width: 170px">
            <Label>Cancel</Label>
        </Button>
        <Button on:click={addUserDevice} variant="unelevated" style="min-width: 170px">
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
